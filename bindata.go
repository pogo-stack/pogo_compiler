// Code generated by go-bindata.
// sources:
// templates/template_psp_function_begin.sql
// templates/template_psp_function_begin_raw.sql
// templates/template_psp_function_begin_view.sql
// templates/template_psp_function_drop.sql
// templates/template_psp_function_end.sql
// templates/template_psp_function_end_raw.sql
// templates/template_psp_function_end_view.sql
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

var _templatesTemplate_psp_function_beginSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x51\x41\x6e\xe3\x30\x0c\x3c\xc7\xaf\xe0\x41\x80\x6c\x20\xd8\xcd\xee\x31\xd9\xed\x47\xd2\x80\x90\x65\xc6\x76\x62\x4b\x06\x45\x05\xed\xef\x0b\xaa\x8d\xd1\xf4\x94\x8b\x31\x1c\x8e\x35\x1c\xf2\x55\x2a\xcf\xe4\x84\xe0\x9c\x83\x97\x31\x06\x30\x77\x64\xa0\x36\x8b\x63\x37\x93\x10\x27\xd3\x54\x4c\x92\x39\x24\x30\x5f\xc0\x80\x4b\x60\x4c\xd5\x91\x9f\x1c\x53\xb5\x41\x8f\x20\xe3\x4c\x49\xdc\xbc\xc0\xfe\x3f\xf8\x29\xfa\x2b\xae\x54\xdd\x1c\xaa\x0d\x0a\x82\xd0\x9b\x28\x9c\x10\xc6\x20\xd4\x13\xab\x7a\xa7\xd4\x0d\xe1\xe6\xd8\x0f\x8e\x8f\x27\xad\xc3\x83\xe4\x8f\x52\xdd\x03\x65\x3a\x6a\x73\x6f\x4a\x23\x21\x5c\x52\x0c\xad\xf2\x05\x60\x9b\xc7\xa9\xc3\xd8\x5e\xc8\x8b\xda\x03\x00\xa8\xae\x7d\x52\x7a\xd6\xfc\xd8\xd1\x22\xc3\xea\x5a\x42\x0c\x4c\xae\xc3\xb1\x5b\xc7\xd5\x87\x6a\x1f\xdd\x44\xc9\x53\x1d\xf2\x34\x8d\xe7\xda\x67\x66\x0a\x82\x89\x44\xc6\xd0\xd7\x16\x71\x89\x7d\xfc\xb5\xfe\x6e\xb7\x20\x9c\xa9\xd9\x82\xb5\xfa\xf9\x17\x7e\xbb\x17\xdb\x94\x45\x0d\x0f\x39\xff\xee\xca\x7e\x88\xf9\x87\xa9\xb5\x65\xa0\xe7\x76\x3f\x7f\x8b\x6d\x8f\x27\xbb\xdf\x97\x52\x7b\x8c\xa0\xb3\x21\x7e\x9e\x17\x51\xde\x17\x3a\x54\x9b\x9b\x7a\x46\xc6\x24\xce\x5f\xcb\xe9\xee\xae\x1f\x01\x00\x00\xff\xff\xd5\x41\xfa\xb1\x3e\x02\x00\x00")

func templatesTemplate_psp_function_beginSqlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTemplate_psp_function_beginSql,
		"templates/template_psp_function_begin.sql",
	)
}

func templatesTemplate_psp_function_beginSql() (*asset, error) {
	bytes, err := templatesTemplate_psp_function_beginSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/template_psp_function_begin.sql", size: 574, mode: os.FileMode(438), modTime: time.Unix(1559978435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTemplate_psp_function_begin_rawSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8d\xc1\x4a\x04\x31\x0c\x86\xcf\xf6\x29\x72\x28\xb4\x85\x45\xef\xae\xfa\x24\x42\xc8\xb4\x99\x99\x2e\x35\x1d\xd2\xd4\xe7\x97\x61\xd5\xeb\x5e\xc2\x77\xf8\xf2\xfd\x9f\xe6\xb2\x32\x19\xc3\x3a\x25\x5b\xed\x02\xfe\x8f\x3c\x44\x7f\x90\xd2\x17\x1b\xeb\xf0\xc9\x29\xdb\x54\x19\xe0\x7f\xc1\x03\x0d\xf0\xde\x15\xce\x8d\x94\xdd\x13\x96\x81\x70\x1b\x5d\x16\x78\x7d\xbf\x03\x2e\xb3\xb6\x82\x7d\xb9\x71\xb6\x98\xae\x77\x69\x79\xec\x01\x00\xe0\x7a\xae\x63\xe1\xc3\x76\x84\x2a\xc6\x1b\xeb\x99\xb0\x5d\x99\x0a\xd6\x82\xf0\x4d\x9a\x77\xd2\x33\x14\x73\xa7\xc6\x23\x73\x94\xd9\x5a\x5d\x63\x9e\xaa\x2c\x86\x83\xcd\xaa\x6c\x31\x20\x1e\x7d\xeb\xcf\xff\xef\xe1\x02\xa6\x93\xd3\x05\x42\x38\xcf\x9b\xbc\xd0\x47\x48\xe9\xea\x7e\x02\x00\x00\xff\xff\x44\x52\x60\x81\x19\x01\x00\x00")

