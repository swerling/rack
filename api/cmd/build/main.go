package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/convox/rack/api/cmd/build/source"
	"github.com/convox/rack/api/structs"
	"github.com/convox/rack/manifest"
	"github.com/convox/rack/provider"
)

var (
	flagApp    string
	flagAuth   string
	flagCache  string
	flagID     string
	flagConfig string
	flagMethod string
	flagPush   string
	flagUrl    string

	currentBuild    *structs.Build
	currentLogs     string
	currentManifest string
	currentProvider provider.Provider

	event *structs.Event
)

func init() {
	currentProvider = provider.FromEnv()

	var buf bytes.Buffer

	currentProvider.Initialize(structs.ProviderOptions{
		LogOutput: &buf,
	})
}

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.StringVar(&flagApp, "app", "example", "app name")
	fs.StringVar(&flagAuth, "auth", "", "docker auth data (json)")
	fs.StringVar(&flagCache, "cache", "true", "use docker cache")
	fs.StringVar(&flagConfig, "config", "docker-compose.yml", "path to app config")
	fs.StringVar(&flagID, "id", "latest", "build id")
	fs.StringVar(&flagMethod, "method", "", "source method")
	fs.StringVar(&flagPush, "push", "", "push to registry")
	fs.StringVar(&flagUrl, "url", "", "source url")

	if err := fs.Parse(os.Args[1:]); err != nil {
		fail(err)
	}

	if v := os.Getenv("BUILD_APP"); v != "" {
		flagApp = v
	}

	if v := os.Getenv("BUILD_AUTH"); v != "" {
		flagAuth = v
	}

	if v := os.Getenv("BUILD_CONFIG"); v != "" {
		flagConfig = v
	}

	if v := os.Getenv("BUILD_ID"); v != "" {
		flagID = v
	}

	if v := os.Getenv("BUILD_PUSH"); v != "" {
		flagPush = v
	}

	if v := os.Getenv("BUILD_URL"); v != "" {
		flagUrl = v
	}

	event = &structs.Event{
		Action: "build:create",
		Data: map[string]string{
			"app": flagApp,
			"id":  flagID,
		},
	}

	if err := execute(); err != nil {
		fail(err)
	}

	if err := success(); err != nil {
		fail(err)
	}

	event.Status = "success"
	event.Data["release_id"] = currentBuild.Release
	if err := currentProvider.EventSend(event, nil); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	}
	time.Sleep(1 * time.Second)
}

func execute() error {
	b, err := currentProvider.BuildGet(flagApp, flagID)
	if err != nil {
		return err
	}

	currentBuild = b

	if err := login(); err != nil {
		return err
	}

	dir, err := fetch()
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir)

	data, err := ioutil.ReadFile(filepath.Join(dir, flagConfig))
	if err != nil {
		return err
	}

	currentBuild.Manifest = string(data)

	if err := build(dir); err != nil {
		return err
	}

	return nil
}

func fetch() (string, error) {
	var s source.Source

	switch flagMethod {
	case "git":
		s = &source.SourceGit{flagUrl}
	case "index":
		s = &source.SourceIndex{flagUrl}
	case "tgz":
		s = &source.SourceTgz{flagUrl}
	case "zip":
		s = &source.SourceZip{flagUrl}
	default:
		return "", fmt.Errorf("unknown method: %s", flagMethod)
	}

	var buf bytes.Buffer

	dir, err := s.Fetch(&buf)
	log(strings.TrimSpace(buf.String()))
	if err != nil {
		return "", err
	}

	return dir, nil
}

func login() error {
	var auth map[string]struct {
		Username string
		Password string
	}

	if err := json.Unmarshal([]byte(flagAuth), &auth); err != nil {
		return err
	}

	for host, entry := range auth {
		out, err := exec.Command("docker", "login", "-u", entry.Username, "-p", entry.Password, host).CombinedOutput()
		log(fmt.Sprintf("Authenticating %s: %s", host, strings.TrimSpace(string(out))))
		if err != nil {
			return err
		}
	}

	return nil
}

func build(dir string) error {
	dcy := filepath.Join(dir, flagConfig)

	if _, err := os.Stat(dcy); os.IsNotExist(err) {
		return fmt.Errorf("no such file: %s", flagConfig)
	}

	data, err := ioutil.ReadFile(dcy)
	if err != nil {
		return err
	}

	m, err := manifest.Load(data)
	if err != nil {
		return err
	}

	errs := m.Validate()
	if len(errs) > 0 {
		return errs[0]
	}

	s := make(chan string)

	go func() {
		for l := range s {
			log(l)
		}
	}()

	defer close(s)

	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("apps/%s/cache.tgz", flagApp)
	cacheDir := filepath.Join(tmp, "cache")
	cacheArchive := filepath.Join(tmp, "cache.tgz")

	if err := os.Mkdir(cacheDir, 0755); err != nil {
		return err
	}

	if currentProvider.ObjectExists(cacheKey) {
		r, err := currentProvider.ObjectFetch(cacheKey)
		if err != nil {
			return err
		}

		fd, err := os.Create(cacheArchive)
		if err != nil {
			return err
		}

		defer fd.Close()

		if _, err := io.Copy(fd, r); err != nil {
			return err
		}

		if err := fd.Close(); err != nil {
			return err
		}

		if err := exec.Command("tar", "xzf", cacheArchive, "-C", cacheDir).Run(); err != nil {
			return err
		}
	}

	err = m.Build(dir, flagApp, s, manifest.BuildOptions{
		Environment: env,
		Cache:       flagCache == "true",
		CacheDir:    cacheDir,
	})
	if err != nil {
		return err
	}

	if err := m.Push(flagPush, flagApp, flagID, s); err != nil {
		return err
	}

	if err := exec.Command("tar", "czf", cacheArchive, "-C", cacheDir, ".").Run(); err != nil {
		return err
	}

	fd, err := os.Open(cacheArchive)
	if err != nil {
		return err
	}

	if _, err := currentProvider.ObjectStore(cacheKey, fd, structs.ObjectOptions{}); err != nil {
		return err
	}

	return nil
}

type ChanWriter chan string

func (cw ChanWriter) Write(data []byte) (int, error) {
	cw <- string(data)
	return len(data), nil
}

func success() error {
	_, err := currentProvider.BuildRelease(currentBuild)
	if err != nil {
		return err
	}

	url, err := currentProvider.ObjectStore(fmt.Sprintf("build/%s/logs", currentBuild.Id), bytes.NewReader([]byte(currentLogs)), structs.ObjectOptions{})
	if err != nil {
		return err
	}

	currentBuild.Ended = time.Now()
	currentBuild.Logs = url
	currentBuild.Status = "complete"

	if err := currentProvider.BuildSave(currentBuild); err != nil {
		return err
	}

	return nil
}

func fail(err error) {
	log(fmt.Sprintf("ERROR: %s", err))
	if e := currentProvider.EventSend(event, err); e != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", e)
	}

	url, _ := currentProvider.ObjectStore(fmt.Sprintf("build/%s/logs", currentBuild.Id), bytes.NewReader([]byte(currentLogs)), structs.ObjectOptions{})

	currentBuild.Ended = time.Now()
	currentBuild.Logs = url
	currentBuild.Reason = err.Error()
	currentBuild.Status = "failed"

	if err := currentProvider.BuildSave(currentBuild); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	}

	os.Exit(1)
}

func log(line string) {
	currentLogs += fmt.Sprintf("%s\n", line)
	fmt.Println(line)
}
