// Code generated by go-bindata. DO NOT EDIT.
// sources:
// compiler/templates/entrypoint.go.tmpl
// compiler/templates/preload.gs
// compiler/templates/vm_file.go.tmpl
package compiler

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

var _entrypointGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\xcf\x4b\xfb\x40\x10\xc5\xef\xf9\x2b\xde\xb7\xe4\x90\xc0\xb7\x01\xaf\x42\x0f\x42\x51\x44\x14\x41\xd4\x63\x59\xb2\xd3\x30\xd8\xce\xc6\xd9\x4d\x7f\xb0\xe4\x7f\x97\x4d\x23\xd8\x7a\xa9\x87\x5e\x42\x66\xf6\xcd\xfb\xbc\x81\x69\x4d\xfd\x61\x1a\xc2\xda\xb0\x64\xd9\xb2\x93\x7a\xf8\x2d\x4a\xc4\x0c\xd8\x18\xc5\xb6\x81\xdf\x4b\x5d\xbd\x1b\x0e\x77\xea\xba\x36\x03\x62\x9c\x42\x8d\x34\x84\x7c\xf1\x1f\x79\xab\x8c\xeb\x19\xf2\xea\x55\xf8\xf3\x59\xd9\x29\x07\x26\x8f\xbe\xff\xad\xdd\xac\x93\xb4\x60\xb1\xb4\x43\x5e\xbd\x38\x0d\x64\xdf\x1e\xfd\xe0\x52\x1e\x46\xb6\x4d\x75\x63\x6d\x71\x55\x66\x40\xe3\x90\x52\x8d\x81\x00\x4b\x4b\x4a\xa1\xaa\xb9\x13\x2a\xca\xa1\x67\x62\x4c\xc6\xd5\xfd\x1c\x7d\x9f\xec\x9f\x68\xfb\xb3\x35\xca\x78\x79\xa2\x9c\xcd\x20\xbc\x1a\x8d\x01\xa5\xd0\xa9\x0c\x45\x3f\x7c\x49\x35\xb9\x1d\x0d\x55\x63\x71\xdb\x49\x1d\xd8\xc9\x03\xed\x31\xe1\x75\xeb\x34\x2c\x8c\xf7\x14\xfc\xe4\x88\x98\x3c\xfe\x9d\xc1\x39\x07\xd3\x2a\xad\x9c\xb1\x97\x03\x8c\x7b\x88\x09\xbc\xa1\x8b\x63\x7c\xad\xdc\x86\xcb\x61\x68\x47\x75\x17\xfe\xbc\xc7\x41\x1c\x23\x48\x2c\xa6\xdf\x07\x99\xce\xff\xf4\xa1\xff\x0a\x00\x00\xff\xff\x87\x68\xa4\x81\x3e\x03\x00\x00")

func entrypointGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_entrypointGoTmpl,
		"entrypoint.go.tmpl",
	)
}