func templatesTemplate_psp_function_begin_rawSqlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTemplate_psp_function_begin_rawSql,
		"templates/template_psp_function_begin_raw.sql",
	)
}

func templatesTemplate_psp_function_begin_rawSql() (*asset, error) {
	bytes, err := templatesTemplate_psp_function_begin_rawSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/template_psp_function_begin_raw.sql", size: 281, mode: os.FileMode(438), modTime: time.Unix(1559978435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTemplate_psp_function_begin_viewSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\x29\xe1\x4a\x29\xca\x2f\x50\x28\xcb\x4c\x2d\x57\xc8\x4c\x53\x48\xad\xc8\x2c\x2e\x29\x56\x50\x49\x2b\xcd\x4b\x2e\xc9\xcc\xcf\x53\xb1\xe6\x4a\x2e\x4a\x4d\x2c\x49\x85\xa8\x40\x88\x73\x25\x16\x73\x01\x02\x00\x00\xff\xff\x4f\xab\x4a\x90\x3d\x00\x00\x00")

func templatesTemplate_psp_function_begin_viewSqlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTemplate_psp_function_begin_viewSql,
		"templates/template_psp_function_begin_view.sql",
	)
}

func templatesTemplate_psp_function_begin_viewSql() (*asset, error) {
	bytes, err := templatesTemplate_psp_function_begin_viewSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/template_psp_function_begin_view.sql", size: 61, mode: os.FileMode(438), modTime: time.Unix(1559978435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTemplate_psp_function_dropSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcc\xc1\x0d\xc3\x30\x08\x05\xd0\x7b\xa7\xe0\x80\xe4\xf6\xda\xab\x87\x41\xad\x01\xcb\x52\xcc\x47\x4e\xb2\x7f\x16\xc8\x29\x0b\x3c\x05\x31\xd3\xdf\xfa\x08\x4a\x5b\x8e\x35\xc9\x45\x17\x52\x12\x1d\xd2\x30\x73\x6c\xa6\xe2\x67\xb4\x63\x20\xde\x25\xf7\xfc\x0a\xc7\x6f\x1a\x97\x4f\x25\x0b\x25\xe6\xfa\x7a\x20\xf9\x0d\x73\x05\x00\x00\xff\xff\x35\xe6\x68\xa2\x91\x00\x00\x00")

func templatesTemplate_psp_function_dropSqlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTemplate_psp_function_dropSql,
		"templates/template_psp_function_drop.sql",
	)
}

