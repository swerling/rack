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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7d\x7b\x73\xdb\x38\x92\xf8\xff\xf9\x14\x28\xfc\xf2\x3b\xd9\x33\xb2\x6c\x67\x6e\xeb\x6e\x39\x97\xab\x72\x64\x67\xc6\xbb\x76\xa2\x93\x9c\x4c\xdd\x26\xae\x2d\x9a\x84\x25\xae\x29\x80\x03\x40\x8e\x3d\x2a\x7d\xf7\x2b\x3c\x48\xe2\x49\xca\x8f\xdc\xcc\x5e\x25\x5b\x9b\xb1\xc9\x46\xa3\x81\x7e\x37\x1a\xcc\x7a\x0d\x72\x74\x5d\x60\x04\x60\x5a\x55\x10\x6c\x36\x2f\x00\x58\xaf\xc1\xcb\xb4\xaa\x40\xf2\x1a\x8c\x8e\xaa\xaa\x7d\xb8\x4c\x71\x71\x8d\x18\x97\x6f\xce\xeb\x5f\xd4\xeb\x17\x00\x00\x00\x8f\x7e\x99\x5d\xa0\x65\x55\xa6\x1c\xbd\x25\x74\x99\xf2\x8f\x88\xb2\x82\x60\x08\x12\x00\x5f\x1d\x1c\x1e\xec\x1d\xfc\x79\xef\xe0\xcf\x70\xa8\xc0\xc7\x04\xe7\x05\x2f\x08\x66\x30\xd1\x28\xe4\x4c\x5c\xe3\x00\xf0\x2a\x2d\x53\x9c\x21\xba\x97\xb5\xa0\xee\xdc\xde\xa0\x8a\x92\x0c\x31\xd6\x37\x06\x9e\x62\x8e\x28\x4e\x4b\x31\x39\x80\x6f\x71\x92\x9c\xfc\xba\x4a\x4b\x41\xcc\x27\xf1\x64\x8a\xae\x61\x62\x80\x81\xcd\x10\xc0\xff\x46\x0c\x82\x4b\xb0\x19\xd6\x58\xde\xa5\xbc\xb8\x45\x67\x64\x3e\x2f\xf0\xbc\x07\x95\x0d\x1b\xc6\x37\xa1\xc5\x6d\xca\x51\x0f\xa6\x1a\x2a\x8c\xe3\x4d\x99\xe2\x9b\x19\xca\x56\xb4\xe0\xf7\x3f\x51\xb2\xaa\x04\x07\xd6\x26\x3a\x90\x80\x4f\x6b\x89\x4d\xf0\xc6\x86\x15\x38\xe1\xa5\xda\x27\x8d\x14\x4e\x52\x9a\x2e\x11\x47\x54\x0e\xed\x66\x56\x25\x60\x1f\xc0\xa8\x20\x7c\xbd\x96\x71\xb9\x62\x1c\x51\x43\x42\x00\x80\x17\xf7\x15\x52\x84\x73\x2a\xb6\x72\xd8\xbe\x3a\x46\xd7\xe9\xaa\xe4\xf2\xad\xfd\x9c\x65\xb4\xa8\x78\x2d\x8e\x50\xbf\x6a\x77\xed\x18\x55\x25\xb9\x5f\x22\xcc\xcf\xd3\xbb\x62\xb9\x5a\x06\xe6\x14\x4c\x5c\x2d\xaf\x10\x0d\x4d\x29\x85\xfc\x20\x36\x69\x02\xa0\xc6\x0b\x2a\x44\x33\x84\x79\x3a\x47\x80\x5c\x03\xbd\x0d\x88\x01\x4e\xc0\x0d\x42\x15\xa0\x2b\x8c\x0b\x3c\x07\x5f\x16\x45\x89\x40\x2e\xe9\x12\xcb\xec\x22\xb9\xc0\x8f\x24\xf9\x4f\x9d\x14\x2b\xb4\xcf\x47\xf1\x09\xbe\x2d\x28\xc1\x82\xe4\x30\xad\x71\x8e\x76\x30\x34\xc8\x4f\x53\xbf\xb7\x9b\xc7\x42\xf8\x1e\x97\xf7\x20\x2d\x4b\xf2\x05\xa4\x99\x58\xae\x58\x2c\x5f\x14\x0c\x08\xeb\x78\x4d\xc9\x12\x14\x98\x15\x39\x02\x7c\x81\xc0\xc7\xc9\x38\x42\xf3\x3b\x62\xbe\x38\x12\x08\x51\xfe\x31\x2d\x57\x48\x29\xb5\x54\xdf\xa1\x84\x03\x97\xde\x22\xfe\x8a\xee\xbf\xf6\x3e\x79\x16\xec\x11\x9b\xf5\x81\x21\x80\x25\x1e\x70\x32\x9e\x81\xf4\x0b\x2b\xc9\x9c\x81\x9c\x16\xb7\x88\x02\x22\xfe\xca\x08\xbe\x25\x77\xfb\xe9\x1c\x61\x0e\xae\x09\x05\xa5\x9e\xf1\x6b\x6c\x9b\x61\x45\x1f\xb9\x9a\xd9\xea\x0a\x23\xce\x34\x22\xc1\x7b\x56\xa1\xac\xb8\xbe\x17\xac\xde\x93\x7c\x2f\x49\x9a\x83\xda\xea\x01\x84\xf3\x8a\x14\x98\xb3\xaf\xb2\xa0\x29\x2a\x51\xca\x42\x0b\x7a\x6e\x33\x38\x45\x15\x61\x05\x27\x34\x24\x78\x4f\x9b\x6c\x46\x56\x34\x43\x20\x23\x39\x02\xb4\x9d\xc6\x23\xc1\x76\x47\xcf\x4d\xc5\xc5\x02\x81\x33\x8b\x75\x4c\xcf\x07\xe6\x62\x42\x29\x9c\xb5\xa2\x07\x88\x53\x82\x11\x21\xeb\xac\x60\xfc\x3f\x8e\x7e\x99\x25\xc9\xc9\xf8\x55\x92\x28\xe0\x24\x39\xcd\xff\xf3\x31\xa4\x7e\x9c\x8c\x01\x53\xf3\x6d\x47\x55\x5c\xee\xbf\x0e\x71\x95\x56\x8f\xed\x88\xbc\x48\xd9\xcd\x94\x94\xcf\x2f\xc5\xa7\x47\xe7\x40\x20\x16\x6a\x9a\x56\x55\x79\x2f\x7e\x10\x86\x48\xcc\xc8\x84\xcb\x8a\x13\x55\xc7\xa8\x16\x4d\x8e\x41\xd8\x99\x9e\xfc\xd7\x87\xd3\xe9\xc9\xf1\x2e\x38\x4b\x97\x57\x79\x0a\xc6\x2b\xc6\xc9\xf2\x82\x54\x45\x06\x7e\x4e\x71\x5e\x22\x0a\xb4\x8e\x82\x1a\xa3\x41\xf0\x79\x81\xcf\x10\x9e\xf3\x85\x24\xf7\xd0\x7c\xe5\x58\x25\x9f\xbe\xc9\x38\xb2\x5f\x2d\x27\x3f\x4e\xc6\x82\x8d\x8f\xe5\x62\x37\xd7\x3e\x4e\xc6\xe3\xd3\xe3\xe9\xb3\x33\x4d\xcc\x2c\x10\x87\xa7\xb7\xa2\xcf\xf3\xb4\xaa\x0a\x3c\x37\x95\x0e\x4e\x08\xe5\x13\x4a\x38\xc9\x88\xe3\xe2\x17\x9c\x57\x2a\x7e\x16\x02\x8f\x30\xa2\x06\x1c\xfc\xf9\xe2\x62\x22\xec\xec\x29\x66\x5c\xa8\x7f\xe8\x9d\x34\x40\x28\x06\x31\x83\xed\xee\xe8\xe9\x58\xf7\x7c\xb3\x27\x4f\x68\xcd\xc8\xb3\x8e\xf5\x5d\x8c\xa3\xcb\xd3\xaf\xe2\x93\xcd\x66\x67\xee\x54\x65\xc7\xd2\x04\xf8\xd3\xa6\x02\x9b\x20\xbf\xa7\x88\x49\x57\x61\x31\xdc\x50\xb9\x88\x11\xa9\x75\xe2\xf4\xe8\x3c\x49\x24\x8c\xb1\x92\x09\x25\x15\xa2\xbc\x40\xb6\xe9\x16\xbe\x98\xb1\xd5\x12\x09\xf8\x09\x29\x8b\xec\xfe\x98\x64\x2b\x2f\x40\x75\x6c\x85\x48\x67\x5f\xed\x1d\x1e\xec\x1d\xfe\x9b\x31\x89\x04\x9a\xf1\x94\x23\x3d\xfe\x93\xf5\x0a\x38\xf8\x24\xf8\xc9\xf5\x35\xca\x64\x84\x20\x63\x02\x07\x9b\x26\xbd\xc0\x59\x51\xd5\xa9\xea\x0c\xd1\xdb\x22\x43\x2a\x6a\x28\xa5\x3d\x1a\xa5\xcb\xf4\x37\x82\xd3\x2f\x6c\x94\x91\xa5\x95\x0d\x9a\x0b\xcd\xb4\x41\xfb\x04\x20\xe3\x2c\x69\x17\xde\x86\x1c\xf5\x9f\x8d\xf5\xbb\xf9\xd6\xc2\x0c\x27\x29\x5f\x08\xe2\xf7\x75\x7c\x07\xed\xb7\x62\x43\xd5\x96\xdb\x5b\xe1\x6e\x84\x82\xbc\x7f\x97\x2e\x15\x1b\xf3\x65\x81\x0b\xc6\x69\xca\x09\xf5\xb6\x04\xf6\xf0\x09\x6c\xcb\x2b\xe0\xf1\x4b\xec\xaf\xc7\x11\x63\xe7\xe0\x77\xe2\xd7\x5a\x3e\xd5\x03\xb0\xe9\xd9\x3d\xf3\xb7\x16\x72\xe3\x59\x5a\x43\xc2\x3b\xa4\x5b\x79\xa0\x24\x79\xbb\xc2\x8a\xaa\xad\x84\x7c\x4c\x72\xe4\x0b\xf4\xec\x87\x37\xab\xec\x06\xf1\xb6\xdc\xf0\x17\x52\x68\x09\xd9\x83\x43\xf1\x1f\xc5\x57\x38\x34\xaa\x0f\x92\x8c\x29\x9a\x4b\x4b\xbe\x01\x97\xbe\xb8\xc1\xd9\x0f\x3a\x73\x71\xb1\x2a\xa4\x54\xb9\xca\x7d\x0b\x6d\x53\x2d\xda\x0c\x01\xdc\x57\x82\xbd\x7f\x2d\x0b\x49\x05\xc1\xa3\xdf\x8a\x0a\xaa\xb9\xa2\xc2\xa8\x3d\xb1\x40\x56\xe0\x1c\xdd\x8d\xd0\x9d\xce\x01\x2d\xb0\x73\xb4\x24\xf4\x7e\x56\xfc\x26\x37\xf5\xf0\xd5\xbf\xdb\xaf\x6b\xeb\xa2\x48\xff\x09\xf1\x23\xae\x64\xc3\x33\x41\x42\x32\x28\xf6\xd4\x0d\x4e\x57\x98\x17\x4a\x92\x31\xc9\xd1\x3f\xd8\xbf\x8e\x7e\xb0\xe7\xb8\x28\x96\x88\xac\xa4\x90\xfd\x70\x70\x00\xe3\x42\x11\x2e\xb1\xd0\xc6\x40\x82\x51\xa4\xba\x92\x51\x82\xff\x41\xae\xb6\x01\x95\x39\xd9\x36\x80\x75\xc5\xc6\x04\xdd\xb2\xc8\xc3\x94\xd1\xea\x40\x4e\xd1\x5c\x68\xfc\x7d\x14\x7b\x68\x50\x1d\xba\xc3\x08\x52\xc6\x55\x99\xcc\xf6\x2f\xef\x57\xbc\x5a\xf1\xfe\xb2\x23\xd1\x70\x60\xd4\xbd\xb8\x16\xae\x67\x37\x9a\x35\x86\x47\xb4\x09\x10\xe7\x4e\xbc\x23\x2c\x9a\x48\x16\x95\x60\x6a\x8d\x69\xe0\x5c\x3f\xfa\x42\xfc\x7f\xbd\x16\x49\xa9\xc4\x6b\x54\x7a\x43\xe5\xd1\xba\xc6\x4b\x53\x3c\x47\xe0\xe5\x8d\x2c\xf1\x9e\x60\x4e\xa5\x41\x66\xf5\x62\xe0\x09\x4e\xaf\x4a\x94\xaf\xd7\x60\x55\x55\x88\x0a\xc8\xcd\xa6\x55\x95\x77\x44\xea\x49\xb0\x70\x29\x9e\xcc\x50\xa9\x0c\xeb\x27\x70\x60\x2a\xbe\x8d\xef\x6d\xad\xf1\xca\xb6\x08\x63\xb0\x77\x28\x75\x4c\xab\x59\xbb\xae\xee\x15\xd6\x75\x45\x67\x75\x48\xae\x4e\x3b\xd1\x76\x6d\x2d\x11\x68\x24\x56\x6d\x51\x62\x84\x21\xb5\x2d\x1e\x93\xe5\x32\x3d\x46\x65\xb1\x2c\x38\xca\x45\x78\x04\x8d\xba\x5c\x93\xf6\xaf\xd7\x02\xa1\x7e\x20\xab\xa8\x62\x4a\x13\xd4\xca\x34\x54\xa9\xce\x2b\xb2\xd1\x15\x1e\x82\xf1\xe4\x03\x58\xe1\x82\xab\x27\x48\x68\x14\x1a\x82\x14\xe7\xe0\xfc\x8d\x18\x31\x3d\x3a\x37\xde\xc0\x56\xe2\xb7\xdd\xb0\x46\x28\xe5\x9e\xc0\x33\x32\xb7\x33\xf0\x80\x04\x36\x30\x4a\xe6\x86\x3d\x33\x18\xaa\x1d\x9b\xc3\xf6\x75\x64\xce\xe4\xdf\x0a\x68\x9b\x29\x5a\x43\xb3\xd5\xc9\x45\xe4\xb4\xa3\xb8\x6e\x87\x8d\x7e\x4e\xd9\xa4\xe1\x86\x96\x17\x47\x9e\x5a\x60\x57\xb0\xc2\xa2\x75\x32\x9e\x89\xdc\xf4\x58\x10\x5f\xf0\x40\xfe\x59\x21\x9c\xb3\xf7\xd2\x69\x5a\x71\xc1\xb0\x89\xff\xa4\x07\xba\x0c\x64\x92\x0a\x5c\xa4\x86\xee\x1c\x06\xb0\x11\x1e\x1d\x8e\x0e\xb6\x8b\x21\xcc\xfc\xbd\x91\x80\xe6\xa1\xe3\x03\x35\x95\x17\xe4\x06\xe1\x5e\x6f\x1a\xf5\xa4\x3a\x20\x8c\x04\x27\x4e\x48\x32\xe3\x69\x76\x23\x47\x48\xab\xa1\x54\x4f\x6f\x38\xf4\xc3\x14\xb3\xa8\xd6\x20\xaa\x9f\x39\xa0\x4e\xdd\xba\x01\x37\x9f\x3b\x43\x9a\x00\x48\x83\x8a\xdf\x1d\x10\x59\x9e\xe8\x8f\x8d\xeb\xa8\xd8\x5e\x90\x17\x15\x9f\x2e\xd3\xb9\x01\x27\x7f\x0d\x01\xae\xd7\x42\xba\xd1\x48\x5a\x30\x9c\x8f\x8e\x28\x4d\xef\x37\x1b\x3f\x32\xd6\x00\x81\x3c\x06\x58\x1a\x20\x63\xad\x21\x78\x89\x4a\x19\x47\x4b\x7d\xe8\x47\x6f\x12\x23\x31\x6c\x36\xc3\xf5\x1a\x95\x0c\x6d\x36\xeb\x35\xc2\x79\x74\x0c\x5c\xaf\xeb\xb9\x36\x1b\x18\x24\x2d\x3c\xfc\xd2\xdf\x0a\x31\x9f\xd0\x76\x8c\x4c\x9a\x55\x55\x03\x40\xd8\xbd\x2d\xeb\x35\xb8\x15\x26\x31\x30\x74\xe3\x25\x60\x61\xa2\xe0\xb8\x5a\xb5\x02\x6e\x78\xc8\xc3\xb0\x87\x0c\x38\x27\xed\x26\x5d\xc4\x2a\xca\x0d\xe2\x7e\xf5\x54\xdc\xb1\x63\x9c\x06\xe0\x68\x32\xa9\x25\x51\xd8\xd5\xa8\xd0\x0a\x2d\x3c\x1a\xff\x55\xc3\x22\x7c\xab\x7f\x8f\xc0\x1e\xfd\x32\xfb\xfb\xf4\xe4\xa7\xd3\xf7\xef\xcc\x11\xc6\xd3\xc8\xb8\xb3\xf7\x3f\xfd\xfd\xa7\xe9\xfb\x0f\x93\x76\x3b\x4e\xaf\x95\x25\xb1\x8f\x3f\xfc\xb1\xa0\xde\xaa\xa6\xe4\xf6\x8e\x28\x3f\x18\xca\xb1\x41\xd4\x37\xda\x7f\x82\x09\xba\xe1\x53\x46\xc6\x0e\x83\xa0\x22\x28\x27\x85\x11\x18\x49\x6a\x00\x84\x61\x38\xe5\x7f\x1a\x06\x18\x42\xab\x07\xfa\x92\xaa\xd1\x6b\x27\xfb\xa0\x37\x75\x04\x89\xee\x87\xe0\xa5\x9a\x46\xb8\xd6\xb3\x02\xdf\x7c\x4c\x29\x0b\x93\x28\xc5\xe4\x06\xdd\x37\xf4\xe9\x91\x21\xca\xe2\xb3\xc3\xc9\xf4\xfd\xf8\x64\x36\xf3\x2d\xa5\x9b\x9c\x7b\xe2\xfc\x91\x94\xab\x65\xa0\x4a\x01\x1c\xa6\x9c\x93\x15\xe6\x22\x04\xd6\x03\xba\x38\x33\x3a\x65\xb3\x7b\xc6\xd1\xb2\x83\x2d\xa3\x9f\x09\xe3\x9b\x4d\xb2\x5e\x8f\xc6\x04\xf3\xb4\xc0\x88\x06\x05\x58\xad\x5b\x98\xaa\x08\xb2\x48\x9e\xbd\x7f\xab\x08\xdd\xf7\xf3\x77\xc7\x59\xee\x0b\x9b\x2a\x77\x4c\x58\xdf\x08\x61\xa1\x54\xbf\x8f\x2d\x1d\x6f\xa2\x7a\xe5\x80\x7a\x66\x1b\x9e\xdc\x71\x9a\x0a\x1a\xfb\x78\xe6\x08\xa2\x50\xac\x66\xe8\x79\x5a\x45\x18\x18\xe6\x97\x18\x64\xba\x62\x2d\xb1\xa1\xed\x10\xde\xb8\x3a\xca\x73\x8a\x18\xab\xc1\x6b\x99\x0e\x39\xac\x07\x09\xfa\x13\xf6\xad\x0e\x4e\xc3\xbb\xf6\x78\xbc\x13\x42\xb9\x51\xa4\xef\xe0\xc8\x48\x80\x76\x2a\x8e\x48\x68\x76\xd0\xaf\x60\x54\x97\x8b\x55\xc1\x7b\x17\xec\xbc\x44\x22\x14\x7f\xa3\xb3\xf3\xdd\x08\x12\x57\x13\x12\xa1\x0a\x31\xa5\x89\xfb\x40\x41\xa7\x30\x9c\xf5\x74\x60\xb3\x11\x02\x10\xb1\xfa\xa0\xb6\xb3\x8d\xba\x80\xcd\x66\x5f\x3c\x68\x56\x11\xe6\x7c\xa7\x46\x75\x28\x3c\x74\x68\x4b\x7a\x27\xff\x03\xa8\xed\x84\x16\xb7\x45\x89\xe6\x28\x6f\x8d\x74\xfb\x2c\x18\xd4\x9e\x91\xf9\x98\xe0\xeb\x62\xbe\xa2\x4d\x2a\xfe\x30\x07\x1e\xd2\x66\x81\xf6\x58\x36\x27\x08\x42\x74\xbb\x42\x70\x87\xe0\xfb\xca\x6d\x51\x73\x00\xf4\xe8\x3d\xaa\x0a\xa2\x49\xb4\x50\x1a\x31\x14\xf5\xf8\x79\x9d\x07\x87\x02\x88\x9e\xb1\x8c\x53\x94\x2e\xf7\x2a\x8a\xae\x8b\x3b\x31\x54\x57\x6d\x43\xa6\xc6\x7b\x16\x32\x3e\x5b\xf3\xf8\x91\xa5\x6e\x6d\x88\x02\x7a\x67\x27\xc2\x4d\x8f\xa0\x4a\xb6\x9c\xf2\x53\x28\x19\xb2\x73\x67\x83\x1c\x9d\xd5\x5b\x36\xc4\x95\x6f\x68\x96\x0b\x1a\xf5\xaa\xcb\xfd\x72\xb2\x48\x06\x16\xd2\x17\x3b\xc3\x0d\x24\xc7\x32\x6f\x6e\x9f\x87\x92\x79\x7d\xc8\x2b\x1c\xb6\x3e\xf6\xd9\xae\xdc\xdf\xb6\xcd\x35\xe2\x54\x3f\x73\x92\xd0\xb6\x89\xcc\x53\x34\x7b\x6f\x74\x33\xd8\xcf\x28\x2d\xf9\xe2\x7e\xa2\x5a\xc2\xcc\x1a\x97\xd3\x8c\xe6\xeb\x73\xdd\x01\xd7\x35\x56\xf7\xc8\xd9\xd6\xd2\xa5\x98\x15\x14\xe5\x63\x11\x8d\x05\xf3\x9b\x48\x75\x71\xab\xfc\x66\x2b\x31\x39\x23\x69\x5e\xbf\x0c\x39\xbc\x40\x46\xd4\xd8\xe6\xed\xb2\x79\x73\x84\x70\x45\x7a\xc4\x8e\xcc\x94\xc5\xc0\x8b\xf1\x44\xb9\xd2\x83\x5d\xcb\xee\x07\xd3\x1f\x83\xdc\xb6\xa2\xd2\xee\xcf\xf6\x22\xef\x19\x00\xe7\x9c\xd0\xe1\x77\xfc\x28\xc5\xd4\x80\x48\xf1\x27\xa8\x52\x7e\xd5\xac\x8b\xd1\x7e\x09\xcc\x20\xd8\xb1\x4a\xe6\x74\x7d\x15\xd3\x60\x63\xb2\x5d\x67\x6e\xb6\xd2\x2c\x19\xbe\xd4\x55\x4a\x49\x5e\xf2\x5a\xd3\x3b\x9a\x18\x4f\x0d\xe0\x7a\x96\x89\x34\xeb\x02\xbe\xa2\x05\xe6\xd7\x00\xd6\xb8\xff\x3f\x83\x36\x4e\xb7\x3a\x39\x32\x23\x33\xa3\x24\x29\x5b\x84\x03\x73\x04\xe3\x9e\xb1\xb0\x31\xd7\x45\xe6\x75\x16\x45\xfb\x93\xdd\xa5\xf6\xa2\x95\x39\x88\xd7\xf8\xf6\x28\x96\x84\xcb\xfe\x61\x76\x34\x2d\x60\x22\x49\xdd\x7a\xf3\x5a\x41\xab\xc7\x3b\x1c\x7c\xc8\x1e\x7e\x95\x26\xbe\xc7\x50\x28\xa3\xdb\xc7\x90\x26\x2c\xa6\xb2\x4a\xcd\x64\xd3\x14\xe7\x64\xc9\xc0\x4e\xc1\x49\xda\xce\xb2\xeb\xb9\xea\xce\x85\x3c\x8a\xfd\xf6\x21\x46\xac\xbe\xaf\x19\x7c\xee\xda\xbd\x7e\xe9\x68\x74\xaf\xd9\x63\x67\x6b\x9d\x7d\xec\x0e\x61\x9c\xb1\xed\xb9\x90\x71\xd4\xe2\x9a\x4e\xc1\x37\xcb\x3e\x8b\x71\x00\x1e\xbf\x9b\xa9\x74\xfe\xd2\xee\xa6\xf9\x2a\xe2\x5c\xff\xf8\x90\x68\x2d\x82\xdd\x3a\x95\xd0\xab\x76\xb3\x9b\xe7\x91\x70\xd7\x05\x7e\x05\xc2\x9d\xdc\xa4\xe9\x41\xb7\xfc\x5a\xe4\x30\xc3\x90\xb8\x91\xeb\xb1\x01\xa7\x2b\x24\x45\xb9\x76\x6d\x43\x00\x71\xa0\x20\xf3\x40\x2c\xcd\xc8\xcb\x67\x71\x87\xee\xf9\xde\x57\x50\xbf\x80\xf4\xc7\xfa\x85\x9f\xc8\x56\x37\x06\x7f\x25\x62\x4c\x73\xa6\x06\x30\x12\x87\x43\x09\x66\x9f\xf0\x7a\xcc\x07\x5b\x9c\x64\xed\xd5\xa4\x7a\x0c\xb7\x7b\xa5\x4f\xf1\x5c\x57\x9c\x9c\xc4\xa7\xd3\x00\x68\x28\x37\x84\x55\x55\xcc\xc9\xea\xaa\x2c\x32\xbf\x0e\x00\xc7\x45\x4e\x4f\xc5\x6e\xc3\x83\x91\xfc\xdf\xfe\x41\xe0\xa4\x29\x52\xc3\x68\x47\x1b\xed\x3f\xba\xcf\xd4\xcf\x47\x63\xb5\x08\x78\x5a\x99\x2d\x85\x7d\x05\x0f\xf8\x96\x92\xa5\x11\x4b\x5b\x06\xc6\x03\xbe\x20\x31\x50\x3b\xe1\xed\x8b\x58\x1d\xce\x06\x32\x69\x33\xed\xfb\x58\x65\xa7\xb9\xbb\x2d\x5e\x7b\xc7\x30\xaa\x0a\xa1\x4e\x05\x25\xbe\x65\xca\x78\x91\xb5\x16\xa1\xc0\xf3\x24\x31\x0d\x44\x2b\xce\x8f\xf3\x58\x56\xde\xbd\x85\x9e\xb6\xeb\x8e\xe9\x8f\x12\x41\xf4\x2b\x18\xcd\xb2\x05\x5a\x22\x00\x8b\xf6\x32\x9e\x95\x15\xa8\xf7\xaa\xef\x2b\xd4\xf1\x65\x34\xed\xdb\x06\xba\x6e\x98\xb7\xd9\x6f\x74\xd8\xd8\x7d\xf5\xae\x6c\x7a\x80\x76\x82\x64\xa9\x6a\x50\x19\x5a\xca\xa3\x9e\xc3\x5c\x53\x5c\x9a\xbc\xc3\xed\xe8\x92\x4f\x43\xd8\xfc\x75\x06\xd7\xe6\xaf\xc8\x16\x77\x21\x3a\x18\xc9\xb6\xc4\x63\x9a\x16\xb8\xc0\x73\xd5\xab\xa9\xc8\xd0\xb2\x04\x13\xe9\x88\x86\x56\x1b\x9c\x90\x97\x7a\x8c\x7e\xac\x12\xce\x61\x08\xbb\xd9\x21\x05\xe0\x69\x5e\x22\x07\x95\xf1\xc8\x47\x43\x09\x63\x7f\x23\x18\xd5\x84\xb4\xaf\x54\x71\x63\xbc\x40\xd9\x8d\x5b\x52\xd1\x75\x8f\x8b\x05\x45\x6c\x41\x4a\x59\xc2\x7c\x65\x8b\x99\xdc\xda\xdb\xb4\xb1\x46\x6a\x48\xfd\xd4\x35\x33\xf0\x22\xa5\xf3\x70\xef\xa5\x57\xb2\x36\xd0\x19\x26\x2e\x89\xca\x6d\x4c\x5d\xeb\xa8\x48\xa3\x22\x94\xc7\xea\xda\xe6\x8c\x29\x5f\x38\x86\xcf\x6f\xa6\x70\xf6\x5f\x8d\x34\x38\x60\x01\x7f\xc0\x8b\xe0\x6e\xb6\xb9\xb9\xc1\x93\xba\xf5\xfc\x39\xfd\x9a\xe5\xfc\xd5\x76\x8e\xc2\x87\x8a\xb6\x83\x31\x62\x2b\xa7\x21\x5e\x8e\xdf\xde\x03\xda\xa8\x1d\x15\x95\xd9\xb9\x97\x67\x3c\x32\xdb\x1c\xb6\xdd\xf8\xb3\xb3\x60\x6b\x78\xd4\xbf\x9a\xee\x61\x6b\xc7\x19\xea\xf6\xb7\x76\xce\x05\x08\xef\x5c\x8b\x47\x4d\x1c\xaa\xf3\x3c\x30\xb9\x0d\x9c\xfc\xce\x66\x67\xc6\x5e\xd5\xae\xf7\xeb\xf1\xc2\x93\x82\xa8\x41\xef\x02\x7d\x2a\x19\xfd\x65\xfc\xe7\x8d\x6a\x22\xd7\x00\xb6\x54\x60\x5f\x61\xef\xee\xbb\xb4\x36\x50\x7c\xb5\x6f\x17\x28\x37\x64\xe1\x09\x5e\xbb\x90\x83\xea\x28\xca\x02\x37\x5e\x85\xba\x54\x38\xa7\xc5\xd5\x8a\xab\x05\x47\x8e\x8d\x6a\x62\xfa\xc8\x00\x56\x5e\x2c\xdc\x95\x7f\xaa\xb3\xf1\x8e\xda\x1c\xfd\x61\xba\x2f\xf7\xe9\x1a\xe4\xdd\x74\x18\xba\xcc\xf2\x65\xe5\xc9\xf2\x73\xf6\x66\x4c\xc8\x4d\x81\x66\xbc\xc8\x6e\x0a\x8c\x18\x6b\xa2\x0a\xb1\x2a\x9b\xbb\xe9\xb5\x2c\xf6\xde\x43\x6b\x5b\x22\x35\xf0\x27\x24\xeb\x4f\xca\xd1\x1f\x98\x9a\xc7\x12\x3e\xfd\x21\x87\x66\x1d\xc0\x56\xb0\xd0\x97\x20\x2c\x42\x9a\x53\xbe\xde\x68\x7d\x13\x1e\xe7\x00\xb5\x34\x37\x82\x62\xa4\x2c\x7d\x65\x84\xc0\x95\x05\xa3\x1f\x57\xf6\x91\x8d\x29\xc1\x7f\x21\x57\xcc\xef\xb4\x17\x51\x1d\x76\x2e\x86\xf5\x5d\x0b\x8b\x26\xee\x5b\x5e\x09\xdb\xe2\x92\x51\xc7\x75\x30\xaf\xcf\xb3\xef\x2a\xd8\xf3\x5c\x04\x7b\xc0\x35\xb0\xc8\x69\xae\x69\xd9\xe3\xd7\xbf\xa2\x56\xdf\x0e\x33\x6d\xe5\xd5\xfc\x75\xcf\x0d\x7b\x2f\x7c\x6d\x79\xdd\xab\xf3\x72\x5e\xa4\x55\xa0\xff\x82\x9e\xb9\xa7\x10\x65\x2c\x99\xae\xf0\x45\xca\x6e\xc2\xa0\xf6\xe5\xb1\x20\x88\x99\x80\x47\xdc\xc7\x11\xc5\xcd\xc9\x4b\x18\x04\x28\x5a\x32\xf3\x24\xb8\x27\xcb\xb0\x06\xa7\x14\x27\xe9\x17\x96\x08\x24\x11\xbf\x04\x7c\x4b\xde\xd3\xed\x20\x31\x3f\x00\xdd\x51\x96\x91\x15\xe6\xa7\x79\x0f\x46\xbd\xca\xfd\x0e\xcc\x4d\xbb\xe8\xf8\xec\xc3\xec\xe2\x64\x0a\x23\x1d\x39\xa0\x4e\x6f\x82\xef\x42\x4f\xb7\x6b\xa9\x78\xba\x6c\x85\x99\x05\x4b\x32\x67\xc9\x98\xa2\x94\xa3\xa6\x63\x24\x12\x47\xd8\xa0\x33\xd9\x35\xd2\x09\x3b\x59\xf1\x33\x32\x3f\xb9\x45\x98\xb3\x60\xfb\x52\x9f\x88\x47\x1a\x13\x6b\xe1\x92\x93\x74\x5c\x2e\x1c\xaa\xfe\xad\x2e\xa9\x00\x50\x60\x51\xed\x33\xc9\x7e\xfa\x85\xd5\x57\x07\xbf\xf3\xaf\x0b\xaa\x3f\xbf\x23\x77\x7e\xbf\x2d\x0f\x1c\x24\xb5\xd2\x62\x1c\xc0\x03\x98\xc4\x37\xce\x4d\x1b\xa2\xde\xc2\x88\x04\xfc\x38\x40\x3b\xeb\xe6\x02\x6b\xcc\x61\x47\x6f\xba\xba\xb5\xc6\xc6\xf5\xf7\xd7\x14\xfd\x3b\xa2\x0b\xfd\xc0\xf0\x5c\x1d\x37\x40\xeb\x99\x82\x3d\x0b\x7d\x17\x3f\x8d\x22\xc9\x9f\x0e\xac\xba\x96\x77\x39\x17\xfe\xad\xa8\xde\x16\x65\x80\x9f\xf0\x33\xf6\xcb\x43\x83\x15\x43\x80\x71\x5a\x64\x7c\xf0\xa3\xeb\x3d\x6f\x53\x0a\xd2\x2f\x0c\xbc\x06\x14\xfd\xba\x2a\x28\xda\x19\xa4\x5f\xd8\x1e\xcb\x6f\x06\xbb\x41\x60\x94\x09\x60\x8c\xbe\x88\x61\xa3\x93\xf1\x6c\x67\xbd\x4c\xef\xa6\x88\xd3\x02\xb1\xe4\xf0\x60\x13\x1e\x26\xbf\xf9\xd3\x8e\x1b\x97\x64\x95\xff\x92\xf2\x6c\x71\x46\xe6\x6c\x27\x3c\x46\x1b\x6e\xf0\x1a\x0c\x02\xf6\xd9\x5b\x4b\xc4\x9c\xe8\xd9\xa5\x34\x0b\x54\x96\xc9\x30\xdb\xe8\x00\x1c\xfc\x18\x6c\x63\xee\x40\xac\x2f\x2b\x7b\x78\x8d\xab\x42\x51\xb4\x12\x01\xb7\x5a\x50\xc4\x16\xad\xdd\x65\xf9\x47\x62\xfe\x3d\xb2\x1e\x52\xc5\xab\x81\xd1\xea\x3f\x48\x1c\x7a\xdb\x9a\x60\x47\x77\x8c\x58\xca\x30\xbc\x43\xc1\x2a\x91\x9a\x76\x90\x0c\x06\x2e\x77\xbd\xe6\x2f\x74\x57\x89\x54\xb4\x56\x38\xf0\x1a\x5c\x6b\xc5\xde\x41\xc2\xda\x0d\x41\x46\x30\x47\x77\x7c\xd7\xdb\x1f\x39\x8b\x14\x17\x75\xcb\x06\xbc\x06\x72\xc8\x48\xff\x3e\xa2\xa8\x2a\xd3\x0c\xed\xec\xff\xcb\xff\xdb\xf9\xfc\x39\xff\x7e\xf7\xc7\xfd\xf9\xb0\xc5\xbf\x14\x52\x38\x04\x39\xca\x22\xb8\x01\xa0\x88\xaf\x28\x06\xaa\xd5\x61\x74\x4d\xc9\x72\xbc\x48\xa9\xd0\xcc\x1d\x31\xcc\x13\x5e\xf1\x57\x40\x0f\x6a\x42\x55\x17\x4a\x80\xd5\xb0\xfe\x81\xf1\x94\x72\x94\xbf\xb9\x4f\xc0\x40\x64\x3e\x83\x61\x0c\xd2\x96\x9f\xc4\x95\xa7\x4f\x6a\x2b\x74\xbf\xcd\x65\x14\x8d\x56\xb5\xa4\xfe\x21\x0e\x28\x9c\x6b\x02\x0e\xa3\x00\xe4\x16\x51\x5a\xe4\x88\x25\xf1\xe5\x29\x44\xba\x2f\xed\x7d\x3b\xe0\x53\xd7\x00\x20\xc5\x1b\xa7\x4b\x94\x58\x8b\x1a\xd6\x8c\x4f\x3e\x81\x01\x5b\x0c\x86\x60\xb0\x97\x0d\x9a\xa7\x42\x58\xbb\xd0\x5e\xc6\x5e\x06\x47\x6d\xa2\x4c\x65\x37\xe8\x0b\x78\x0d\xce\x53\xbe\x18\x5d\x97\x84\xd0\x1d\xf9\x23\x95\xdd\x2e\x3b\xbb\xdf\x1d\x1e\x1c\x1c\x1c\x84\x65\xa2\x24\xf3\x1d\x6b\x49\xe0\x7b\x30\x48\x06\xe0\xfb\xc6\xbc\x7c\x0f\x06\xfb\x42\x0e\xe4\x2c\xaf\xc5\x1b\x39\xdd\xf7\x60\xb0\x64\xf5\x42\xe5\x63\x4b\xf2\x0d\x21\x47\x94\x46\xa5\x5b\xf2\x82\x91\x12\x8d\x04\x21\x03\x44\xe9\xab\xc1\x10\x88\x11\x41\x6a\xc5\x1f\x86\xb8\x76\x57\x3b\xcd\x14\xbb\x60\x2d\x9c\xc3\x88\xaa\x04\x67\x47\x49\x79\xa3\xb8\xa3\x9c\x60\xb4\x2b\x8c\x88\x20\x7d\x6b\x9d\xf1\x37\xbc\x9e\x50\x6e\xdb\x12\x31\x96\xce\xd1\x10\x64\x57\x1d\x96\x41\xb5\x40\x0b\x23\x2d\x36\x71\x4f\x6c\xd4\x8e\xf0\x44\xc7\x29\x47\x3b\xbb\xbb\xa3\xb9\x5a\x4e\xc0\x0d\x81\xad\x55\xb6\x76\x31\xc2\x7e\x26\xcd\x6f\x51\x35\x29\xeb\x78\x4f\xc1\xb3\x50\xec\xa7\xf6\x24\x22\x31\x6c\x94\xd9\x81\x63\xb3\xe1\x8f\x64\x7a\x1f\xcf\xb7\xdb\x06\x4d\x9d\x0a\x51\xb7\xd2\x68\xcd\xc2\x04\x34\xbc\x14\x61\x12\xe3\xe9\xb2\x4a\x22\x6c\xea\xd1\xe8\xe8\xa6\x83\x87\xf3\x09\x3c\x84\x57\x20\xc6\x2f\xe0\xec\xb7\x5e\x6a\x7c\xb3\x25\x83\x2b\x23\xda\x6f\xd5\xe9\x2a\xa6\x3b\xbe\x93\x35\x4b\x34\x56\xd0\x1e\xe8\x19\x70\x7a\x7a\x74\xc8\xd1\x53\x57\x13\x21\xc5\x6c\x41\x28\xd7\x51\xc3\x74\xd5\x51\x63\xd3\x32\x91\x48\xa0\xde\x58\xdc\x68\x85\x1e\x9d\x11\x3c\xd7\x33\xec\xb1\x6c\x81\xf2\x95\xfd\xfd\xa6\x99\x7e\x76\x72\x57\x51\xc4\xea\x5a\x8f\x24\x4e\xbf\x71\x5a\xb2\xd4\xe9\xaa\x57\x82\x97\x61\x7b\x34\xb6\x6f\x53\x8d\xc8\x85\xf4\xd3\x3c\x40\xb0\x3e\xc8\x75\xce\x82\x2b\x7d\x1a\xfa\xb9\xfe\x36\xc1\x67\x98\x80\xcf\x75\x9b\x8a\xf4\x01\x9b\xcd\x67\x38\x04\x9f\xa1\x36\xe6\x2d\x80\xbe\x4d\x2c\x01\x0c\x1e\x87\xaa\xaa\x01\x16\xa9\xc4\x69\x82\xe8\xb2\x60\x2c\x94\x61\x01\x37\xc5\x32\x60\x43\x5c\x03\x76\x75\x34\x6b\x3a\x36\x55\xaa\x9d\x9c\xe2\x5b\x72\x83\x42\x5f\x24\xaa\x9f\xa9\x66\xa6\x47\xee\xbb\x51\xf8\x14\x93\x4a\x07\xc8\x9c\x52\xa7\x29\x2a\x32\x13\x96\x68\xa2\x0d\x94\x9e\x44\x1b\x13\x77\x29\x4e\xb8\x72\xed\x7d\x41\xc7\xad\x5b\x1f\x09\x80\x90\x76\xc9\x17\xdf\xca\xd6\xff\xe4\x65\x6b\xc9\xc5\x7f\xfe\xba\x75\xa4\x1c\xfd\xad\x62\xfd\xad\x62\x5d\x93\xf4\xad\x62\xfd\xad\x62\xfd\xbf\xb4\xe5\x7f\xac\x8a\xb5\x34\xf1\xcf\x57\xb2\x6e\xfd\xfe\x57\xaf\x59\xb7\x53\xfd\xb1\x8a\xd6\xeb\x35\x78\x29\xbf\x05\xae\x6f\x07\x8c\xac\x2d\x16\xe8\x43\xad\x40\xea\xa6\x91\x1c\x07\x17\x28\xcd\xbd\xcf\x9e\xd8\x89\x4d\x2c\xf0\x32\xe0\x43\xc2\x07\x41\xb4\x78\xaa\xc3\x97\xe7\x2c\xa2\xba\xeb\xba\x22\xf9\xbd\xb3\xae\x27\xe4\x76\xd1\xd8\xf3\x77\xc8\xec\xb2\x85\xa0\xcb\x94\x2e\x89\x77\x92\x72\x8e\xa8\x1b\x3c\xc0\xc6\x40\x38\xa7\x1d\x22\x4e\x44\x99\x65\x6b\x2c\x1b\x03\x73\xc4\xd3\xa2\xdc\xe3\x6a\x01\xce\xe8\x93\xf1\x0c\xb4\x37\x66\xeb\x36\x29\x20\x43\x2c\x30\x76\x09\x6c\x86\x08\x2e\xdb\x40\x3d\xd3\xfb\x37\xa8\xb5\xcf\x57\x19\xa8\x77\x67\x78\xab\x60\xa7\x2f\xc8\xd9\x3e\xb8\x09\x06\x35\x0f\x0a\x66\x3a\x83\x98\xd0\xbd\x73\xd7\x18\x7b\x26\xfb\x32\x68\x90\x37\x4f\xc9\xea\x6d\xbb\xfd\x98\xb4\xfe\xff\x76\xf6\xbd\xc5\xfe\xfc\xe1\xd3\xef\xd0\xa7\x63\xd5\xf7\x25\xa7\xfa\x4d\xe0\x5f\xb5\x08\x7c\x2f\x71\x6a\x80\x69\xe6\x06\xbf\x93\x18\x36\x7b\x4f\xff\xe4\xa1\xf1\x8f\x6f\x78\xd7\xf4\xbd\x7b\x61\x2f\xea\x4d\x8a\x7e\x12\xd3\xff\xda\xac\xbd\x29\xaa\xdd\x59\xaf\xd3\xfb\xb4\xa7\x67\x00\x5e\x68\x6e\x74\x6e\x6a\xc7\xbd\xd5\xc0\xb0\xa1\xb7\x64\x2d\x00\xd1\x35\x39\x1f\xfc\xb5\xbe\x23\x10\xfe\x86\xa8\x5b\x3f\x89\x30\x6f\xab\xda\x49\x34\x0f\x77\x0c\x7d\x34\xc4\x76\x2b\x17\xb6\xf5\x73\xbd\x4e\x67\x34\x6f\x97\x55\xbc\x82\x41\x5b\x64\xf1\xed\x36\xca\x5c\xf5\x7d\x4c\x98\x6c\xa8\x7c\xb8\x8a\x51\x2b\x72\x6b\xd3\x62\x65\x97\x70\xd1\xc5\x32\xf0\x76\xc1\xc5\xfa\x7c\x84\x0f\x16\xff\x84\xff\xf3\x7f\x9d\xbf\x33\x9d\x82\x48\xdd\xbb\x2b\x49\x9a\x5f\x35\xf7\xee\xd4\x95\xd0\x2b\x14\x29\xa6\x44\xc6\x28\x75\x46\xb4\x8e\x58\xd8\x5b\x4a\x96\xc1\x1b\x7c\xfd\xd8\xa6\x2e\xae\x5f\x0a\xbe\xd8\x02\x57\xf6\xaa\x97\xf8\xec\x55\x72\xb4\xe2\x0b\x42\x8b\xdf\x50\xf0\x76\xaa\x37\x2a\xd4\xbd\x6e\x64\x87\xc1\x7d\xfd\x2e\x80\xc6\x79\xb2\x55\x68\xf1\xc2\x7c\xdb\x6d\x79\x8c\x4f\x89\xfb\x5f\xe8\xb6\x6d\xce\xec\x87\x24\xd1\x5f\xd6\xd7\x46\xe7\x18\x95\x48\xc8\x49\xd3\xb8\x0e\xa7\x22\x48\xc4\x3d\x46\x49\xfe\xcb\x63\x22\x58\xa5\xea\x22\x8d\x7b\x5f\x11\x5e\xa4\xce\x17\xd7\xd6\xf5\x97\x67\x21\x93\x9f\x1f\x14\x36\xb6\xb9\x3e\xa0\x3f\x0b\x65\x5d\xe2\x6b\xe0\xd3\xaa\x32\x81\x3b\x5c\x4f\x68\xdb\x8c\x5d\xfb\x9f\x00\x00\x00\xff\xff\x38\x4a\x19\xa9\x58\x72\x00\x00")

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

