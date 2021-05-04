// Code generated by go-bindata. DO NOT EDIT.
// sources:
// pkg/security/tests/schemas/container.json
// pkg/security/tests/schemas/container_context.json
// pkg/security/tests/schemas/container_path.json
// pkg/security/tests/schemas/datetime.json
// pkg/security/tests/schemas/event.json
// pkg/security/tests/schemas/exec.schema.json
// pkg/security/tests/schemas/file.json
// pkg/security/tests/schemas/process.json
// pkg/security/tests/schemas/process_context.json
// pkg/security/tests/schemas/usr.json
// +build functionaltests

package schemas

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataContainerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\xae\xc2\x20\x10\x45\xf7\xfd\x8a\x09\x79\xcb\xd7\x52\x59\xf2\x2b\xc6\x05\xd2\xb1\xa5\x89\x80\xc3\xb8\x30\x86\x7f\x37\x50\xaa\xe9\x8e\x7b\x2f\xe7\x64\xde\x1d\x00\x80\xf8\x4b\x76\xc1\xbb\x11\x1a\xc4\xc2\x1c\x93\x96\x72\x4d\xc1\xf7\x5b\x3d\x04\x9a\xe5\x44\xe6\xc6\x52\x8d\x6a\xec\x4f\x4a\xb6\xff\xff\x0d\x77\x53\x41\x6d\xf0\x6c\x9c\x47\x1a\x0a\xbc\x8f\xfc\x8a\x58\xd6\x70\x5d\xd1\xf2\xde\x46\x0a\x11\x89\x1d\x26\xa1\x61\xbb\xa2\xf6\xd5\xf4\xcb\x07\x43\x62\x72\x7e\x16\xdf\x31\xd7\x57\x6e\x46\xc2\xc7\xd3\x11\x16\xfe\x7c\xf0\xd5\x70\xe9\xf2\x27\x00\x00\xff\xff\x12\xe6\xde\xdb\xec\x00\x00\x00")

func bindataContainerJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataContainerJson,
		"/container.json",
	)
}

func bindataContainerJson() (*asset, error) {
	bytes, err := bindataContainerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/container.json",
		size:        236,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataContainercontextJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xd1\x0a\x82\x40\x10\x45\xdf\xfd\x8a\x61\xe9\x31\x1d\xf3\x71\x7f\x25\x22\x4c\xc7\x5c\x29\x77\x9b\x9d\xa0\x08\xff\x3d\xc6\x55\x21\xe8\x6d\x38\x9c\xb9\xf7\x7e\x32\x00\x00\xb3\x8b\x4d\x4f\xf7\xda\x58\x30\xbd\x48\x88\x16\x71\x88\x7e\xcc\x13\x2e\x3c\x5f\xb1\xe5\xba\x13\xac\xca\xaa\xcc\x0f\x15\x2e\xfe\x7e\x79\x77\xad\xbe\x36\x7e\x94\xda\x8d\xc4\x67\xbd\xe8\x25\x85\x86\xac\x92\xbc\x03\xa9\xe5\x2f\x03\x35\xb2\xd2\xc0\x3e\x10\x8b\xa3\x68\x2c\xa4\x35\x33\xdf\xb2\x7e\x70\x6a\x63\xea\x34\xa8\x73\x37\xb2\x88\xb8\xa9\xa9\x6e\x93\xa7\xf9\x9a\x96\x22\xa6\xc7\xd3\x31\xe9\xd0\xe3\xbf\x9a\x99\x9d\xb2\xe9\x1b\x00\x00\xff\xff\xb9\xfa\x42\x76\x12\x01\x00\x00")

func bindataContainercontextJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataContainercontextJson,
		"/container_context.json",
	)
}

func bindataContainercontextJson() (*asset, error) {
	bytes, err := bindataContainercontextJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/container_context.json",
		size:        274,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataContainerpathJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8c\x41\xae\xc2\x30\x0c\x05\xf7\x3d\x85\x15\xfd\x55\xf5\x1b\x97\x2c\x7b\x19\x64\x25\xa6\x09\x08\x27\x72\xbc\x41\x88\xbb\xa3\x88\x76\x3b\x6f\xe6\xbd\x27\x00\x00\xf7\xd7\x63\xe6\x27\xb9\x0d\x5c\x36\x6b\x7d\x43\xbc\xf7\x2a\xcb\x0f\xfb\xaa\x3b\x26\xa5\x9b\x61\x58\xc3\xba\x5c\x02\x1e\xfe\xff\x91\x97\x34\xd2\x58\xc5\xa8\x08\xeb\xb5\x91\x65\x3f\x1e\x4e\xc3\x5e\x8d\x87\xd2\x4d\x8b\xec\x27\x6d\x64\xc6\x2a\x63\xf0\x73\xaa\xf1\xc1\xea\x67\x37\x7d\xbe\x01\x00\x00\xff\xff\x66\xad\xfc\xe4\x96\x00\x00\x00")

func bindataContainerpathJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataContainerpathJson,
		"/container_path.json",
	)
}