func templatesTemplate_psp_function_dropSql() (*asset, error) {
	bytes, err := templatesTemplate_psp_function_dropSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/template_psp_function_drop.sql", size: 145, mode: os.FileMode(438), modTime: time.Unix(1559978435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTemplate_psp_function_endSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x4d\x8f\xdb\x38\x12\x3d\x8b\xbf\xa2\x62\x68\x41\xc9\xb1\xe3\xf6\x2e\x72\x88\xd4\xf6\x22\x18\xcc\x25\x18\x24\x87\x99\xc3\x00\x49\x50\xa1\xa5\xb2\xcc\x34\x4d\x2a\x24\xe5\xb4\x81\xfe\xf1\x03\x92\xf2\x47\xd2\x9d\x00\x93\x43\x9b\x54\x91\xac\x7a\xf5\x8a\xaf\xd8\x72\x0b\xd8\x22\xac\xe1\x06\xfc\x8e\x34\xcb\xf0\x80\xef\x51\xe3\x47\xa8\x56\xc0\x6f\x9f\xcd\xe7\x40\xba\xad\x20\xd7\x62\x4f\x39\xcc\xe7\x6b\x5e\x03\x6a\x0c\xcb\x61\x78\x0e\xcb\xfa\xa9\x43\x1c\x1e\x1e\xa0\x68\x94\x69\xee\xd0\xcb\x3d\x39\x2f\xf6\x7d\x51\xc2\x1c\xb0\xc1\x12\xaa\x0a\xa4\xf6\x64\x0f\x42\x85\x8d\xfc\x47\x8e\x03\x3c\xbf\x47\x78\xb6\x02\xfe\xfe\x23\xaf\xaa\xcf\xce\xe8\xcd\x08\x35\x8b\x4b\x61\x7f\x18\x1f\x1e\x20\x2e\xe2\x66\x90\xaa\x45\xb3\xf9\x4c\x8d\x2f\xf8\x1f\x52\x53\x4c\x81\xcf\xa0\x78\x0c\x68\x8e\xbe\xc1\xb2\xaa\x4e\x68\xca\xb2\x0e\x8e\x9f\xce\x27\xf9\xef\x2d\x79\x7f\x2c\x42\xd0\xf2\xa7\xe0\x49\xb7\x20\xb7\x35\x3b\x8d\xd8\xc4\x0d\x8f\x30\xd4\x8c\xb1\xad\xb1\x20\x41\x6a\x58\xc2\x8b\x17\xd1\x87\x32\xa6\x4f\x04\x1c\xf0\xbd\xfc\x08\xd2\x81\x1e\x94\x8a\xb9\x9f\x6c\x01\x1e\xaf\xe1\x3a\x50\x38\x56\x33\x86\x3e\xc6\x12\xd6\x8a\x23\x7a\x83\xce\x5b\xa9\xbb\x02\x0f\x38\x03\xce\x43\x48\x47\x8a\x1a\x0f\xb8\xc3\x19\xa0\x47\x96\x2d\xa6\x21\x99\x7f\x59\x3d\x96\x45\x06\x16\xe9\x08\x6a\xfc\xd6\xe0\xad\xdc\x17\xde\x60\xb3\x13\xb6\x50\xa4\x3b\xbf\x2b\xd0\x63\x39\x03\xfe\xea\xd5\xab\xd9\xd5\x1f\x2f\xaf\xc8\x9c\x2e\x98\xd4\xde\x00\x5a\xac\x99\x25\x3f\x58\x9d\xe6\x8c\xee\x1b\xea\xbd\x34\x1a\xbe\x06\x1e\x8c\xdf\x91\x75\xe3\x7d\xe8\xc8\x83\xf3\xa2\xb9\xa3\x16\x5a\x29\x3a\x6d\x9c\x97\x8d\x83\x03\x92\xb5\xc6\x62\x5c\x83\x15\xf4\x1d\x9e\xdd\x60\x63\xb4\xa7\x7b\x5f\x33\x96\xe1\x2e\x72\x76\x13\xe6\x89\xfd\x93\x29\x0c\xa9\xa6\x19\xdd\x4b\x9f\x62\x07\xe3\x1a\xfe\x1b\x8c\x1b\xea\x64\xb8\x90\x99\xd4\x8e\xac\x87\x04\x1e\x7b\xd3\x19\x4c\xc1\x5d\x58\x2d\xc2\x4f\x26\xdb\x59\x1c\x03\xb1\x18\x99\x4d\xdf\x83\x23\x8b\xa7\xc5\xed\xa0\x9b\x88\x2f\xe8\x2e\x99\x7a\x61\xc5\xde\xa5\xb9\x92\x9a\x50\x0f\xfb\x0d\xd9\x64\x88\xb9\x8d\xd3\x2f\x2a\xb8\xf5\x74\xf9\x8c\x10\xc2\x57\x19\x7e\x0e\x42\x0d\x74\x05\x68\xdf\xbe\x7c\x5c\xeb\xaa\x3a\x08\x1b\xca\x16\xeb\xaa\xb0\x4c\xde\x9a\xc1\x5a\xd2\xfe\xb2\x33\x99\x97\x69\xe0\xf9\x09\x77\xce\x93\x25\xdf\x1a\xbb\xcf\xd3\x1c\x15\xa6\xc9\x37\x05\x39\x07\x3a\xc3\xfd\x16\x3c\x59\xbb\x3f\x43\x4f\x77\x41\xea\x0e\x64\x3b\xb2\x4c\xd6\xca\x16\x43\x15\x62\x6d\xc6\xc9\x8f\x6e\x49\xd4\x50\xac\xa3\x6e\x47\x89\x8e\x8a\xc9\xe6\x73\x78\xfb\xee\xaf\xdf\x2b\xf8\x4d\x68\xee\xa1\xd9\x09\xdd\x11\x48\x0f\xde\xc0\x9b\x3f\xdf\xbd\x85\x1d\x59\x02\x27\x75\x13\xad\x5f\xa5\x52\xb0\xb1\x24\xee\x82\x63\xd0\x46\xc3\xeb\x37\xaf\xff\x86\xc1\x51\x23\x1c\xb1\x6c\x14\xd8\xcb\x9b\x9b\x19\xf0\xdb\xde\x12\x34\x4a\x38\xb7\x9a\xf4\x96\xe6\x47\x52\xca\x7c\x9d\x47\x1e\x26\xe0\xfc\x51\xd1\x6a\xb2\x11\xcd\x5d\x67\xcd\xa0\xdb\x79\x63\x94\xb1\x55\xda\x55\x43\xfa\xea\xa5\xbe\xab\x27\xeb\xe9\x74\x0a\xf1\x1c\xec\x84\x03\xd3\xc4\x92\xb4\xc0\xd9\x7c\x9e\x25\xf1\x49\x0d\x9f\x2e\x95\xf8\x04\xe1\xb6\x8c\x02\x55\x78\xd9\x16\x79\x4e\xf6\x0b\xcf\x51\xf8\xc2\x51\x62\x6e\x19\x7c\x15\xcb\x32\x71\xc7\x43\xe8\x5b\x01\x5e\xd8\x8e\xfc\x6a\x82\x1b\x25\xf4\xdd\x04\x76\x96\xb6\xab\x49\xc4\xf4\xff\x1e\x65\xbb\x4a\xb1\x52\x65\xa2\xaa\x27\xeb\x96\xbc\x90\xca\xdd\x2e\xc4\x9a\x03\x29\x47\xc0\x79\x68\x5c\x25\x1b\xd1\x4c\xa7\xd3\xdb\x45\x6f\x69\xcd\x19\xcb\x2e\xea\x67\xd9\xb5\xfe\x63\xd5\xf2\x9c\x29\xa1\xbb\x41\x74\x04\xbd\xea\x3b\xf7\x45\x41\x7e\x30\x4a\x78\xa9\xa4\x3f\x62\x23\x3c\x75\xc6\x1e\xf3\x9a\xb1\x0f\x9e\xb1\x96\x14\x79\x82\xad\x35\xfb\xb3\x28\x3d\x39\x1f\xd2\x74\x21\x4f\x4b\xd0\x8b\x8e\xa2\xd6\x60\x05\x3c\x3d\x76\xbc\x66\xec\x29\x3d\x9f\x8f\x9e\x5a\x68\x2b\x9d\x97\xba\xf1\xe7\x83\x33\x18\xb4\x26\xe7\x0b\x4b\x1d\xdd\xf7\xe8\x7a\x25\x7d\xe8\xc2\xb1\x1d\x17\xf9\xd9\x43\x3e\x03\x3e\xe3\xe1\xcd\x79\x12\x64\x4b\x3d\xe9\x96\x74\x23\x7f\x05\xe7\xf5\xe9\x5f\x86\x7a\xed\xe4\x1a\xad\x81\x3c\x67\x63\xdb\x7b\x2a\x78\x63\xf6\xbd\x54\xd4\xa2\x33\x83\x6d\x08\x8a\xd8\xc7\x40\x89\x90\x3b\x1e\xc8\x3a\x69\xf4\x0c\xa4\x43\x6d\xc4\xe0\x77\x25\x1b\xdb\x12\x14\x17\x68\xcb\x19\xe4\xe7\x1d\x79\x59\xff\xb4\xfb\x0f\x7d\x1b\xae\xf3\x0f\x00\x04\x39\xfa\xef\xc2\xc3\xea\x7b\xc3\x73\x58\xb2\x2c\x11\xfd\x88\xe3\xf1\xea\xd5\x8c\x2d\xa6\x29\xfd\x96\x1a\x25\x2c\xb1\xcc\x43\x78\x42\xe2\xdb\xbe\xb3\xc5\xf2\x7f\xe1\x7f\x08\x0b\x09\x47\xba\xbd\x88\xfe\xd8\x53\x7d\x62\xcc\x86\xbd\xbd\xeb\x31\x79\x2f\x78\x54\xcd\x32\xa4\x1c\xce\x46\x57\x3e\x08\x67\x7c\x32\xed\x0b\x4b\xae\x37\xda\x11\x86\x48\x65\x55\xc5\x80\x0f\x0f\xa7\x80\x57\xd3\x10\x5b\x48\x47\x20\xf5\xd6\x00\xff\x0f\x9f\x81\xbf\xa0\x9f\x2e\x18\xcb\x5b\xda\x0c\x5d\x47\x16\x63\xfb\xea\x8d\xd4\xde\xe5\x8c\x7d\xf8\xc2\xfe\x09\x00\x00\xff\xff\xd1\x30\xbb\x2e\x0f\x0a\x00\x00")

