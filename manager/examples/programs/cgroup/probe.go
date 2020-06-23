// Code generated by go-bindata.
// sources:
// ebpf/bin/probe.o
// DO NOT EDIT!

package main

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

var _probeO = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x3d\x6f\x13\x41\x10\x7d\x6b\x1b\xc7\x24\x29\x02\x05\x0a\xa7\x14\x27\x90\xd2\x80\x36\x09\x42\x08\xba\xc8\x12\xa6\x49\x81\x50\xe8\x90\x8e\xcb\x79\x13\x2c\xdb\x67\x73\x77\xe6\xcb\x48\x88\x82\x3e\x0d\x35\xe2\x1f\xd0\x85\x8e\xbf\x90\x92\x92\x12\x89\x26\x48\x48\x74\x1c\x9a\xf5\xac\xbd\x5a\xdf\xc9\x19\x69\xbd\x33\x6f\x77\x66\xde\x3c\xad\xef\xdd\xfd\xbd\x56\x45\x08\x18\x13\xf8\x8b\x59\x34\xb3\xaf\x95\x99\xbf\xcb\xbf\xab\x10\x38\x11\xc0\x32\x80\xae\xf7\x27\x27\x94\xe2\x44\xa5\x69\xe4\x9d\xe9\x78\x5d\x00\x47\xc9\x60\x44\xfe\xd0\x57\x47\x63\xef\xe7\x14\x4f\x54\xdb\x27\x7f\x10\xfb\xd1\xd8\xfb\x31\xc5\x55\xe6\x47\xe4\x87\xc3\x6c\x34\xf6\xbe\x4f\xf1\x58\xbd\xd4\xf7\x87\x61\xd4\x1d\x7b\xa7\x1a\xff\xf6\x79\xc2\x69\x49\x00\xa7\x79\x9e\x9f\x54\x80\x4d\x00\x1f\x00\xd4\x89\x8f\x9e\x09\xf8\xc8\xdc\xa9\x06\xe5\x53\x0f\xaa\x4f\x1c\xa8\x3f\x71\x24\x7e\xc4\x7d\x19\x0f\x1e\xee\x01\xf8\x97\xe7\xf9\xa7\x5f\x02\xeb\x8e\x16\x5a\x1f\xfb\xa0\x66\xad\x06\x80\xb5\xd9\x5d\xe3\x8a\x37\x8f\xd0\x78\xbb\x22\x56\x69\x16\x5e\xc6\xbe\x14\xe8\xed\xda\x63\x5d\xfe\x77\x5e\x74\x56\x45\x75\x0e\xbb\x01\xe0\x12\x2e\x4c\xe3\x1a\xef\xd7\x34\x5e\x9f\xc3\x9b\x00\x2e\x5b\x75\x0c\xa7\x0d\x6b\x66\xfd\x04\x64\xa6\x5e\x65\x90\xcd\xfd\x96\x24\x67\x22\xdc\x56\xda\x3d\xd8\x9a\x88\x87\xe0\x85\x4a\xd2\xce\x20\x46\xd0\xeb\x44\x2a\x4e\x15\x64\xa2\x7a\x52\x3d\x0b\x0e\x93\xb0\xaf\x38\x23\x98\xdc\x0e\x0e\x47\x71\x84\x7e\xd8\x89\x65\x04\x99\x66\x49\x16\x1e\x40\xa6\xaf\xfb\x7a\x6f\xee\xb7\x20\x93\x41\x3b\xcc\x42\x3a\xdb\x91\x3b\x77\xce\x21\xd5\xb9\xec\x89\xd6\x6d\xde\xb6\xf9\x9d\xbf\x77\x70\xf7\x3f\x21\x78\xd5\x1d\x7c\xb7\xa4\x5f\xcd\x89\xd7\x16\xe4\xbb\x6f\xa2\xe1\xc4\xcf\x39\xff\x96\x83\x9f\xf1\xbe\x59\xd0\xcf\x9e\xe3\x26\xfb\xae\x06\x57\x44\x31\x5f\x77\xfe\xeb\x25\xf9\x1b\x25\xf9\x6e\xdc\x2b\xa8\x49\xe6\x33\x78\x75\x41\xff\xa5\x92\xfc\x7b\x0c\xfa\x0b\xf2\xef\xda\x6f\xda\xb2\xa7\x7c\x71\xdb\xc1\x5d\xfd\x6f\x03\xb8\x58\xc4\x9f\x0b\x1a\xbd\x57\xf8\x9e\xc9\x37\x78\xbb\xa0\x37\xd9\x31\xf7\x3f\xb6\x78\x57\xad\x7c\xf3\xdd\xf8\x1f\x00\x00\xff\xff\x8b\x0d\xd4\x7f\xb8\x05\x00\x00")

func probeOBytes() ([]byte, error) {
	return bindataRead(
		_probeO,
		"probe.o",
	)
}

func probeO() (*asset, error) {
	bytes, err := probeOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "probe.o", size: 1464, mode: os.FileMode(420), modTime: time.Unix(1592915230, 0)}
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
	"probe.o": probeO,
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
	"probe.o": &bintree{probeO, map[string]*bintree{}},
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