func bindataContainerpathJson() (*asset, error) {
	bytes, err := bindataContainerpathJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/container_path.json",
		size:        150,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataDatetimeJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\x4d\xca\x83\x30\x10\x86\xf7\x9e\x22\x0c\xdf\xd2\x98\x31\x7c\x9b\xe6\x12\x3d\x80\x58\x08\x35\xfe\x14\x8d\x92\xcc\xa6\x48\xee\x5e\xa2\xd6\x62\x4b\x69\xc8\xe2\xcd\xcc\xf3\x90\x77\x4e\x18\x63\x0c\xfe\xfc\xb5\x35\x83\x06\xc5\xa0\x25\x9a\xbc\x12\xe2\xe6\x47\xcb\xd7\x71\x36\xba\x46\x54\x4e\xd7\x24\x24\x4a\xe4\xb9\x14\x1b\x9f\x6e\x7a\x57\x45\xb5\xd2\x64\xa8\x1b\x4c\x16\xdd\xe7\x4e\xf7\xfd\xb9\x06\xc5\x8a\xe5\x19\xcf\xbc\xa7\x05\xa0\xfb\x64\xa2\xed\xc9\x75\xb6\xd9\xb4\x7d\x3b\x69\x22\xe3\x6c\x04\x2e\x05\xf2\x53\x39\xff\x07\xbe\x06\xf9\x0a\xb0\x4b\x21\xfd\xf6\x8f\x1d\x09\xd4\xdb\xf0\x77\x81\x8f\x12\x88\x98\xf3\xe5\xc2\x01\x0c\xc9\x31\x95\x49\x78\x04\x00\x00\xff\xff\x99\x6a\xef\xa9\x5d\x01\x00\x00")

func bindataDatetimeJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataDatetimeJson,
		"/datetime.json",
	)
}

func bindataDatetimeJson() (*asset, error) {
	bytes, err := bindataDatetimeJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/datetime.json",
		size:        349,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataEventJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\xc1\x6e\x85\x20\x10\xbc\xf3\x15\x84\xf4\xf8\xde\xe3\xd5\xa3\xbf\xd2\xf4\x40\x71\xab\x98\x08\x74\x59\x4d\x4c\xe3\xbf\x37\x88\x9a\x58\xc5\x4b\x7b\x9d\xd9\x99\x9d\x81\xfd\x66\x9c\x73\x2e\x5e\x82\x6e\xa0\x53\xa2\xe4\xa2\x21\xf2\xa1\x94\xb2\x0d\xce\xde\x13\xfc\x70\x58\xcb\x0a\xd5\x27\xc9\xe2\x59\x3c\xef\xaf\x85\x5c\xe6\x6f\x8b\xdc\x54\x51\x0a\x03\x58\x7a\x44\xe1\x4a\xd0\xe8\x21\x32\xee\xa3\x05\x4d\x2b\xea\xd1\x79\x40\x32\x10\x44\xc9\x53\x82\x19\x87\x81\x76\x40\xde\x62\x63\x33\x56\x1b\x6f\x55\x07\xa7\xcc\xce\x3b\x10\x1a\x5b\x8b\xc3\xd0\x74\x3b\x3a\x6a\x45\x50\x3b\x1c\xff\xd7\xd5\xf5\xa4\xdd\x9f\xa2\xb2\x8b\x15\x02\xe1\xab\x37\x08\xf1\x97\xde\x32\x8f\x74\x55\xf5\x22\xf0\x8e\x79\x67\x27\x01\x44\xa5\xe8\x58\x2c\x5b\x28\x15\x59\xf4\xe7\xc1\xe7\x43\xf9\xbd\x80\xa5\x00\xd3\x4f\x00\x00\x00\xff\xff\x80\x89\xa3\x61\xd2\x02\x00\x00")

func bindataEventJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataEventJson,
		"/event.json",
	)
}

