package handler

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/convox/rack/api/crypt"
	"github.com/convox/rack/api/models"
	awshelp "github.com/convox/rack/provider/aws"
)

// Parses as [host]:[container]/[protocol?], where [protocol] is optional
var portMappingRegex = regexp.MustCompile(`^(\d+):(\d+)(?:/(udp|tcp))?$`)

func HandleECSService(req Request) (string, map[string]string, error) {
	switch req.RequestType {
	case "Create":
		return "invalid", nil, fmt.Errorf("creation of Custom::ECSService no longer supported")
	case "Update":
		return "invalid", nil, fmt.Errorf("updating Custom::ECSService no longer supported")
	case "Delete":
		fmt.Println("DELETING SERVICE")
		fmt.Printf("req %+v\n", req)
		return ECSServiceDelete(req)
	}

	return "invalid", nil, fmt.Errorf("unknown RequestType: %s", req.RequestType)
}

func HandleECSTaskDefinition(req Request) (string, map[string]string, error) {
	switch req.RequestType {
	case "Create":
		fmt.Println("CREATING TASK")
		fmt.Printf("req %+v\n", req)
		return ECSTaskDefinitionCreate(req)
	case "Update":
		fmt.Println("UPDATING TASK")
		fmt.Printf("req %+v\n", req)
		return ECSTaskDefinitionCreate(req)
	case "Delete":
		fmt.Println("DELETING TASK")
		fmt.Printf("req %+v\n", req)
		return ECSTaskDefinitionDelete(req)
	}

	return "invalid", nil, fmt.Errorf("unknown RequestType: %s", req.RequestType)
}

func ECSServiceDelete(req Request) (string, map[string]string, error) {
	cluster := req.ResourceProperties["Cluster"].(string)

	// arn:aws:ecs:us-east-1:922560784203:service/sinatra-SZXTRXEMYEY
	parts := strings.Split(req.PhysicalResourceId, "/")
	name := parts[1]

	// Set Desired to 0
	_, err := ECS(req).UpdateService(&ecs.UpdateServiceInput{
		Cluster:      aws.String(cluster),
		Service:      aws.String(name),
		DesiredCount: aws.Int64(0),
	})

	// go ahead and mark the delete good if the service is not found
	if ae, ok := err.(awserr.Error); ok {
		if ae.Code() == "ServiceNotFoundException" {
			return req.PhysicalResourceId, nil, nil
		}
	}

	// signal DELETE_FAILED to cloudformation
	if err != nil {
		fmt.Printf("ECS UpdateService error: %s\n", err)
		return req.PhysicalResourceId, nil, err
	}

	if err := waitForServiceStop(req, cluster, name); err != nil {
		return req.PhysicalResourceId, nil, err
	}

	_, err = ECS(req).DeleteService(&ecs.DeleteServiceInput{
		Cluster: aws.String(cluster),
		Service: aws.String(name),
	})

	return req.PhysicalResourceId, nil, err
}

