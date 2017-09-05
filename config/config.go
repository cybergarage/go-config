// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Config loads JSON files which has comments.

Config gets a setting value in the specified file by the given path like Path.

	config, err := NewConfig()
	if err != nil {
		t.Error(err)
	}

	err = config.ParseFromFile("/etc/profile.conf")
	if err != nil {
		t.Error(err)
	}

	name, err := config.GetKeyStringByPath("/organizer/name")
	if err != nil {
		t.Error(err)
	}

	age, err := config.GetKeyStringByPath("/organizer/age")
	if err != nil {
		t.Error(err)
	}

The configuration file fomat is based on JSON as the following.

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
package config

import (
	"bytes"
	"strings"
)

const (
	Comment = "#"
)

type Config struct {
	Parser
}

// NewConfig returns a new Config.
func NewConfig() (*Config, error) {
	config := &Config{}
	return config, nil
}

// NewConfigFromFile returns a new Config from the given file.
func NewConfigFromFile(file string) (*Config, error) {
	config := &Config{}
	err := config.ParseFromFile(file)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// NewConfigFromString returns a new Config from the given string.
func NewConfigFromString(s string) (*Config, error) {
	config := &Config{}
	err := config.ParseFromString(s)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// ParseFromString parses the given string.
func (config *Config) ParseFromString(source string) error {
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

	return config.ParseFromBytes(strippedSource.Bytes())
}