func bindataEventJson() (*asset, error) {
	bytes, err := bindataEventJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/event.json",
		size:        722,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataExecSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x54\xc1\x6e\x83\x30\x0c\xbd\xf7\x2b\xa2\xa8\xc7\xb6\xe9\x38\xf2\x13\xfb\x80\x09\x4d\x2c\x98\x91\x8a\x25\x99\xe3\x4e\xad\x2a\xfe\x7d\x0a\x94\x56\x90\x20\xaa\x09\x69\x87\x2d\x17\x82\xfd\x9e\xf3\xf0\x8b\xb9\xac\x18\x63\x8c\xaf\x9d\xac\xe0\x23\xe7\x29\xe3\x15\x91\x75\xa9\x10\x07\x67\xf4\xb6\x0b\xef\x0c\xbe\x8b\x02\xf3\x92\x44\xb2\x4f\xf6\xdb\xa7\x44\x5c\xf1\x9b\x2b\x5d\x15\x9e\x0a\x27\x90\x3b\xcf\xeb\xe3\x74\xb6\xe0\x13\xe6\xed\x00\x92\xfa\x68\xae\xcf\xcf\x25\x4f\xd9\x4b\xfb\xea\xd7\xe5\xb6\xeb\x00\x75\x3d\x02\xc4\x81\x37\xc2\x1a\xc1\xe3\x79\xa9\x6a\x48\x85\x10\xf0\x05\x9a\x3a\x25\x01\xa1\xd9\xfc\xb4\xaa\x7f\x2e\x5e\xf4\xe8\x70\xf1\x9a\x16\x8d\x04\xe7\x5e\xa5\xd1\x04\xa7\xe5\x1b\xe1\xeb\xe6\x4a\x03\xce\x9d\x30\x88\x64\xab\xc8\xc9\xff\xce\xff\x86\xf3\x53\xbe\xb4\xbb\xac\x1f\xd3\xc0\x8b\x91\x59\xca\x67\x43\x61\xdc\xa2\xb1\x80\xa4\xc0\x45\xf3\x2d\xc6\x0b\x9e\xcc\xb2\xc9\x7f\x47\x14\x89\xf0\x79\x54\x08\x45\xf4\xde\x0c\x90\xf7\x8b\x6b\x73\xaa\xc2\xc6\xf4\x2b\x8b\x66\x9a\x99\x46\x8e\x6c\xe3\x54\x81\x8e\x77\x68\x56\xf1\x5d\x69\x28\x72\x28\xae\x79\x60\xaa\xa6\x8c\x5a\x52\xc6\xc3\xdf\xfe\x67\x6f\xc7\x78\xcc\x56\xcd\x77\x00\x00\x00\xff\xff\x43\x27\x19\xe1\x7d\x07\x00\x00")

func bindataExecSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataExecSchemaJson,
		"/exec.schema.json",
	)
}

func bindataExecSchemaJson() (*asset, error) {
	bytes, err := bindataExecSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/exec.schema.json",
		size:        1917,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataFileJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\xc1\x6e\x83\x30\x0c\x86\xef\x3c\x45\x14\xf5\xd8\x36\x1d\x47\x5e\x65\x9a\x50\x06\x06\x5c\x8d\x84\x39\xe6\x50\x4d\xbc\xfb\x94\x52\x3a\x31\x16\xc6\xb4\xf4\x88\x9d\xff\xb7\xb1\xbf\xe4\x23\x11\x42\x08\xb9\x73\x45\x03\xad\x96\x99\x90\x0d\x73\xe7\x32\xa5\xce\xce\x9a\xc3\x18\x3e\x5a\xaa\x55\x49\xba\x62\x95\x9e\xd2\xd3\xe1\x29\x55\xb7\xf3\xfb\x9b\x1c\x4b\x2f\xad\xf0\x0d\x8e\x5e\x37\xc5\xf9\xd2\x81\x4f\xd8\xd7\x33\x14\x3c\x45\x3b\xb2\x1d\x10\x23\x38\x99\x89\xb1\x81\x6b\xdc\xcb\x67\x91\xb0\xc7\x3d\x1b\xf0\xfa\xca\x6b\x6e\x7e\xcc\xcc\xbc\x1d\x13\x9a\x5a\x2e\x0e\x0d\xfb\xa5\xa3\xd1\xed\xb2\xcb\x7f\x39\x16\xd6\xb0\x46\x03\x94\xaf\x77\xbb\x23\xa8\xa6\x31\x67\x4a\xa9\xb9\x6e\x1c\xfc\xa6\x82\x68\x6c\xb9\xe1\x1f\xd0\x30\xd4\x40\xdb\x3c\xdb\x47\x58\xf6\x86\xf3\x2b\x5a\x31\x6d\xfd\xf8\xdc\xc5\x31\xb4\x71\xd7\xd8\x3b\xa0\xb8\x8e\x35\xd9\xbe\x8b\x6b\xd9\xda\x12\x2b\x2c\x34\xa3\x35\x39\xe3\x1a\xca\xdf\x71\x2b\x35\x83\x57\xfc\x01\xb4\xa2\xd1\xa6\x86\xc8\x75\x92\x95\xaa\x92\xe0\xbd\x47\x02\xcf\xcc\x73\xe0\x35\x08\xdd\xe9\xe0\x45\x09\xd1\xbe\x86\xec\x3a\x77\x21\x7a\x82\x0c\x6c\xda\xe4\x2f\x0b\x98\x65\x5f\xee\x5f\xe3\x3c\x87\x64\xf8\x0c\x00\x00\xff\xff\x6a\x8d\x18\xb0\x0b\x06\x00\x00")

func bindataFileJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataFileJson,
		"/file.json",
	)
}

