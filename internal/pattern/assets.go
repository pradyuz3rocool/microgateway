// Code generated by go-bindata.
// sources:
// DefaultHttpPattern.json
// DefaultChannelPattern.json
// DO NOT EDIT!

package pattern

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

var _defaulthttppatternJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\xb8\x0a\x41\x53\x03\xa9\x92\xfd\x79\x32\xe0\x87\x38\x0b\xd6\x75\x0b\x56\xb4\xee\xfa\x30\xec\x81\x96\xce\x16\x63\x8a\xd4\xc8\xa3\x33\xaf\xc9\x77\x1f\x44\x49\x0e\xf5\xc7\xb2\x9d\xc0\x01\x86\xbe\x24\x11\xef\x78\xfc\xdd\xef\x8e\x77\xbc\x7c\x1d\x00\x04\x92\xa5\x18\x8c\x20\xf8\x09\xe7\xcc\x0a\x7a\x47\x94\x7d\x60\x44\xa8\x65\x70\x96\xcb\x0d\x61\x66\x82\x11\xfc\x39\x00\x00\xf8\xea\x7e\x02\x04\x7c\x9e\x6f\x3a\x09\x23\x25\xe7\xa1\x35\xf8\x91\x11\xfe\xc6\x53\x4e\xa8\x61\x3c\x06\xd2\x16\xdd\x7e\xa7\x6c\x50\xaf\x78\xe4\x8e\xf1\xf4\x1e\xe5\x5c\x66\x96\x82\xd1\xc6\x3a\x40\x40\x6a\x89\x32\xdf\xb0\x10\x6a\xc6\x44\x50\x4a\x1e\xdc\xef\x87\xb3\x7e\x30\xef\xbf\x4c\x7b\x41\xbc\xff\x32\xfd\x83\x09\x1e\x33\x52\x7b\xa2\x18\x9f\x84\x19\x5b\x0b\xc5\xe2\x30\x41\x16\xa3\x36\xe1\xa5\xa5\x44\x69\xfe\x2f\x23\xae\xe4\xc6\x4a\x7e\x10\x5f\x48\x2e\x17\x37\x48\x89\x8a\xcb\xcd\x0e\xd9\xed\x1d\x7d\xaa\xc9\xbc\x4d\x4b\x5c\x37\x54\x7f\xc5\xb5\xaf\xc0\x6c\xd3\xd6\xa5\xad\x59\xe0\xc6\x34\x14\x7e\x31\xa6\x86\xcb\xce\x9a\x68\xec\xec\x30\x62\xaf\xb8\x8e\x2c\xa7\x89\x46\xb6\xdc\x11\xe8\xba\x6a\xd0\x63\xfe\x4d\x2b\x70\x73\x26\x0c\xc2\xfd\x3d\x9c\x84\x7e\xa8\x42\x65\x29\xb3\x64\xc2\x55\xbe\x52\x9d\x3e\x84\xd7\xaf\xa1\x34\x91\x3a\x5e\x73\xc9\xe9\xcf\xd7\xd3\xd3\x00\xba\x90\xe5\x29\x3e\x61\xd1\x12\x65\x7c\xf9\x08\x3d\x61\x82\x4a\x34\xbe\x42\x88\x5a\x2b\x0d\xaf\xc6\x20\xb9\x70\x47\xbd\x72\x2b\x21\x37\x12\xc9\xfd\xd9\xb9\x63\x78\x6c\x8f\xdf\xb4\x5d\xfe\xf0\x79\x7a\x5a\xd8\x68\x49\x7e\xff\xb4\x55\x74\x39\xbd\x7a\x77\x3a\xec\x8c\xa1\xe7\xd6\xa4\xf7\x9e\x44\x4a\x12\x4a\xf2\xd3\xab\x5a\xaa\xf2\x6b\x07\xd1\x93\x83\x89\x9e\xec\x41\xf4\xae\xcc\x75\x44\xf6\x45\x7c\x7b\x6e\x0d\xef\xef\xa1\xcf\x85\x23\xc4\x68\xd8\x1d\xa5\xc6\x4d\xeb\x8b\x93\xca\x50\x17\xf5\x6a\x94\x07\xcd\x4a\xaa\xae\xe6\xf3\x2a\x40\x0f\x8f\xe3\x27\xf3\x38\xfe\x9f\xf0\xa8\xd1\x20\x35\x58\x1c\x00\xfc\xe5\xda\xa7\x46\x93\x29\x69\xb0\xa7\x85\x7a\x2d\x71\x73\xdf\x45\xfe\xfd\x11\x59\x94\x60\xdc\x2e\xb3\x8e\x9f\x60\xe4\x56\x37\x8b\xc5\xd6\xe6\xbd\x8c\x73\xd7\x7e\xbc\xf8\xc1\xeb\x03\x31\x23\x56\x53\x73\x5d\x9e\x91\x35\x55\x83\x06\x07\x07\xae\xff\x89\x10\x63\x8c\xe1\x2d\x4c\x13\x84\x92\x2a\x58\x2b\x0b\x09\x5b\x21\x68\xfc\xdb\xa2\x21\x8c\x81\x1b\x50\x2b\xd4\x40\x09\x02\x13\x42\xdd\x61\x0c\xce\x85\x30\xd8\x9c\xf2\xf0\xd4\x16\x5e\xe4\x4e\x7f\x59\x74\xf5\xf3\x39\x04\x7d\xb7\x83\x20\x67\xf1\x06\x8d\x61\x0b\x2c\xcb\x5c\x27\xa2\x9a\xde\x99\x6f\x61\x55\xe8\x72\x25\xf7\x30\xd3\x56\x7e\x22\x8f\xdb\xef\xeb\x49\x58\x97\x6d\xce\x26\xcd\xb3\xec\x65\xd2\xae\xb2\x18\x44\x05\x14\x98\x95\x38\x4b\x0c\x87\x79\x7d\xc4\x87\x44\x8b\x03\x67\x79\x1f\x12\xbe\xbf\xb8\x68\x93\x90\xc7\xbd\x56\x2b\x2b\x40\x4e\xfe\x42\xae\x1e\xed\x05\x71\x14\x8e\x26\x07\x70\x74\x78\xa2\x76\x9d\x7f\xed\xac\x6c\xa9\xeb\x65\x31\xec\x28\xeb\xd5\x40\xd5\x39\xe9\xc4\x68\x22\xcd\xb3\xaa\x75\xb8\x62\x2b\x9a\x4a\x1a\x5d\x8c\x17\x9c\x12\x3b\x0b\x23\x95\x9e\x67\x5a\xdd\x62\x44\x6f\xe7\x42\x2d\xd4\x79\xca\x23\xad\x16\x8c\xf0\x8e\xad\xcf\x59\x44\x7c\xc5\x69\x7d\xae\x19\x61\xcb\x94\x41\x22\x2e\x17\xa6\xee\xb6\x53\xf3\x1f\x6b\xba\xc2\xda\xcf\x6b\xe5\x5a\xf7\xfc\xd4\xf0\xad\xd4\x40\x30\x2a\x45\x70\x93\x94\x79\xae\x8f\xb7\x77\xd4\xfd\xdc\xab\x90\x6d\xeb\xe1\x0d\x6c\x57\x8d\x72\x53\xbd\x02\x9e\x09\xaf\xac\x62\xb3\xe6\xe9\xdd\x51\x48\x8b\xe4\xdb\x04\xc1\x7d\x7b\x79\x48\x89\x46\x93\x28\x51\x9b\xfb\x1e\x17\x7d\x4d\x9e\xa2\xb2\xb5\x88\x56\x4b\x9e\x56\x86\x9a\xd7\x07\xd2\x72\x65\xaf\x98\x77\x4f\x4d\x0d\x5e\x6f\xd8\x12\x81\x49\x48\x88\x32\x88\x98\x10\x40\x2a\x7f\x2f\x68\x98\x15\x5b\xf7\x67\x38\x9f\x1f\x34\x9f\x79\xf9\x8d\x86\x76\x52\xda\x9a\xb9\xd3\xd6\xa4\x6d\x35\xf7\x15\x4a\x60\x9f\xf5\x8e\xff\x31\x74\xf0\x30\xf9\x26\x78\xc8\x4b\xde\xe0\x61\xf0\x5f\x00\x00\x00\xff\xff\x30\x31\xfa\x4b\x2d\x12\x00\x00")

