// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package views generated by go-bindata.// sources:
// web/views/errors/404.html
// web/views/errors/500.html
// web/views/index.html
// web/views/layouts/footer.html
// web/views/layouts/header.html
// web/views/layouts/layout.html
package views

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
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
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _errors404Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xc1\x6a\xc3\x30\x0c\x86\xef\x79\x0a\xcd\xe7\x79\x69\xa0\x87\x1d\xec\xc0\xe8\x5a\xd8\x65\xdb\xa1\x85\xed\xe8\xda\x3f\xb5\xc0\xb1\xb3\x54\x4d\xd9\xdb\x8f\x34\x1d\x74\x3b\x09\x7d\xfa\x3f\x81\x64\xee\x9e\xdf\x56\xdb\xcf\xf7\x35\x45\xe9\x52\x5b\x99\xa9\x50\x72\xf9\x60\x15\xb2\x9a\x00\x5c\x68\x2b\x22\x22\xd3\x41\x1c\xf9\xe8\x86\x23\xc4\xaa\xdd\x76\xa3\x1f\xd5\xed\x28\x8a\xf4\x1a\x5f\x27\x1e\xad\xfa\xd0\xbb\x27\xbd\x2a\x5d\xef\x84\xf7\x09\x8a\x7c\xc9\x82\x2c\x56\xbd\xac\x2d\xc2\x01\x7f\xcc\xec\x3a\x58\x35\x32\xce\x7d\x19\xe4\x26\x7c\xe6\x20\xd1\x06\x8c\xec\xa1\x2f\xcd\x3d\x71\x66\x61\x97\xf4\xd1\xbb\x04\xdb\x3c\x2c\x7e\x57\x09\x4b\x42\xbb\x5c\x2c\xe9\xb5\x08\x6d\xca\x29\x07\x53\xcf\xb0\x32\xf5\x7c\x88\xd9\x97\xf0\x7d\xcd\xc7\xa6\x35\x1e\x59\x30\xfc\x97\xae\xd4\xd4\xb1\x99\xd4\xd9\x31\xf5\xe5\x47\x3f\x01\x00\x00\xff\xff\xf0\x70\x71\x97\x33\x01\x00\x00")

func errors404HtmlBytes() ([]byte, error) {
	return bindataRead(
		_errors404Html,
		"errors/404.html",
	)
}

