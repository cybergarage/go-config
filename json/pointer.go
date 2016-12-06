// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package pointer loads JSON files.

Pointer gets a setting value in the specified file by the given path like Path.

	pointer, err := NewPointer()
	if err != nil {
		t.Error(err)
	}

	err = pointer.ParseFromFile("/etc/profile.conf")
	if err != nil {
		t.Error(err)
	}

	name, err := pointer.GetKeyStringByPath("/organizer/name")
	if err != nil {
		t.Error(err)
	}

	age, err := pointer.GetKeyStringByPath("/organizer/age")
	if err != nil {
		t.Error(err)
	}

The pointeruration file fomat is based on JSON as the following.

	#
	#  /etc/profile.conf
	#

	{
		"organizer": {
			"name": "John Smith",
			"age": 33
		}
	}
*/
package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	//	"reflect"
	"strconv"
)

const (
	PathSep = "/"
	LineSep = "\n"
	Comment = "#"
)

const (
	errorPointerKeyNull        = "Path is null"
	errorPointerKeyNotFound    = "Key (%s) is not found"
	errorPointerKeyTypeInvalid = "Key (%s) type is invalid"
)

type Pointer struct {
	FileName   string
	rootObject interface{}
}

// NewPointer returns a new Pointer.
func NewPointer() (*Pointer, error) {
	pointer := &Pointer{}
	return pointer, nil
}

// NewPointerFromFile returns a new Pointer from the given file.
func NewPointerFromFile(file string) (*Pointer, error) {
	pointer := &Pointer{}
	err := pointer.ParseFromFile(file)
	if err != nil {
		return nil, err
	}
	return pointer, nil
}

// NewPointerFromString returns a new Pointer from the given string.
func NewPointerFromString(s string) (*Pointer, error) {
	pointer := &Pointer{}
	err := pointer.ParseFromString(s)
	if err != nil {
		return nil, err
	}
	return pointer, nil
}

// ParseFromFile parses the given file.
func (pointer *Pointer) ParseFromFile(file string) error {
	pointer.FileName = file

	_, err := os.Stat(file)
	if err != nil {
		return err
	}

	sourceBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return pointer.ParseFromString(string(sourceBytes))
}

// ParseFromString parses the given string.
func (pointer *Pointer) ParseFromString(source string) error {
	lines := strings.Split(source, LineSep)

	// Strip comment and null lines
	var strippedSource bytes.Buffer
	for _, line := range lines {
		if len(line) <= 0 {
			continue
		}
		commentIdx := strings.Index(line, Comment)
		if 0 <= commentIdx {
			continue
		}
		strippedSource.WriteString(line + LineSep)
	}

	return json.Unmarshal(strippedSource.Bytes(), &pointer.rootObject)
}

// getKeyObjectFromObject returns a object the given key.
func (pointer *Pointer) getKeyObjectFromObject(key string, obj interface{}) (interface{}, error) {
	switch obj.(type) {
	case map[string]interface{}:
		jsonDir, _ := obj.(map[string]interface{})
		keyObj, hasKey := jsonDir[key]
		if !hasKey {
			return "", errors.New(fmt.Sprintf(errorPointerKeyNotFound, key))
		}
		//fmt.Println("%s = %s", key, reflect.TypeOf(keyObj))
		switch keyObj.(type) {
		case string:
			return keyObj, nil
		case float64:
			return keyObj, nil
		case map[string]interface{}:
			return keyObj, nil
		default:
			return "", errors.New(fmt.Sprintf(errorPointerKeyTypeInvalid, key))
		}
	}
	return "", nil
}

// GetKey returns a object the given paths.
func (pointer *Pointer) getPathObjectFromObject(paths []string, rootObj interface{}) (interface{}, error) {
	obj := rootObj
	for _, path := range paths {
		keyObj, err := pointer.getKeyObjectFromObject(path, obj)
		if err != nil {
			return nil, err
		}
		obj = keyObj
	}

	return obj, nil
}

// GetKeyObjectByPaths returns a key object by the given paths.
func (pointer *Pointer) GetKeyObjectByPaths(paths []string) (interface{}, error) {
	var keyObj interface{} = nil

	keyObj, err := pointer.getPathObjectFromObject(paths, pointer.rootObject)
	if err != nil {
		return "", err
	}

	return keyObj, nil
}

// GetStringByPaths returns a key string by the given paths.
func (pointer *Pointer) GetKeyStringByPaths(paths []string) (string, error) {
	keyStr := ""

	keyObj, err := pointer.getPathObjectFromObject(paths, pointer.rootObject)
	if err != nil {
		return "", err
	}

	switch keyObj.(type) {
	case string:
		keyStr, _ = keyObj.(string)
	case float64:
		keyValue, _ := keyObj.(float64)
		keyStr = strconv.FormatFloat(keyValue, 'g', -1, 64)
	default:
		return "", errors.New(fmt.Sprintf(errorPointerKeyTypeInvalid, (PathSep + strings.Join(paths, PathSep))))
	}

	return keyStr, nil
}

// GetObjectByPath returns a key object by the given Path.
func (pointer *Pointer) GetKeyObjectByPath(path string) (interface{}, error) {
	paths := strings.Split(path, PathSep)
	if len(paths) <= 0 {
		return "", errors.New(errorPointerKeyNull)
	}
	return pointer.GetKeyObjectByPaths(paths)
}

// GetStringByPath returns a key string by the given Path.
func (pointer *Pointer) GetKeyStringByPath(path string) (string, error) {
	paths := strings.Split(path, PathSep)
	if len(paths) <= 0 {
		return "", errors.New(errorPointerKeyNull)
	}
	return pointer.GetKeyStringByPaths(paths)
}
