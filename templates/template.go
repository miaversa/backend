// Code generated by go-bindata.
// sources:
// templates/cart.html
// templates/login.html
// templates/payment.html
// templates/register.html
// templates/shipping.html
// DO NOT EDIT!

package templates

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

var _cartHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4d\x8f\xdb\x36\x10\x3d\x5b\xbf\x62\x4a\xb8\xb7\xae\xe8\x28\xe9\x25\xa0\x04\x14\x69\x0b\x14\xfd\xc8\xb6\xbb\x2d\xd0\x53\x40\x49\x63\x8b\x58\x8a\x64\xc9\x91\x11\x47\xf0\x7f\x2f\x48\x59\xda\xcd\xc2\xb5\x93\xb6\x17\xcb\x22\xdf\x23\xdf\x3c\x3d\x0e\xc5\x17\xdf\xbe\x7d\x73\xff\xe7\xed\x77\xd0\x51\xaf\xab\x4c\xcc\x0f\x94\x6d\x95\xad\x04\x29\xd2\x58\xdd\x63\x20\x14\x7c\x7a\xc9\x56\x42\x2b\xf3\x00\x1e\x75\xc9\x02\x1d\x34\x86\x0e\x91\x18\xd0\xc1\x61\xc9\x08\xdf\x13\x6f\x42\x60\xd0\x79\xdc\x96\x8c\xcb\x10\x90\x02\x27\xd9\x74\x07\x6b\x42\xde\x2b\x93\xa7\x79\x5e\x65\x82\x4f\x1b\x89\xda\xb6\x87\x2a\xcb\xc6\xb1\xc5\xad\x32\x08\xac\x91\x9e\x7e\x20\xec\xd9\xf1\x18\x65\xf8\x2a\x5b\xad\x04\xb5\xd0\x68\x19\x42\xc9\x9c\x2c\x80\x1a\x16\x47\x57\x62\x6b\x7d\x0f\x3d\x52\x67\xdb\x92\x39\x1b\x68\x1a\x5f\x09\x65\xdc\x40\x27\x5d\x9d\x6a\x5b\x34\x0c\x8c\xec\xb1\x64\xef\x26\x38\x83\xbd\xd4\x03\x96\xac\x45\x8d\x84\xd7\x79\xe1\x61\x58\x38\xe3\x98\xdf\x7a\xdb\x0e\x0d\xe5\x77\x3f\xfe\x7e\x3c\xce\xec\x7a\x20\xb2\xe6\x44\x0f\x43\xdd\x2b\x62\xd5\x7b\xc1\xa7\xf1\x49\x32\x8f\x9a\x53\x4d\x9c\xda\x4b\xb5\x05\xf2\xd6\xec\xaa\x27\x7b\xfd\x22\x7b\x3c\x1e\x05\x3f\xcd\x88\xda\xf3\x04\x1d\x47\x2f\xcd\x0e\x61\xfd\xf0\x15\xac\xf7\xf0\xba\x84\x85\xf2\xd6\x91\xb2\x26\x44\x2b\x13\x10\xd4\x16\xf0\x2f\x58\xef\xd3\x62\xc0\x82\xfa\x80\x0c\x4e\xd3\xf7\xb2\x97\xa6\xb3\xaf\xe1\x04\x46\xd3\x2e\xc4\xf5\x3e\xff\x23\x16\x3f\x0d\x3c\xce\x5d\xa8\xe2\x89\xf2\x5b\xaf\x9a\x24\xfd\x12\xf8\xd7\x41\x1a\x52\x74\xb8\x86\xbb\xb7\x24\xf5\x02\x12\x3c\x46\x64\x16\xf4\x2c\x47\x77\x9d\x72\x4e\x99\xdd\xe5\x2c\x41\x63\x75\x70\xd2\x94\xec\x15\xab\xbe\xf7\x98\x12\xff\xcf\x02\x9c\x57\x86\xb6\xc0\xbe\xcc\x8b\x2d\x83\xfc\xd3\x94\x24\xd1\x9f\x21\x23\xe1\xff\x7f\x19\x51\x81\x20\x59\x6b\x9c\x17\x6c\xac\xd6\xd2\x05\x84\x5a\x42\xed\x0b\xa8\x6f\x6e\x6a\x2d\x9b\x87\x9b\x17\x1b\x70\xfb\x02\x5c\xf7\x32\x06\xf2\x4c\xc6\xe2\x21\x9d\x92\x35\x8e\x84\xbd\xd3\x92\x9e\x1e\x5f\x58\xef\xe3\xe4\x12\x95\xe7\xa0\xe5\xdb\x40\x3e\xff\x3d\x07\x9b\x8c\x83\xf9\xab\x67\x82\x27\xf9\x1f\x55\x18\x53\xbd\x23\x58\x12\x04\x9b\x18\xe9\xe7\x2b\x25\x93\x22\x4d\x87\x18\x63\xe1\xaa\x3b\x1c\xa0\x91\xde\x2b\xd3\x59\xc0\x40\x12\xf6\xf2\x83\xb2\x82\xbb\xc7\xe5\xcf\xf7\x98\x4f\x6c\x14\x2f\x8a\x97\xa9\xdb\x5d\xc2\xc7\xdf\x85\xf0\x8d\x41\x0d\xbf\x61\x6b\x4d\x6b\xe1\x27\xdc\xc5\xd2\xaf\xf0\x5d\x3c\x5b\x8f\x3b\x6e\x36\x79\x71\x95\x63\x53\x5b\x78\x37\x9d\xfe\x99\xf9\xf5\x89\x76\xb6\x8b\x49\x83\xfa\xb1\x91\xcd\x5d\xec\xbf\xb9\xf3\x6a\x73\x5d\xe9\x47\xee\xbc\xb1\x5a\x7a\xf8\x59\x1a\x54\xde\x7e\xa6\x31\xc5\x66\x93\x6f\xfe\x95\x31\xc5\x45\x63\x9a\xa8\xe9\x9c\x33\xae\x12\x72\xbe\x09\x9d\xdc\x49\xcf\x59\x95\x9e\x82\xcb\x2a\x45\x2c\x13\x7c\xba\xfd\x04\x4f\x97\xef\xdf\x01\x00\x00\xff\xff\x98\x4c\xda\x84\x93\x07\x00\x00")

