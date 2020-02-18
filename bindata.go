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

var _templatesTemplate_psp_function_endSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x5d\x92\xdb\x28\x10\x7e\x86\x53\x74\x5c\xda\x42\x9a\xd8\xc9\xcc\x56\xe5\x45\x1a\x7b\x2f\xb0\x37\xc8\xa6\xba\xb0\x68\xcb\x64\x30\x28\x80\x9c\x71\xd5\x1c\x7e\x0b\x90\x1d\x27\xe3\xc9\x43\x1e\x6c\xa0\x81\xfe\xbe\xfe\xe1\x93\xde\x01\x2a\x84\x0d\xdc\x43\xdc\x93\xe5\x0c\x8f\xf8\x19\x2d\x7e\x81\x76\x0d\xe2\xf1\xdd\x6a\x05\x64\x55\x0b\x95\x95\x07\xaa\x60\xb5\xda\x88\x0e\xd0\x62\xda\x4e\xc3\x7b\x78\xe8\x6e\x5d\x12\xf0\xf2\x02\x75\x6f\x5c\xff\x84\x51\x1f\x28\x44\x79\x18\xeb\x06\x56\x80\x3d\x36\xd0\xb6\xa0\x6d\x24\x7f\x94\x26\x1d\x14\x6f\x39\x4e\xf4\xe2\x01\xe1\xdd\x1a\xc4\xe7\x2f\xa2\x6d\xbf\x06\x67\xb7\x33\x55\x96\xb7\xd2\xf9\x34\xbe\xbc\x40\xde\xc4\xed\xa4\x8d\x42\xb7\xfd\x4a\x7d\xac\xc5\xbf\xda\x52\x0e\x41\x2c\xa1\x7e\x4d\x68\x85\xb1\xc7\xa6\x6d\xcf\x6c\x9a\xa6\x4b\x8e\x6f\xc7\x53\xfc\x8f\x9e\x62\x3c\xd5\x09\xb4\xf9\x2d\x79\xb2\x0a\xf4\xae\xe3\xe7\x11\xfb\x7c\xe0\x15\x87\x8e\x73\xbe\x73\x1e\x34\x68\x0b\x0f\xf0\xe1\x43\xf6\x61\x9c\x1b\x4b\x02\x8e\xf8\x59\x7f\x01\x1d\xc0\x4e\xc6\xe4\xd8\xcf\xb6\x44\x4f\x74\x70\x0d\x94\xae\x75\x9c\x63\xcc\x58\xd2\x7b\x79\xc2\xe8\x30\x44\xaf\xed\x50\xe3\x11\x97\x20\x44\x82\xf4\x14\x27\x6f\xa1\xf6\xee\x7b\x8d\x11\xdb\x36\xd2\x73\x5c\x72\xc6\x18\xa6\x38\x3d\xce\xc9\x2e\xa6\xed\x29\x92\x4c\xb6\x3c\x29\x36\xa9\xd4\xe5\x54\x32\x34\x9c\xb1\xe4\x99\x9e\x7b\x1a\xa3\x76\x16\xbe\x27\xae\x2e\xee\xc9\x87\xb9\x66\x03\x45\x08\x51\xf6\x4f\xa4\x40\x69\x39\x58\x17\xa2\xee\x03\x1c\x91\xbc\x77\x1e\xf3\x1e\xac\x61\x1c\xf0\xe2\x06\x7b\x67\x13\xb9\x8e\x73\x86\xfb\x1c\xd7\x7d\x9a\x97\x0c\x9d\x4d\x69\x28\x79\x67\xf4\xac\x63\xc1\x4e\xc6\x0d\xfc\x9d\x8c\x5b\x1a\x74\x6a\x1a\xa6\x6d\x20\x1f\x53\xff\x39\x40\x1c\xdd\xe0\xb0\x80\x07\xce\x18\x67\x75\x3a\xc3\xb4\xca\x31\xb2\x54\x26\xcc\x75\x2a\xeb\x29\x90\xc7\xf3\xe6\x6e\xb2\x7d\x26\x98\x1e\x47\x31\x8d\xd2\xcb\x43\x28\x73\xa3\x2d\xa1\x9d\x0e\x5b\xf2\xc5\x90\x83\x9b\xa7\xdf\x4c\x72\x1b\xe9\xc7\x32\x73\x38\xe7\x91\x1d\xa5\x99\x28\x31\x9a\x09\x1d\xd4\xa7\xd7\xdd\xdb\xb6\x47\xe9\xfb\xbd\xf4\xa9\x0f\xd1\x60\x53\xbc\xf5\x93\xf7\x64\xe3\x8f\x93\xc5\xfc\x50\x06\x51\x9d\x79\x57\xa2\x58\xaa\x9d\xf3\x87\xaa\xcc\xd1\x60\x99\xfc\x54\x91\x0b\xd0\x85\xee\xcf\xe4\xc9\xfb\xc3\x85\x7a\xe9\x2c\x6d\x07\xd0\x6a\x4e\x33\x79\xaf\x15\x76\x79\x57\xea\x40\x60\x5d\xd4\x3d\x81\x08\xdf\x0c\x64\x9c\x16\xfe\x82\x55\xf9\x89\x25\x5c\x73\x9c\x3b\x13\x66\x98\xf3\xf2\x17\x7e\xa5\x41\x58\x29\xfe\x3c\x79\xab\x0d\xf3\x43\xca\x8d\x62\xd5\xfc\x4e\xe7\x67\xc3\xae\x5f\x45\x0e\x8e\x89\xc7\xd1\x13\xf4\x46\x86\xb0\x5e\x8c\x9e\x56\x27\x32\xc6\x7d\x5f\x65\xf4\x05\x84\x78\x32\xb4\x5e\x6c\x65\xff\x34\x78\x37\x59\xb5\xea\x9d\x71\xbe\x2d\xa7\x3a\x28\xab\x51\xdb\xa7\x6e\xb1\xb9\xbb\xbb\x2b\xd1\xc2\x5e\x06\x70\x7d\x2e\x94\x02\x51\x90\x18\x63\x59\x34\x65\xa0\x42\xf9\x21\x09\x42\xfd\xd0\x14\xd2\x22\xdd\x7e\x94\x10\xa5\x1f\x28\xae\x17\xb8\x35\xd2\x3e\x2d\x60\xef\x69\xb7\x5e\x64\xb7\xff\x8c\xa8\xd5\x3a\x6b\xd5\x9c\xf2\x2c\x50\x8b\x8d\xa2\x28\xb5\x09\x8f\x1f\xe5\x46\x00\x99\x40\x20\x44\x92\x8d\xe6\x1a\x59\xc0\xdd\xdd\xdd\xe3\xc7\xd1\xd3\x46\x5c\xc9\x01\x63\x37\x74\xf5\x72\x4f\xa4\x46\x98\x42\xaa\x59\xa6\x30\xb7\x54\xde\x2a\x05\xd2\x4a\x2c\xcf\x74\xce\x7b\xcd\xb5\xb8\x30\x96\xea\x71\x2d\x2d\x37\x21\xc5\x3e\xc6\x11\x7b\xa7\x48\x2c\xe1\xd3\xfd\x7d\xe1\x9e\xff\xb3\x84\xe6\x62\x56\x15\x37\xd2\x0e\x93\x1c\x08\x46\x33\x0e\xa9\xbf\xaa\xa3\x33\x32\x6a\xa3\xe3\x09\x7b\x19\x69\x70\xfe\x54\x75\x9c\xff\x17\x39\x57\x64\x28\x12\xec\xbc\x3b\x5c\xc4\x20\x52\x88\xa9\x08\x21\x55\xc1\x13\x8c\x72\xa0\xfc\xc4\x61\x0d\xa2\x7c\x08\x45\xc7\xf9\x2d\x1d\xb9\x5c\xe5\x81\x0c\xf5\x11\x94\x0e\x51\xdb\x3e\x5e\x2e\x2e\x61\xb2\x96\x42\xac\x3d\x0d\xf4\x3c\x62\x18\x8d\x8e\x49\xa1\xb3\x54\xd7\xd5\xc5\x43\xb5\x04\xb1\x14\xe9\x7b\x74\x93\xa4\xa2\x91\xac\x22\xdb\xeb\x3f\xe1\x79\x7d\xfb\x8f\xa9\x5e\x3b\xb9\x66\xeb\xa0\xaa\xf8\x2c\xb7\xb7\xc0\x7b\x77\x18\xb5\x21\x85\xc1\x4d\xbe\x27\xa8\xb3\x7c\x82\x91\x29\x76\x3c\x92\x0f\xda\xd9\x86\xcf\x12\x08\xf5\x0f\x3e\x0f\x4d\xf7\xdb\xaf\xcb\x34\x2a\x19\xe9\x2d\x20\xce\x02\xc5\x5f\x60\x60\xfd\xab\xe1\x3d\x3c\x70\x56\x12\xfa\x2a\x97\x73\x8b\x75\x9c\xff\x1e\x09\x12\x90\x0e\x68\x9d\x9c\xe2\x1e\xd6\x50\x5d\x16\x15\xbc\xe1\x9b\xf3\x4a\xd1\x76\x1a\x06\xf2\xb8\xf5\x24\x9f\x46\xa7\x6d\x0c\x15\xff\x3f\x00\x00\xff\xff\xc8\x22\xa9\x27\xa3\x09\x00\x00")

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

	info := bindataFileInfo{name: "templates/template_psp_function_end.sql", size: 2467, mode: os.FileMode(438), modTime: time.Unix(1582024131, 0)}
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

