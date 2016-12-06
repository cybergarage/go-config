// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
The Parser loads JSON files.

Parser gets a setting value in the specified file by the given path like Path.

	parser, err := NewParser()
	if err != nil {
		t.Error(err)
	}

	err = parser.ParseFromFile("/etc/profile.conf")
	if err != nil {
		t.Error(err)
	}

	name, err := parser.GetKeyStringByPath("/organizer/name")
	if err != nil {
		t.Error(err)
	}

	age, err := parser.GetKeyStringByPath("/organizer/age")
	if err != nil {
		t.Error(err)
	}

The sample JSON is defined as the following.

	{
		"organizer": {
			"name": "John Smith",
			"age": 33
		}
	}
*/
package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	PathSep = "/"
	LineSep = "\n"
)

const (
	errorKeyNull        = "Path is null"
	errorKeyNotFound    = "Key (%s) is not found"
	errorKeyTypeInvalid = "Key (%s) type is invalid"
)

type Parser struct {
	FileName   string
	rootObject interface{}
}

// NewParser returns a new Parser.
func NewParser() (*Parser, error) {
	parser := &Parser{}
	return parser, nil
}

// NewParserFromFile returns a new Parser from the given file.
func NewParserFromFile(file string) (*Parser, error) {
	parser := &Parser{}
	err := parser.ParseFromFile(file)
	if err != nil {
		return nil, err
	}
	return parser, nil
}

// NewParserFromString returns a new Parser from the given string.
func NewParserFromString(s string) (*Parser, error) {
	parser := &Parser{}
	err := parser.ParseFromString(s)
	if err != nil {
		return nil, err
	}
	return parser, nil
}

// ParseFromFile parses the given file.
func (parser *Parser) ParseFromFile(file string) error {
	parser.FileName = file

	_, err := os.Stat(file)
	if err != nil {
		return err
	}

	sourceBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return parser.ParseFromString(string(sourceBytes))
}

// ParseFromBytes parses the given bytes.
func (parser *Parser) ParseFromBytes(source []byte) error {
	return json.Unmarshal(source, &parser.rootObject)
}

// ParseFromString parses the given string.
func (parser *Parser) ParseFromString(source string) error {
	return parser.ParseFromBytes([]byte(source))
}

// getKeyObjectFromObject returns a object the given key.
func (parser *Parser) getKeyObjectFromObject(key string, obj interface{}) (interface{}, error) {
	switch obj.(type) {
	case map[string]interface{}:
		jsonDir, _ := obj.(map[string]interface{})
		keyObj, hasKey := jsonDir[key]
		if !hasKey {
			return "", errors.New(fmt.Sprintf(errorKeyNotFound, key))
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
			return "", errors.New(fmt.Sprintf(errorKeyTypeInvalid, key))
		}
	}
	return "", nil
}

// GetKey returns a object the given paths.
func (parser *Parser) getPathObjectFromObject(paths []string, rootObj interface{}) (interface{}, error) {
	obj := rootObj
	for _, path := range paths {
		keyObj, err := parser.getKeyObjectFromObject(path, obj)
		if err != nil {
			return nil, err
		}
		obj = keyObj
	}

	return obj, nil
}

// GetKeyObjectByPaths returns a key object by the given paths.
func (parser *Parser) GetKeyObjectByPaths(paths []string) (interface{}, error) {
	var keyObj interface{} = nil

	keyObj, err := parser.getPathObjectFromObject(paths, parser.rootObject)
	if err != nil {
		return "", err
	}

	return keyObj, nil
}

// GetStringByPaths returns a key string by the given paths.
func (parser *Parser) GetKeyStringByPaths(paths []string) (string, error) {
	keyStr := ""

	keyObj, err := parser.getPathObjectFromObject(paths, parser.rootObject)
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
		return "", errors.New(fmt.Sprintf(errorKeyTypeInvalid, (PathSep + strings.Join(paths, PathSep))))
	}

	return keyStr, nil
}

// GetObjectByPath returns a key object by the given Path.
func (parser *Parser) GetKeyObjectByPath(path string) (interface{}, error) {
	paths := strings.Split(path, PathSep)
	if len(paths) <= 0 {
		return "", errors.New(errorKeyNull)
	}
	return parser.GetKeyObjectByPaths(paths)
}

// GetStringByPath returns a key string by the given Path.
func (parser *Parser) GetKeyStringByPath(path string) (string, error) {
	paths := strings.Split(path, PathSep)
	if len(paths) <= 0 {
		return "", errors.New(errorKeyNull)
	}
	return parser.GetKeyStringByPaths(paths)
}