func cartHtmlBytes() ([]byte, error) {
	return bindataRead(
		_cartHtml,
		"cart.html",
	)
}

func cartHtml() (*asset, error) {
	bytes, err := cartHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cart.html", size: 1939, mode: os.FileMode(436), modTime: time.Unix(1517544759, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loginHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xbd\x4e\x03\x31\x10\x84\x6b\xfb\x29\xcc\x3e\x40\xdc\xa3\xb5\x1b\x48\x0d\x12\x69\x28\x7d\xb9\x0d\xb6\xf0\x9f\x6e\xf7\x80\x7b\x7b\x94\x38\x11\x48\x88\xca\x1a\x8d\xe7\x1b\xcd\xe2\xdd\xe3\xd3\xc3\xe1\xf5\x79\x6f\xa2\x94\xec\x35\xde\x1e\x0a\xb3\xd7\x0a\x25\x49\x26\x7f\x20\x16\x42\x3b\x84\x56\x98\x53\x7d\x37\x0b\x65\x07\x2c\x5b\x26\x8e\x44\x02\x46\xb6\x4e\x0e\x84\xbe\xc4\x1e\x99\xc1\xc4\x85\x4e\x0e\x6c\x60\x26\x61\x2b\xe1\x18\xb7\x56\x79\x57\x52\xdd\x5d\x7c\xeb\x35\xda\x51\x84\x53\x9b\x37\xaf\xb5\xc2\x53\x5b\x8a\x29\x24\xb1\xcd\x0e\x7a\x63\x01\xaf\x95\xc2\xee\xf7\x25\xa4\x7c\x8f\xd3\x72\xce\x29\xa5\x30\xd5\xbe\xca\xaf\x52\x30\x35\x14\x72\x40\xe7\x8f\x60\x3e\x42\x5e\xc9\xc1\x88\xdb\x7e\xa5\xbc\x50\x8d\xe1\x3f\x4a\x0f\xcc\x9f\x6d\x99\x6f\xa4\x1f\xfd\x07\x36\x68\x38\xad\x22\xad\x5e\xe3\xbc\x4e\x25\x09\xf8\xdc\xde\x52\x45\x3b\x3c\x3f\xba\xd1\x9e\x87\x79\xad\xd1\x8e\xa9\x68\x2f\x97\xfe\x0e\x00\x00\xff\xff\x59\xf9\x53\xdb\x80\x01\x00\x00")

func loginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_loginHtml,
		"login.html",
	)
}

