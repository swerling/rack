// Code generated by go-bindata.
// sources:
// models/templates/app.tmpl
// DO NOT EDIT!

package models

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7d\x7b\x73\xdb\x38\x92\xf8\xff\xf9\x14\x28\xfe\xf2\x3b\xd9\x13\x59\xb6\x33\xb7\x75\xb7\xdc\xcb\x55\x39\xb2\x33\xf1\xae\x9d\xe8\x2c\x27\x53\xb7\x89\x6b\x0b\x26\x21\x89\x6b\x0a\xe0\x00\x90\x13\x8f\x4a\xdf\xfd\x0a\x0f\x92\x78\x92\xf2\x23\x3b\xb3\x5b\xc9\xd5\xcd\xda\x64\xa3\xd1\xe8\x6e\x34\xfa\x05\x7a\xbd\x06\x39\x9a\x15\x18\x81\x04\x56\x55\x02\x36\x9b\x67\x00\xac\xd7\xe0\x39\xac\x2a\x90\xbe\x02\xa3\xa3\xaa\x6a\x1f\x2e\x21\x2e\x66\x88\x71\xf9\xe6\xbc\xfe\x45\xbd\x7e\x06\x00\x00\xc9\xd1\xcf\xd3\x4b\xb4\xac\x4a\xc8\xd1\x1b\x42\x97\x90\x7f\x44\x94\x15\x04\x27\x20\x05\xc9\xcb\x83\xc3\x83\xbd\x83\x3f\xee\x1d\xfc\x31\x19\x2a\xf0\x31\xc1\x79\xc1\x0b\x82\x59\x92\x6a\x14\x72\x26\xae\x71\x80\xe4\x1a\x96\x10\x67\x88\xee\x65\x2d\xa8\x3b\xb7\x37\xa8\xa2\x24\x43\x8c\xf5\x8d\x49\x4e\x31\x47\x14\xc3\x52\x4c\x0e\x92\x37\x38\x4d\x4f\x7e\x59\xc1\x52\x10\xf3\x49\x3c\xb9\x40\xb3\x24\x35\xc0\xc0\x66\x08\x92\xff\x45\x2c\x01\x57\x60\x33\xac\xb1\xbc\x83\xbc\xb8\x45\x67\x64\x3e\x2f\xf0\xbc\x07\x95\x0d\x1b\xc6\x37\xa1\xc5\x2d\xe4\xa8\x07\x53\x0d\x15\xc6\xf1\xba\x84\xf8\x66\x8a\xb2\x15\x2d\xf8\xdd\x4f\x94\xac\x2a\x21\x81\xb5\x89\x0e\xa4\xe0\xd3\x5a\x62\x13\xb2\xb1\x61\x05\xce\xe4\x4a\xf1\x49\x23\x4d\x26\x90\xc2\x25\xe2\x88\xca\xa1\xdd\xc2\xaa\x04\xec\x3d\x04\x15\x84\xaf\xd7\x32\x2e\x57\x8c\x23\x6a\x68\x08\x00\xc9\xe5\x5d\x85\x14\xe1\x9c\x0a\x56\x0e\xdb\x57\xc7\x68\x06\x57\x25\x97\x6f\xed\xe7\x2c\xa3\x45\xc5\x6b\x75\x4c\xf4\xab\x96\x6b\xc7\xa8\x2a\xc9\xdd\x12\x61\x7e\x0e\xbf\x16\xcb\xd5\x32\x30\xa7\x10\xe2\x6a\x79\x8d\x68\x68\x4a\xa9\xe4\x07\xb1\x49\x53\x90\x68\xbc\xa0\x42\x34\x43\x98\xc3\x39\x02\x64\x06\x34\x1b\x10\x03\x9c\x80\x1b\x84\x2a\x40\x57\x18\x17\x78\x0e\xbe\x2c\x8a\x12\x81\x5c\xd2\x25\x96\xd9\x45\x72\x81\x1f\x48\xf2\x1f\x3a\x29\x56\x68\x9f\x8e\xe2\x13\x7c\x5b\x50\x82\x05\xc9\x61\x5a\xe3\x12\xed\x10\x68\x50\x9e\xe6\xfe\xde\x6e\x1e\x0b\xe1\x7b\x5c\xde\x01\x58\x96\xe4\x0b\x80\x99\x58\xae\x58\x2c\x5f\x14\x0c\x08\xeb\x38\xa3\x64\x09\x0a\xcc\x8a\x1c\x01\xbe\x40\xe0\xe3\x64\x1c\xa1\xf9\x1d\x31\x5f\x1c\x09\x84\x28\xff\x08\xcb\x15\x52\x9b\x5a\x6e\xdf\xa1\x84\x03\x57\xde\x22\xfe\x82\xee\xbe\x35\x9f\x3c\x0b\xf6\x00\x66\x7d\x60\x08\x60\x89\x07\x9c\x8c\xa7\x00\x7e\x61\x25\x99\x33\x90\xd3\xe2\x16\x51\x40\xc4\x7f\x32\x82\x6f\xc9\xd7\x7d\x38\x47\x98\x83\x19\xa1\xa0\xd4\x33\x7e\x0b\xb6\x19\x56\xf4\x81\xab\x99\xae\xae\x31\xe2\x4c\x23\x12\xb2\x67\x15\xca\x8a\xd9\x9d\x10\xf5\x9e\x94\x7b\x49\x60\x0e\x6a\xab\x07\x10\xce\x2b\x52\x60\xce\xbe\xc9\x82\x2e\x50\x89\x20\x0b\x2d\xe8\xa9\xcd\xe0\x05\xaa\x08\x2b\x38\xa1\x21\xc5\x7b\xdc\x64\x53\xb2\xa2\x19\x02\x19\xc9\x11\xa0\xed\x34\x1e\x09\xf6\x71\xf4\xd4\x54\x5c\x2e\x10\x38\xb3\x44\xc7\xf4\x7c\x60\x2e\x26\x94\xca\x59\x6f\xf4\x00\x71\x4a\x31\x22\x64\x9d\x15\x8c\xff\xd7\xd1\xcf\xd3\x34\x3d\x19\xbf\x4c\x53\x05\x9c\xa6\xa7\xf9\x7f\x3f\x84\xd4\x8f\x93\x31\x60\x6a\xbe\xed\xa8\x8a\xeb\xfd\xb7\x21\xae\xd2\xdb\x63\x3b\x22\x2f\x21\xbb\xb9\x20\xe5\xd3\x6b\xf1\xe9\xd1\x39\x10\x88\xc5\x36\x85\x55\x55\xde\x89\x1f\x84\x21\x12\x33\x32\x71\x64\xc5\x89\xaa\x7d\x54\x8b\x26\xc7\x20\xec\x5c\x9c\xfc\xcf\x87\xd3\x8b\x93\xe3\x5d\x70\x06\x97\xd7\x39\x04\xe3\x15\xe3\x64\x79\x49\xaa\x22\x03\x6f\x21\xce\x4b\x44\x81\xde\xa3\xa0\xc6\x68\x10\x7c\x5e\xe0\x33\x84\xe7\x7c\x21\xc9\x3d\x34\x5f\x39\x56\xc9\xa7\x6f\x32\x8e\xf0\xab\x95\xe4\xc7\xc9\x58\x88\xf1\xa1\x52\xec\x96\xda\xc7\xc9\x78\x7c\x7a\x7c\xf1\xe4\x42\x13\x33\x0b\xc4\xe1\xe9\x2d\xef\xf3\x1c\x56\x55\x81\xe7\xe6\xa6\x13\xd1\x03\xe3\x10\xbb\x3b\xb1\xb1\xa8\xed\x33\xe9\x76\x52\x88\xe7\xa8\x75\x31\x47\xaf\xf5\xe6\x67\xad\x73\xda\xc0\x3e\x6f\x2c\x83\x08\x75\x02\x00\x1a\xd9\x84\x50\x5e\x93\xe6\x42\x01\x90\xac\xd7\x60\x55\x55\x88\xb6\xf8\x46\x13\xe5\x39\xbd\x83\x4b\x04\x36\x1b\x31\x7e\xbd\x06\x0d\x2d\x60\xb3\x79\x4b\x98\x3c\x2d\xd6\x6b\x50\xe0\x1c\x7d\x35\xc6\x5e\x40\x9c\x93\x25\x03\x3b\x05\x27\xb0\x1d\xb4\x0b\x36\x1b\x83\xcd\x9a\x42\x84\x73\x9b\xa2\xd0\xb3\xe4\x1d\x91\xec\x52\xc1\x86\x0e\x2d\xa4\x56\xd5\x2f\x40\x0b\xbe\xf1\x0f\x59\x42\xf9\x84\x12\x4e\x32\xe2\x38\x59\x0b\xce\x2b\x85\x54\x98\x1c\x84\x11\x35\xe0\x92\xb7\x97\x97\x13\x71\xd2\x9d\x4a\x01\x66\x28\xf4\x4e\x1e\x01\x28\x06\x31\x4d\x5a\x32\xf4\x74\xac\x7b\xbe\xe9\xa3\x27\xb4\x66\xe4\x59\xc7\xfa\x2e\xc7\xd1\xe5\xe9\x57\xf1\xc9\xa6\xd3\x33\x77\xaa\xb2\x63\x69\x02\xfc\x71\x53\xb5\x72\xb5\x76\xdc\x05\x62\xf2\xb0\xb6\xb7\x5c\x6b\xf4\x22\x66\xbc\xd6\x9f\xd3\xa3\xf3\x34\x95\x30\xc6\x4a\x26\x94\x54\x88\xf2\xc2\xdb\x9e\xc9\x11\x63\xab\x25\x12\xf0\x13\x52\x16\xd9\xdd\x31\xc9\x56\x5e\x88\x00\x6c\x6b\x9d\xbc\x3c\x38\x7c\xb9\x77\x78\xb0\x77\xf8\x1f\x8e\xfa\x27\x53\x0e\x39\xd2\xe3\x3f\x39\xbb\x72\xed\xfc\x2e\x62\x92\xd9\x0c\x65\x72\xd7\x49\xaf\xcc\xc1\xa6\x49\x2f\x70\x56\x54\x75\xb2\x60\x8a\xe8\x6d\x91\x21\xe5\xb7\x95\xf2\x44\x18\xc1\x25\xfc\x95\x60\xf8\x85\x8d\x32\xb2\xb4\xe2\x71\x73\xa1\x99\x3e\x52\x3e\x81\x84\x71\x96\xb6\x0b\x6f\x9d\xbe\xfa\x9f\x6d\x4f\xcc\xb7\x16\xe6\x64\x02\xf9\x42\x10\xbf\xaf\x3d\xec\xc4\x7e\x2b\x18\xaa\x58\x6e\xb3\xc2\x65\x84\x82\xbc\x13\xa6\x49\xb2\x22\x5f\x16\xb8\x60\x9c\x42\x4e\xa8\xc7\x92\xa4\x47\x4e\x60\x5b\x59\x01\x4f\x5e\x82\xbf\x9e\x44\x0c\xce\x25\x3f\x88\x5f\x6b\xfd\x54\x0f\xc0\xa6\x87\x7b\xe6\x6f\x57\x1d\x26\xcd\xd0\xf0\x0e\xed\x56\x3e\x40\x9a\xbe\x59\x61\x45\xd5\x56\x4a\x3e\x26\x39\xf2\x15\x7a\xfa\xe3\xeb\x55\x76\x83\x78\x9b\xf0\xf9\x33\x29\xb4\x86\xec\x25\x43\xf1\x3f\x4a\xae\xc9\xd0\x35\xd2\x17\x68\x2e\xcf\xd2\x0d\xb8\xf2\xd5\x2d\x99\xfe\xa8\x63\x47\x17\xab\x42\x4a\x95\xb3\xb2\x6f\xa1\x6d\xf2\x75\x9b\x21\x48\xf6\x95\x62\xef\xcf\x64\x2a\xaf\x20\x78\xf4\x6b\x51\x25\x6a\xae\xa8\x32\x6a\x5f\x48\x20\x93\xe7\xd7\x08\x7d\xd5\x51\xb8\x05\x76\x8e\x96\x84\xde\x4d\x8b\x5f\x25\x53\x0f\x5f\xfe\xa7\xfd\xba\xb6\x2e\x8a\xf4\x9f\x10\x3f\xe2\x4a\x37\x3c\x13\x24\x34\x83\x62\x6f\xbb\x25\x17\x2b\xcc\x0b\xa5\xc9\x98\xe4\xe8\xef\xec\xdf\x47\x3f\xda\x73\x5c\x16\x4b\x44\x56\x52\xc9\x7e\x3c\x38\x48\x3a\x94\x42\x99\x50\x23\x73\xf1\xdd\xf8\x35\xc6\x0f\x65\x6c\x8f\x0b\x27\xfb\x5f\xcd\xfe\x79\x42\x57\xaf\xff\x81\x96\xd0\xc7\xb3\x95\xb8\x2c\x2e\x07\xde\x02\x90\xdc\x2c\x59\x7a\x82\x33\x7a\x57\xf1\x20\x02\x0d\x72\x8c\x14\x48\x00\xe2\x2a\x38\xaf\x61\x94\xc3\x33\xb7\x96\x46\xd8\x26\xe0\x7a\xcc\xc0\x91\xb3\xfa\xe7\x42\x3d\xc8\xb2\x87\x33\xd5\xb4\xf1\x72\x4c\x37\xdf\x82\xcd\x28\xc1\x7f\x27\xd7\xdb\x80\xca\xd4\xd6\x36\x80\x75\xe2\xdb\x04\xdd\x32\x57\xce\xd4\xe6\xeb\x40\x4e\xd1\x5c\x1c\xdb\x77\x51\xec\xa1\x41\x75\x06\x24\x89\x20\x65\x5c\x55\x1b\x6c\x27\xf1\xfd\x8a\x57\x2b\xde\x5f\xbd\x21\x1a\x0e\x8c\xba\x17\xd7\xc2\xf5\x70\xa3\x59\x63\x78\x44\x6b\xb8\x39\x77\xc2\xc6\x3a\x42\xb4\x42\x9e\x06\xce\x75\x86\x9f\x89\xff\x6f\x23\xa7\x67\x46\xc1\x2c\x54\x65\xaa\x4b\x65\x2a\x3c\x7c\x7e\x23\xc3\xc7\x13\xcc\xa9\xb4\x2a\x4d\x88\x98\x9c\x60\x78\x5d\xa2\xbc\x8d\x0f\x6f\x44\x0c\xd7\x9c\x77\xef\x88\xb4\xdd\xc1\xfa\x8f\x78\x32\x45\xa5\xb2\x01\x9f\xc0\x81\x79\x7a\xdb\xf8\xde\xd4\xc7\xb6\x72\x10\xc4\x89\xbe\x77\x28\xed\xb2\x36\xcd\xed\xba\xba\x57\x58\x97\x67\x9c\xd5\x21\xb9\x3a\x7d\x18\xb4\x6b\x6b\x89\x40\x23\x1d\xe6\xb6\x94\x18\xb1\x44\x7d\x62\x8e\xc9\x72\x09\x8f\x51\x59\x2c\x0b\x8e\x72\x11\xe3\x24\x46\x79\xa3\xc9\x9e\x8a\xa8\x1c\x8d\xf4\x03\x59\x8c\x62\x66\xdc\xeb\x26\x6c\x54\xc5\xc3\xab\x55\xd0\x15\x1e\x82\xf1\xe4\x03\x58\xe1\x82\xab\x27\x48\xec\x28\x34\x04\x10\xe7\xe0\xfc\xb5\x18\x71\x71\x74\x6e\xbc\x49\x5a\x8d\xdf\x96\x61\x8d\x52\x4a\x9e\x24\x67\x64\x6e\x27\x32\x03\x1a\xd8\xc0\x28\x9d\x1b\xf6\xcc\x60\x6c\xed\xd8\x1c\xb6\xc3\x4a\xe6\x4c\xfe\x57\x01\x6d\x33\x45\x6b\x68\xb6\x2a\x00\x47\x8a\xc6\xc5\xac\x1d\x36\x7a\x0b\xd9\xa4\x91\x86\xd6\x17\x47\x9f\x5a\x60\x57\xb1\xc2\xaa\x75\x32\x9e\x5e\x42\x76\x73\x2c\x88\x2f\x78\x20\x8d\x57\x21\x9c\xb3\xf7\xf2\x2c\xb4\x9c\xfb\x61\xe3\xc7\x48\xe7\xe3\x2a\x90\x90\x53\xe0\x69\xea\xcf\x61\x00\x1b\x27\xfb\xe1\xe8\x60\x2b\x87\x4f\xb3\x05\x8d\x3e\x30\xe4\x79\x19\x6e\x1a\xc9\xca\x99\xfa\xee\x70\xd8\x35\x8d\x38\xc5\x42\xe0\x25\x43\x3d\x53\x68\x8d\x6c\x1e\xfa\x38\xbc\x34\x92\xe6\xe4\x25\xb9\x41\xb8\xd7\x6d\x8f\xba\xec\xda\xf3\x8a\x44\x41\x4e\xec\x33\xe5\x30\xbb\x91\x23\xa4\x65\x53\xe6\x41\x2b\x45\xe2\xc7\x43\x66\xfd\xa4\x41\x54\x3f\x73\x40\x9d\x12\x65\x03\x6e\x3e\x77\x86\x34\x91\x96\xe5\xdd\xd8\xe1\x86\x70\x92\xfb\x9d\xd0\xda\xfd\xb4\x17\xe4\x39\x9d\xa7\x4b\x38\x37\xe0\xe4\xaf\x41\x40\x4f\x41\x04\x9d\x3d\x0a\xc8\xe9\x0a\xb5\xba\x32\x83\x25\x43\x8d\xd8\xdd\x09\xd6\x6b\x85\x4a\x9a\x71\x9c\x8f\x8e\x28\x85\x77\x5e\x2a\x14\x68\x3b\x8f\xf3\x88\x6b\xd8\x9a\x01\x19\x35\x0e\xc1\x73\x54\x4a\x3f\x58\x1a\x85\x7e\xf4\x26\x31\x12\xc3\x66\x33\x5c\xaf\xc5\x0a\x36\x9b\xf5\x1a\xe1\x3c\x3a\x26\x59\xaf\xeb\xb9\x36\x9b\x90\xab\x1b\x1b\xee\xb9\xc0\x6a\x3e\xc1\x5a\x8c\x4c\x9a\x55\x86\x1c\x24\x49\x37\x5b\xd6\x6b\x70\x2b\xce\x85\xc0\xd0\x10\xdb\x43\x44\x25\xe3\x6a\xd5\xee\x20\xc3\x4d\x38\x0c\xbb\x09\x81\x13\x5a\xfb\x0a\x2e\x62\x15\xaf\x07\x71\xbf\x7c\x2c\xee\x58\x4b\x40\x03\x70\x34\x99\xd4\xaa\x2e\x0e\x97\xe8\xae\x10\xdb\xfc\x68\xfc\x17\x0d\x8b\xf0\xad\xfe\x3d\x02\x7b\xf4\xf3\xf4\x6f\x17\x27\x3f\x9d\xbe\x7f\x67\x8e\x30\x9e\x46\xc6\x9d\xbd\xff\xe9\x6f\x3f\x5d\xbc\xff\x30\x69\xd9\x71\x3a\x53\xa6\xca\x2e\xa5\x87\x62\x24\xcd\x2a\xe0\x27\xda\x3b\xa0\x3d\x07\xc1\xfe\x17\x0c\xb5\x8d\x83\x75\xd4\x71\xbe\x34\xc0\x4a\x6d\x47\x92\x1a\x90\x24\x61\x38\x75\x08\x37\x02\x30\x94\x56\x0f\xf4\x35\x15\x44\x0e\x8d\x2d\xde\xd4\x6e\x34\xba\x1b\x82\xe7\x6a\x1a\xe1\x5f\x9c\x15\xf8\xe6\x23\xf4\x0b\x37\x2d\x81\x62\x48\x43\x9f\x1e\x19\xa2\x2c\x3e\x7b\x32\xb9\x78\x3f\x3e\x99\x4e\x7d\x53\xec\x06\xa3\x9e\x3a\x7f\x24\xe5\x6a\x19\xc8\x37\x00\x47\x28\xe7\x64\x85\xb9\x88\x03\xf4\x80\x2e\xc9\x8c\x4e\xd9\xf4\x8e\x71\xb4\xec\x10\xcb\xe8\x2d\x61\x7c\xb3\x49\xd7\xeb\xd1\x98\x60\x0e\x0b\x8c\x68\x50\x81\xa3\x8e\x40\xf3\x3a\x9c\x31\xdc\xbf\x55\x84\xee\xfb\x99\x48\xe7\x34\xde\x17\x36\x55\x72\x4c\x58\xdf\x08\x61\xa1\xa4\x65\x9f\x58\x3a\xde\x44\xf7\x95\x03\xea\x99\xed\xe4\xe4\x2b\xa7\x50\xd0\xd8\x27\x33\x47\x11\xc5\xc6\x6a\x86\x9e\xc3\x2a\x22\xc0\xb0\xbc\xc4\x20\xf3\xac\xd7\x1a\x1b\xc9\xc6\x9c\x56\x47\x79\x4e\x11\x63\x35\x78\xad\xd3\xa1\x03\xeb\x5e\x8a\xfe\x08\xbe\xd5\x1e\x7a\x98\x6b\x0f\xc7\x6b\x56\x55\xbb\x25\x22\x0b\xb0\x9d\x1b\x47\x44\x75\x3b\xe8\x17\x30\xaa\x0b\x5f\xaa\x74\xb7\x0b\x76\x9e\x23\x11\x8f\x98\x45\xd4\xb0\xfd\xb5\x77\x42\x2a\xb6\x42\xd7\xa6\x79\x53\xe0\xfc\x14\x9f\x43\xd9\x71\xf9\xc9\xac\x52\x0f\x9b\xda\xf4\x30\x78\x4e\xc6\x8a\xc1\xb1\x5d\x02\x6a\x6b\xdc\x6c\x2a\xb0\xd9\xec\x8b\x07\xcd\x5a\xc3\xfa\xd1\xb9\xef\x3a\xcc\x42\xe2\x50\x97\xf6\x4e\xfe\x3b\xd8\xdc\x13\x5a\xdc\x16\x25\x9a\xa3\xbc\x35\xe5\xed\xb3\xa0\xcb\x7c\x46\xe6\x63\x82\x67\xc5\x7c\x45\x9b\xac\xc5\xfd\x8e\xf9\x60\xe6\xf6\x8c\xcc\x8f\x65\x3b\x9c\x20\x44\x37\xc8\x85\xd3\xb7\xef\x2b\xb7\x29\xda\x01\xd0\xa3\xf7\xa8\x2a\x00\x79\xd5\xfb\xa6\x30\x14\x31\x27\xf5\xf8\x79\x9d\x32\x08\xb9\x19\x3d\x63\x19\xa7\x08\x2e\xf7\x2a\x8a\x66\xc5\x57\x31\x54\x57\xa9\x42\x06\xc9\x4f\xe9\x06\x78\xb6\xb5\x8c\x1f\x58\xda\xd3\xe6\x2a\xb0\xf3\xec\x9c\x41\xd3\x95\xae\x62\x3e\x27\x53\x17\x8a\xc9\xec\x34\x43\x28\xd2\x37\x2c\x8d\x17\x7d\x9b\x99\x95\x66\x7b\xd5\x99\x74\x39\x59\x24\x10\x0c\x86\xe2\x56\xa0\x1d\x88\xd1\x65\x48\xdf\x3e\x0f\xe5\x3d\x74\x5b\x91\x38\xd6\x75\xa5\x67\xbb\xf2\x66\xdb\xa8\xdd\xa8\x53\xfd\xcc\x89\x85\xdb\xb6\x65\x6f\xa3\xd9\xbc\xd1\xed\xc7\x6f\x11\x2c\xf9\xe2\x6e\xa2\x9a\x90\xcd\x74\xa0\xd3\xfe\xec\xef\xe7\xba\xe7\xba\x6b\xac\xee\xca\xb6\xad\xa5\x4b\x31\x2b\x28\xca\xc7\xc2\x67\x0b\x46\x41\x91\x44\xec\x56\x51\xd0\x56\x6a\x72\x46\x60\xde\xf4\x2c\x6d\x57\xda\x6b\x6c\xf3\x76\x49\x05\x73\x84\x38\x8c\xf4\x88\x1d\xdd\x87\x84\x46\x97\xe3\x89\x3a\x70\x0f\x76\x2d\xbb\x1f\x0c\x92\x0c\x72\xdb\xc4\x4e\xcb\x9f\xed\x55\xbe\xa7\xfe\xe3\x98\xfc\x8e\xd2\xb1\xb9\x03\x3a\x32\x64\xde\x96\xf2\x13\x8c\x5d\x82\xf6\xb3\x85\xf1\x1e\x2a\x73\xba\xbe\xe4\x72\xf0\x2a\x8c\x9d\x92\xf7\x7b\xda\x04\xa3\x2b\xa3\xe5\x2c\x7d\xa5\xe9\x75\x1a\xd1\x1a\xe0\x7a\x96\x89\x34\xeb\x02\xbe\xa2\x05\xe6\x33\x90\xd4\xb8\xff\x3f\x4b\x6c\x9c\x6e\x22\x37\xd2\x15\xa7\x2e\xa5\x04\xe6\x08\x7a\x3e\x63\x61\x63\x66\x45\xe6\xf5\xb2\x46\x6f\xc4\xb8\x4b\xed\x45\x2b\x23\x15\xaf\xd5\xfa\x41\x22\x09\x57\x48\xc2\xe2\x08\xb6\x16\xf6\x32\xef\x21\xad\x84\x51\x1e\x3e\x7d\xdb\xf8\x83\xb8\x66\x97\x49\x62\x15\x84\xde\xbe\xcd\x28\x53\x1b\x95\x6d\x98\xe7\xf0\x4c\xf7\x5a\x9a\x1d\xa5\xf1\x93\xdf\x19\xdb\x56\x9e\x8c\x62\x8e\x6b\x71\x84\x40\x2c\xb3\x26\xc6\x81\xe4\xf8\xdd\x54\xc5\xca\x57\x76\xd3\xdd\x37\xd1\x82\xfa\xc7\xfb\x38\x39\x11\xec\x56\xdd\x43\xaf\xda\x0d\x0a\xfc\xf6\xe1\xc7\x10\x5d\x9f\x1c\xdf\x80\x70\xc7\xa5\x6f\x2e\x0b\x59\xc7\x41\xa4\x14\x61\x68\xdc\xc8\x3d\xe8\x00\xa7\x2b\x24\x55\xb9\x3e\x11\x86\x20\xc1\x81\x6c\xc7\x3d\xb1\x34\x23\xaf\x9e\xe4\x14\x71\x2b\x88\xdf\x60\xfb\x05\xb4\x3f\x76\xb1\xe3\x91\x62\x75\x5d\xd7\x97\xc2\x35\x33\x67\x6a\x5b\x96\xc3\xee\x6b\x22\xc1\xec\x1a\xb2\x27\x7c\xb0\x45\x1d\x6a\xaf\x26\xd5\x13\xb8\x7d\xa9\xe5\x14\xcf\x75\x3a\xc7\x89\x17\xb6\xe8\x2c\x77\x3d\x3d\x9d\x22\x9c\xac\xae\xcb\x22\x0b\xb4\xa1\x8f\x8b\x9c\x9e\x0a\x6e\x27\x07\x23\xf9\x7f\xfb\x07\x9e\xc7\x16\x0d\xfd\xdb\xd1\x46\x97\xa0\xbe\x10\xe0\x87\x71\xb1\x10\x3e\x39\xad\xcc\xce\xe3\xbe\x3c\x41\xf2\x86\x92\xa5\xe1\x82\x5a\x06\xc6\x03\xbe\x24\x31\x50\x3b\x4e\xec\x73\xf4\xfa\x5b\xe0\xcd\x68\xe9\x63\x95\x9d\xe6\x2e\x5b\xbc\x06\x92\x61\x74\x2b\x84\x7a\x21\x94\xfa\x96\x90\xf1\x22\x6b\x2d\x42\x81\xe7\x69\x6a\x1a\x88\x56\x9d\x1f\x76\x62\x59\xe1\xea\x16\xfb\xb4\x5d\x77\x6c\xff\x28\x15\x44\xbf\x80\xd1\x34\x5b\xa0\x25\x02\x49\xd1\xde\x9a\xb6\x4b\xc5\xf2\xbd\x6a\x0f\x0d\x35\x86\x1a\xb7\xab\x6c\x03\x5d\xdf\x6c\x72\x6e\x3a\xb4\x3d\x3c\xf6\x05\x28\x57\x37\x3d\x40\x3b\xae\xb0\xb6\x6a\x70\x33\xb4\x94\x47\x4f\x0e\x73\x4d\x71\x6d\xf2\x4a\xd3\xd1\x25\x9f\x86\xb0\xf9\xeb\x0c\xae\xcd\x5f\x91\xad\xee\x42\x75\x30\x92\x7d\x82\xc7\x14\x16\xb8\xc0\x73\xdd\xe0\x28\xc9\xd0\xba\x94\xa4\xf2\x20\x1a\x5a\xdd\xb2\x42\x5f\xea\x31\xfa\xb1\x8a\xd3\x86\x21\xec\x66\x0f\x16\x48\x4e\xf3\x12\x39\xa8\x8c\x47\x3e\x1a\x4a\x18\xfb\x2b\xc1\xa8\x26\xa4\x7d\xa5\x72\x02\xe3\x05\xca\x6e\xdc\x4c\x84\x4e\x17\x5c\x2e\x28\x62\x0b\x52\xca\xcc\xdf\x4b\x5b\xcd\x24\x6b\x6f\x61\x63\x8d\xd4\x90\xfa\xa9\x6b\x66\x92\x4b\x48\xe7\xe1\x16\x6d\x2f\x1f\x6c\xa0\x33\x4c\x5c\xea\xeb\xed\x43\xb2\xc5\x11\xcf\x49\x4f\x47\x28\x8f\x27\x8d\x4d\xba\x20\x5f\x38\xe6\xd1\x6f\x98\x70\xa4\xa4\x46\x1a\x72\xb2\x80\x3f\xe0\x45\x90\xe7\x6d\xe0\x6b\x48\xae\xbe\xc7\xf2\x94\xa7\x9f\xe5\x22\x28\xa6\x7b\xf7\xb6\x80\x36\x60\xe6\x31\x64\x78\x60\xce\xed\x1a\x39\x7e\xfb\x73\xd2\x46\xed\x6c\x64\x19\xfa\x7a\xd1\xc8\x03\x43\xb9\x61\x7b\xb5\x67\x7a\x16\xac\x0e\x44\x4f\x61\xf3\x10\xd9\xfa\x78\x0d\x5d\x1d\xb2\x38\xe7\x02\x84\x39\xd7\xe2\x51\x13\x3f\xb4\x5e\x72\xcf\xab\x74\xc1\x8e\x83\xe9\xf4\xcc\xe0\x68\x7d\x8c\x7f\x3b\x89\x79\xba\x12\x3d\x1c\xba\x40\x1f\x4b\x46\x7f\x26\xfd\x69\x3d\xa4\x48\xe7\xfd\x96\xdb\xdc\xdf\xd6\x5f\xef\xba\xf6\x76\x20\xff\x69\x37\xf4\xab\x23\xcd\xc2\x13\xbc\xec\x20\x07\xd5\x1e\x99\x05\x6e\xbc\x0a\xb5\x93\x70\x4e\x8b\xeb\x15\x57\x0b\x8e\x54\x6e\x6a\x62\xfa\xc8\x00\x56\x8c\x2d\x8e\x3e\xbf\xb0\xb2\x09\x34\xe1\x5b\xbb\x8c\x19\x5d\xc4\xff\x88\x7d\xe6\x5d\xe0\x18\xba\x22\xdd\xea\x56\xea\xfd\xb4\xec\xec\xf5\x98\x90\x9b\x02\x4d\x79\x91\xdd\x14\x18\x31\xd6\xf8\x31\x62\xed\xb6\x0e\xc0\x99\xcc\xca\xde\x25\x16\xf3\x22\xc9\xea\x47\xa4\x07\x1e\x95\x15\xb8\x67\x32\x20\x16\x62\xea\x6f\xfc\x34\xeb\x00\xf6\x36\x0c\x7d\x24\xc8\x22\xa4\x29\xc7\xf5\xc6\x07\x9b\xf0\x38\x07\xa8\xa5\xb9\x51\x14\x23\x48\xea\x4b\x5c\x04\xae\x61\x18\x3d\xc6\xb2\x2d\x6c\x4c\x09\xfe\x33\xb9\x66\xfe\xed\x01\xe1\x47\x62\xe7\xd2\x56\xdf\x95\xad\x68\xaa\x60\xcb\xeb\x5a\x5b\xdc\xf9\xe9\xb8\xaa\xe5\xf5\x85\xf6\xdd\xfb\x79\x9a\x1b\xaa\xf7\xb8\x9f\x15\x29\xbb\x9a\xf6\x3f\x7e\x2f\x2b\x7a\x36\xd8\x2e\xab\xbd\x79\xb5\x7c\xdd\x02\x5f\xef\xfd\xab\x2d\x6f\x5f\x75\x5e\x9c\x8b\xd4\xf4\xef\x77\x1b\x2b\x41\x19\x4b\x2f\x56\xf8\x12\xb2\x9b\x30\xa8\x7d\xab\x35\x08\x62\x86\xfc\x91\x43\xe6\x88\xe2\xa6\x44\x12\x06\x01\x8a\x96\xcc\x2c\xd9\xf6\xc4\x35\xd6\x60\x48\x71\x0a\xbf\xb0\x54\x20\x89\x9c\x5e\xc0\xb7\xe4\x3d\x6d\x09\x12\xf3\x3d\xd0\x1d\x65\x19\x59\x61\x7e\x9a\xf7\x60\xd4\xab\xdc\xef\xc0\xdc\x74\x7f\x8e\xcf\x3e\x4c\x2f\x4f\x2e\x92\x48\xeb\x0c\xa8\x43\xa5\xe0\xbb\xd0\xd3\xed\x7a\x1f\x1e\xaf\x5b\x91\x9b\x7e\x25\x99\xb3\x74\x4c\x11\xe4\xa8\x69\xed\x88\x78\x1b\x36\xe8\x54\xb6\x77\x74\xc2\x4e\x56\xfc\x8c\xcc\x4f\x6e\x91\xf0\x22\x02\x70\xbd\x77\x04\x23\x7d\x86\xb5\x72\xc9\x49\x3a\x6e\x3d\x0f\x55\x3b\x56\x97\x56\x80\x44\x60\x51\x7d\x2e\xe9\x3e\xfc\xc2\xea\x3b\xcd\x3f\xf8\xf7\x98\xd5\xbf\xdf\x50\x3a\xbf\x1d\xcb\x03\xa5\xab\x56\x5b\x8c\x4a\x39\x48\xd2\x38\xe3\xdc\xe0\x22\x7a\x5a\x18\x9e\x80\xef\x07\xe8\xc3\xba\xb9\x59\x1f\x3b\xb0\xa3\x57\xf0\xdd\xec\x66\x73\xf4\xf7\x67\x31\xfd\xcb\xeb\x0b\xfd\xc0\x38\xb9\x3a\xae\xa6\xd7\x33\x05\x9b\x0b\xfa\x6e\xa4\x1b\x09\x97\x3f\x1c\x58\x99\x34\xef\xab\x01\xc9\x5f\x8b\xea\x4d\x51\x06\xe4\x99\x7c\xc6\x7e\x42\x6a\xb0\x62\x08\x30\x4e\x8b\x8c\x0f\xfe\xe4\x9e\x9e\xb7\x90\x02\xf8\x85\x81\x57\x80\xa2\x5f\x56\x05\x45\x3b\x03\xf8\x85\xed\xb1\xfc\x66\xb0\x1b\x04\x46\x99\x00\xc6\xe8\x8b\x18\x36\x3a\x19\x4f\x77\xd6\x4b\xf8\xf5\x02\x71\x5a\x20\x96\x1e\x1e\x6c\xc2\xc3\xe4\xe7\xe0\xda\x71\xe3\x92\xac\xf2\x9f\x21\xcf\x16\x67\x64\xce\x76\xc2\x63\xb4\xe1\x06\xaf\xc0\x20\x60\x9f\xbd\xb5\x44\xcc\x89\x9e\x5d\x6a\xb3\x40\x65\x99\x0c\xb3\xdf\x0d\x24\x83\x3f\x05\xbb\x92\x3b\x10\xeb\xaf\x28\x78\x78\x8d\xab\x45\x51\xb4\x12\x01\xb7\x7a\x45\x04\x8b\xd6\xfe\xc7\x7c\xdc\x22\x9c\x7f\x37\xae\x87\x54\xf1\x6a\x60\x74\xee\x0f\x52\x87\xde\x36\xc3\xd8\xd1\xc6\x22\x96\x32\x0c\x73\x28\x98\x71\x52\xd3\x0e\xd2\xc1\xc0\x95\xae\xd7\xa5\x85\xbe\x56\x22\x60\xad\x37\x1c\x78\x05\x66\x7a\x63\xef\x20\x61\xed\x86\x20\x23\x98\xa3\xaf\xe2\x87\xeb\x5d\x8f\x47\x72\x26\xa9\x32\xea\xe2\x0c\x78\x05\xe4\xb0\x91\xfe\x7d\x44\x51\x55\xc2\x0c\xed\xec\xff\xdb\xff\xdb\xf9\xfc\x39\x7f\xb1\xfb\xa7\xfd\xf9\xb0\x9d\x63\x29\x34\x71\x08\x72\x94\x45\x70\x03\x40\x11\x5f\x51\x0c\x54\xcb\xc4\x68\x46\xc9\x72\xbc\x80\x54\xec\xce\x1d\x31\xcc\x53\x60\xf1\x9f\xc0\x5e\xa8\x09\x55\x2d\x23\x01\x71\x27\xf5\x0f\x8c\x43\xca\x51\xfe\xfa\x2e\x05\x03\x11\xfd\x0c\x86\x31\x48\x5b\x87\x52\x57\xa7\x3e\x29\x56\xe8\xe6\x98\xab\x28\x1a\xbd\xdd\xd2\xfa\x87\x38\xa0\x38\x60\x53\x70\x18\x05\x20\xb7\x88\xd2\x22\x47\x2c\x8d\x2f\x4f\x21\xd2\x4d\x64\xef\xdb\x01\x9f\xba\x06\x00\xa9\xe2\x18\x2e\x51\x6a\x2d\x6a\x58\x0b\x3e\xfd\x04\x06\x6c\x31\x18\x82\xc1\x5e\x36\x68\x9e\x0a\x85\xed\x42\x7b\x15\x7b\x19\x1c\xb5\x89\x0a\x95\xdd\xa0\x2f\xe0\x15\x38\x87\x7c\x31\x9a\x95\x84\xd0\x1d\xf9\x23\x95\x1f\xee\xda\xd9\xfd\xe1\xf0\xe0\xe0\xe0\x20\xac\x13\x0c\x71\x7d\x00\xec\x34\x5a\x19\xd5\x45\x20\xac\xf0\x88\xaa\x48\x62\x47\xa9\x52\xab\xcd\x60\x07\x51\x3a\x04\x14\xb1\x0e\x04\xe2\x5f\x31\x93\xa0\xbb\xb5\x6e\x67\xd7\xf2\xd7\x20\x7d\xf5\x3f\xb1\xcc\x25\x62\x0c\xce\x51\xb3\xc9\xb4\x10\xc0\x0b\x30\x48\x07\xe0\x45\x63\x14\x5f\x80\xc1\xbe\xd0\x5c\xc9\x97\x57\xe2\x8d\x64\xd0\x0b\x30\x58\xb2\x5a\x34\xf2\xb1\xb5\x57\x3b\xa7\x17\x24\x53\xc4\x46\x33\x58\x94\x2b\xf1\x43\x29\xbf\xaa\xb7\xdb\xa7\x34\x35\xc9\x2f\x5e\x81\x81\xe0\xcc\xaa\xe4\xaf\x34\x0e\x40\x11\x64\x04\xbf\x52\x94\xb7\xa8\x3f\x1d\x5c\x8d\xd4\xab\x4e\x8a\x50\xc9\xdc\x74\xe0\x56\xb3\xb3\x95\xfe\xa8\x2c\x64\x37\xcd\xdc\xf2\x4b\x2a\x62\x62\xf1\xc3\x11\xed\x9e\xb9\x24\xf3\x1d\x8d\xd9\x30\x65\x52\xa0\x9d\x62\x07\x6a\xe7\x31\x52\xa2\x91\xc0\x31\x40\x94\xbe\x1c\x0c\x41\xaf\xec\xe5\xc0\xeb\xc0\x61\x6d\xfe\x12\xb1\x79\x1d\xaf\x36\x43\xa9\x17\xfe\x4b\x7f\xa3\x35\x2a\x6e\xad\xbd\xf3\x44\x50\x7d\xea\xe2\x80\x16\xaa\xb8\x27\x38\xbd\x23\xbc\x90\x63\xc8\xd1\xce\xee\xee\x68\xae\x36\x5e\x64\x55\xdb\x99\xea\xda\xbd\x10\x67\x67\xda\xfc\x16\x35\x8f\x65\xed\xeb\x2b\x78\x16\xf2\xfb\x15\x67\x42\x0f\x85\x27\x35\xca\xec\xa0\xc1\xb3\x01\x3d\x6a\xe0\x8a\xbf\x4f\xfa\xdb\xb1\x41\x53\xa7\xc2\x93\xad\x2c\xb9\x16\x61\x0a\x1a\x59\x0a\x17\x99\x71\xb8\xac\xd2\x88\x98\x7a\x2c\x79\x94\xe9\xe0\xfe\x72\x02\xf7\x91\x15\x88\xc9\x0b\x38\xfc\xd6\x4b\x8d\x33\x5b\x0a\xb8\x32\x22\xbd\x46\xba\xd9\x75\xcc\xcf\xf0\x1d\x2c\x33\x3d\x67\x05\x6c\x81\x0e\x15\xa7\x83\x4c\xbb\x9b\x3d\x39\x55\xe1\x4e\x4e\x17\x84\x72\xed\x31\x5e\xac\x3a\xf2\xab\x5a\x27\x52\x09\xd4\x1b\x87\x19\xfd\xea\xa3\x33\x82\xe7\x7a\x86\x3d\x96\x2d\x50\xbe\xb2\xbf\xab\x35\xd5\xcf\x4e\xbe\x56\x14\xb1\x3a\xcf\x27\x89\xd3\x6f\x9c\x06\x40\x55\xcb\xf7\x8a\x34\x32\x64\x8b\xc6\x75\x6d\x98\x19\xf9\x78\xc1\x69\x1e\x20\x58\xb7\x0d\x38\x9d\x07\x95\xae\xaa\x7f\xae\xbf\xb5\xf1\x39\x49\xc1\xe7\xba\x29\x4a\x1e\x0a\x9b\xcd\xe7\x64\x08\x3e\x27\xfa\x48\x6c\x01\xf4\xc5\x70\x09\x60\xc8\x38\x94\x51\x0f\x88\x48\x05\xcd\x13\x44\x97\x05\x63\xa1\xe8\x1a\xb8\xe1\xb5\x01\x1b\x92\x1a\xb0\x33\xe3\x59\xd3\x1f\xac\xd2\x2c\xe9\x29\xbe\x25\x37\x28\xf4\x99\xbc\xfa\x99\x6a\x9d\x7b\x20\xdf\x8d\xa4\xb7\x98\x54\xba\x11\xee\x87\xc8\x4c\x55\x91\x59\x10\x89\x26\xda\xae\xeb\x69\xb4\x31\x71\xd7\xc6\x09\x57\x2d\xbc\x2f\x42\xb9\x35\x8b\x23\x01\x10\xda\x5d\xf2\xc5\xf7\x92\xc5\x3f\x79\xc9\x42\x4a\xf1\x9f\xbf\x66\x11\x29\x45\x7c\xaf\x56\x7c\xaf\x56\xd4\x24\x7d\xaf\x56\x7c\xaf\x56\xfc\x83\x58\xfe\xfb\xaa\x56\x48\x13\xff\x74\xe5\x8a\xf6\xdc\xff\xe6\xf5\x8a\x76\xaa\xdf\x57\xc1\x62\xbd\x06\xcf\xe5\x9f\x08\xd1\x77\x51\x46\x16\x8b\x05\xfa\x50\xb3\x98\xba\xa4\x2a\xc7\x25\x0b\x04\x73\xef\x0b\x36\x76\x60\x13\x73\xbc\x0c\xf8\x90\xf2\x25\x20\x9a\x38\xd7\xee\xcb\x53\x26\xd0\xdd\x75\x5d\x93\xfc\xce\x59\xd7\x23\x62\xbb\xa8\xef\xf9\x1b\x44\x76\xd9\x42\xd0\x65\x6a\x97\xc4\x3b\x81\x9c\x23\xea\x3a\x0f\x49\xe4\x2b\xb3\x89\xf0\x13\x51\x66\xd9\x1a\xcb\xc6\x24\x39\xe2\xb0\x28\xf7\xb8\x5a\x80\x33\xfa\x64\x3c\x05\xed\xb5\xe6\xba\x91\x0e\x48\x17\x0b\x8c\x5d\x02\x9b\x21\x42\xca\x36\x50\xcf\xf4\xfe\x35\x77\x7d\xe6\xab\x08\xd4\xbb\xd8\xbd\x95\xb3\xd3\xe7\xe4\x6c\xef\xdc\x04\x9d\x9a\x7b\x39\x33\x9d\x4e\x4c\xe8\xe3\x00\xae\x31\xf6\x4c\xf6\x55\xd0\x20\x6f\x1e\x13\xd5\xdb\x76\xfb\x21\x61\xfd\xbf\x76\xf4\xbd\x05\x7f\x7e\xf7\xe1\x77\xe8\x53\xc8\xea\x7b\xa9\x17\xfa\x4d\xe0\x8f\x5d\x05\xbe\xff\x79\x61\x80\x69\xe1\x06\xbf\xfb\x19\x36\x7b\x8f\xff\x3c\xa6\xf1\x37\xb9\xbc\x6f\x29\x78\xb7\x10\x9f\xd5\x4c\x8a\x7e\xe2\xd5\xff\x7a\xb2\xcd\x14\xd5\x10\xaf\xd7\x19\xfe\xfb\x30\xa6\x01\x78\xa6\xa5\xd1\xc9\xd4\x8e\x5b\xd2\x81\x61\x43\x6f\xc9\x5a\x01\xa2\x6b\x72\x3e\x60\x6d\x7d\xec\x21\xfc\x4d\x5c\x37\x7f\x12\x11\xde\x56\xb9\x93\x68\x1c\xee\x18\xfa\xa8\x8b\xed\x66\x2e\x6c\xeb\xe7\x9e\x3a\x9d\xde\xbc\x9d\x56\xf1\x12\x06\x6d\x92\xc5\xb7\xdb\x28\x73\xb7\xef\x43\xdc\x64\x63\xcb\x87\xb3\x18\xf5\x46\x6e\x6d\x5a\x2c\xed\x12\x4e\xba\x58\x06\xde\xfd\x6a\xbf\xf1\x8d\x0f\x1f\x2c\xfe\xa7\x15\x9e\xfe\xaf\x26\x74\x86\x53\x09\x52\xb7\x3c\x4b\x02\xf3\xeb\xe6\x96\xa7\xba\x80\x7c\x8d\x22\xc9\x94\xc8\x18\xb5\x9d\x11\xad\x3d\x16\xf6\x86\x92\x65\xf0\xbe\x68\x3f\xb6\x0b\x17\xd7\xcf\x05\x5f\x6c\x81\x2b\x7b\xd9\x4b\x7c\xf6\x32\x3d\x5a\xf1\x05\xa1\xc5\xaf\x28\x78\x17\xda\x1b\x15\xba\xdf\xd0\xf3\x27\x06\x92\x1f\x02\x68\x9c\x27\x5b\xb9\x16\xcf\xcc\xb7\xdd\x96\xc7\xf8\x34\xbe\xff\xc5\x79\xdb\xe6\x4c\x7f\x4c\x53\xfd\xe7\x5e\xb4\xd1\x39\x46\x25\x12\x7a\xd2\x5c\x5a\x48\x2e\x84\x93\x88\x7b\x8c\x92\xac\x1d\x0b\x67\x95\xaa\x0b\x59\xee\xed\xd8\xe4\x12\x3a\x1f\xcf\x5b\xd7\x5f\x29\x4e\x98\xfc\x92\x64\x73\xef\xc3\xf8\x76\x97\x75\x65\xb4\x81\x87\x55\x65\x02\x77\x1c\x3d\x21\xb6\x19\x5c\xfb\xbf\x00\x00\x00\xff\xff\xd1\x81\x56\x58\x6f\x7a\x00\x00")

func templatesAppTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAppTmpl,
		"templates/app.tmpl",
	)
}

func templatesAppTmpl() (*asset, error) {
	bytes, err := templatesAppTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/app.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/app.tmpl": templatesAppTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"app.tmpl": &bintree{templatesAppTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