func defaulthttppatternJsonBytes() ([]byte, error) {
	return bindataRead(
		_defaulthttppatternJson,
		"DefaultHttpPattern.json",
	)
}

func defaulthttppatternJson() (*asset, error) {
	bytes, err := defaulthttppatternJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "DefaultHttpPattern.json", size: 4653, mode: os.FileMode(420), modTime: time.Unix(1550181788, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _defaultchannelpatternJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4d\x6f\xdb\x30\x0c\xbd\xe7\x57\xb0\x42\xd0\x53\xe7\x74\x1f\xa7\x00\x39\xb4\xdd\xa9\x43\x81\x01\x2b\xd6\xc3\xb0\x03\x23\x33\xb1\x1a\x45\x32\x24\x2a\x45\xd6\xe6\xbf\x0f\x96\x3f\xe2\x18\x89\xeb\xa2\xbd\x24\xb0\xf4\x48\x3d\x3e\x3e\x4a\xcf\x23\x00\x61\x70\x4d\x62\x0a\xe2\x3b\x2d\x30\x68\xbe\xc9\xd0\x18\xd2\x3f\x91\x99\x9c\x11\x17\x05\xc4\x33\xe5\x5e\x4c\xe1\xcf\x08\x00\xe0\x39\xfe\x02\x08\xb5\x28\xe2\xc6\x89\xb4\x66\x91\x04\x4f\xb7\x0f\xf7\x30\x9b\x01\xbb\x40\x31\x2e\x82\x3c\xb9\x8d\x92\xf1\x84\xdb\x87\xfb\xdf\xa8\x55\x8a\x6c\xdd\x1e\xa0\x4c\x1e\x58\x4c\x9b\xb4\x00\x82\xed\x8a\x4c\x11\x31\x1b\x27\x39\x6e\xb5\xc5\x34\xc9\x08\x53\x72\x3e\xb9\x0a\x9c\x59\xa7\xfe\x21\x2b\x6b\x9a\x2c\x00\x62\x45\xdb\x2a\x24\xf2\x79\x7c\xe2\x1f\xb4\x15\xd5\xfe\x2e\xfe\xef\x2e\xfa\x0b\xb8\x51\x4e\x06\xc5\xd7\x8e\x70\x45\xae\xb7\x96\x43\xa8\x18\x94\xbe\xd2\x67\x81\xda\x13\xbc\xbc\xc0\x38\x69\x2b\x92\xd8\xc0\x79\x60\x9f\x6c\x8a\x95\xde\xc3\xaf\x64\x51\xfc\x35\xca\x15\x99\xb4\x4f\x49\x00\x21\xcb\x7e\x16\x61\x4c\x9e\x5b\x8a\x01\x88\x0d\xea\x40\xcd\x56\xad\x55\x93\x30\x43\x5d\xe4\x13\x91\xf1\xb0\x1a\x8f\x4b\x08\xe7\xe7\x30\x4e\x0e\x68\x27\xe4\x9c\x75\x70\x36\x03\xa3\xf4\x10\x89\xfb\x0c\x63\x73\x72\xa5\x21\xa6\x20\xa4\x0d\x86\xeb\x9e\xbc\xaf\xf5\xa7\x79\xcf\x3e\x9e\xb7\x23\x4f\xdc\x61\x3d\x02\xf8\x1b\x47\xd0\x91\xcf\xad\xf1\xf4\xb6\x31\x2c\xf9\xf7\xbb\xac\x6c\x6e\xc3\x32\x96\x27\xa6\x31\xbc\x59\x2c\x83\x0e\xb9\x4b\x9b\x16\xe5\x7e\xbb\xfc\xdc\x1a\xc2\x14\x19\xbb\x0e\xac\x33\x16\xa3\x79\x94\x4a\x04\xdc\x91\xf7\xb8\x24\xd1\x44\xee\x3e\xac\x7f\x87\x7b\xcd\xb1\xec\x54\x9e\xd3\x91\x49\x7b\xbb\x04\x5f\x87\x4a\x20\x4b\x2a\x30\xaf\x78\x56\x1c\x86\x56\x5d\xe7\x89\x3d\x1b\x42\xed\xcb\xe5\xe5\x2b\xd4\x6a\x63\x15\xec\x7e\x05\x29\xc9\xfb\xb3\x93\x74\x1a\x3b\x56\x76\x3f\xe2\xc6\xfa\x39\x39\x7e\xd9\xa7\xe4\xa5\x53\x79\x6d\xf9\x0a\x41\xe0\xed\x9a\x20\x5e\xfb\x7e\x0f\x76\x14\x9b\xbc\x54\x9c\x85\x79\x22\xed\x7a\x92\x3b\xfb\x48\x92\x3f\x2d\xb4\x5d\xda\xc9\x5a\x49\x67\x97\xc8\xf4\x84\xdb\x09\x4a\x56\x1b\xc5\xdb\xc9\xe3\x13\x1f\xbf\xaa\x6a\x66\xa7\xe6\xb3\xc3\xed\xa6\xd3\xaa\x7a\xc2\xdf\x49\xaf\x72\xc0\xbc\x7b\xba\x27\x66\x65\x96\xfe\xb0\x8d\xeb\xb2\x8d\xcd\xab\x16\xbf\x5b\x1d\xe5\xcc\x91\xcf\xac\x4e\xdb\xa0\xfd\x62\x1b\xa9\xd6\x64\xa3\x4b\xf6\xb8\x6a\xa9\x85\xca\xc9\x29\x7b\x90\xac\x5a\xe9\xb5\x65\xad\xec\x89\x77\xa9\x23\xec\x1d\xae\x08\xd0\x00\x46\x34\x48\xd4\x1a\xd8\xc2\xd6\x06\x07\xf3\x6e\xec\x6b\x22\x4b\x6b\xd8\xa9\x79\x4b\xdf\xea\xbd\xdb\x1b\x76\xb4\x1b\xfd\x0f\x00\x00\xff\xff\x6d\x96\xfa\xa0\xe9\x08\x00\x00")

func defaultchannelpatternJsonBytes() ([]byte, error) {
	return bindataRead(
		_defaultchannelpatternJson,
		"DefaultChannelPattern.json",
	)
}

func defaultchannelpatternJson() (*asset, error) {
	bytes, err := defaultchannelpatternJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "DefaultChannelPattern.json", size: 2281, mode: os.FileMode(420), modTime: time.Unix(1550184410, 0)}
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
	"DefaultHttpPattern.json": defaulthttppatternJson,
	"DefaultChannelPattern.json": defaultchannelpatternJson,
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
	"DefaultChannelPattern.json": &bintree{defaultchannelpatternJson, map[string]*bintree{}},
	"DefaultHttpPattern.json": &bintree{defaulthttppatternJson, map[string]*bintree{}},
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