func loginHtml() (*asset, error) {
	bytes, err := loginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "login.html", size: 384, mode: os.FileMode(436), modTime: time.Unix(1517613131, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _paymentHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8e\xb1\x0e\x83\x30\x10\x43\x67\xf2\x15\x69\x3e\x80\x13\xfb\x91\xa5\xed\x5c\x06\x96\x8e\x29\xb8\x0d\x6a\x08\x15\x77\x43\xf3\xf7\x15\xa0\x4e\x96\xf5\x2c\xdb\x7c\xba\xdc\xce\xfd\xbd\xbb\xda\xa8\x73\xf2\x86\xff\x82\x30\x7a\x53\xb1\x4e\x9a\xe0\x7b\x88\x82\xe9\x30\xa6\xe2\x34\xe5\xb7\x5d\x91\x5a\x27\x5a\x12\x24\x02\xea\xac\x96\x0f\x5a\xa7\xf8\x2a\x0d\x22\xce\xc6\x15\xcf\xd6\x51\x10\x81\x0a\x69\x18\x62\x59\xb2\xd4\xf3\x94\xeb\x9d\x93\x37\x4c\xc7\x10\x3f\x96\xb1\x6c\xcd\xb1\xf1\x5d\x78\x85\x19\x59\x17\xa6\xd8\x6c\x91\x83\x31\xed\xd7\x7e\x01\x00\x00\xff\xff\x68\x73\xfc\x57\xb1\x00\x00\x00")

func paymentHtmlBytes() ([]byte, error) {
	return bindataRead(
		_paymentHtml,
		"payment.html",
	)
}

func paymentHtml() (*asset, error) {
	bytes, err := paymentHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "payment.html", size: 177, mode: os.FileMode(436), modTime: time.Unix(1517691276, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _registerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x4d\x4e\xc4\x30\x0c\x85\xd7\xed\x29\x42\x0e\x30\xbe\x80\x93\x0d\xcc\x1a\x16\xb3\x61\x99\xb6\x1e\x25\x22\x7f\x8a\x3d\x82\xde\x1e\xa5\xa5\x82\x41\x48\x6c\x12\x39\xcf\xef\xb3\xf2\x8c\x0f\x4f\xcf\x8f\x97\xd7\x97\xb3\xf2\x92\xa2\x1d\xf1\xb8\xc8\x2d\x76\x1c\x50\x82\x44\xb2\x17\x62\x21\x84\xbd\x18\x07\x8c\x21\xbf\xa9\x46\xd1\x68\x96\x35\x12\x7b\x22\xd1\x4a\xd6\x4a\x46\x0b\x7d\x08\xcc\xcc\x5a\xf9\x46\x57\xa3\xc1\x31\x93\x30\x88\x9b\xfd\x5a\x32\x9f\x52\xc8\xa7\x4d\x07\x3b\x22\xec\x83\x70\x2a\xcb\xda\xc9\xd7\xd2\x92\x4a\x24\xbe\x2c\x46\xd7\xc2\xa2\xfb\x6b\xb5\xb9\x24\x52\x73\x49\x35\x92\x14\x9c\x1a\xd8\x71\x18\x30\xe4\x7a\x93\x1f\x73\xb5\xca\x2e\x91\xd1\xfd\xdc\xf8\x03\x42\xdd\x01\xe7\xe4\x42\xfc\xd7\x48\xbd\xeb\x97\xb3\x3a\xe6\xf7\xd2\x96\x3f\xcd\x87\x78\x00\xbe\xeb\x3b\xc6\x9d\x87\x6f\x53\x0a\xd2\x3b\xbe\x74\xe8\xbf\xee\x61\xec\x29\x20\x6c\x4b\xf8\x0c\x00\x00\xff\xff\x9b\xc0\x4c\x1e\x9b\x01\x00\x00")

func registerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_registerHtml,
		"register.html",
	)
}

func registerHtml() (*asset, error) {
	bytes, err := registerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "register.html", size: 411, mode: os.FileMode(436), modTime: time.Unix(1517772806, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _shippingHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xce\xc8\x2c\x28\xc8\xcc\x4b\x07\x04\x00\x00\xff\xff\x24\x17\x1c\x2d\x08\x00\x00\x00")

func shippingHtmlBytes() ([]byte, error) {
	return bindataRead(
		_shippingHtml,
		"shipping.html",
	)
}

func shippingHtml() (*asset, error) {
	bytes, err := shippingHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "shipping.html", size: 8, mode: os.FileMode(436), modTime: time.Unix(1517702087, 0)}
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
	"cart.html": cartHtml,
	"login.html": loginHtml,
	"payment.html": paymentHtml,
	"register.html": registerHtml,
	"shipping.html": shippingHtml,
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
	"cart.html": &bintree{cartHtml, map[string]*bintree{}},
	"login.html": &bintree{loginHtml, map[string]*bintree{}},
	"payment.html": &bintree{paymentHtml, map[string]*bintree{}},
	"register.html": &bintree{registerHtml, map[string]*bintree{}},
	"shipping.html": &bintree{shippingHtml, map[string]*bintree{}},
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