func bindataFileJson() (*asset, error) {
	bytes, err := bindataFileJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/file.json",
		size:        1547,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataProcessJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x4b\x6e\xdb\x30\x10\xdd\xfb\x14\x82\x90\x65\x12\x25\x5e\xfa\x06\x5d\xf5\x00\x86\x61\x30\xd4\xd0\x66\x1a\x7e\x42\x8e\x8a\x1a\x81\xef\x5e\x48\x76\x6c\x4b\xe4\xe8\x6f\xb4\x5b\x91\xf3\xe6\xfb\x86\xa3\xf9\x5a\x24\x49\x92\xa4\x0f\x9e\xef\x41\xb1\x74\x95\xa4\x7b\x44\xeb\x57\x59\xf6\xee\x8d\x7e\x3a\x7d\x7e\x36\x6e\x97\xe5\x8e\x09\xcc\x96\x2f\xcb\x97\xa7\xd7\x65\x76\xbe\xff\x78\x16\x97\x79\x29\x6a\x9d\xe1\xe0\xfd\x73\x29\xfa\x7d\x84\x07\x0b\xe5\x99\x79\x7b\x07\x8e\xdf\x5f\xad\x33\x16\x1c\x4a\xf0\xe9\x2a\x39\xd9\x70\xba\x5d\x01\x5d\x3f\xd4\x20\xa4\x46\xd8\x81\x4b\x2f\xa7\xc7\xc7\xab\x64\x31\x5a\x72\x37\x5e\xa7\x07\x47\x8b\x7a\x74\x52\xef\x08\x9d\xce\x14\x76\x9c\x28\xfc\x01\x5e\x20\x7b\xfb\x80\x2d\x37\x1a\x99\xd4\xe0\xb6\x96\xe1\x3e\x84\x7b\x70\x20\x4a\x38\x21\x3f\x60\x95\x65\x59\xfd\xfe\x29\x4d\x5d\x3a\xe2\xc8\x7d\x0c\xe5\x46\xa9\xc9\x2e\x4a\x6d\x72\x18\x99\x9e\x1b\x18\x65\x0a\x8d\xdb\xd1\x89\x46\x3c\x8c\xf3\x44\x18\xf7\x6b\x8b\x52\x45\x5c\x68\x26\x27\x67\x08\xe5\xcd\x8e\xb4\xcc\x87\xc6\x1d\xe4\xa0\x51\xb2\x0f\x4f\x3b\x57\xa3\xed\xe5\x94\xa0\xef\xe5\x3c\x46\xc6\xee\x88\x47\x8c\xbc\x22\xc6\xa8\x16\x40\x36\x53\xd1\x86\x18\xa3\xfd\x34\x1b\xe3\xa4\x9e\x64\x24\xcc\x1f\x49\x98\x3f\x94\x30\x7f\x2c\xe1\x0e\xc1\x14\x7e\xfe\x68\x0a\x3f\x7f\x38\x85\x9f\x3f\x9e\xc2\xdf\x21\xa0\x9c\xd9\x2d\x08\x01\x1c\xe5\xef\xb0\x27\x05\xd0\xcc\x39\x76\x48\x43\x9c\xea\x92\x44\x50\xf1\x6e\xd2\xdb\x44\xc2\xcc\x4a\xb6\xd0\xf2\xb3\x80\x1f\x67\x25\xe8\x0a\xe8\xed\xa1\x05\xa7\x24\x22\xf4\x48\xc9\xff\xee\xe1\xa2\x05\x28\x75\xf0\x59\x48\x57\xb9\xb9\x8e\xb7\x74\xa2\x8b\x52\x7d\x8b\x6c\x3e\x54\x0b\xa1\xfb\x00\xc9\x66\x9a\x92\x24\xb1\x5a\xd8\xd1\x59\xe3\x9d\x25\x52\x3b\xdf\x74\x8c\x25\x73\x3f\xbb\xd1\x59\x2d\xc0\x1e\xc2\x70\xcd\x22\xc3\xc6\x24\xc4\x8e\x99\xf5\x72\x6f\xd0\xec\xda\xa6\x30\x3e\x41\x06\x3e\x0c\xea\xa6\xea\x1e\x90\xc4\x84\x3a\x0d\xb6\x0c\x9f\x3f\x78\x84\x70\x14\x0f\x80\x87\xa4\x71\xfe\x97\x74\xfe\x37\x4f\x99\x5c\x0a\xc9\x19\x4a\xa3\xe3\x63\x73\x80\x3f\xa8\x92\xf7\x4c\xef\x60\x2a\xf0\x84\x96\x5c\x11\x88\x22\x2d\xc9\x04\xaa\x9c\xdb\x6a\xb2\xbd\xb0\x88\xf2\xe8\xff\x5a\x84\x89\xea\x08\x77\x8f\x2e\x7b\xe9\x17\x03\x7e\x8b\x9b\x3f\x4b\x8b\x1b\xd0\xd4\x68\xf8\x29\x6a\x89\xf8\x1a\xd8\x9d\x87\x14\x38\xf1\xc8\x83\x2e\x4a\x22\xaf\x5f\x37\x13\x2a\xa9\xb2\x24\x1a\xb5\x7f\xe0\x91\x92\x5a\xaa\xca\xa9\x65\x2f\xd6\xd9\x41\x4a\xa7\x46\xe9\xf1\xac\x70\xd3\x28\x8a\xcd\xb9\x28\xa2\x04\xad\xb6\x58\x8d\xd5\x54\x63\xdf\xd4\xb2\x6b\x69\x6e\x50\xda\xf6\x22\x1d\xcb\x0e\xe2\x67\x3f\x3e\x8b\x44\xb7\x16\xd1\xe5\x43\x2d\x02\x39\x58\xd0\x39\x68\x1e\xec\xf2\xda\x16\x55\xeb\x1b\x7e\x6e\x08\xd6\xae\x5b\x10\x4e\xf9\x38\x2e\x8e\x7f\x03\x00\x00\xff\xff\xb7\x7e\x45\x3e\xc0\x14\x00\x00")

func bindataProcessJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataProcessJson,
		"/process.json",
	)
}