func errors404Html() (*asset, error) {
	bytes, err := errors404HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "errors/404.html", size: 307, mode: os.FileMode(420), modTime: time.Unix(1651999235, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errors500Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x4e\xc3\x30\x0c\x86\xef\x79\x0a\x93\xf3\x42\xd9\x8d\x43\x52\x09\x8d\x21\x71\x1a\x87\x4d\x82\xa3\xd7\x5a\x8d\xa5\xd4\x29\xad\xb7\x6a\x6f\x8f\xba\x0c\x69\xdc\x38\x25\xbf\xfc\xf9\xb3\x6c\xff\xf0\xba\xdb\xec\xbf\x3e\xb6\x10\xb5\x4f\xb5\xf1\xcb\x03\x09\xa5\x0b\x96\xc4\xd6\xc6\xf8\x48\xd8\xd6\x06\x00\xc0\xf7\xa4\x08\x4d\xc4\x71\x22\x0d\xf6\xb0\x7f\x73\xcf\xf6\xbe\x14\x55\x07\x47\xdf\x27\x3e\x07\xfb\xe9\x0e\x2f\x6e\x93\xfb\x01\x95\x8f\x89\x2c\x34\x59\x94\x44\x83\x7d\xdf\x06\x6a\x3b\xfa\xd3\x29\xd8\x53\xb0\x67\xa6\x79\xc8\xa3\xde\xc1\x33\xb7\x1a\x43\x4b\x67\x6e\xc8\x5d\xc3\x0a\x58\x58\x19\x93\x9b\x1a\x4c\x14\xd6\x8f\x4f\xbf\x2a\x65\x4d\x54\xef\x4e\xc3\x04\x53\xee\x49\x23\x4b\x07\x33\x89\xc2\x3c\x66\xe9\x56\xa0\xe3\x05\xb0\x43\x16\x5f\x15\xd6\xf8\xaa\xec\x67\xfc\x31\xb7\x97\x9b\x27\xae\xcb\xe7\x1a\x1a\x12\xa5\xf1\x9f\xd6\x1b\x5c\x34\xd5\xe2\xf1\x55\x11\x2f\x93\x96\x13\xff\x04\x00\x00\xff\xff\x89\x94\x40\x58\x72\x01\x00\x00")

func errors500HtmlBytes() ([]byte, error) {
	return bindataRead(
		_errors500Html,
		"errors/500.html",
	)
}

func errors500Html() (*asset, error) {
	bytes, err := errors500HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "errors/500.html", size: 370, mode: os.FileMode(420), modTime: time.Unix(1651999235, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\x2c\x2a\xc9\x4c\xce\x49\x55\x48\xce\x49\x2c\x2e\xb6\x55\xca\x4d\x2c\xca\x4e\xc9\x2f\xcf\xd3\x4d\xca\x4f\xa9\x54\xb2\xe3\x52\x50\x50\x50\xa8\xae\xd6\x73\x84\xa8\xaa\xad\xe5\xb2\xd1\x87\xea\xb0\x03\x04\x00\x00\xff\xff\x82\x8b\x94\x11\x3b\x00\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 59, mode: os.FileMode(420), modTime: time.Unix(1668859849, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutsFooterHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func layoutsFooterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_layoutsFooterHtml,
		"layouts/footer.html",
	)
}

func layoutsFooterHtml() (*asset, error) {
	bytes, err := layoutsFooterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layouts/footer.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(1667724476, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutsHeaderHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x41\x6a\xc4\x30\x0c\x05\xd0\x7d\x4f\x21\xb4\x17\xd9\xb7\x8e\xef\xa2\x49\x64\x47\xa9\x62\x0d\xb6\x63\x98\xdb\x97\x52\x3a\x84\x42\x61\x96\x12\xff\xff\x17\x56\x1d\xb0\x18\xb7\x36\xe3\xcd\xfd\x93\x36\xe1\x55\x2a\x82\xae\x7f\x1e\xd5\x4d\x66\x2c\x3c\x34\x73\x57\x2f\x18\xdf\x00\x00\x02\x3f\xeb\xbd\xc0\xfd\x34\x23\x93\xd4\x61\x6f\xd4\xdd\xed\xc6\x95\x78\xf9\xce\x43\xf7\x9c\x4d\xa8\x9d\xc7\xc1\xf5\x81\xb0\x55\x49\x33\xee\x3c\xb8\x2d\x55\xef\xfd\xfd\x03\x63\xd0\xdf\xb5\xc4\x90\x98\xd8\x34\x17\xda\xcf\xd6\x35\x3d\x30\x86\x49\x63\x98\xf8\x5f\xb9\x6a\xde\x7e\xe8\x4d\x0e\x79\xc2\x97\xe3\x35\xb6\x9d\x85\xfc\xc2\x85\x69\xd5\x11\xbf\x02\x00\x00\xff\xff\x39\xdf\x3f\x2c\x2c\x01\x00\x00")

func layoutsHeaderHtmlBytes() ([]byte, error) {
	return bindataRead(
		_layoutsHeaderHtml,
		"layouts/header.html",
	)
}

func layoutsHeaderHtml() (*asset, error) {
	bytes, err := layoutsHeaderHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layouts/header.html", size: 300, mode: os.FileMode(420), modTime: time.Unix(1667728423, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutsLayoutHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x5f\x6f\xd4\xba\x12\x7f\xde\x4a\xfd\x0e\x43\xee\x43\x77\xa5\x26\xa1\x3c\x5d\xd1\x64\x11\xb4\x5c\x54\x09\xc1\xbd\x50\xa4\x7b\x84\xaa\xa3\xd9\x78\x36\x71\xeb\xd8\x39\xb6\x77\x97\x25\xec\x77\x3f\xb2\x9d\x6c\xb2\xb4\x45\x70\x78\x8a\xed\xf9\xcd\x1f\xff\xc6\x33\x93\xec\xc9\xe5\xfb\x8b\xeb\x3f\xfe\xfb\x1a\x2a\x5b\x8b\xf9\xf1\x51\xe6\xbe\x20\x50\x96\x79\x44\x32\xf2\x27\x84\x6c\x7e\x7c\x34\xc9\x6a\xb2\x08\x45\x85\xda\x90\xcd\xa3\x4f\xd7\xff\x89\xff\x1d\x0d\x82\xca\xda\x26\xa6\xbf\x56\x7c\x9d\x47\xff\x8f\x3f\xbd\x8c\x2f\x54\xdd\xa0\xe5\x0b\x41\x11\x14\x4a\x5a\x92\x36\x8f\xae\x5e\xe7\xc4\x4a\x1a\xe9\x49\xac\x29\x8f\xd6\x9c\x36\x8d\xd2\x76\x04\xdd\x70\x66\xab\x9c\xd1\x9a\x17\x14\xfb\xcd\x29\x70\xc9\x2d\x47\x11\x9b\x02\x05\xe5\x67\xc9\xd3\x60\xc8\x72\x2b\x68\xde\xb6\x90\x5c\xbb\x15\xec\x76\x59\x1a\xce\x8e\x8f\x9c\x5c\x70\x79\x07\x9a\x44\x1e\x19\xbb\x15\x64\x2a\x22\x1b\x01\x67\x79\x64\x2b\xaa\x29\x2e\x8c\x89\xa0\xd2\xb4\xcc\xa3\xd4\x58\xb4\xbc\x48\x0b\x63\xd2\x92\xdb\x6a\xb5\x88\x6b\xd4\x77\x4c\x6d\xa4\x83\xa5\x0c\xf5\x5d\xe2\xf0\xf3\xc7\x0d\x3f\x68\x69\xa1\xd4\x5d\xec\xdd\xa5\x1e\xfb\xbb\x46\x36\xb4\x30\xdc\xfe\xb6\x99\xa2\xc2\xc6\x92\x8e\x97\x4a\xb0\xdf\xbe\x57\x23\xb8\xb5\xa4\x7f\xdd\x4e\xc5\xcb\x4a\xf0\xb2\xb2\x9d\x25\x3c\x3b\xdb\xc6\xff\x8c\xea\x1a\xb9\xec\xb5\x8e\x8f\x26\x6d\x0b\x7c\x09\xc9\x1b\x6e\x51\xdc\x25\x17\x82\x93\xb4\x57\x97\xb0\xdb\xfd\xe2\x35\x51\xdc\x75\x9f\x7d\x48\x6d\x4b\x92\x39\x43\x7e\xed\xbc\xbc\x94\x28\xb6\x5f\x49\x27\xaf\x90\xb3\x55\xf0\xf1\x24\x8e\xc1\x6f\x01\x9d\xd4\xf2\xc2\x40\x1c\xfb\x2b\x99\x42\xf3\xc6\xba\xe5\x64\x8d\x1a\xfe\xac\x6a\x0b\x79\xf8\x7c\xfb\x06\x9f\x6f\xce\x9d\x64\xba\x5c\xc9\xc2\x72\x25\xa7\x33\x68\xdd\x01\x80\x03\x57\x35\xe4\xc0\x54\xb1\xaa\x49\xda\xa4\xd0\x84\x96\x5e\x0b\x72\xbb\x69\x14\x0c\x47\xb3\xf3\x80\xaf\xea\xc4\xe8\x02\x72\x88\x5c\x91\x9a\xe7\x69\x5a\xd5\xc9\xc2\xc5\x94\x14\xaa\x76\x9b\x5b\xf3\xa2\x6d\xef\x85\x1f\x9d\x0f\xfe\xcc\xd8\x5d\x49\xb6\xf3\x65\x5e\x6d\xaf\xb1\x7c\x87\x35\x0d\x5e\x3f\x3f\xbd\xe9\x14\x4d\xd2\xa0\x26\x69\xdf\x29\x46\x09\x97\x86\xb4\x7d\x45\x4b\xa5\x69\x5a\xd5\xa7\x60\x42\x7c\xbb\xd9\xd4\x2f\xb2\x74\xe0\xe3\x07\xd4\xbe\x51\xaa\x14\x34\x70\x1b\xf6\x60\xb1\x84\x69\x69\xb1\x4c\x6e\xcd\xec\x90\x60\x40\xb3\x95\x05\x18\x5d\xe4\x7b\x02\x36\x9b\x4d\x52\x7a\x4d\x8b\x65\x8d\x12\x4b\xf7\x6c\x55\x9d\x3a\x13\xe9\xad\x79\xc1\x59\x3e\x66\xa4\xf7\x1a\xcd\xc7\x61\x8e\x33\xb8\xe1\x92\xa9\x4d\xc2\xd0\xe2\x5b\xdc\x92\x86\x1c\xee\x1d\x0d\x59\xed\x93\x0a\xce\xdf\x74\xd6\xee\x31\x49\xb3\x32\xd5\x14\x75\xe9\x99\x36\xb3\x73\x77\xd1\x89\x47\x9d\xdc\x9a\x93\x53\x90\xb4\x81\x4b\xb4\x34\x9d\x39\xd6\x06\x61\xa1\xe4\x92\x97\x27\xa7\x70\xf2\x50\xdc\x27\x3f\xa2\x38\x4b\xbb\x1e\x9f\x2d\x14\xdb\xfa\x8b\x31\xbe\x86\x42\xa0\x31\x79\xe4\x0a\x1c\x96\x4a\xda\xd8\xf0\xaf\x14\x3f\x0b\xeb\x25\xd6\x5c\x6c\xe3\x33\x28\x94\x50\x3a\x94\x6d\xfc\xcc\xd7\xc5\x3d\xed\xd8\xac\xea\x1a\xf5\x36\x48\x27\x59\x41\xd2\x92\x0e\x9b\x49\xd6\xf4\x50\xa1\x4a\x15\x81\x56\x82\xba\xf5\x3c\xc3\xae\x18\xa3\xef\x9a\x3b\xce\xb3\xb4\xe9\xac\xa5\x63\x73\x99\xc4\x75\x67\x42\xe2\x9a\x97\xe8\x48\x8e\x7a\x57\x2b\xd1\xfb\x3a\x8c\x68\x32\x69\x5b\x8d\xb2\x24\x48\xde\xe1\xda\xbf\xad\x70\x08\x96\xea\x46\xa0\x25\x70\xe6\x4c\xe2\xa6\x63\x04\xc9\x08\xd1\x91\xe8\x77\x99\xe0\xbd\x7d\xc6\xd7\x9c\x91\x76\xcf\x45\xf0\xde\x7d\xba\x12\x7d\xcc\x12\xd7\x81\xaa\x94\xf1\xf5\xc3\xa4\xb9\x5c\xf4\x8c\x1d\xc8\xd8\x36\xe6\x52\x3a\xe3\xc1\x6e\xdb\x82\x26\xc9\x48\x43\x24\x70\xab\x56\xd6\xf8\x7c\x92\xee\xc2\xed\xe3\x1b\x5b\x69\xb0\xa4\x78\xa3\xb1\x69\x48\x47\x60\x71\xc1\x25\xa3\x2f\x79\x14\x9f\xf5\x19\x70\x7d\x74\x4f\xcf\x3d\xd5\x83\x00\x26\xa1\x12\x3f\x84\x20\x6c\x45\x50\xac\xb4\xab\xfb\x81\xbe\x8a\x34\x75\x55\xd9\x53\xbb\xe5\x24\xd8\x3e\xb8\xce\x87\x9b\xc9\xa1\xcd\x3a\xea\x7a\x6e\xbc\xb8\x7f\x0b\xff\xf2\xec\x84\x1b\x46\x7d\x50\x43\xb2\xe3\x52\xc5\x56\x35\xd1\x3c\xdb\x27\x63\x89\xb0\xc4\x18\xb5\x56\x9b\x78\xe5\x24\x29\x9f\xbb\x37\xd4\x5f\x6e\xe4\x67\xb4\x1e\xa5\x66\x58\x2d\x95\x72\x6f\xed\x01\xca\x83\x64\x4f\x79\xd6\x1d\xf8\x6a\xea\xf4\xb3\xb4\xab\xaf\xbe\x31\xf9\x96\xd4\x0f\x9a\xdb\xd1\x18\x4c\x6e\xcd\x41\xa7\x79\x4c\xe1\xf6\x7f\x2b\xd2\xdb\x9f\x45\x1f\xcc\xfa\x9f\xd4\xd9\xcf\xf4\x9f\xc4\xfb\xf1\xfb\x1d\xf6\x07\x33\xf8\x31\x33\xdd\xa8\xad\xef\x1b\x1b\xf5\xdc\x42\x49\x63\x21\x40\x21\xf7\x9d\x31\xf8\x98\xfa\x41\x59\x74\x8e\x9e\xfb\x96\xf8\x9d\xf7\xdd\xee\xe4\x74\x00\x7d\xa4\x42\x93\x7d\x00\x18\x04\x3d\x58\x53\xa3\x0e\x40\x1f\xa8\x51\xbd\x50\x6d\x24\xe9\x03\xe9\x7b\x77\xd2\x8b\x91\xd5\x5c\x3e\x87\xcf\x0f\xc8\x6f\x3c\x80\xb3\x03\xe5\x2b\xd6\x6b\x0a\x5c\x90\x30\x4e\x35\x5c\xf5\xe4\xe6\xf8\x68\xb2\x9b\xf9\xe6\xdf\xf1\x14\x1e\xe3\xb4\x07\xcc\xdc\x5b\x1b\xd1\x3f\xee\xf4\xe1\x3f\xdf\x27\x85\xd1\x92\xcb\xc3\xd6\xe6\x73\x32\xb4\xb1\xee\xc5\x80\x1b\x29\x85\xe5\x6b\x3f\x01\xdd\x8b\xee\x8b\xb1\x6d\x93\xb7\x5c\xde\x75\xc7\x93\xb6\x4d\x3e\x56\x6a\xe3\xfe\x07\x42\x61\x87\xe1\x7d\x65\x2e\xb9\xde\xed\x86\x8a\xa4\x2f\x45\x6c\x35\x2f\x4b\xd2\xb0\xc4\x50\x91\x43\x2f\xed\x6a\x33\xe8\x5e\x54\x5c\x30\x4d\x32\x48\x86\x0e\x8e\xda\xf2\x42\x90\xe9\x1d\x77\xfd\xfb\x00\xfe\x78\x07\x87\x5e\x3e\xf8\x0c\xcd\x79\x7f\xd2\x35\xee\x6e\xff\x77\x00\x00\x00\xff\xff\xbd\x04\x13\x76\x29\x0d\x00\x00")

func layoutsLayoutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_layoutsLayoutHtml,
		"layouts/layout.html",
	)
}

func layoutsLayoutHtml() (*asset, error) {
	bytes, err := layoutsLayoutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layouts/layout.html", size: 3369, mode: os.FileMode(493), modTime: time.Unix(1669530533, 0)}
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
	"errors/404.html":     errors404Html,
	"errors/500.html":     errors500Html,
	"index.html":          indexHtml,
	"layouts/footer.html": layoutsFooterHtml,
	"layouts/header.html": layoutsHeaderHtml,
	"layouts/layout.html": layoutsLayoutHtml,
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
	"errors": &bintree{nil, map[string]*bintree{
		"404.html": &bintree{errors404Html, map[string]*bintree{}},
		"500.html": &bintree{errors500Html, map[string]*bintree{}},
	}},
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"layouts": &bintree{nil, map[string]*bintree{
		"footer.html": &bintree{layoutsFooterHtml, map[string]*bintree{}},
		"header.html": &bintree{layoutsHeaderHtml, map[string]*bintree{}},
		"layout.html": &bintree{layoutsLayoutHtml, map[string]*bintree{}},
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