func templatesTemplate_psp_function_endSqlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTemplate_psp_function_endSql,
		"templates/template_psp_function_end.sql",
	)
}

func templatesTemplate_psp_function_endSql() (*asset, error) {
	bytes, err := templatesTemplate_psp_function_endSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/template_psp_function_end.sql", size: 2575, mode: os.FileMode(438), modTime: time.Unix(1560864146, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTemplate_psp_function_end_rawSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\x4a\xcd\x4b\xb1\xe6\x52\x51\xe1\xca\x49\xcc\x4b\x2f\x4d\x4c\x4f\x55\x28\xc8\x29\x48\x2f\x2e\xcc\x51\x50\x29\xcb\xcf\x49\x2c\xc9\xcc\xc9\x2c\xa9\x8c\x4f\x4e\x2c\x49\x4d\xcf\x2f\xaa\x54\xb1\xe6\xe2\x8a\x29\xe1\x8a\x29\xe4\x02\x04\x00\x00\xff\xff\x43\xbb\x67\x3e\x38\x00\x00\x00")

func templatesTemplate_psp_function_end_rawSqlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTemplate_psp_function_end_rawSql,
		"templates/template_psp_function_end_raw.sql",
	)
}

func templatesTemplate_psp_function_end_rawSql() (*asset, error) {
	bytes, err := templatesTemplate_psp_function_end_rawSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/template_psp_function_end_raw.sql", size: 56, mode: os.FileMode(438), modTime: time.Unix(1559978435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTemplate_psp_function_end_viewSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\xe2\x02\x04\x00\x00\xff\xff\x6b\x13\xe3\x5b\x02\x00\x00\x00")

