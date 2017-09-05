// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import "testing"

func ParseConfigTest(t *testing.T, s string) {

	config, err := NewConfig()
	if err != nil {
		t.Error(err)
	}

	err = config.ParseFromString(s)
	if err != nil {
		t.Error(err)
	}

	CheckParserMembers(t, &config.Parser)
}

func TestLoadingCommentedConfig(t *testing.T) {
	TEST_COMMENTED_CONFIG :=
		"####\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)

	TEST_COMMENTED_CONFIG =
		"####\n" +
			"####\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)

	TEST_COMMENTED_CONFIG =
		" ####\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)
}

func TestLoadingBlankConfig(t *testing.T) {
	TEST_COMMENTED_CONFIG :=
		"\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)

	TEST_COMMENTED_CONFIG =
		"\n" +
			"\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)
}

func TestLoadingCommentAndBlankConfig(t *testing.T) {
	TEST_COMMENTED_CONFIG :=
		"####\n" +
			"\n" +
			TEST_CONFIG
	ParseConfigTest(t, TEST_COMMENTED_CONFIG)
}