func bindataProcessJson() (*asset, error) {
	bytes, err := bindataProcessJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/process.json",
		size:        5312,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataProcesscontextJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\x41\x6e\xc3\x20\x10\x45\xf7\x3e\x05\x42\x5d\x26\x21\xf5\x92\x4b\xf4\x00\x55\x55\x51\x3c\x6e\x88\x1c\xa0\xc3\x54\x6a\x14\xe5\xee\x15\x06\x5b\x8d\x8c\x9d\x34\xf1\xf2\xf3\xff\xd7\x1b\x18\x9f\x2a\xc6\x18\xe3\x4f\x41\xef\xe0\xa0\xb8\x64\x7c\x47\xe4\x83\x14\x62\x1f\x9c\x5d\x27\x79\xe3\xf0\x53\x34\xa8\x5a\x12\xf5\xb6\xde\xae\x9f\x6b\x91\xfd\xab\x1c\x37\x4d\x8c\x7a\x74\x1a\x42\x78\xd7\xce\x12\xfc\xd0\x26\x56\x0c\x16\x3a\x7a\x88\x1e\xf7\xb1\x07\x4d\x83\xea\xd1\x79\x40\x32\x10\xb8\x64\x89\x65\xd0\x63\xd3\x85\xd8\x1f\xa8\xae\x7b\x69\xb9\x64\xaf\x17\x72\xfc\x4e\x13\x25\xa1\x21\x44\x3f\x6f\x4d\x07\x52\x08\x91\x9b\x13\xdb\x24\x72\x5e\xdd\xda\x3b\x43\x3e\xf5\x29\x04\x4b\x8b\x9e\x3b\x38\x17\x78\xc7\x46\x65\x35\x04\x72\xb8\x0c\xc8\xfe\x3e\x8e\x42\x54\x47\x3e\xdf\xd9\x9b\x0d\xc1\xe1\x7a\xe7\x03\x43\xf5\x83\xdd\x35\x72\x5c\x3c\x65\x2c\xe0\xff\xef\x7b\x8c\x5e\xbb\xf1\xe2\xc9\x0c\x14\x47\xf8\xfa\x36\x08\x4d\x71\x63\x47\x57\xde\x91\x9b\xde\xb2\xe8\x79\x9b\x2e\x72\x55\x3e\x4f\x7a\xc6\x2d\xe3\x8d\x3f\x5f\x95\xa2\xe7\xdf\x00\x00\x00\xff\xff\x69\x16\x74\x58\x24\x04\x00\x00")

func bindataProcesscontextJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataProcesscontextJson,
		"/process_context.json",
	)
}