func entrypointGoTmpl() (*asset, error) {
	bytes, err := entrypointGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "entrypoint.go.tmpl", size: 830, mode: os.FileMode(420), modTime: time.Unix(1530571510, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _preloadGs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x3f\x6f\xea\x30\x14\xc5\xf7\x7c\x8a\x23\x26\xfb\xc1\x8b\xc2\x5b\xde\x90\x52\x89\x3f\x5d\x3a\x94\x01\xb6\xb6\x83\x81\xeb\xe0\xd6\xd8\xd1\xb5\xd3\x2a\xaa\xf8\xee\x15\x09\xa0\x06\x51\xb5\x63\x6e\xce\xf9\xdd\xa3\x7b\xac\x2b\xb7\x8e\xc6\x3b\x2c\x22\x1b\x57\x2c\xfd\xa4\x8e\x34\x66\x56\xb5\x08\x12\x1f\x09\xf0\xa6\x18\x1b\x15\x15\x46\x78\x7c\xce\x13\x40\x7b\x86\x38\x4c\x0d\x46\xc8\x72\x18\xdc\x20\xa4\x96\x5c\x11\xb7\x39\x4c\xbf\xdf\xfa\xd0\xb8\xd2\xb2\x0a\x5b\x11\xd2\xf5\x56\xf1\xd4\x6f\x68\x1c\x85\x91\xf2\x80\xd9\x27\x00\x53\xac\xd8\x35\xc2\x3c\xd9\x27\xc9\x39\xcd\x39\xc5\xd2\xb7\xc1\x84\x6a\xa9\x47\x47\x3b\x4c\x35\xfb\xdd\xf4\x48\x4e\x55\x59\xda\x5a\xb4\x7f\x06\x50\xb2\x4b\x9c\x55\xbb\x52\xf8\xd5\x4b\x07\xd3\x7b\x72\x3d\xf4\x71\xbf\x98\x3f\xa4\xa1\x31\x1a\x5d\x1f\x54\x03\xb8\xca\xda\x01\xfe\x5d\x50\x26\xa4\x3d\xd3\x8c\x4a\xeb\x6b\xd1\x41\x45\xae\xe8\x62\xe3\x15\x95\x56\x36\x5c\xc8\xc6\x3a\x12\xff\x8e\x38\x77\x77\xcc\x9e\x7f\x44\x2e\x2c\x51\x29\x02\xad\xbd\xdb\x7c\x69\x31\x44\xc5\x11\x23\x38\x7a\xc7\x4c\x45\x12\x32\x2d\x28\x2e\xcd\x8e\x84\xfc\xae\xd8\x21\xfd\xef\x74\x6a\x34\x84\xb8\x06\xc0\xdf\x96\x2f\x71\x8b\xd3\x6a\xfc\xc1\x30\xcb\x32\x79\x32\x03\x2b\x26\xf5\x9a\x37\x1f\xfb\xe6\x0d\x74\x0f\xb6\xaa\x8a\xa9\x77\xc1\x5b\xba\x7a\x8a\xcf\x00\x00\x00\xff\xff\xf8\xdd\x34\xe2\xad\x02\x00\x00")

func preloadGsBytes() ([]byte, error) {
	return bindataRead(
		_preloadGs,
		"preload.gs",
	)
}

func preloadGs() (*asset, error) {
	bytes, err := preloadGsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "preload.gs", size: 685, mode: os.FileMode(420), modTime: time.Unix(1530143358, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vm_fileGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xdd\x53\x1b\xc9\x11\x7f\x46\x7f\x45\x67\x03\xa9\xd5\xd5\xb2\xd8\x97\x54\x1e\x94\x38\x55\x80\x74\x2e\x62\x0c\x14\xe2\x7c\x95\xa2\x28\x3c\xda\xed\x5d\x8d\xd9\x9d\xd9\xcc\x8c\x00\x99\xd2\xff\x9e\xea\x99\xd9\x2f\x21\xd9\x70\xe7\x87\x3c\x81\xe6\xa3\x3f\x7e\xdd\xfd\xeb\x9e\xad\x58\x72\xc7\x72\x84\x92\x71\x31\x18\xf0\xb2\x92\xca\x40\x38\x00\x08\x66\x4b\x83\x3a\x18\xec\x04\x89\x2c\x2b\x85\x5a\x1f\xe4\x5f\x79\x15\xd0\x56\xa2\x96\x95\x91\x07\x8c\xf6\xdb\x9f\x09\xaf\xe6\xa8\xe8\x06\x8a\x44\xa6\x5c\xe4\x07\x33\xa6\xf1\xef\x7f\xb3\x87\xb2\xd2\xd0\x16\x97\xc1\x80\x7e\xe6\xdc\xcc\x17\xb3\x38\x91\xe5\x41\x8e\xe2\x4d\xc2\x53\x3c\xc8\x75\xa2\x78\x65\x0e\x50\xe4\x5c\x60\xb0\x76\x4c\xc9\x19\x2a\x73\xa7\x78\x89\xe2\x40\x1a\xe3\x04\x3d\x3d\x81\x62\x22\x47\xd8\x15\x3a\x82\xdd\xea\x2e\x87\xd1\x3b\xd8\x8d\xdf\xcb\x0b\xe7\xd9\xd1\xf2\x8c\x95\xa8\x2b\x96\x20\xac\x56\x03\x80\x83\x03\x38\xb1\x6e\x72\x91\x43\x8a\x15\x8a\x14\x45\xb2\x84\x1a\x89\xa7\x27\x2b\x26\xa6\x6b\xb0\x5a\x41\x26\x15\x08\x66\xf8\x3d\x76\x4e\x93\x6d\xf5\x41\x27\xed\x82\x99\x39\xac\x56\x81\xb3\x09\x45\x4a\xda\x86\x83\x41\x22\x85\x26\x44\x9f\x9e\xf6\x6b\x4b\x6f\x23\xd8\xc5\x72\x86\xa9\x33\x75\x42\xff\xea\xc6\x38\x12\x6b\x77\xe3\x93\x31\xe9\x9f\xcb\x22\xd5\x60\xe6\x08\x29\x33\xcc\x9a\x63\xb7\x53\x4c\x21\xe3\x05\x76\x2e\x9c\x2b\x9e\x7b\xb3\x9d\x19\x3d\x41\xef\xe0\x73\xbb\x34\x26\x59\xab\xd5\xe7\x81\x35\xac\x35\xd7\x1b\xe0\xaf\x3c\x28\x56\x39\xdd\x39\x0a\xd4\x5c\xc3\xa7\x8f\xd6\x02\x7b\xa6\x56\x65\x96\x15\x76\x6f\x69\xa3\x16\x89\x81\xa7\x01\xc0\x04\x7e\x72\xd1\x8c\x27\xf6\xcf\x00\xe0\x03\x5c\xdf\x50\x6e\x0d\x56\x56\xdb\x19\x3e\x74\xae\x26\x0a\x99\xc1\xbe\x4a\xb5\x10\x86\x97\x68\xf5\xba\x0c\xe9\xab\xcf\x16\x22\xe9\x8b\x09\x87\xf0\x53\x47\x28\x19\x62\x90\xc0\xf6\xb6\x9c\xe1\x43\x18\x74\x65\x04\x91\x0b\xa7\xbb\x10\x44\xf0\xd7\x37\x11\x04\x93\x47\x4c\x16\x06\x83\xe1\x00\x40\xd2\xf5\xbf\xb4\x67\x48\x26\xc0\x64\x04\x06\xa3\x01\x00\xe1\xad\xd0\x2c\x94\x00\xe9\x1d\xb3\x67\x7f\x59\x88\xc4\x70\x29\x3e\xe0\x12\x02\x57\x5d\xb7\x4c\x6b\x34\x3a\x20\xcb\xdc\x8a\x06\xb7\x04\x5c\x18\xf9\x7a\xd7\x43\xd9\xf5\x76\xf8\x12\xc5\xe1\x10\x50\x29\x8a\x63\xaf\x84\x90\xa7\x1b\x53\x73\xbf\x97\x9b\x3c\x75\xd9\xb4\x2d\xef\x64\x3c\xf1\x35\xa1\xaf\x83\xf6\x46\x70\x03\xef\x80\x0c\x0e\x87\x3e\x03\xc0\x81\x78\x70\x00\xbf\x0a\xaa\x3e\x54\x36\xe1\x2a\xfa\xcb\xcd\x1c\x5c\xe9\xa4\x98\x14\x4c\x61\x0a\x6c\x26\xef\xd1\xde\xf0\x48\xb7\x4e\x8f\xc3\xb5\x64\x1f\xfa\x98\xf8\x52\x74\x0e\xf8\x6b\x82\x17\x14\xa2\x2d\x41\xaa\x14\x16\x92\xa5\xbd\xf0\x50\x48\xea\x50\xf8\x7d\x28\xf8\x4c\x31\xb5\xfc\xa3\xa1\xe9\xa8\xeb\x05\xa5\x4e\xa6\x78\x12\x9f\x4a\x96\x4e\xad\x86\xb0\x3e\x1e\x7f\xd1\x41\xd4\x07\xfa\xf6\xf6\xe2\x72\x72\x7a\x7e\x38\x0e\x6e\xc2\xe1\xf0\x7b\x39\xe8\x2c\xb6\x4e\x92\x3c\xe7\xe2\x06\x37\x9e\xa7\xe4\xa7\x8f\xaf\xcc\xba\x56\xd5\x4b\x1c\x5c\xab\xca\x35\x17\x27\x67\x57\x97\xff\xb9\x38\x3f\x39\xbb\xfa\xb6\x97\xe8\x2b\x97\x3c\xf0\xff\x3f\xa3\x31\x66\xec\x0a\x0a\xa3\x96\x50\x49\x2e\x0c\x84\x29\x66\x5c\x60\x0a\xb3\x25\xdc\xa3\xd2\x5c\x8a\x08\x16\x7a\xc1\x8a\x62\x09\x63\xac\x0a\xb9\x24\xad\x2f\xf4\xbf\x63\x44\xcf\xf3\xdb\x88\x7e\x50\x81\x91\x77\xc7\xac\x28\xea\x6b\xbf\x71\x33\xbf\xe2\x25\xca\x85\x09\x03\xa7\xcf\x92\x8f\xc7\x0a\x95\xfa\x5e\x5c\x5d\xa3\x72\xc9\x2b\xbe\x60\xe2\x93\x37\x5d\x0a\x56\xf2\xc4\xfa\x51\x70\x71\x87\x69\xdd\xd2\x32\x2f\x64\x03\xf9\x78\xaa\xff\xe3\xbc\xd3\x1a\xd5\xc3\xe1\x9e\x29\x0b\x84\x5d\x19\x80\xfd\xff\x9d\xad\x4e\x47\x36\xfb\xfb\x70\x34\x79\x7f\x72\x06\x67\x87\x57\x27\x9f\x26\x70\x71\x78\xfc\xe1\xf0\xfd\x04\x4e\x3e\x5e\x9c\x5f\x5e\x4d\x5f\xdb\xfc\xf7\x9f\x77\xff\x8d\x1d\x1c\x16\x22\x45\x6a\xf9\xf5\x45\x3a\x25\x7c\x7f\xbe\x15\x55\xf3\xd3\x76\x84\xba\xa3\x58\x17\xbd\x56\x47\x6b\xad\xd0\xd1\xb6\x59\x21\xb2\x07\xc9\xc4\xce\x91\x26\xf9\xed\xe6\x74\x59\xce\x64\x71\xc5\x66\x05\x8e\xa0\x64\xd5\xb5\x36\x8a\x8b\xfc\xe6\xa7\x9e\x62\x42\xfd\x69\x15\xb5\xa4\xe7\x61\xc9\x78\xfa\x18\xc1\x6e\x66\x61\x21\xe1\xa7\x36\xf8\x74\xbc\xc3\xea\x87\x29\x0d\x6b\x4d\x2a\xb8\x62\x40\xd5\x9d\x7c\x9a\xbd\x35\x2b\x63\xfa\x9d\x35\x51\x27\x54\x8c\x04\x6d\x8d\x06\x43\x56\x5b\xd0\x32\x7b\x8c\xa7\x8f\x1b\x61\xb3\xd6\xaf\x43\xd1\x13\xea\xc1\xa0\x95\x11\xc8\xb8\xd3\x86\xed\x49\xfb\x6f\xed\x7d\x37\x44\x71\x07\xbe\xeb\x0d\x62\xa9\x25\xf5\xad\x5b\xef\x19\x94\x2f\xb6\x90\x08\x21\x0f\x46\x3d\x29\x6e\xe2\xc6\x3a\x8d\x5b\xe2\xea\xa5\x86\x23\x38\x6b\x5b\x10\xf5\x4c\xa5\x3a\xe7\x99\xbd\xfd\x27\x5b\x05\xbe\x3b\x76\x6a\x7f\x43\x4b\x73\x65\x32\x39\x1b\x6f\x2f\x92\xcd\xe4\x61\x21\x3b\x1f\x7f\xf0\xdb\x9e\x26\xd0\x8e\xf2\x04\xce\x1d\xba\xd6\xb6\x4e\x9a\x76\x84\x70\xf5\xdf\x93\xd3\x6f\xea\x5e\xa7\x5b\xf1\xa4\x3e\x6e\x84\x13\x37\xac\x56\xc1\xf0\x99\x3d\x63\xe0\x8d\x21\xb2\x9f\x93\xbd\xb1\xd7\x8f\x4b\xdf\x64\xa7\xee\x70\xa0\xc1\x55\x4d\xcf\xc6\x59\x43\xc3\x0c\x35\x8d\x84\xc7\xf6\xfd\x12\xae\xb9\xf5\x9d\xb8\x38\x81\x4f\x2b\x1f\x9c\x74\xf6\x96\x24\x0a\x7c\x08\xed\xfb\x29\x3e\x5a\x64\x19\xaa\xa1\xdd\xfa\x79\xf3\xd6\x8e\x56\x09\xed\xb8\xd5\x33\x7c\xb8\x44\x96\xa2\x0a\x3d\x7a\xda\x9a\x40\x64\xc9\xef\xe1\x9a\x6c\x3d\x2a\x64\x72\x37\xe5\x5f\xd1\xcd\xd1\x40\xde\x21\x2b\x49\x86\x7b\x83\x91\x90\xf3\x5f\x8e\xc2\x59\x04\xfc\xfe\x7a\x74\x33\x1c\xec\x58\x44\xd1\xba\xeb\x1e\x65\x74\x66\xec\x16\x43\xbf\x32\x35\xe9\xc4\xbf\xdc\x22\xd0\x2a\x21\xbd\x28\x12\x67\x8e\x2d\x5b\x2f\x7e\x6a\xf5\xb9\xf5\xa7\xe9\xc8\xeb\x8f\xe0\x72\x04\x5e\xcf\xca\xa1\xd6\xb6\x3a\x2e\xe3\x63\x59\x2d\xc3\x74\xf6\x36\x6a\x85\x0e\xff\xf1\x32\x60\x77\xf2\xaf\xaa\x11\x45\xcf\xd0\x0e\x4c\xe9\xec\xed\x70\xb0\xb3\x1e\xa2\x9d\x9d\x75\x39\x3b\xab\xc1\x8e\xb7\xa7\x6b\xce\xcf\x11\xe4\x5f\xd5\x8b\x25\xe4\x5f\x55\x7c\x5c\x48\x8d\xe1\x70\x50\x1f\x48\x67\x3f\xc7\x47\x14\xbc\xb0\x4e\xe9\xb1\x9d\x59\x29\x81\xc7\x9d\xce\xeb\xc8\x17\x1a\x02\xfa\xcd\x0d\xbb\x7a\xd0\x12\xf6\x6d\xc3\xd6\x8e\xaa\x55\xec\xb8\x7a\xb5\xea\x57\x4a\x87\xf7\xea\x9a\xb1\x7d\x5d\xf5\x2b\xc6\x1d\x6b\xda\xe1\x76\xd6\xfe\x46\x47\x5f\xd7\x17\x92\x33\x40\x0f\xf0\x46\x00\xcd\x30\x43\xb7\xf4\x89\x15\x0b\x57\x5f\xd4\x58\x54\xbe\x28\x51\x18\x98\x32\xc1\xcd\x12\x8e\xe7\x98\xdc\x69\x97\x1a\x05\x0a\x2b\x29\xae\x0f\x9d\x72\x6d\x86\xf0\x2f\x52\x5a\xa0\x70\x76\x1f\xaa\x5c\xd7\x2f\x38\xf0\xb3\x62\x9e\xa3\x8a\x27\x34\x33\x64\x61\x60\xa4\x84\x92\x89\x25\x30\x2f\x45\x43\x45\xec\x90\x52\x1b\x6a\xa0\xd8\xd3\x34\xed\xed\xe9\xfa\x99\xb7\xd6\x5b\xc0\xda\x41\x4e\xa0\x3a\x95\x09\xa3\x0d\x57\xf8\x4d\x36\xda\x13\xe7\xe4\xe0\x47\x76\x87\xc7\x0b\x6d\x64\x69\x6d\x08\x83\x46\x8b\x9d\x63\x48\xc5\x76\xa3\x6c\xc7\xa8\x2f\x04\xf5\x23\x65\x3b\x1a\xff\x7c\x3d\x1a\x19\x3e\xfc\xdf\x81\xb1\xd1\xa6\x4d\x58\xf4\x86\x17\xe6\x86\x17\x66\xcb\xa1\xf1\xbf\xe9\x7b\xae\xaf\xb6\xd5\xd4\x24\xdb\x9f\xc9\x2b\xd6\xb4\x73\x62\x4e\xd6\x59\xb2\x59\xcd\xe2\xc9\xa3\x99\xf2\xdc\x1d\x51\xdd\xfd\x86\x67\x7a\xf1\x08\x3b\x27\x86\xf1\xe4\x91\x3a\x7b\xb8\xad\x37\x6c\x08\x4d\x22\x17\x45\x0a\x42\x1a\x40\x7b\xb7\xc1\x03\xf6\x52\x90\xd9\xc6\xf0\xf4\x8c\x7a\x6d\xac\xd6\x6d\x08\x51\xa9\xdf\x1d\xc5\xac\x34\xf1\xb4\x52\x5c\x98\x9e\x2b\x46\x31\xa1\x0b\x66\xb0\xe7\x8d\x0d\x6c\x2e\xe1\x9e\xa8\xa0\xef\xc6\xb0\x4e\x79\xfd\xc0\x4d\x32\x87\x7b\xc2\xb9\x87\x7e\x1c\x9a\x65\x85\x43\x0b\x64\xc2\x34\xae\x47\x6b\x64\x5d\xe8\xc5\xf3\x99\x84\xb5\x2b\xb6\xf9\x62\xc6\x16\x85\x71\xb7\x51\xa9\x8f\xda\x3e\x15\x7a\x7e\x35\x09\x64\xbf\x67\x95\x5c\x97\xcc\x24\xf3\x11\x05\x0c\x13\x83\x29\xec\xe9\x08\x72\x69\x60\xef\xaa\xae\x9d\x8e\x96\x20\x82\xfb\xcd\xc8\x67\x61\xc0\x7a\xa2\x13\x29\xfc\xeb\xd2\x01\x3c\x72\xe1\x76\x66\xfd\xee\x18\xb5\xd7\x57\xbd\x6f\x90\xdd\x92\x52\x04\x51\x04\xbb\xaa\xa9\xa9\x4b\x3f\xfe\x6d\x2d\x2b\x77\xc0\x15\x95\xea\x15\x95\xea\x2c\xd9\x30\xa9\x5e\x51\xf5\xe6\x54\xbb\x8d\xe6\x14\x05\x69\x6e\x18\xad\xd6\x5e\x9f\xe1\x19\xe4\xa6\x39\xf9\xa6\xf9\x26\x4a\x29\x6e\x9b\xdc\xda\x73\xc4\x4b\xce\xe2\xf7\x28\x50\x31\x83\x4e\xe0\xd4\x4e\x7b\x10\xa8\xa0\xfd\x52\xf5\xa2\x36\x18\xf6\xa5\x1d\xaa\xbc\x16\xc5\x02\x9f\x49\xce\x4a\xfc\x6f\x63\xe5\xdb\xc6\xca\x29\x17\x79\x81\x4d\xec\xa4\x30\xf8\x68\x60\x1f\x98\x31\x58\x56\x86\x58\xd8\xef\x99\x39\xb3\x3f\xed\x83\x41\xa1\xf9\xc4\x8a\x3e\xed\xd8\x98\x5f\x49\xdb\x4d\x43\xf5\x66\x1b\xd1\x6c\xc9\x64\xaf\xc5\x16\x60\x37\xdb\x32\xc6\x8b\x85\xc2\x36\xdf\x3c\x35\x6c\xa1\x8c\x2c\x7c\x5d\x4a\xae\xe7\xf5\xb3\xa4\xf4\x22\x9c\xc3\x3e\x45\x0a\xdd\x7e\x97\xbf\x9a\x73\xdd\x32\xe1\x9c\x69\x28\x17\x85\xe1\x55\x03\xaa\x86\x7d\xff\x5d\x83\x82\x62\x89\x86\xc1\xbf\xa7\xc0\x94\x62\xfe\x53\xdc\xc6\x18\x24\xb2\xac\x98\xe1\x33\x5e\x70\xb3\x1c\x00\x7c\xd1\xe7\xb3\x2f\x1b\x10\x3f\x9f\x91\xec\xf0\xf3\xf5\xcd\xe7\x0d\x13\xe1\x37\x10\x6f\x39\xb1\x64\x77\x08\x4c\x78\x93\xa4\x15\x68\x2d\x6b\x5c\x61\x5a\xf3\x5c\x58\x3a\x70\x66\x06\x3f\x06\xfe\xfa\x0b\xe5\x26\xec\x9f\xb1\x40\x04\xbb\xb7\xdb\x49\xc0\x17\xbd\x1b\xe6\xd6\x4a\xdf\x0d\xd0\x16\xc1\x78\x8a\xfe\x93\x9d\xdf\x0e\xa2\x1e\x2b\xbc\x32\x6d\x5b\x10\x59\x9a\x02\x16\xd8\xe9\x8f\x96\x38\xf7\xae\xa8\x68\xbe\xe8\x1a\x0c\x02\x88\x92\x79\xe8\xbb\x8c\xaa\x9b\xa5\xea\xfd\xfa\xe1\x89\xfe\x2d\xa4\xfb\xd4\xdb\x62\x59\x3f\xa4\x89\x57\x52\x9f\x17\x6d\x49\x38\x38\x5d\xb9\x0f\x37\x90\xe7\x77\xca\x44\xd4\xcc\xa2\xdd\x4c\x60\xb5\x35\x27\x98\x48\x6b\x3d\x0b\xe1\x3f\x71\xc6\x2d\x75\xfe\x10\x5a\xac\x3f\xe7\x12\x5e\xbf\xd6\x4a\x36\xfa\xb3\x1a\x34\x00\xfd\x2f\x00\x00\xff\xff\xae\x12\x87\x5d\x7d\x1c\x00\x00")

func vm_fileGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_vm_fileGoTmpl,
		"vm_file.go.tmpl",
	)
}

func vm_fileGoTmpl() (*asset, error) {
	bytes, err := vm_fileGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vm_file.go.tmpl", size: 7293, mode: os.FileMode(420), modTime: time.Unix(1530571529, 0)}
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
	"entrypoint.go.tmpl": entrypointGoTmpl,
	"preload.gs": preloadGs,
	"vm_file.go.tmpl": vm_fileGoTmpl,
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
	"entrypoint.go.tmpl": &bintree{entrypointGoTmpl, map[string]*bintree{}},
	"preload.gs": &bintree{preloadGs, map[string]*bintree{}},
	"vm_file.go.tmpl": &bintree{vm_fileGoTmpl, map[string]*bintree{}},
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

