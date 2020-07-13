// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ebpf/bin/probe.o

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

var _bindataProbeo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\x77\xf5\x71\x63\x62\x64\x64\x80\x01\x46\x86\xef\x0c\x08\x1e\x02\x64" +
	"\x30\x21\xd8\x0e\x50\x92\x93\x81\x11\xac\x96\x05\x8a\xb9\xb0\xe8\x1b\xc9\xc0\x3d\xc0\x87\xe1\xdf\xff\xff\xff\xe7" +
	"\xbf\x66\x64\x90\x40\x93\x03\x87\x31\xb2\x04\x0b\x12\xe6\x60\x60\x60\x10\xc0\x6f\xb6\x2e\x58\xe9\xc7\xff\xe8\xe2" +
	"\x92\x0c\x0c\x0c\x82\x60\x43\x10\xc6\x32\x40\x8d\x13\x64\x60\xc5\x10\x57\x07\x8b\x33\xc3\xc5\x65\x60\x91\xaf\x57" +
	"\x92\x5a\x51\xc2\xa0\xe7\x14\xe2\xa6\x07\x62\xc4\x97\xa5\x16\x15\x67\xe6\xe7\x31\xc4\xe7\x64\x26\xa7\xe6\x15\xa7" +
	"\x32\xe4\x26\x16\x14\xeb\x27\x27\x26\x67\x80\x98\x99\x79\x7a\xc9\x0c\x7a\xc5\x25\x45\x25\x89\x49\x0c\x7a\xc5\x95" +
	"\xb9\x60\xda\x29\xc4\x8d\x94\xf0\xc2\x07\x4c\x18\x18\x90\x5c\x89\x00\x12\xd0\x74\xe9\x89\x26\x8e\x9e\x86\x19\xa1" +
	"\x98\x0d\x4d\xdc\x01\x87\x7d\x2c\x68\x7c\x25\xa8\x7e\x74\x37\xc0\xf4\xcb\x30\xe2\xd7\x2f\x85\x43\x7f\x0c\x23\x76" +
	"\xf5\xe8\xee\x17\xc4\xa1\x3f\x01\x87\x7e\x74\xbe\x0b\x16\x33\x41\x20\x05\x2a\x28\x49\xc0\x7e\x76\x1c\xfa\x6b\xa1" +
	"\x82\x0a\x04\xf4\xdb\x30\x30\x30\x30\x61\xd1\xbf\x00\xaa\xb0\x02\x49\x1f\x13\x34\x0b\x80\x00\x2c\x7b\x00\x02\x00" +
	"\x00\xff\xff\x8f\x5f\x41\x85\xa8\x04\x00\x00")

func bindataProbeoBytes() ([]byte, error) {
	return bindataRead(
		_bindataProbeo,
		"/probe.o",
	)
}



func bindataProbeo() (*asset, error) {
	bytes, err := bindataProbeoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "/probe.o",
		size: 1192,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1594365727, 0),
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
	"/probe.o": bindataProbeo,
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
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
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
		"probe.o": {Func: bindataProbeo, Children: map[string]*bintree{}},
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