func bindataProcesscontextJson() (*asset, error) {
	bytes, err := bindataProcesscontextJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/process_context.json",
		size:        1060,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataUsrJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8f\x4d\x0e\x82\x30\x10\x85\xf7\x3d\x45\xd3\xb8\x04\x8a\x2c\xb9\x8a\x71\x81\x30\x40\x89\xd2\x3a\x33\x2c\x8c\xe1\xee\xa6\xfc\x25\x0d\xe2\xc6\xed\xf7\xfa\xbe\x79\x7d\x0b\x29\xa5\x54\x27\x2a\x5b\x78\x14\x2a\x97\xaa\x65\x76\x94\x6b\xdd\x91\xed\xe3\x19\x27\x16\x1b\x5d\x61\x51\xb3\xce\xd2\x2c\x8d\xcf\x99\x5e\xde\x47\x4b\xdd\x54\xbe\x5a\x9b\x3b\x24\xbe\xb7\x72\x7e\x39\xf0\x81\xbd\x75\x50\xf2\x4a\x1d\x5a\x07\xc8\x06\x48\xe5\x72\x1e\x30\xf1\x81\x30\x00\xc7\x8a\x2d\x3d\x50\x6d\xf9\xb4\x6b\xcf\x03\x33\x31\x9a\xbe\x51\xbb\x47\x63\xb4\xf7\x35\x68\x07\xf7\x87\x52\xfc\x38\xa0\x10\x9e\x83\x41\xf0\x93\x2f\x5f\xbf\x72\x38\x28\xe0\x57\x11\x9e\x1b\xc5\xf8\x09\x00\x00\xff\xff\xd2\xf0\x69\x81\xe6\x01\x00\x00")

func bindataUsrJsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataUsrJson,
		"/usr.json",
	)
}

func bindataUsrJson() (*asset, error) {
	bytes, err := bindataUsrJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/usr.json",
		size:        486,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"/container.json":         bindataContainerJson,
	"/container_context.json": bindataContainercontextJson,
	"/container_path.json":    bindataContainerpathJson,
	"/datetime.json":          bindataDatetimeJson,
	"/event.json":             bindataEventJson,
	"/exec.schema.json":       bindataExecSchemaJson,
	"/file.json":              bindataFileJson,
	"/process.json":           bindataProcessJson,
	"/process_context.json":   bindataProcesscontextJson,
	"/usr.json":               bindataUsrJson,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"": {Func: nil, Children: map[string]*bintree{
		"container.json":         {Func: bindataContainerJson, Children: map[string]*bintree{}},
		"container_context.json": {Func: bindataContainercontextJson, Children: map[string]*bintree{}},
		"container_path.json":    {Func: bindataContainerpathJson, Children: map[string]*bintree{}},
		"datetime.json":          {Func: bindataDatetimeJson, Children: map[string]*bintree{}},
		"event.json":             {Func: bindataEventJson, Children: map[string]*bintree{}},
		"exec.schema.json":       {Func: bindataExecSchemaJson, Children: map[string]*bintree{}},
		"file.json":              {Func: bindataFileJson, Children: map[string]*bintree{}},
		"process.json":           {Func: bindataProcessJson, Children: map[string]*bintree{}},
		"process_context.json":   {Func: bindataProcesscontextJson, Children: map[string]*bintree{}},
		"usr.json":               {Func: bindataUsrJson, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