func templatesTemplate_psp_function_end_viewSqlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTemplate_psp_function_end_viewSql,
		"templates/template_psp_function_end_view.sql",
	)
}

func templatesTemplate_psp_function_end_viewSql() (*asset, error) {
	bytes, err := templatesTemplate_psp_function_end_viewSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/template_psp_function_end_view.sql", size: 2, mode: os.FileMode(438), modTime: time.Unix(1559978435, 0)}
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
	"templates/template_psp_function_begin.sql": templatesTemplate_psp_function_beginSql,
	"templates/template_psp_function_begin_raw.sql": templatesTemplate_psp_function_begin_rawSql,
	"templates/template_psp_function_begin_view.sql": templatesTemplate_psp_function_begin_viewSql,
	"templates/template_psp_function_drop.sql": templatesTemplate_psp_function_dropSql,
	"templates/template_psp_function_end.sql": templatesTemplate_psp_function_endSql,
	"templates/template_psp_function_end_raw.sql": templatesTemplate_psp_function_end_rawSql,
	"templates/template_psp_function_end_view.sql": templatesTemplate_psp_function_end_viewSql,
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
		"template_psp_function_begin.sql": &bintree{templatesTemplate_psp_function_beginSql, map[string]*bintree{}},
		"template_psp_function_begin_raw.sql": &bintree{templatesTemplate_psp_function_begin_rawSql, map[string]*bintree{}},
		"template_psp_function_begin_view.sql": &bintree{templatesTemplate_psp_function_begin_viewSql, map[string]*bintree{}},
		"template_psp_function_drop.sql": &bintree{templatesTemplate_psp_function_dropSql, map[string]*bintree{}},
		"template_psp_function_end.sql": &bintree{templatesTemplate_psp_function_endSql, map[string]*bintree{}},
		"template_psp_function_end_raw.sql": &bintree{templatesTemplate_psp_function_end_rawSql, map[string]*bintree{}},
		"template_psp_function_end_view.sql": &bintree{templatesTemplate_psp_function_end_viewSql, map[string]*bintree{}},
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

