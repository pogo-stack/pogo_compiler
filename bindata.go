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

var _templatesTemplate_psp_function_beginSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x51\xd1\x8e\xd3\x30\x10\x7c\xbe\x7c\xc5\x3e\x58\x72\x22\x45\x50\x78\xec\x01\x3f\x02\x68\xb5\xb1\xb7\x89\xaf\x89\x1d\xad\x37\x15\xf7\xf7\xc8\x2e\x04\x15\xa8\xae\x2f\xc9\x78\x76\x32\xe3\x9d\x7c\xd3\xc6\x09\x93\x32\x9c\xb6\xe8\x34\xa4\x08\xe6\x37\x32\xd0\x9a\x95\x84\x16\x56\x96\x6c\xba\x46\x58\x37\x89\x19\xcc\x2f\x60\x80\x32\x18\xd3\x78\x76\x33\x09\x37\x4f\xe8\x10\x34\x2c\x9c\x95\x96\x15\x8e\x9f\xc1\xcd\xc9\x9d\x71\xa7\xda\xee\xb9\x79\x42\x45\x50\xfe\xa1\x05\xce\x08\x21\x2a\x8f\x2c\x45\x7d\x28\xd4\x05\xe1\x42\xe2\x26\x92\xaf\xdf\xcb\x39\xde\x48\x3e\x14\xca\xdf\x50\xc6\xf3\xb0\x8d\xa6\x0e\x32\xc2\x4b\x4e\x71\x28\x7c\x05\x38\x6c\x61\xf6\x98\x86\x17\x76\x7a\x8d\xf7\x19\x87\x47\x74\xa7\xb2\x39\x7a\x5e\x75\xda\xf3\xea\xf5\x27\x61\xf2\x18\xfc\x7e\xd1\xe2\xd2\xba\x44\x33\x67\xc7\x6d\xdc\xe6\x39\x9c\x5a\xb7\x89\x70\x54\xcc\xac\x1a\xe2\xd8\x5a\xc4\x35\x8d\xe9\xdd\xfe\xb9\xed\x41\x65\xe3\xae\x07\x6b\xcb\xe3\x53\x7c\x4f\x5f\x6c\x57\xb3\x59\xe4\x2f\x7f\x6b\x6b\xf6\x63\x05\x2f\x77\xd6\x23\x11\x7a\xbd\x6a\xa6\x7f\x9a\xbf\x94\xd4\x24\x98\x95\xdc\xb9\xfe\xa1\x3f\xb9\xc5\x44\x1e\xe8\x6c\x78\x55\x26\x41\xa8\xef\x22\x2c\x65\x1c\x8f\xf5\x58\xe6\xe4\xfd\x1b\x36\x76\x52\x5d\xd1\x25\xcf\xb6\x87\x8f\x87\x43\x0f\x76\x62\xf2\x2c\xd9\xf6\xff\xdb\xa5\x07\xeb\x52\x3a\x07\xbe\x33\xef\x9e\x7f\x06\x00\x00\xff\xff\x38\xd0\xbc\xf1\xe3\x02\x00\x00")

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

	info := bindataFileInfo{name: "templates/template_psp_function_begin.sql", size: 739, mode: os.FileMode(438), modTime: time.Unix(1580730084, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTemplate_psp_function_begin_rawSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8d\xc1\x4a\x04\x31\x0c\x86\xcf\xf6\x29\x72\x28\xb4\x85\x45\xef\xae\xfa\x24\x42\xc8\xb4\x99\x99\x2e\x35\x1d\xd2\xd4\xe7\x97\x61\xd5\xf3\x5e\xc2\x47\xf8\xfe\xff\xff\x34\x97\x95\xc9\x18\xd6\x29\xd9\x6a\x17\xf0\x7f\xe4\x21\xfa\x83\x94\xbe\xd8\x58\x87\x4f\x4e\xd9\xa6\xca\x00\xff\x0b\x1e\x68\x80\xf7\xae\x70\x6e\xa4\xec\x9e\xb0\x0c\x84\xdb\xe8\xb2\xc0\xeb\xfb\x1d\x70\x99\xb5\x15\xec\xcb\x8d\xb3\xc5\x74\xbd\x4b\xcb\x23\xde\x7a\x4e\x63\xe1\xc3\x76\x84\x2a\xc6\x1b\xeb\xf9\xb7\x5d\x99\x0a\xd6\x82\xf0\x4d\x9a\x77\xd2\xb3\x25\xe6\x4e\x8d\x47\xe6\x28\xb3\xb5\xba\xc6\x3c\x55\x59\x0c\x07\x9b\x55\xd9\x62\x40\x3c\xfa\xd6\x9f\xff\xe3\xe1\x02\xa6\x93\xd3\x05\x42\x38\xcf\x9b\xbc\xd0\x47\x48\xe9\xea\x7e\x02\x00\x00\xff\xff\xf8\xc6\x18\x60\x16\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/template_psp_function_begin_raw.sql", size: 278, mode: os.FileMode(438), modTime: time.Unix(1580730084, 0)}
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

var _templatesTemplate_psp_function_endSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\xd1\x92\xdb\x2a\x0f\xbe\x86\xa7\x50\x33\x9e\xc1\xde\x26\xed\xee\x3f\xd3\x1b\x7b\x93\xff\x05\xce\x1b\xf4\x74\x34\xc4\x28\x0e\x5d\x02\x2e\xe0\x74\x33\xb3\x0f\x7f\x06\x70\xb2\x69\x37\xdb\x8b\xde\x18\x10\x42\xfa\xa4\x4f\x92\xf5\x0e\x50\x21\x6c\xe0\x1e\xe2\x9e\x2c\x67\x78\xc4\xaf\x68\xf1\x1b\xb4\x6b\x10\x8f\x1f\x56\x2b\x20\xab\x5a\xa8\xac\x3c\x50\x05\xab\xd5\x46\x74\x80\x16\xd3\x75\x5a\x3e\xc2\x43\x77\xeb\x91\x80\x97\x17\xa8\x7b\xe3\xfa\x27\x8c\xfa\x40\x21\xca\xc3\x58\x37\xb0\x02\xec\xb1\x81\xb6\x05\x6d\x23\xf9\xa3\x34\x49\x51\xbc\x67\x38\xc1\x8b\x07\x84\x0f\x6b\x10\x5f\xbf\x89\xb6\xfd\x1e\x9c\xdd\xce\x50\x59\xbe\x4a\xfa\x69\x7d\x79\x81\x7c\x89\xdb\x49\x1b\x85\x6e\xfb\x9d\xfa\x58\x8b\x7f\xb4\xa5\x1c\x82\x58\x42\xfd\x16\xd0\x0a\x63\x8f\x4d\xdb\x9e\xd1\x34\x4d\x97\x0c\xdf\x8e\xa7\xd8\x1f\x3d\xc5\x78\xaa\x93\xd3\xe6\x8f\xe0\xc9\x2a\xd0\xbb\x8e\x9f\x57\xec\xb3\xc2\x1b\x0c\x1d\xe7\x7c\xe7\x3c\x68\xd0\x16\x1e\xe0\xd3\xa7\x6c\xc3\x38\x37\x96\x04\x1c\xf1\xab\xfe\x06\x3a\x80\x9d\x8c\xc9\xb1\x9f\x65\x09\x9e\xe8\xe0\xda\x51\x7a\xd6\x71\x8e\x31\xfb\x92\xde\xcb\x13\x46\x87\x21\x7a\x6d\x87\x1a\x8f\xb8\x04\x21\x92\x4b\x4f\x71\xf2\x16\x6a\xef\x7e\xd6\x18\xb1\x6d\x23\x3d\xc7\x25\x67\x8c\x61\x8a\xd3\xe3\x9c\xec\x22\xda\x9e\x22\xc9\x24\xcb\x9b\x22\x93\x4a\x5d\xb4\x92\xa0\xe1\x8c\x25\xcb\xf4\xdc\xd3\x18\xb5\xb3\xf0\x33\x61\x75\x71\x4f\x3e\xcc\x9c\x0d\x14\x21\x44\xd9\x3f\x91\x02\xa5\xe5\x60\x5d\x88\xba\x0f\x70\x44\xf2\xde\x79\xcc\x77\xb0\x86\x71\xc0\x8b\x19\xec\x9d\x4d\xe0\x3a\xce\x19\xee\x73\x5c\xf7\x1d\x67\x5e\xea\x40\x60\x5d\xd4\x3d\x81\x08\x3f\x0c\x64\x13\x22\xa9\x95\xe4\x9d\xb5\xd3\x52\x28\x61\xf4\xac\x63\x81\x95\x84\x1b\xf8\x5f\x12\x6e\x69\xd0\xa9\x9e\x98\xb6\x81\x7c\x4c\xa5\xe9\x00\x71\x74\x83\xc3\x82\x2b\x70\xc6\x38\xab\x93\x0e\xd3\x2a\x87\xcf\x12\x83\x98\x29\x2c\xe7\x29\x90\xc7\xf3\xe5\x6e\xb2\x7d\xc6\x9e\xfa\xa6\x88\x46\xe9\xe5\x21\x94\xbd\xd1\x96\xd0\x4e\x87\x2d\xf9\x22\xc8\x71\xcf\xdb\x1f\x26\x99\x8d\xf4\x7a\xcc\x18\xce\x29\x66\x47\x69\x26\x4a\x88\x66\x40\x07\xf5\xe5\x6d\x61\xb7\xed\x51\xfa\x7e\x2f\x7d\x2a\x51\x34\xd8\x14\x6b\xfd\xe4\x3d\xd9\xf8\xaa\x59\xc4\x0f\x65\x11\xd5\x19\x77\x25\x8a\xa4\xda\x39\x7f\xa8\xca\x1e\x0d\x96\xcd\x2f\x64\x5d\x1c\x5d\xe0\xfe\x0a\x9e\xbc\x3f\x5c\xa0\x97\xa2\xd3\x76\x00\xad\xe6\x34\x93\xf7\x5a\x61\xa2\x21\x93\x33\x6f\xde\xab\xa0\xdc\x03\x99\x48\xab\xe6\x16\x9b\x2b\x9e\xb3\xeb\x8a\xce\xde\x99\x78\x1c\x3d\x41\x6f\x64\x08\xeb\xc5\xe8\x69\x75\x22\x63\xdc\xcf\x55\x86\xbf\x80\x10\x4f\x86\xd6\x8b\xad\xec\x9f\x06\xef\x26\xab\x56\xbd\x33\xce\xb7\x45\xab\x83\x72\x1a\xb5\x7d\xea\x16\x9b\xbb\xbb\xbb\x52\x60\xb0\x97\x01\x5c\x9f\x33\xa9\x40\x14\x4f\x8c\xb1\x3c\xf0\x64\xa0\x82\xf9\x21\x35\x73\xfd\xd0\x14\xd4\x22\xbd\x7e\x94\x10\xa5\x1f\x28\xae\x17\xb8\x35\xd2\x3e\x2d\x60\xef\x69\xb7\x5e\x64\xb3\xff\x1f\x51\xab\x75\x9e\x33\x73\x4e\xf2\x70\x59\x6c\x14\x45\xa9\x4d\x78\xfc\x2c\x37\x02\xc8\x04\x02\x21\x52\xcb\x37\xd7\x9e\x05\xdc\xdd\xdd\x3d\x7e\x1e\x3d\x6d\xc4\x55\x2b\x33\x76\x63\x26\x5e\xde\x89\xc4\xd4\x14\xc4\x12\x44\x69\x9d\xe5\xeb\x55\x61\x58\x2b\xb1\x3c\xc3\x39\xdf\x35\xd7\x83\x81\xb1\x44\xc8\xf5\x58\xb8\xe9\x52\xec\x63\x1c\xb1\x77\x8a\xc4\x12\xbe\xdc\xdf\x17\xec\xf9\x9b\xc7\x5f\x66\xb3\xaa\xb8\x91\x76\x98\xe4\x40\x30\x9a\x71\x48\x2d\x5d\x1d\x9d\x91\x51\x1b\x1d\x4f\xd8\xcb\x48\x83\xf3\xa7\xaa\xe3\xfc\xdf\xc8\xb9\x22\x43\x91\x60\xe7\xdd\xe1\xd2\xad\x91\x42\x4c\x24\x84\xc4\x82\x27\x18\xe5\x40\xb9\x07\x61\x0d\xa2\xfc\xc4\xd2\x80\xb8\xd5\xe8\x97\xa7\x3c\x90\xa1\x3e\x82\xd2\x21\x6a\xdb\xc7\xcb\xc3\x25\x4c\xd6\x52\x88\xb5\xa7\x81\x9e\x47\x0c\xa3\xd1\x31\x4d\xd7\x3c\x66\xeb\xea\x62\xa1\x5a\x82\x58\x8a\xf4\x2f\xb9\x09\x52\xd1\x48\x56\x91\xed\xf5\xdf\xe0\xbc\x7e\xfd\xd7\x50\xaf\x8d\x5c\xa3\x75\x50\x55\x7c\x9e\x87\xb7\x9c\xf7\xee\x30\x6a\x43\x0a\x83\x9b\x7c\x4f\x50\xe7\xf9\x06\x46\xa6\xd8\xf1\x48\x3e\x68\x67\x1b\x3e\xcf\x28\xa8\x5f\xf1\x3c\x34\xdd\x1f\xff\x0c\xd3\xa8\x64\xa4\xf7\x1c\x71\x16\x28\xfe\xe6\x06\xd6\xbf\x0b\x3e\xc2\x03\x67\x25\xa1\x6f\x72\x39\x97\x58\xc7\xf9\x9f\x3d\x41\x72\xa4\x03\x5a\x27\xa7\xb8\x87\x35\x54\x97\x43\x05\xef\xd8\xe6\xbc\x52\xb4\x9d\x86\x81\x3c\x6e\x3d\xc9\xa7\xd1\x69\x1b\x43\xc5\xff\x0b\x00\x00\xff\xff\x11\x28\x90\x99\x5f\x09\x00\x00")

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

	info := bindataFileInfo{name: "templates/template_psp_function_end.sql", size: 2399, mode: os.FileMode(438), modTime: time.Unix(1580731861, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTemplate_psp_function_end_rawSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\x4a\xcd\x4b\xb1\xe6\x52\x51\xe1\xca\x49\xcc\x4b\x2f\x4d\x4c\x4f\x55\x28\xc8\x29\x48\x2f\x2e\xcc\x51\x50\x29\xcb\xcf\x49\x2c\xc9\xcc\xc9\x2c\xa9\x8c\x4f\x4e\x2c\x49\x4d\xcf\x2f\xaa\x54\xb1\xe6\xe2\x8a\x29\x01\x04\x00\x00\xff\xff\xe5\xcd\x1c\xfa\x34\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/template_psp_function_end_raw.sql", size: 52, mode: os.FileMode(438), modTime: time.Unix(1580730084, 0)}
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

