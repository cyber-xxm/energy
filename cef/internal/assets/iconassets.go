// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package assets generated by go-bindata.
// sources:
// assets/icon.ico
// assets/icon.png
package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

var __assetsFile__ = &assetOperator{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

func (f *assetOperator) ReadFile(name string) ([]byte, error) {
	return Asset(name)
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() *assetOperator {
	return __assetsFile__
}

var _assetsIconIco = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x0b\x50\x54\x57\x9a\x6e\x59\xcb\xda\xb2\x28\xc7\xb5\x2c\xca\x4d\xa5\x12\xe7\xde\xdb\x0f\xfa\xc5\x43\x04\xd4\xe6\x21\x1a\xc1\x47\x1c\x35\x89\xa1\x5c\xd7\x50\x8e\x93\x90\x59\xcb\xc9\x24\xd6\x48\x94\x73\xbb\x11\x5b\x5e\x12\x45\xa4\x1b\x04\xed\x38\x0e\x45\x1c\x42\xf7\x3d\x34\xcd\x43\x10\x54\x50\xc4\x00\x2d\x82\x36\x84\x31\x60\x5c\xd2\x21\xae\xeb\x3a\x8e\xe9\x41\x42\xee\xd6\xb9\x8f\xe6\x76\x03\x92\x8d\xa9\xe8\xd6\xf4\x57\x75\xeb\xdc\xf3\x9f\xff\xfc\xf7\x7c\xff\xf9\xef\x39\xff\x3d\xdd\x22\xd1\x0c\xd1\x0c\x51\x6c\x2c\x2a\x17\x8a\x64\x71\x22\x51\x80\x48\x24\x92\x89\x44\xa2\x58\x91\x48\xa4\x13\xb1\x72\x06\xb1\x22\x1f\x7c\xf0\xc1\x07\x1f\x7c\xf0\xc1\x07\x1f\x7c\xf0\xc1\x87\x7f\x78\x98\x81\x54\x43\x91\x44\xd9\x94\x17\x10\x93\x48\x8f\x02\xe2\xd4\x71\x39\x2b\xe3\x51\x06\xc2\xfc\x2c\x40\xbc\x1d\xdd\x57\x00\x39\x31\x8d\x3d\x23\xd2\x2b\x07\x4a\x7f\xa1\xdc\x02\x24\x29\xbc\x3d\x0a\x88\x93\x78\x79\x76\xea\xaf\xfd\x78\xb9\x05\x48\x24\x90\xc4\xcb\x20\x89\xf7\x41\x12\x2f\x37\x93\xb2\x17\x26\xe3\x44\x91\xe2\x1c\x2f\xdb\xdb\xa7\xe2\xdf\x79\x52\xb5\xad\x39\x2f\x90\xae\x24\x31\xfa\xc2\x21\x09\x6d\x37\x29\x69\xbb\x49\x45\x5f\x3c\xc2\xca\x20\xc0\xee\x20\x3d\xbb\x49\xd9\xd4\x98\x23\xa3\x1b\xb3\xc5\xa8\xbd\x49\x68\xc3\x02\xc4\x71\x90\xc4\xbb\xd0\x7d\x5b\x49\x68\x64\xab\x41\xc1\xf4\x6d\xc8\x14\x73\xf6\x94\xf4\xe5\x02\x05\x6f\x6f\xec\x0c\x08\xf2\x6f\x2d\x59\x34\xaf\xbd\x58\xc9\xc8\x6a\xd3\x71\xba\xe3\xa4\xaa\x86\xb7\x07\x01\xde\x86\xe4\x9d\xc5\x72\xba\x30\x63\xfb\x2c\x8e\x3b\x01\x01\x76\xb7\x31\x53\x4c\xf7\x96\xa9\x87\xaf\x14\xa8\xc6\x20\x89\x39\x28\x40\x64\x4c\xe0\x64\x52\x77\x35\x1d\x92\x31\xb6\x5b\x8e\x48\x91\x6d\xe3\x54\xfc\xff\xb3\x79\xcd\x3f\xd7\xe7\xa8\x7f\xcb\xf0\xcf\x56\xa5\x0c\x5f\x8a\x0f\xf8\xa6\x35\x21\x60\xa4\x63\x5d\x6c\x7b\xe1\xe2\x3e\x08\x30\xfa\x0c\x08\x9e\xe3\xbc\xb4\x7a\x6e\x6d\x86\x3c\xf7\x5c\x86\x3c\xd7\xd9\x12\x3f\x57\x68\x03\x02\xbc\x18\x02\xcc\x8e\xee\xbf\xbc\xb0\x76\x66\x8b\x31\x62\x15\xb2\x77\x56\x2f\x2b\x46\xf6\x86\x2f\x27\x04\xb8\x3a\xd6\x84\x39\x4e\x2f\x6b\x46\x72\x0b\x90\x44\x7e\xd1\xb8\xce\x6f\xb0\xfe\x35\x0c\xd5\x6b\xf6\x4b\xea\x9d\xcd\xac\xcd\x72\xa0\x9c\x8b\x7c\xd4\x90\x2e\xbb\x7d\xaf\x29\xee\x85\xa2\xa3\xbb\xf8\x67\x9c\xaa\xd6\xe1\xf4\xd0\xe9\xc5\xc9\x15\x45\xdb\x45\xdf\xd4\xbd\x16\x6e\xdb\x2f\x19\x83\x00\x1f\xf2\xe6\xf4\xf5\xa5\xf8\x79\xf5\x39\xea\x2c\x64\xfb\xe2\x47\xca\x2d\xce\x96\x78\xff\xa9\xf8\x73\xf3\xb7\x85\x99\x1b\x6d\x60\xa2\x50\x3e\x64\xdb\x2c\xaf\xde\x2f\x1d\xb5\x00\x49\x38\x37\x86\x54\xab\x56\xfc\x9e\x50\xe7\x0c\x08\x9a\x05\x01\x76\x1f\x02\xbc\x83\x97\x99\x81\x4c\x89\xec\x51\x00\x4b\x13\xea\xde\x6b\xfc\xf7\x79\x0d\x59\xaa\xbb\x14\x20\x76\xa0\x3a\xf9\x87\x0f\xfc\x18\x3d\x12\x3f\xc3\xeb\x50\x80\xd8\xc0\xc4\xce\x01\x85\x7b\xce\x2a\x80\x3c\x00\x02\x6c\xb4\x51\x1f\xd8\x2a\xb4\x67\xd5\x49\xbb\x20\xc0\x87\x27\xe3\x54\xa9\x93\xee\xe4\x7c\xbd\xf2\x49\xdc\x39\x5e\xdb\x90\x6e\xa5\x56\xf6\xba\x77\xdb\xd9\xac\xa0\x1a\x0a\x10\x49\xac\x1e\x96\x56\xa9\x95\xec\x14\xb6\x53\x40\xbc\x81\x8d\x6b\xfc\x33\x81\x3f\xd5\xcc\xb3\x53\x7f\x99\xe6\x6d\xaf\xe9\xa3\xc5\x39\x10\xe0\xf9\xe8\x7e\xdf\x6e\x8e\x3f\xc0\xcb\x05\x63\xc9\x47\xb2\xcb\xd9\x8b\x36\x08\x7c\x92\x88\x64\x8d\x7a\x95\x87\xef\x21\x49\x34\x43\x80\xdd\x9b\x94\x93\x56\x96\xcc\xf1\x5f\x33\x1d\x7f\x34\x1f\x48\xd7\xa6\x93\xad\xf7\x6e\xbb\x62\x08\xdf\x79\xf3\x84\x3a\x85\xd5\xc3\xd3\x21\x29\xf9\xad\xc7\x73\x00\x5e\x5e\xbd\x1f\xa7\x2b\xb5\x44\xdb\xb8\x3d\x9c\xe3\x8f\x91\xde\xf6\x5a\x8b\x5f\xd1\x74\x1d\x57\x32\x7c\x33\xf7\xec\x11\x4d\xc2\xff\x86\x4d\x27\x1e\xbd\x55\xf4\xea\x5c\x81\xac\x18\xe9\x9d\xcf\x0c\x8f\x14\xda\xaa\xd4\x12\x17\x50\xec\x4d\xca\x5f\x27\xdd\xc1\xcd\xcb\x04\x4e\x13\xf9\x8b\x19\x5f\x55\xea\xe4\x13\x7c\x75\xbf\x76\xb9\xff\x80\x39\x9e\x59\x83\x29\x92\xc8\x80\xa4\xe4\x6d\xbe\xad\x02\xc8\xe7\x40\x80\xb9\x6c\x3a\xe9\x28\x05\x88\xd6\xf1\xf1\x62\xc1\x2c\x7f\x3c\xd5\xdb\xde\x80\x75\xe3\xcc\xbf\xd6\x69\x66\xa3\xfb\xbd\x1c\x7f\xc8\xf1\xaf\x00\xf2\x05\x68\xbd\xa9\xdb\xef\x19\xe7\x10\xe0\x97\x2a\x49\x9c\x6e\xcf\x7f\x63\xb6\x87\x5c\x4b\x9c\x83\x00\x7f\x38\x19\x27\x6b\x9a\x34\x89\xf5\x2d\xb1\x69\x3a\xfe\x10\xe0\xbb\xd8\xf7\x5f\xbe\x4a\xe0\x93\x30\x0b\x19\xb8\xd0\xd3\x4f\x44\x0e\x24\xc5\x3b\x04\x75\xe6\xbd\xa9\x3f\x10\xd4\x44\x01\xe2\x92\x40\x1e\x8a\xe4\x66\x01\x7f\x33\x90\xbe\x64\x06\xd2\x70\xa1\xbd\x94\x14\x4f\xfe\x14\x20\xb6\x70\x71\x9e\xe6\x35\xbe\xbb\x95\x5a\xe2\xee\x84\x71\x93\x44\x13\x04\xd8\xc8\x29\x10\xed\x37\xa1\x4d\x27\x63\x6c\x59\x48\x71\xe2\x74\xfc\x29\x80\xbf\xcf\xc6\xbf\x3c\x56\x20\x2b\xa8\xd4\xca\xb6\x7a\x8d\x23\xd7\xaa\x15\x27\x09\xea\x75\x6c\x5c\x2e\x49\x86\x00\x6f\xe6\xe5\x16\x20\xe1\xf8\x13\x7b\xc7\x65\xe2\xcd\x14\x20\x26\xec\x43\x1c\xff\x0a\xce\xde\x09\x54\x6f\xff\x68\xa9\x86\x6f\xcf\x4c\x4d\xf6\x43\xfb\x01\xbf\xbf\x22\x9c\xd1\x2e\xf2\xa7\x00\x5e\x81\xb8\x73\xfd\xcf\x99\x81\x14\xf3\xe0\x44\x8a\x77\x70\xf3\xbf\x75\x5a\xfe\x24\x91\x82\x74\xab\x74\x4a\x8d\x80\x9b\x15\x6a\x65\x89\x9e\xfc\xb1\x3c\xab\x56\xb2\x4d\xc4\xc4\x6a\x60\x00\x24\xb1\xd1\xb3\xe9\x8a\x21\x8b\x56\xa1\x84\x24\x7e\x81\xd7\x33\x03\x69\x98\x37\x7f\x0a\x10\xef\x51\x80\x28\xf0\x7e\x36\x04\xbf\x44\x39\x41\x05\x5a\x0b\x20\xc0\x6f\xd7\xa4\x49\x1f\xf6\x1d\x7f\x73\x16\xdf\xfe\x27\x10\x39\x9b\x5b\x23\xdc\xf1\xd5\x52\xb2\x6c\x9e\xfd\xa4\x92\x16\x5e\x6d\x25\x21\x61\x9e\x76\x59\x5f\xa2\x35\x6b\x5a\xfe\x00\x4f\x65\xfc\x48\x4a\x52\xd0\xfb\x42\x01\x62\x17\x24\xb1\x11\xe8\xb5\x1f\x50\x00\xcf\x87\x5a\xe9\x16\x8e\x0f\xb3\xbf\x9c\x3f\x18\x92\x67\x06\x12\x25\x04\x58\xd3\x38\x7f\x59\x38\xe7\xfb\x52\xc6\x1e\x49\x6c\x81\x00\xbb\xf5\x04\xfe\x43\x10\xe0\x90\x8d\x41\x49\x9f\x19\x48\x17\xf0\xed\x15\xe9\x11\xf3\xb9\x39\x76\xc7\xd7\xad\xc6\x75\x7e\xc3\x97\xe3\x17\x08\xaf\xa1\xe6\xd5\x33\xbd\xe6\xaa\xbf\xe6\x80\x0c\xf5\x6b\x9b\x8e\x3f\x04\x78\x3a\x3b\xff\x18\x6d\x4b\xc3\x99\xbc\x89\xad\x4b\x37\x78\xf0\x27\x09\x23\x45\x4a\x13\xb9\x3e\xad\x56\x12\xa7\xbb\xf2\x56\x44\xa2\xf5\x1e\x92\xf8\x38\x7f\x92\xe5\x6f\xd5\xb2\xf6\xac\x3a\xd6\xde\xe4\xfc\xb1\xb1\xfa\x4c\x31\x7d\xcd\xa4\xa2\xdb\x0a\x95\x4c\x1f\x48\x62\xb7\xcd\x64\x20\xe3\x83\xaa\xc2\xb7\xfc\xb9\xbc\xb1\x63\x3a\x1e\x3c\x3e\x05\xf2\x17\xd0\x3a\x5a\xa3\x0f\xad\xb7\xa6\x49\x47\x3f\x05\x8a\x79\x4f\xd2\x47\xeb\x3a\x7a\xc6\x50\x6d\x64\xc6\x77\xdd\x31\xbb\xfe\xa7\x79\x65\x46\x4d\xba\x74\xd4\xaa\x0b\x5c\xe7\xa1\x07\x88\x62\x8a\x94\xbc\x6e\x06\x32\x0c\xa2\x75\x2f\x5d\x31\x50\x95\xfb\x1b\xe4\x8b\x60\xf4\x0e\x0a\xf8\x47\x22\x7b\x6d\x45\x4a\x88\xec\x3d\xbe\x16\xb7\xbb\xf9\x48\xf0\x1d\xb4\xa6\x4c\xc6\xdf\xaa\xc3\x6d\x77\xaf\xc4\xcf\xff\xb6\x63\xb5\xac\xbf\x2c\xc6\xc6\xc5\xed\x69\xd4\x9e\x9b\xf9\x21\xda\xe7\x46\x21\xc0\x1d\x3f\x94\x3f\x45\xb2\xf9\x42\xdd\x81\x90\x9d\xb6\x83\x41\xdd\x14\x20\x26\xe4\x35\x9e\xbc\xf0\x5c\x66\xce\xd3\x42\xe5\xbc\xcc\x9a\x26\x3d\x07\xb5\xf2\x04\xaf\xb1\x9e\xb0\x6a\xc5\x9b\x28\x12\xdf\xcb\xe4\xcb\x19\x21\x59\x22\x6e\xbd\x83\x24\x51\xef\xe6\x0f\xa4\x4b\xd9\xdc\x83\xd8\xcd\xcb\xaa\xf5\xaa\xf7\xa7\xe2\x0f\x01\x66\xe1\xeb\xce\xfa\x24\xff\xba\x83\x8a\xbb\x28\xdf\xfb\x33\x50\xfb\xb3\x63\x91\x0d\x43\x80\x4d\x9a\xe7\x4d\x06\x88\xd6\x6e\xf4\x6e\x66\xc7\x84\xd5\x64\x2d\xce\x47\xf9\xc3\x93\xf5\xb1\x3c\x8e\x3f\x21\xb0\x51\x01\xb5\x81\x2b\xbd\xf4\x4e\x55\x69\x25\xeb\x51\x8e\x52\xa5\xc5\xe9\x9e\xfc\x95\xc1\x1c\xdf\x30\x8a\x24\xea\x78\x3d\x0b\x90\x68\xbc\xf9\x53\xa4\x24\x91\xe2\xf2\xbe\x49\xf8\x43\xa1\xcc\x76\x40\x51\xc6\xe5\x6e\x1a\xae\x8e\xf2\x1c\xba\x0c\x2c\x9a\xfd\x03\xf9\x77\x57\xe9\x24\x8f\xda\x8d\xef\xce\x6c\xc8\x5f\xb1\x09\x02\x7c\x20\x77\xff\xef\xa7\xd4\x47\xef\x25\x1b\x2f\x61\xee\xfd\xde\x65\x8f\xd9\xf4\xa8\x7d\xb9\xd7\xfe\x8f\x97\x52\xa4\x74\x2f\x97\x9f\x3b\x32\x33\xf7\x30\x72\x14\xef\x14\x20\x6a\xc6\xf9\x13\x1a\x2e\x86\xdf\xe7\x65\x0f\xdb\x63\xe5\xdf\x76\xc6\xac\xfa\x21\xfc\x2b\x75\xd2\x0c\x61\xee\x52\x93\xa1\x2e\x60\xfd\x21\x55\x4f\x32\x76\x58\xb5\x3f\x38\x80\xaf\xa3\x77\x1d\xd9\xac\x4e\x0b\x64\xf6\xa3\x96\x4f\x92\xe7\x5b\xd3\xa4\x63\x15\x40\x26\x99\x9a\x3f\x9b\x5f\xda\x0e\x2c\x7b\x71\x42\x1b\x49\x9c\x6a\xce\x5e\x36\x9f\xbb\x2f\x83\xa4\xb8\x1b\xe9\x36\x67\x86\xb8\x73\x5b\x33\x29\x5d\x0a\x49\xdc\x26\xb0\x17\xed\xcd\x7f\xdc\x9e\x38\xa9\x4a\x17\xf8\xfa\x93\xf8\x43\x80\x67\xb0\xf9\x98\x74\x33\xc3\xa1\x30\x2e\x91\xf3\xc7\x2e\xa1\xde\x19\x10\x84\xf6\xc6\x91\xb3\xd9\xeb\xfc\xc7\x9f\x2d\x5e\xcf\x7c\x53\xef\x57\x64\xf1\xb2\xea\x83\x41\x5d\x68\xbf\x9a\x8a\x3f\x04\x98\x89\x79\x5e\x7a\x54\x80\x50\x6e\x01\x92\x30\x9b\x4e\x32\xd2\x63\x78\x73\x36\xc7\xab\xdc\xa6\x13\xd3\x36\x2d\x41\xf7\x16\xac\x94\x09\xf4\x34\x90\xc4\xad\x02\x7b\xb1\xdc\xf7\x9f\xc7\xf7\x4a\x29\x08\xf7\x43\xb1\xd9\xa0\x0f\xda\x26\xd0\x45\x6b\x9b\xd5\x8b\x7f\x05\x13\x8f\xe9\x6a\x26\x5e\x1c\xd6\xb7\xe7\xd4\x1c\x08\x74\x09\xd7\x58\x8e\xeb\x9a\x4a\x92\xb8\xef\x29\x63\xd7\xb2\x06\x7d\xa8\x7b\xef\xaa\xcd\x0e\x3b\xec\xfd\x0c\xaf\x3e\xa5\x4c\x1e\x9b\xb5\xc2\xbd\x4f\x7c\x02\x82\x67\xa2\x3d\xb7\x36\x2d\xb0\x43\x30\x56\x66\x5c\xe7\x0e\x28\xec\x9e\xfd\x89\x68\xe1\x1c\x52\x00\x8f\xe5\xbe\xff\xde\xf3\xd2\xdb\x8e\xe4\xad\xb9\x31\xf2\x71\x99\x27\x7f\x33\x90\xbe\x08\x49\xec\x51\xb5\x8e\x18\x73\x18\x12\xdc\x79\xc0\xe5\x82\x25\x65\x68\x0d\xb0\x90\xd2\x30\x81\x9f\x2c\x56\xad\xb8\xcb\xd3\x77\x58\x07\xda\x97\xdb\xf3\xc6\xfb\x36\x1c\x8d\x5d\x0f\x01\xf6\xb0\x1c\xa8\x26\x9c\x03\x50\x80\x58\x03\x01\xd6\xcd\xe6\x3f\xe2\xdd\x68\x8c\x14\x20\xde\xe6\x73\xdb\xb3\xe9\x0a\x23\xca\xcd\x90\x1c\x02\xac\x8f\xf9\x36\xcd\x0a\xf9\x83\xa0\x7f\x38\x93\x17\x91\x98\xc3\x02\xc4\x49\x28\x0f\xa5\xb8\xf5\x17\xe5\x34\x9c\x3d\x74\xe9\x51\xbe\x6a\xd3\x49\x1e\x5e\x39\x96\x34\xb3\x3c\x35\x70\x16\x05\x88\x64\xc4\x89\xcb\x8d\xb6\xa3\x3c\x14\xdd\xa3\xbe\xf6\x23\x6a\x9b\x70\x9c\x5f\x55\x6f\x96\x34\x64\xca\x47\x21\x89\xf5\x5b\x48\x49\x2c\x45\x12\x7b\x51\x5f\xab\x56\xec\xf6\x9d\x85\x14\xef\x46\xb2\x4a\x12\x77\xf1\x79\xaf\x59\xab\x98\x43\x91\x62\xf6\x7d\x22\xf1\x33\x66\x20\x95\x09\xed\xd6\xea\xc5\x17\x6a\xf5\x04\x3d\xd5\x75\x21\x2b\x74\x5b\x16\xf9\x1f\x7e\x7c\xbd\xfe\xa0\x84\xfe\xa2\x70\xe5\xc2\xf1\xfe\x44\x1e\xdf\x56\xa3\x17\x8f\xd8\xf4\xb2\xcd\x4f\xb2\xd7\x90\x21\x67\xf2\x58\xab\x5e\x3e\xc7\xbb\xad\xee\x20\x41\x37\x1f\x96\xd0\xfd\xa7\xd4\x76\x17\xd4\x4c\x38\xdb\xfb\xba\x6a\xf5\xce\x8b\x87\xd9\x73\xb4\x3a\x3d\x41\x5f\x37\x05\xb9\x6a\xd2\x03\xdd\xdf\x14\x35\x7a\x89\x6b\xdc\x9e\xf8\x01\x92\x55\x67\x2f\x26\x84\xcf\xa8\xd6\x4b\x3d\xbe\xdf\xc7\xba\xa3\x13\xc6\x7a\xa2\x93\xa6\xba\xfe\xd6\xba\x22\x20\x2f\xe7\x7d\x11\x5f\x1f\xb5\xc7\x6c\xf6\xea\x1f\xee\x6e\xbb\x1e\xb3\xed\xef\xd7\x96\x2f\x7c\x92\xbd\xc7\xd7\x62\x99\x3d\xed\x51\xc7\xf2\x59\x63\x3d\x51\x49\x63\xdd\xe8\xd2\x30\xd7\x77\x5d\x9a\xad\x23\x57\x97\x6a\xee\xd4\xbe\x32\xe1\x7b\x8e\xc7\x48\xc7\xda\xd8\xfb\x2d\x9a\x5c\xd7\xc5\x88\xa4\xe1\xb3\x1b\x67\x0f\x57\x6d\x88\xe6\xdb\x46\xaf\x47\x6f\x65\x6c\xb2\x76\xb7\x70\xcf\xf1\x77\xcb\x7a\xa2\x92\x46\xba\x62\x64\x53\xd9\xf6\xc1\x07\x1f\x7c\xf0\xe1\xa7\xc1\xd5\x63\x51\x09\x0e\xa3\xa2\xcc\x61\x94\x77\x38\x8c\x0a\x6b\x8f\x41\xa5\x69\x38\x1c\xe7\x71\x4e\x51\x94\xfe\xae\xe8\x6a\xfe\xb2\x2d\x0e\xa3\xa2\xce\x61\x54\xd8\x1d\x46\x45\x7e\x97\x21\x24\xe0\xd9\x8d\xfa\xa7\xc1\xb9\x43\x31\xf9\xbd\x46\x39\xdd\x57\x12\x7c\xf8\xe1\xf5\xc4\x2d\x7f\x6b\x59\x7b\xa3\xbf\x58\x35\xe6\x30\x2a\x06\x78\x9d\x23\xba\x64\x51\xe3\xa1\xe8\xd3\xbd\x46\xf9\xc8\xad\xb2\x25\x49\x7f\xbf\xf1\x66\xd6\x37\x95\x51\x63\xbd\x46\xf9\xad\xeb\x86\xe0\x05\xcf\x96\xc1\x8f\x87\xe5\xe0\x8a\x55\xd7\x8e\x05\xd1\xbd\x46\xf9\x0d\x98\xfb\x2a\x23\xbb\x6d\x59\xbf\xa1\xdf\xb4\x98\x76\x18\xe5\x8f\xdc\x7a\xfa\xb8\xad\xdd\x05\x2a\xda\x61\x1c\xff\xad\xe7\x71\xf3\xc6\x9c\xc1\x53\x8b\x90\xcc\xf2\x8c\x86\xff\xd4\xa8\x3f\xb4\xac\x1c\xcd\xbd\xc3\xa8\xd0\xf3\xb2\xee\xd2\xb5\x91\x8e\x13\xe1\xc8\x27\x74\x55\xee\x5a\xe6\x7c\xb3\x29\x37\xb2\xa3\xaf\x50\x4e\xdf\x30\xa8\xdc\xe7\x2c\xc3\xd5\xaf\x2d\x18\x2a\x5f\x3e\x8a\xf4\xba\x0d\xea\x09\xdf\xbb\xcf\x3b\xb2\xd2\x76\x89\x5a\xf2\x22\xee\x33\xbc\x8c\x4a\xf7\x77\xee\x67\xa6\x35\xb2\x9e\x92\x08\x1a\xc9\x3b\x8e\x85\x05\x50\x39\xf1\x73\x5b\x8f\x86\x31\xf5\x6e\x43\x90\xc7\xb7\xf9\x60\x79\xbc\x1d\xc9\x6f\x1a\x15\x13\x7e\x3f\x7b\xde\x51\x77\x34\x7e\x41\xbb\x61\x31\xc3\xab\xc7\xa0\x5e\x8a\x64\x57\x4d\x6b\xe7\xd8\x8d\xe1\xc9\x5d\x86\x60\x9a\xe5\xa5\x4c\x6f\x3e\x16\xb5\xd9\x6e\x58\xc4\xd4\xdb\xf2\x23\x3c\xde\xf5\xbe\x4f\x56\x9b\x7a\x0b\x15\x28\x7e\xce\x3d\x33\x22\x3f\x12\xcd\x45\xab\x89\xeb\xc7\xc3\x19\x5e\xd7\x0d\xc1\xcc\xef\x75\xd7\x4a\xd7\x26\x0c\xc2\x15\xad\x77\x60\x0c\xba\x2e\xdd\x81\xd1\x4d\xf6\x13\x51\xbb\xbb\x8b\x58\xfe\xf6\x82\x50\x8f\x73\x97\xee\x3f\x25\xe8\x7b\x8b\xd4\x88\xff\xad\x67\x46\xe4\x47\xe2\xea\xc7\xaf\xbe\xe8\x30\x2d\xa5\xb9\xf8\x4f\x98\x4a\xcf\x7e\x2a\x41\x73\xb3\x24\x9c\x8f\x13\x8f\x6f\x92\x8e\x8f\xe3\x77\xde\x2c\x0a\x46\x6b\xc5\xa4\xbf\xed\x3d\xcf\xa8\x34\xbc\xe5\xf7\x97\xd2\xd8\x11\x2e\xce\xa7\xfc\xdf\x49\x57\xe9\xfa\x97\x3e\xff\xe3\x32\x9a\x5b\xff\xe2\x84\x6d\xad\xc5\xaf\xec\xe8\x29\x64\xf8\x3f\xf8\x59\x06\xfd\x13\x63\xc8\xb2\xa6\x0d\xf1\x72\x18\x15\x79\x53\xe9\x54\x17\x25\xf9\xdd\xfe\xf3\x8a\x07\x9c\x9f\x92\x85\x6d\xe7\x0b\xe2\x92\xbb\x8d\x4c\xfc\xff\xe0\x33\xff\xe7\x09\xff\x5d\xb7\x31\x65\xe0\x63\x66\x0f\xbf\x51\xa0\x7f\xd7\x2d\xff\x24\x63\xed\xcc\x96\x23\xe1\x9f\xdd\x34\x2a\x99\xb3\x99\xff\xaa\x5e\x5f\xfe\xf9\x71\x25\xd2\x2b\x15\xf6\xaf\xfd\x28\x36\xf5\x86\x91\x91\x97\x3d\x83\xe1\x3f\x35\xee\xd5\xbd\x31\xff\x41\xf5\xaa\x07\xec\xbb\xad\x72\xff\x86\x40\x65\x2c\x8f\xee\xcc\x0f\x46\xf3\xcd\x9c\x7f\xfe\xb5\x61\x53\xdc\x37\x15\x51\x28\xce\x1f\xd9\x0b\x42\xdd\x67\x8f\x75\x39\x4b\x2d\x68\xff\xbf\x61\x54\x6e\x78\x56\x1c\x9e\x16\x8f\x2f\xbf\xbe\xd3\xf9\xe9\x52\xc4\x6d\xe0\xba\x21\xf8\x85\xac\xb4\xdf\x8b\xce\xe5\x68\x58\x5e\x06\x95\xfb\xcc\x66\xe4\xd2\xc6\xf2\xc1\x3f\x86\xa2\xb9\xae\x68\x3f\xb6\x78\xd6\xa7\xfa\xf5\x0b\xaf\xe6\x85\xba\x1c\x46\xf9\xa5\xcb\x47\x97\x4c\x79\x0e\xf4\xbc\x23\x33\x73\x8f\x68\xc4\xfe\x46\xfa\x70\x75\xcc\x68\xdf\x71\xe5\xfd\x9b\x46\x65\x5b\xbf\x29\xc4\x35\x70\x3a\x94\xee\x2a\x08\x79\x89\xd7\x7b\xd0\x9a\x38\xdb\xd5\xbe\xb1\xfc\x2b\xb8\x8c\xfe\xbc\x58\x75\xb7\xf7\xb8\xfa\x5e\x6f\xa1\x12\x76\x1b\xd4\xff\x6f\xf3\x7f\x21\x5c\xfd\xdb\xb0\x91\x2f\x92\x12\xbf\xbe\xf8\xab\xe8\xc1\x0b\xff\x36\xcb\x75\xed\x8d\x0d\x0d\x47\x56\x4d\x98\xd7\x47\xfd\x49\xd8\xc8\x97\xbf\x89\xbb\xdf\xf7\xce\xa4\xff\x69\xf4\xc1\x07\x1f\x7c\xf0\xc1\x07\x1f\x7c\xf0\xe1\xe9\x40\xff\x4c\x18\x5c\xf2\xfd\x8c\xc3\x1a\xf2\xe5\xc1\xb9\xdf\xce\x38\x18\xf9\xd2\x2f\x5c\x45\xce\x25\x9d\x1f\x98\x62\x5c\x03\xce\xa8\xce\x0f\x4c\x4d\xae\x7e\xa7\xa6\xdd\x65\x6a\x72\x8d\xb6\x92\xed\x4e\xa6\xd4\x7a\x94\xfb\x8c\xfb\xda\x9c\xbf\xe0\xcb\x97\x5d\xfb\x8c\x8b\x2e\x38\xe9\x97\x07\xf7\x45\xfc\xd3\x79\x27\xfd\xf2\x97\xfb\x22\x16\xb5\xbf\x43\xbf\xe5\xdc\x17\xb1\xaf\xf3\x9d\xef\x7f\xe5\xdc\xbb\x07\x95\xbf\x73\x3e\xde\xa3\xed\x7c\xe7\x8b\xdf\x39\x7b\x5d\xda\xce\x3d\x4c\x09\x3a\x3f\xf8\x8b\xd3\x59\x88\xca\x12\x97\xf3\x5f\xbf\x0a\x69\xff\xb0\xe4\xce\xf9\x90\x8b\x33\x74\xdf\x2d\xf8\x35\x33\xd8\xef\xff\x4f\xd4\x5a\x1f\xb0\xe5\xbf\x5c\x61\xcb\x25\xed\x5a\xc6\xc8\x92\x93\x6c\xf9\xe1\xc9\x18\xa6\x4c\xe1\xca\xf7\xb8\x52\x6b\xf2\x2c\x63\xd8\x92\x8e\x3a\xcc\x95\x9d\x4c\x7f\xba\xe8\x32\x5b\xa6\x0d\x3e\xfd\x34\x4c\xc0\xff\x06\x00\x00\xff\xff\x0d\xe2\xf3\x89\x3e\x42\x00\x00")

func assetsIconIcoBytes() ([]byte, error) {
	return bindataRead(
		_assetsIconIco,
		"assets/icon.ico",
	)
}

func assetsIconIco() (*asset, error) {
	bytes, err := assetsIconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/icon.ico", size: 16958, mode: os.FileMode(438), modTime: time.Unix(1695106044, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsIconPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\xb0\x0f\x4f\xf0\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x40\x00\x00\x00\x40\x08\x06\x00\x00\x00\xaa\x69\x71\xde\x00\x00\x00\x01\x73\x52\x47\x42\x00\xae\xce\x1c\xe9\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8f\x0b\xfc\x61\x05\x00\x00\x00\x09\x70\x48\x59\x73\x00\x00\x0e\xc3\x00\x00\x0e\xc3\x01\xc7\x6f\xa8\x64\x00\x00\x0f\x45\x49\x44\x41\x54\x78\x5e\xed\x5a\x79\x7c\x55\xc5\xbd\x9f\x73\xce\xdd\x12\x12\xc2\x56\x14\x14\xb4\xb6\x15\x2a\x45\xe8\x43\x41\x2b\x36\x60\x51\x12\x16\x03\x02\x49\xab\xef\x81\xb4\xb5\x56\xac\xb5\xa5\xad\x16\x5b\x7d\xf5\x5a\xfc\xd0\x45\x5b\x8b\x16\xb5\xf5\x69\x9f\xbb\x4f\x8a\x50\x44\x08\xb2\x27\x2c\x4a\x0c\x4a\xf6\x7d\xb9\x4b\x42\x72\x73\x73\xd7\xb3\xce\x4c\x7f\x33\xe7\xdc\x43\x2e\xa1\xcf\xfe\xc1\xed\x3f\xbd\xdf\x0f\x93\x99\xf9\xcd\x72\x66\xbe\xf3\x9b\xdf\xfc\x66\x2e\x28\x8b\x2c\xb2\xc8\x22\x8b\x2c\xb2\xc8\xe2\xdf\x15\x82\x15\x67\x1c\x6b\x9a\xe5\xcf\x09\x0e\xf1\xda\x70\xf7\x60\x30\x9f\x26\x8f\xa9\x05\x13\x16\x47\x0e\x55\x6c\x2f\xbf\x7f\x01\xb1\xaa\x70\xac\x6d\x4e\x5c\xe1\xf0\x38\x2e\x97\x92\xc9\x86\x67\xa7\x8c\x0e\x58\xe2\x8c\xe1\x5f\x42\xc0\xaa\x6a\xd5\x7b\x26\x18\xd9\x10\xf2\x0d\xc6\x04\x6c\x34\x3b\x73\x5c\x57\x23\x87\xc3\x43\xa3\x91\xcb\xaa\xd7\xcd\xe8\x62\x75\xbe\x71\x3c\x92\xab\x38\xdd\x2f\xf5\xf9\xc3\xab\x22\x67\x62\x7d\x02\xa5\x22\x21\xb4\x42\xc2\xfa\x77\x3f\x5e\x37\xbd\x87\x77\x94\x01\x88\x56\x9c\x31\xac\x3a\xa6\x7f\xcf\xdf\x16\x7e\xb8\xbf\x6b\xa0\xdb\x61\xa8\xd3\x4f\xdf\x73\xd5\x9c\xdc\xfe\xce\xf7\x94\xb8\x8a\x0c\xc1\x71\xb9\x55\x0d\x29\x38\xe7\x2f\x9d\x75\xbd\xab\x06\xfd\x83\x5b\x3d\x6a\x74\xd2\x67\xfb\x6a\x66\xb9\x95\xe8\x42\x43\x74\xbe\xfd\x95\xcd\x47\x25\xab\xda\x05\x47\x46\x09\x28\x2d\x1f\x18\x17\x0d\xcb\x1b\x07\xfb\x22\x48\xa2\xf8\x3b\xa7\xd6\x4d\xf7\x33\xb9\xcb\x88\x3d\xe5\x32\x12\x88\x22\x81\x13\xb0\xe2\xfd\xe8\x4d\xa1\x9e\xc8\x4a\x25\x21\x27\xdc\x44\xb9\xeb\xc3\x7b\xaf\x51\xb7\x3f\xb4\xa2\xd3\xa3\x85\xdf\x83\x3a\x37\xc4\x1c\xa3\x96\xb2\x7a\x99\x40\x46\x09\xd0\xa9\x70\x57\x7f\x5f\x68\xa4\x80\x68\xed\x27\xf7\x4c\x7b\xdf\x12\xa3\x5c\xaa\x1e\x75\x19\xf2\x49\xd8\x7f\x94\xe5\x0d\x5d\xbf\x7b\x20\x34\x88\x44\x4a\xdf\xa9\xba\xf7\xcb\x61\x5e\x09\xe0\x74\xd0\x6a\x91\x62\x44\x04\xb1\xd4\x12\x5d\x70\x64\x94\x00\x4d\x4e\x96\xc8\xb2\x86\x60\x3f\xdb\x93\x67\xd8\xb9\xa1\x4c\xff\xe0\x07\x37\x5c\x0b\xdb\xe1\xe5\x3b\x9f\xdb\x25\x25\x62\xf1\x85\x84\x30\x5b\x48\x2b\xcd\x1a\x26\x24\xb7\xa7\x4f\x42\xdc\x46\x7e\x99\x0b\x32\x80\x8c\x11\xb0\x66\xcb\x0e\x29\x91\x90\x67\x9a\x39\x5a\x63\xc6\xc3\x31\x90\x77\xe9\x25\xb2\xa2\x15\xb0\xb4\x88\x68\x23\x17\x5a\x90\x9c\x6e\x2c\x88\xdc\x4e\x4f\xe0\x82\x0c\x20\x63\x04\x0c\xe4\x4c\x98\xa8\x69\xba\x9b\xa5\x61\x0a\x3e\x2e\x3c\x0f\x34\x03\x4f\xd6\x75\x83\xa7\x61\xab\xa4\x1d\x7b\x82\xd3\xed\x41\x02\x1f\x62\xc6\xc6\x99\xb1\x8e\xb1\xa6\xb8\xb1\xae\xf3\x34\x18\xc0\x08\x8b\x97\xbc\x7a\xaa\xf8\x6b\xdb\x3b\x8f\x17\x6e\xef\x3e\x5a\xb8\xdd\x57\x09\xe1\x60\xc2\x10\xae\x23\x86\x49\x80\x83\xe8\x71\x9e\xb0\x20\x38\x5c\xe3\x89\x79\x52\xf7\x73\x41\x06\x90\x39\x02\x54\x39\x41\x0d\x8d\xa7\x89\x60\x2e\x23\xd5\xb4\xa3\xb1\xe0\xc0\x4b\x51\xff\xc0\xf5\x7d\xc1\xc4\x57\x42\x81\xe8\x61\x24\xc7\xbb\x28\x31\x09\x70\x52\x8b\x31\x0b\xa2\xd3\x79\x29\xa6\x8c\x00\xda\x66\x4a\x2e\x3c\x32\x46\xc0\x9e\xfb\x6e\x09\x8a\x86\x32\xc8\xd2\x14\x89\x7c\x0f\xbf\xbb\x76\x76\x64\x9c\xdb\x38\xe4\x76\x39\x58\x16\xe5\x10\xe5\xe9\x5c\xa5\x7f\x8f\xa8\x2b\xfc\x34\x90\xc5\x9c\x51\xbc\xc0\x02\x16\xdd\x33\x29\x94\x08\x14\x55\x58\xa2\x0b\x8e\x8c\x11\xc0\xe0\xd0\x65\x6e\xfd\xe1\x2c\xb7\x8c\x21\x4c\xc6\xe5\x29\xa0\xa2\xe9\xd7\x5c\xa4\x04\x22\x7b\x7e\xbc\x2c\x22\x19\xca\x47\x2c\x0f\xc7\xdd\x67\x79\x01\x60\xc5\xae\xde\x8b\x35\xc1\xf9\x25\x48\x52\x89\x1a\xdb\x4c\xe9\x85\x47\x46\x09\x70\x6b\xf1\x3f\x39\x0d\x05\x51\x41\x58\x76\xeb\x13\xdb\xf9\x66\x96\x72\x0b\x26\x10\x41\x62\x06\x4f\xde\xf9\xa3\x25\x2a\x93\xe5\x2a\x91\xdf\x4b\x58\x83\x03\x4f\xbc\x95\xe5\x19\x68\xfe\xe8\x1f\x45\x93\x86\x43\x40\xe4\x1d\x70\x85\x3f\xb6\xc4\x17\x1c\x19\x25\x60\xdf\x86\xe5\xe5\xa3\x94\xd0\x66\x48\x4e\xeb\x2a\x98\xfa\xe4\xd7\x4f\xc7\x6e\xc7\x9f\x19\xf3\x4b\xc5\x80\xc5\x46\xd4\xf6\xef\xf7\x3f\xb4\xec\x95\xd1\x72\xdf\xcb\x48\x10\xbe\x75\xfd\xeb\x6d\x77\x96\xd5\xa9\x9b\xfa\xc2\x83\xeb\x95\xa4\xdc\xea\x24\xfa\x3a\xab\x5a\x46\xc0\x57\x25\xd3\xb8\xf1\x99\x0f\x8b\x06\xa4\xbc\x35\xf0\xb9\x29\xb0\xb6\x7e\x89\x90\x5f\x7d\x46\x0d\x1e\xdb\xf7\xc3\x05\xd8\xaa\x82\xee\x79\xec\x39\xe1\x93\xd1\xff\x71\xfb\xa0\x23\x7f\x35\x64\xc7\x83\xf3\x54\x01\x46\xd1\x5b\xbd\xee\xea\x3e\xb3\x46\x16\x59\x64\x91\x45\x16\x17\x1c\xf6\x29\x50\x58\x83\x8b\x21\x7b\xb1\x95\x1d\x06\x67\x4c\x7e\xf7\xea\xca\x3f\xf7\x55\x15\xdd\x77\x27\xcb\x0b\x06\x89\x1f\x9c\xe9\xf8\x3f\x5e\x08\x80\xf6\xb3\x41\x3a\x8d\xa5\x05\x82\xb1\x48\xf0\x11\xec\x70\xcd\xe3\x85\xe7\x81\x60\xe0\x96\x83\x33\x9d\x47\x6e\xaa\x4a\xba\x0c\xb7\xe7\x0e\xf6\x32\xc0\xdd\x41\x00\x25\xd4\xa0\xaa\xd1\x3e\x62\xe0\xcc\xd1\xf2\xa2\x49\x69\x6f\x86\x29\x2c\xae\x52\xe7\x47\x64\x79\xb1\x41\x3d\x35\x13\x95\xc8\x5b\x54\x13\xae\xd9\xb6\x78\xfc\x61\x56\x56\xf8\x89\xf1\x5f\x54\x14\x4d\x77\x13\xee\x5b\x87\xbf\x24\xbe\x7a\x53\x55\x22\xcf\x70\xe7\xac\xb2\x64\x48\xc4\xfa\xb1\x83\x33\xdc\x0d\x36\x01\x5f\xd8\xb8\xfb\x10\x78\x6c\x5f\xb5\xb2\xc3\x50\xe0\xcc\xb9\xf3\xfa\xc4\xf1\x57\x76\xbb\x66\x70\xc7\x5d\x12\x44\x3a\x79\xec\xe7\xae\xd8\x7b\xf7\x15\x1d\x2c\xff\xf9\x8d\xbb\x7f\x07\xd3\xfa\x01\x4b\x8b\x88\x68\x70\xb5\x5d\x63\x20\xe9\x75\x96\x3f\x1f\x5c\xa2\x54\x59\xb7\x61\xc1\xdc\x69\x1b\x77\x16\xa8\xc8\xc9\x5d\x66\x1b\x82\x00\x0e\x53\x0e\x72\x8f\xb9\xf8\x54\x41\xc1\xa4\x25\x47\x4a\x3c\xfc\x25\x29\x85\xe2\x9d\x3d\xf7\x75\x35\x37\x3d\x25\x87\x07\x05\x11\x2e\x8c\x79\x13\x27\x28\xaa\xcf\xf7\x62\xfd\xc3\xc5\xdc\x67\xb8\xf2\xf1\xdd\x32\xa1\x82\xc7\x9c\x1c\x8d\x35\xff\xac\x68\xe4\xb5\xbf\xde\xf5\xf9\xb0\x2e\x35\x33\x19\x23\x5a\x42\xe4\x7b\x8d\x3f\x2b\x7e\xc6\x76\x84\xa8\xaa\xfc\x0a\x42\x0d\x51\x55\x24\xa8\xca\x4f\x90\xaa\x7c\x1b\xc2\xdd\x20\x2b\x67\x32\x55\x55\xe7\x3c\xfd\xe8\x03\x98\xc9\x41\xd6\xa4\x2b\xb2\xd0\x1f\xe9\x29\xb3\x9a\x23\x90\xbf\x01\xe1\x69\xaa\xaa\x8d\x10\xee\x86\x06\x27\x20\xff\x47\xd6\x16\xea\xef\x60\xed\xac\xb0\x11\x64\x2a\x56\xb5\x99\x6b\x9f\x3e\x2e\x11\x45\x91\x91\x2a\xaf\x03\x19\x85\x7a\xad\xbc\x8e\x22\x6f\x30\x06\xc2\x6d\xf1\x96\xfa\x99\x91\x8e\xfa\xe7\xac\x2f\x70\x94\xee\x0a\x4c\xe9\xac\xa9\x7d\x52\x09\x87\x5b\x9c\x88\xcc\x83\xb6\x3f\x8f\xb6\xb7\x7b\xb0\x61\x4c\xb2\xaa\x20\xa4\x28\x0f\xb3\xfe\x88\xa6\xca\x30\x96\x7b\x99\x48\x4e\x28\xbd\x82\xa6\x6c\xc2\x30\x1e\xf8\xc6\x9b\x10\xf6\x31\xb9\x4d\x40\x8b\x77\xd9\x2e\x60\x87\xbb\x9c\x93\x47\x8c\xf8\x9f\x16\x6f\xc9\x0b\x10\x9e\x1f\x49\x95\xc5\xe0\xb5\x55\x60\x8a\xe7\x98\xf5\x4a\x5e\x00\x06\xf9\x03\x87\xa6\xcb\x5f\x67\x31\x03\xb4\x3f\x01\x34\xc2\x96\xa0\x0d\xcd\xde\x65\x2f\x35\x78\x97\xb7\x83\x78\x2b\x2b\x23\x84\x1c\xb0\xfa\x63\xe1\xe7\xa0\x21\xeb\x74\x62\x8c\x68\x57\xd1\xd4\xfa\xc7\x56\x6a\xd0\x76\x0b\x7c\x83\xa9\x7a\xbd\x55\x67\x93\x03\xe1\x42\xc8\x27\x93\xe1\xe0\xa2\xe2\x3f\xd6\x5f\xc4\xfa\x61\xe8\xee\xe8\xfc\x85\x92\x4c\x48\x0e\x44\xee\x68\xf4\x2e\x3b\x04\x75\x37\x42\xdb\x6d\x30\xa6\xc9\x56\x15\xd4\xe4\x5d\xf6\x5b\x90\xc1\xfd\x42\xc8\xf9\xc2\xd8\x4b\xf6\x32\x59\xcd\xa3\xb7\xc5\x2e\x1d\x9b\x77\x1c\xc6\x17\xcf\x41\xfa\xb7\x1b\xbd\xcb\x1b\x98\x3c\xcd\x15\x86\x11\xf0\xeb\xa8\xae\x27\x9c\x5c\x00\xa8\x7e\xac\xcc\x80\x8f\xad\xc7\x84\x4c\x2f\xdb\x52\x9b\xc3\x64\xd0\x08\x4b\xa2\x44\x55\x5d\x99\x79\xf3\x33\x0d\x53\x79\x45\x80\x64\x4e\xc2\xee\x13\xee\xf2\x7c\xff\x0a\xa0\xd2\x43\x31\x16\x25\x5e\x84\x01\x9e\x1e\x54\xe2\xb3\x2c\x51\x0a\x29\x33\x80\x60\x80\x3e\xa8\xb3\x9b\x10\x43\x08\x84\xba\x66\x30\xd9\x77\xfe\xd6\x90\x3f\x18\x8e\x2c\x83\xed\xb5\x1f\xca\x3f\xe4\x15\x01\x90\x7f\x01\xa2\xb3\x1a\x60\xe2\x20\x85\xab\x64\x6f\x22\x72\xbd\x95\x47\x71\xd5\x98\x07\x23\xd9\x7f\xda\xbb\xd2\x7e\x77\x48\x23\x00\x32\x9c\x00\x8c\x1c\x36\x01\x0c\xec\x63\x94\xe8\x6f\xf8\x63\xe1\x11\x96\x08\x23\x62\xd4\x10\x82\x51\x7f\xbc\xdf\xde\x06\xd0\x39\xf8\xf8\x67\xfb\x84\x09\xf0\x09\x01\x01\x69\x0c\x9c\xf0\xde\x41\x61\xd0\x4f\x02\x81\x49\x4b\x94\x82\x4d\x00\x03\x34\xe2\xab\xa4\x1a\xea\x68\x16\xb7\xf8\x43\x8b\x55\x55\xf3\x40\xb5\xed\x2c\x9f\xc2\x08\xa4\x1d\x20\x94\xe6\x2d\xf9\xcd\xde\x7c\x4b\xc4\x48\x39\xc2\x62\xcd\xd0\xaf\xe3\x02\x80\xa6\xe3\xf9\x30\x26\xae\x11\x29\xa4\x11\x00\xe0\x04\xa8\x82\x33\x65\x41\xd1\xbc\x6a\xe5\xb6\xf9\x55\x89\xcb\x9a\xbd\x25\xab\x2b\x1e\x98\x9b\x7a\x99\x31\x24\x01\xbd\xc9\x12\x9a\xae\xd8\xdb\x00\x06\x96\xa6\x01\x30\x08\x53\x03\xe0\x3e\xc8\x05\x80\xf9\x55\xf1\xab\x0a\x3f\x92\x17\xb2\x6d\x52\xff\x8b\x25\x6f\x5b\x62\x46\x55\x1a\x49\x0c\x92\x68\x3e\xa4\x60\xeb\xc5\x24\x14\x8e\xde\xc0\x62\xd0\x34\x6e\xed\x53\x38\xe5\x2d\x4d\xc2\xad\xf1\x3d\xbf\xac\xe5\x5a\x22\xe4\x42\x06\x7b\x43\x20\xd0\x96\x13\x70\xcf\x1b\x15\xe3\x14\x83\x4c\x07\xbb\xf1\x8f\x09\x48\x6d\x01\x62\x28\xb6\x06\x74\x6f\xdb\x73\x7b\xef\x8e\xf7\xaf\xb4\xb2\x1c\xa0\xda\x98\x20\xa9\x06\x26\x56\xa7\xe8\xca\xd4\x05\x9b\x6b\x53\xf7\x7d\x3c\x54\x03\x80\x00\x73\x45\x29\x2c\xb8\x85\xc0\x8e\xf7\x67\xf8\xdf\x29\x5f\x62\x65\xcf\x02\x94\x04\x38\x48\xd3\x00\xa7\xd3\xfc\xe1\x04\x4a\xce\xb0\xd8\xc0\xe4\x6a\x88\x68\x3e\x52\x5a\x58\x7e\x28\xc0\x8e\x94\x54\x3f\xb2\xb8\xd7\xca\xa2\x5a\xef\xca\x10\xb4\xab\x25\x94\x5c\xb3\x6e\xcb\x49\x67\x73\x9f\x52\x68\x18\xb8\xbb\xde\x7b\x5b\xda\xc3\x6b\x1a\x01\x30\x78\xf3\x49\x4a\x40\x36\x01\x0e\x49\x62\xea\x67\x6b\x04\x03\x74\x0c\xb7\x38\x0a\x4a\x80\x5e\xa3\xb0\xc8\xa1\x78\x88\x6b\x01\xdb\xf3\x30\x03\xbb\x4f\x28\xb7\x34\x80\xfd\x33\xe1\x76\xbb\x26\x40\x5b\x3b\x7f\x0e\x6c\x02\xd6\xee\x0d\xe6\x21\x51\xba\x19\x92\x3a\x18\xad\x6a\x26\xc3\x84\x7e\x91\x91\x71\xd2\xfb\x8d\x04\xcb\xff\x13\x38\xac\x13\x92\x5b\x17\x4b\xcc\x88\xc8\x7a\x21\x7c\xb4\xdc\x92\xdb\x48\x23\x00\x56\x8c\x13\xe0\x9e\x3d\x6b\xf5\xbc\x1a\xfc\xfd\x9b\x2b\x22\x8f\x23\x49\x2a\x44\x82\xed\x54\x70\x40\x47\xb0\xd2\x04\x2c\x31\x66\xe7\x3c\xd5\xb0\x5a\x76\xd7\x13\xef\x82\xc5\x13\xd3\xb6\x00\x14\xf1\x09\xb9\x27\x4d\x9c\x0b\x8e\xd2\xf7\x6f\x3a\xa5\xfd\x58\xcc\xcd\x5f\xcf\x8b\x86\x81\x0a\x82\x28\xba\x16\x1e\xef\x1f\xb7\xa8\x2a\xf9\xc5\xb6\x50\xfc\xad\xa4\xac\x8e\x85\x31\xbd\xf9\xb1\xb7\x34\xfe\xd0\xa6\x27\x04\x8c\x09\x5b\x0c\xfb\x87\x93\x4f\x03\x8c\x93\x3f\xa5\x45\x34\xe5\x3a\xcd\xa0\xc3\xf6\x3f\x43\x1a\x01\xb0\x82\x9c\x80\x33\x27\xaa\x1f\xec\xde\xba\xfb\xa9\xd6\xbd\x87\x37\xe8\x9a\x2e\xc1\xc4\xd2\x08\x00\x18\xd0\xb9\xd4\xe0\xbd\xad\x0d\x3a\x3d\x01\xdb\xe0\xf2\x36\xc7\x24\xd8\x6b\xe0\xc3\xa5\xf7\xc9\x35\x40\x0e\xf4\xdd\xea\x83\xfe\x3a\xdf\x29\xff\x4d\x74\x20\x7c\xc9\x10\x85\x18\x02\x01\x89\x39\xee\x45\x81\x06\x5f\x5f\xd7\xc9\x9a\x3a\x5f\x7d\x43\x31\xf4\xdd\xe5\x44\xf8\x27\xac\xd4\x3f\x6a\x4a\x2e\xa6\x5c\x13\xff\xd9\xd5\x67\x76\xe0\x10\x8b\x45\x91\x94\x28\x3a\x9e\xea\x46\x06\x3f\xfb\x87\x22\x8d\x00\x18\x16\x27\xc0\x41\xf0\x43\xe0\x2a\xae\x14\x29\xb9\x1f\xb2\x1a\x68\xfc\xb9\x04\xc0\x16\x10\xf8\xc3\x1e\x0c\x12\xb6\x01\x45\x03\x89\x81\xb2\x73\x8f\x41\x00\xd7\x00\x01\x1b\xaf\xf2\xfe\x88\xf1\x9f\xf0\x8d\x36\xb6\xda\xbc\x74\x28\x28\x15\x48\x42\xf6\x27\xdb\x3a\x77\x24\xfd\x3d\xcc\xd3\x6c\x84\xc9\xcf\x86\x3d\xcb\x5f\x8e\x1a\xfa\x35\x7e\x04\x03\xec\x93\x63\xe9\xfe\x56\xb1\xe8\x68\xef\xc5\x43\xc3\xa2\x0a\xbf\xbd\x7d\xeb\xbc\x2b\x82\x30\xbe\x96\x04\x41\x0b\xc0\x7e\x54\xd5\x78\x57\x0e\x58\x45\x36\xce\x21\xc0\xdc\x02\xe0\xa6\x56\x80\x95\xde\x0a\x4e\xc6\x1f\x40\x56\x0e\x21\x8d\x00\xd3\x08\x9a\x04\xb8\x10\x7e\x0b\x22\x43\xc3\x5a\xa9\x93\x9f\x76\xa6\xe5\x36\x91\x32\x6a\xb4\x96\xf5\x07\xe1\x55\xe8\x6b\x33\xfb\xd2\xf9\x41\x4f\xb4\x7a\x4b\x4a\xa0\x4e\xa7\x20\x48\x13\x67\x4e\x9a\x66\x0f\x58\x54\x93\x32\x8b\xa1\xa5\x3d\x96\x81\xb6\xe0\x68\x5f\x63\x20\x38\x34\x04\x1b\xfd\xf6\x03\xac\x85\x43\x9a\xa6\xb3\x76\xc3\xd4\x9f\x21\x8d\x00\xb0\xc2\x9c\x00\x8c\x04\xfb\x23\xd0\xb0\x8b\x0e\xb3\x01\x14\xb4\x11\x4e\x23\x00\xb0\xdc\x0b\xf9\x7d\x8a\xae\x4e\x18\x99\x3b\x7a\x2e\xa4\xed\x3e\x41\x1d\xf8\x16\x60\x1d\xf3\x18\x00\xe5\x01\x10\x0c\x67\xc0\x74\x15\x2c\xc2\xd0\x3e\x1d\x6b\xf9\x5d\xb1\x98\x7d\x86\xcf\xa0\x7e\xb6\xf2\x50\x2e\xa4\x7c\x11\xa4\x74\x07\x34\xa5\xbd\x6b\x6b\xb2\xbd\x5b\x83\x80\x20\xbd\x4f\xee\xf4\x85\xac\x62\x0e\xd0\x4a\xf0\xfe\x18\x68\x93\x19\xa7\xe3\x1c\x02\x10\x3f\x6f\xe1\x36\x67\xab\x11\x0c\xf8\x45\x81\xf0\x33\xd5\x06\xc8\x0c\x81\x12\xfb\x37\x7b\xc8\xbf\x4e\xa1\x75\x52\xd7\xd8\xaf\xb8\x76\x9f\xa9\x63\x0d\xfe\xd8\x32\xb8\x84\x1c\x87\x41\xbd\x64\x65\x39\x1e\xfc\xe9\x4f\xf9\xec\xe1\x0f\xaf\x0f\x86\xef\x00\x8b\x63\x4a\x74\x01\x8b\x19\x9e\xfd\xe5\x83\xc0\x1a\x0d\x81\x4f\x65\xff\x4e\x58\xf5\x68\x69\x0c\xb4\x74\x25\xd4\xaf\x84\x32\xed\x12\x14\x59\xc8\xec\x92\x55\xcc\x01\xbe\x84\xc2\x62\x18\x00\x5f\xdc\x73\x91\x46\x40\xea\x14\x80\xd5\xb0\x09\x00\xb5\x3d\x59\x0f\x7e\xfd\xfc\xdd\x83\xf9\x45\xdb\xda\x79\x7d\x18\x28\x86\x05\xb3\xb5\xc2\x8d\x30\x7b\xb7\x97\x65\x43\x9f\x0b\xb1\xdd\xa7\x75\x2a\x30\xd8\x2b\x0e\x03\xec\x02\xcf\xf2\x83\xe5\x7f\xeb\x70\xde\xb8\x27\x6a\x3a\x2e\x67\x95\x86\x03\xf6\x3e\x37\x56\x18\x1b\xec\x18\x1c\x8a\x26\xb0\xb2\xe3\x4a\x37\x57\xd9\x5a\xc0\x21\x70\xa2\xb5\x43\xde\xd5\xf6\x23\x6b\x0a\x29\xed\x85\x0a\x9f\x4e\x00\x80\x57\x82\x81\xdb\x04\xa4\x90\x68\x6b\x58\x13\x08\xcb\x0f\xb0\x34\xac\xac\x6d\x04\x19\x6a\xbd\x2b\xa2\xb0\x02\x3b\xb1\x84\x1c\xec\x2c\xb3\xc4\x30\x6b\x76\x28\xf0\xfa\x36\x01\x29\xc4\x7b\x7a\xae\x0b\x77\x76\xfc\xc5\xcc\x8d\x34\x23\x3e\x4e\xdb\x78\xd5\xe9\xd8\x98\xbd\xf4\xf9\xd6\xa1\xbf\x16\xd5\x82\xa6\x09\x1d\x89\x18\x73\x88\x86\x42\x82\x6f\x98\xbf\xaf\x9d\x0b\x4a\x38\x01\x60\xb3\xce\x5b\x9e\x46\x40\xea\x18\x3c\x8f\xd5\x47\x6a\x2c\xb6\x28\xd9\xd1\xcc\x3d\x2d\xf6\x31\x08\x36\x01\x0c\xa0\x3d\xaf\x11\xdd\x00\x63\xce\xae\x03\x26\x40\xc6\x35\x80\x0d\x9a\x0b\x86\x20\x12\x97\x6f\x55\x7c\xed\x41\x96\xc6\x38\x9e\x2a\xe7\x04\x58\xd8\x87\x89\x21\x05\x23\xbd\xf3\xad\x3c\xeb\x8f\x6f\x0d\x85\x98\xee\xed\x59\x30\xa7\x8c\x39\x67\xe7\x05\x5f\x4c\xd8\x7a\x9f\xae\x01\xb0\x37\x53\x95\xd2\x34\xa0\xec\x5d\xff\x55\x8a\xa6\xdf\x02\x9d\xd4\xb1\x3c\xfb\x18\x8c\x34\xad\xed\x08\xa4\xbe\x07\xb3\x08\x0f\x95\xc3\x80\xf9\x84\x40\x2b\xd2\x08\x58\xbd\x3f\x34\x46\xd6\xf1\x5a\x18\xf8\x69\x96\xa7\x67\x95\xc6\x06\xb4\xdd\xcf\x62\x59\x93\x6d\x3b\x00\x5b\x83\xfd\xd4\x66\x00\xfb\x67\xdf\x21\x00\x92\x28\x32\x15\xfa\x07\x04\xd8\x06\xfd\xff\x27\x60\xd1\x11\xbf\xc7\x93\x9f\xc7\xd5\xcd\x91\x57\x70\x79\x51\x65\xef\xf8\xe2\x63\x67\x2e\x5a\xfa\x91\x3a\xaf\xdd\xd7\xf3\x57\x5d\x37\xc4\x1c\xa4\xd5\x17\x57\xf6\x8c\x72\xe7\x78\x46\x3b\x3d\xb9\x93\x16\x56\x04\x6d\xf5\x3c\xe5\x2d\x53\x61\x2a\x5b\x53\x04\x2c\x39\xd4\xe5\x2c\x98\x78\xd1\x44\x96\x76\xe5\xe4\x4e\x5c\xc8\xfa\x3b\xda\x3b\x7e\x51\x95\x72\x4d\x5b\xe0\xcc\xf6\x44\x3c\x39\x16\x08\x3f\xbd\x74\x7f\x9b\xd8\x71\xcb\x37\xb9\xcf\x2f\x3a\x5c\x63\x8a\x8e\x98\x7d\x7a\x90\x7e\x10\x22\x8c\x29\x5d\xf2\xb5\x03\x03\xe0\x3c\xf1\xad\xd1\x07\xe4\xbf\x06\x17\xb0\x39\xd7\xbe\x1c\xf8\xee\xb7\x9e\xdf\x2a\xac\x28\xef\x9b\x03\xca\xc8\x7e\x43\x1c\xa6\xe2\x0b\x2b\x7b\xc6\xb8\x47\xe4\xf2\xb6\xee\x51\x63\xc7\x17\x55\x04\xed\xdb\x62\x0a\x36\x01\xbe\x96\xde\x52\xc3\xe5\x7e\x86\xa5\x35\x11\x3d\xee\x6b\x0a\xf6\xfa\x1a\xfc\x3d\xad\x47\x3e\x38\xd0\xef\xf7\x4f\x81\x0f\xfb\x61\x92\x51\x90\xff\x15\x3b\x9d\xeb\x0d\x49\x5c\x1f\x68\xee\x49\xfb\xd1\x12\x56\xed\x75\xc1\xf2\x03\x02\x2d\x3d\xb3\xe2\x44\xda\xc3\xd2\x58\x14\xef\xf2\xb3\xfe\x20\x74\x1e\xaf\xfe\x30\xd0\xd4\xcc\x8c\x25\x01\x1f\xbf\x26\xd8\x76\x66\x54\x53\x7b\x7f\x2b\xab\x47\x04\x71\x41\xa0\x39\xc8\x9f\xd1\x4e\x7b\x57\x0d\xc2\x37\xab\x34\x4d\x9e\xdc\xd3\xda\xe1\xfb\xe6\xa6\x67\xb9\x56\x3a\x10\xf1\x12\x62\xf4\x0d\x06\x1a\xb7\x9c\xcc\x9f\xd6\xd3\xdc\xdc\x5a\x69\x18\x7a\x13\x7c\xfb\x7f\x59\xf9\x50\x04\x9b\x83\x07\x54\xd1\xf1\x20\x4b\xeb\x92\xf4\x4a\xa0\x29\xf8\x6b\x5e\x30\x04\xf6\x5e\x57\xdb\x3b\xdb\x30\x12\xdf\x60\xa3\xc7\x91\x41\x1e\x52\x60\x32\x18\x4c\x3d\x4b\x2b\xed\x9d\xfb\x61\xff\x73\x5b\xa0\xc4\xa3\x5c\x96\xc2\x28\x2a\x1f\x8c\x08\x39\x7f\x60\x69\xa5\xb3\xbb\x5f\x43\x0e\xde\x1f\x4d\xc6\xa1\xdd\xd9\xff\xfb\x60\xf5\x17\x86\x6b\x6c\x7c\xfa\x23\x6f\xe7\xc9\xc8\xc9\xeb\x21\x5d\x41\x4a\x47\x97\xfd\x43\x28\x4c\x6a\x0b\x45\xa4\x4d\xed\xea\x46\x22\x8d\x72\x7b\x02\x27\x48\xeb\x95\x8f\x6c\xbb\x11\x27\xa2\xff\x1d\xaf\x8f\xce\x82\xbd\x75\x18\xae\xb8\xf7\x37\x78\x97\x0f\xfb\x4f\x95\x4a\x5b\xd7\x1e\x68\x54\xc7\xfa\x36\x42\xfd\x88\x84\xe8\x07\x66\x49\x16\x59\x64\x91\x45\x16\x59\x64\x91\x45\x16\x59\x64\x91\xc5\xbf\x35\x10\xfa\x3b\x8f\x30\x9b\x17\xe9\xdb\xb9\xa5\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x27\x7f\xdf\xcf\xb0\x0f\x00\x00")

func assetsIconPngBytes() ([]byte, error) {
	return bindataRead(
		_assetsIconPng,
		"assets/icon.png",
	)
}

func assetsIconPng() (*asset, error) {
	bytes, err := assetsIconPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/icon.png", size: 4016, mode: os.FileMode(438), modTime: time.Unix(1695106044, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"assets/icon.ico": assetsIconIco,
	"assets/icon.png": assetsIconPng,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"assets": &bintree{nil, map[string]*bintree{
		"icon.ico": &bintree{assetsIconIco, map[string]*bintree{}},
		"icon.png": &bintree{assetsIconPng, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
