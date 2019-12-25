package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _drop_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x29\x4a\xcc\x2b\x4e\x4c\x2e\xc9\xcc\xcf\x2b\x56\x70\x76\x0c\x76\x76\x74\x71\xb5\xe6\xc2\xaa\xb4\xb4\x38\xb5\x88\x90\x9a\xe2\xfc\xd2\xa2\xe4\x54\x84\x2a\x40\x00\x00\x00\xff\xff\x4e\xbc\xed\xb2\x74\x00\x00\x00")

func drop_sql() ([]byte, error) {
	return bindata_read(
		_drop_sql,
		"drop.sql",
	)
}

var _init_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\xcd\x8e\xda\x30\x10\xc7\xef\x79\x8a\x51\x2e\x2c\x52\xdf\x60\xb5\x87\x34\xf1\x4a\x51\x53\xd3\x3a\x8e\xd4\x3d\x21\x37\x19\x81\x45\x62\x47\xfe\x80\xd2\xa7\xaf\xc8\x07\x24\x40\xb5\x9c\xb8\xfa\xe7\x19\xcd\xff\x37\x13\x33\x12\x71\x02\x3c\xfa\x9a\x11\xf0\x16\x8d\x85\x97\x00\x00\x20\x94\x55\x08\x16\x8d\x14\x35\xd0\x15\x07\x5a\x64\xd9\x97\x9e\x6c\xfc\x89\x79\x2f\xab\x33\x81\x82\xa6\x3f\x0b\x32\x7c\x28\x0d\x0a\x87\x55\x08\xe0\x64\x83\xd6\x89\xa6\x75\x7f\xaf\xbb\x60\x23\x64\x1d\xc2\x5e\x98\x72\x2b\xcc\x35\x3d\x88\xba\x46\x17\x82\xf2\x0d\x1a\x59\x4e\x70\xc7\xe3\x15\xcd\x39\x8b\x52\xca\xfb\xa1\xd7\xed\x0e\x7e\xb0\xf4\x7b\xc4\x3e\xe0\x1b\xf9\x80\x97\x7e\xc8\x65\xb0\x7c\x0d\x82\x59\x46\xab\xbd\x29\xf1\x69\x29\xdd\xb1\xc5\xdb\x90\x50\x6e\xb1\xdc\x0d\x33\x4c\xfe\xbd\xc1\x62\x23\x1a\x5c\x9c\xdf\xb5\x99\x20\x8b\x66\x8f\xe6\x3f\xb0\x15\xc7\x06\x95\xeb\xe9\xf2\x56\xd3\x90\xfb\x71\x51\xce\x08\x65\x45\xe9\xa4\x56\x4f\xb3\x65\x9d\x70\x0f\xe9\x1a\x3e\xbe\xc1\xe2\x20\xd5\xdc\xc8\x05\xd5\xda\x5e\x7c\x74\x65\xa2\xd1\x5e\xdd\xbd\xaa\xbe\x6b\xe7\x68\x7d\x2f\x0b\x23\xef\x84\x11\x1a\x93\x7c\x34\x39\x9a\x83\x15\x85\x84\x64\x84\x13\x60\x24\xe7\x2c\x8d\xf9\xd0\xee\x74\x99\x9f\x35\xeb\xae\xf7\x4e\xab\x38\xca\xe3\x28\x19\xe5\x55\x5a\x61\x08\xbf\xb5\xbe\x78\x87\x84\xbc\x47\x45\xc6\xc1\x19\x8f\xb7\xeb\x9e\x6e\xef\x91\x9d\xa7\x34\x21\xbf\xe6\x55\xb2\xfa\x73\xaa\xd4\x6a\xf6\x3c\x16\xbf\xfe\x0b\x00\x00\xff\xff\x2d\xe4\x9d\x11\x3a\x04\x00\x00")

func init_sql() ([]byte, error) {
	return bindata_read(
		_init_sql,
		"init.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"drop.sql": drop_sql,
	"init.sql": init_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"drop.sql": &_bintree_t{drop_sql, map[string]*_bintree_t{}},
	"init.sql": &_bintree_t{init_sql, map[string]*_bintree_t{}},
}}