// GetS3EnvironmentFromURL gets environment variables that are stored at the provided URL
func GetS3EnvironmentFromURL(req Request, url string) (models.Environment, error) {
	bucket, key, _, err := awshelp.ParseS3Url(url)
	if err != nil {
		return nil, err
	}

	input := s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	output, err := S3(req).GetObject(&input)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	if key, ok := req.ResourceProperties["Key"].(string); ok && key != "" {
		cr := crypt.New(*Region(&req), os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"))
		cr.AwsToken = os.Getenv("AWS_SESSION_TOKEN")

		dec, err := cr.Decrypt(key, data)

		if err != nil {
			return nil, err
		}

		data = dec
	}

	return models.LoadEnvironment(data)
}

func ECSTaskDefinitionCreate(req Request) (string, map[string]string, error) {
	// return "", fmt.Errorf("fail")

	tasks := req.ResourceProperties["Tasks"].([]interface{})

	r := &ecs.RegisterTaskDefinitionInput{
		Family: aws.String(req.ResourceProperties["Name"].(string)),
	}

	// get environment from S3 URL
	// 'Environment' is a CloudFormation Template Property that references 'Environment' CF Parameter with S3 URL
	// S3 body may be encrypted with KMS key
	var env models.Environment

	if taskRole, ok := req.ResourceProperties["TaskRole"].(string); ok && taskRole != "" {
		r.TaskRoleArn = &taskRole
	}

	if envUrl, ok := req.ResourceProperties["Environment"].(string); ok && envUrl != "" {
		var err error
		env, err = GetS3EnvironmentFromURL(req, envUrl)
		if err != nil {
			return "invalid", nil, err
		}
	}

	r.ContainerDefinitions = make([]*ecs.ContainerDefinition, len(tasks))

	for i, itask := range tasks {
		task := itask.(map[string]interface{})

		cpu := 0
		var err error

		if c, ok := task["Cpu"].(string); ok && c != "" {
			cpu, err = strconv.Atoi(c)
			if err != nil {
				return "invalid", nil, err
			}
		}

		memory, err := strconv.Atoi(task["Memory"].(string))
		if err != nil {
			return "invalid", nil, err
		}

		privileged := false

		if p, ok := task["Privileged"].(string); ok && p != "" {
			privileged, err = strconv.ParseBool(p)
			if err != nil {
				return "invalid", nil, err
			}
		}

		r.ContainerDefinitions[i] = &ecs.ContainerDefinition{
			Name:       aws.String(task["Name"].(string)),
			Essential:  aws.Bool(true),
			Image:      aws.String(task["Image"].(string)),
			Cpu:        aws.Int64(int64(cpu)),
			Memory:     aws.Int64(int64(memory)),
			Privileged: aws.Bool(privileged),
		}

		// set Command from either -
		// a single string (shell form) - ["sh", "-c", command]
		// an array of strings (exec form) - ["cmd1", "cmd2"]
		switch commands := task["Command"].(type) {
		case string:
			if commands != "" {
				r.ContainerDefinitions[i].Command = []*string{aws.String("sh"), aws.String("-c"), aws.String(commands)}
			}
		case []interface{}:
			r.ContainerDefinitions[i].Command = make([]*string, len(commands))
			for j, command := range commands {
				r.ContainerDefinitions[i].Command[j] = aws.String(command.(string))
			}
		}

		// set Task environment from CF Tasks[].Environment key/values
		// These key/values are read from the app manifest environment hash
		if oenv, ok := task["Environment"].(map[string]interface{}); ok {
			for key, val := range oenv {
				r.ContainerDefinitions[i].Environment = append(r.ContainerDefinitions[i].Environment, &ecs.KeyValuePair{
					Name:  aws.String(key),
					Value: aws.String(val.(string)),
				})
			}
		}

		// set Task environment from decrypted S3 URL body of key/values
		// These key/values take precident over the above environment
		for key, val := range env {
			r.ContainerDefinitions[i].Environment = append(r.ContainerDefinitions[i].Environment, &ecs.KeyValuePair{
				Name:  aws.String(key),
				Value: aws.String(val),
			})
		}

		// set Release value in Task environment
		if release, ok := req.ResourceProperties["Release"].(string); ok {
			r.ContainerDefinitions[i].Environment = append(r.ContainerDefinitions[i].Environment, &ecs.KeyValuePair{
				Name:  aws.String("RELEASE"),
				Value: aws.String(release),
			})
		}

		// set links
		if links, ok := task["Links"].([]interface{}); ok {
			r.ContainerDefinitions[i].Links = make([]*string, len(links))

			for j, link := range links {
				r.ContainerDefinitions[i].Links[j] = aws.String(link.(string))
			}
		}

		// set portmappings
		if ports, ok := task["PortMappings"].([]interface{}); ok {
			r.ContainerDefinitions[i].PortMappings = make([]*ecs.PortMapping, len(ports))

			for j, port := range ports {
				parts := portMappingRegex.FindStringSubmatch(port.(string))
				host, _ := strconv.Atoi(parts[1])
				container, _ := strconv.Atoi(parts[2])
				protocol := strings.ToLower(parts[3])
				if protocol != "tcp" && protocol != "udp" {
					protocol = "tcp" // default
				}

				r.ContainerDefinitions[i].PortMappings[j] = &ecs.PortMapping{
					ContainerPort: aws.Int64(int64(container)),
					HostPort:      aws.Int64(int64(host)),
					Protocol:      aws.String(protocol),
				}
			}
		}

		// set volumes
		if volumes, ok := task["Volumes"].([]interface{}); ok {
			for j, volume := range volumes {
				name := fmt.Sprintf("%s-%d-%d", task["Name"].(string), i, j)
				parts := strings.Split(volume.(string), ":")

				r.Volumes = append(r.Volumes, &ecs.Volume{
					Name: aws.String(name),
					Host: &ecs.HostVolumeProperties{
						SourcePath: aws.String(parts[0]),
					},
				})

				r.ContainerDefinitions[i].MountPoints = append(r.ContainerDefinitions[i].MountPoints, &ecs.MountPoint{
					SourceVolume:  aws.String(name),
					ContainerPath: aws.String(parts[1]),
					ReadOnly:      aws.Bool(false),
				})
			}
		}

		// set extra hosts
		if extraHosts, ok := task["ExtraHosts"].([]interface{}); ok {
			for _, host := range extraHosts {
				hostx, oky := host.(map[string]interface{})
				if oky {
					r.ContainerDefinitions[i].ExtraHosts = append(r.ContainerDefinitions[i].ExtraHosts, &ecs.HostEntry{
						Hostname:  aws.String(hostx["HostName"].(string)),
						IpAddress: aws.String(hostx["IpAddress"].(string)),
					})
				}
			}
		}
	}

	res, err := ECS(req).RegisterTaskDefinition(r)
	if err != nil {
		return "invalid", nil, err
	}

	return *res.TaskDefinition.TaskDefinitionArn, nil, nil
}

func ECSTaskDefinitionDelete(req Request) (string, map[string]string, error) {
	// We have observed a race condition quickly deregistering then re-registering
	// Task Definitions, where the Register fails. We work around this by not
	// deregistering any Task Definitions.
	// _, err := ECS(req).DeregisterTaskDefinition(&ecs.DeregisterTaskDefinitionInput{TaskDefinition: aws.String(req.PhysicalResourceId)})
	return req.PhysicalResourceId, nil, nil
}

var idAlphabet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateId(prefix string, size int) string {
	b := make([]rune, size)
	for i := range b {
		b[i] = idAlphabet[rand.Intn(len(idAlphabet))]
	}
	return prefix + string(b)
}

func waitForServiceStop(req Request, cluster, name string) error {
	timeout := time.After(4 * time.Minute)
	tick := time.Tick(10 * time.Second)

	for {
		select {
		case <-tick:
			tasks, err := ECS(req).ListTasks(&ecs.ListTasksInput{
				Cluster:     aws.String(cluster),
				ServiceName: aws.String(name),
			})
			if err != nil {
				return err
			}
			if len(tasks.TaskArns) == 0 {
				return nil
			}
		case <-timeout:
			return fmt.Errorf("timeout stopping service")
		}
	}
}
