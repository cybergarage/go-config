// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"errors"
	"fmt"
	"testing"
)

const (
	TEST_KEY_ORGNIZER   = "organizer"
	TEST_KEY_NAME       = "name"
	TEST_KEY_AGE        = "age"
	TEST_KEY_NAME_VALUE = "John Smith"
	TEST_KEY_AGE_VALUE  = "33"

	TEST_CONFIG = "{\n" +
		"    \"" + TEST_KEY_ORGNIZER + "\": {\n" +
		"        \"" + TEST_KEY_NAME + "\": \"" + TEST_KEY_NAME_VALUE + "\",\n" +
		"        \"" + TEST_KEY_AGE + "\": " + TEST_KEY_AGE_VALUE + "\n" +
		"    }\n" +
		"}\n"
)

func TestLoadingSimplePointer(t *testing.T) {

	const TEST_SIMPLE_KEY = TEST_KEY_NAME
	const TEST_SIMPLE_VAL = TEST_KEY_NAME_VALUE

	const TEST_CONFIG = "{\n" +
		"\"" + TEST_SIMPLE_KEY + "\": \"" + TEST_SIMPLE_VAL + "\"" +
		"}"

	pointer, err := NewPointer()
	if err != nil {
		t.Error(err)
	}

	err = pointer.ParseFromString(TEST_CONFIG)
	if err != nil {
		t.Error(err)
	}

	// TEST_SIMPLE_KEY

	keyValue, err := pointer.GetKeyStringByPath(TEST_SIMPLE_KEY)
	if err != nil {
		t.Error(err)
	}

	if len(keyValue) <= 0 {
		t.Error(errors.New(fmt.Sprintf("%s is null", TEST_SIMPLE_KEY)))
	}

	if keyValue != TEST_SIMPLE_VAL {
		t.Error(errors.New(fmt.Sprintf("%s is not equals (%s)", TEST_SIMPLE_VAL, keyValue)))
	}
}

func CheckPointerMembers(t *testing.T, pointer *Pointer) {

	// /organizer/name

	xpath := TEST_KEY_ORGNIZER + "/" + TEST_KEY_NAME
	keyValue, err := pointer.GetKeyStringByPath(xpath)
	if err != nil {
		t.Error(err)
	}

	if len(keyValue) <= 0 {
		t.Error(errors.New(fmt.Sprintf("%s is null", xpath)))
	}

	if keyValue != TEST_KEY_NAME_VALUE {
		t.Error(errors.New(fmt.Sprintf("%s is not equals (%s)", TEST_KEY_NAME_VALUE, keyValue)))
	}

	// /organizer/age

	xpath = TEST_KEY_ORGNIZER + "/" + TEST_KEY_AGE
	keyValue, err = pointer.GetKeyStringByPath(xpath)
	if err != nil {
		t.Error(err)
	}

	if len(keyValue) <= 0 {
		t.Error(errors.New(fmt.Sprintf("%s is null", xpath)))
	}

	if keyValue != TEST_KEY_AGE_VALUE {
		t.Error(errors.New(fmt.Sprintf("%s is not equals (%s)", TEST_KEY_AGE_VALUE, keyValue)))
	}
}

func ParsePointerTest(t *testing.T, s string) {

	pointer, err := NewPointer()
	if err != nil {
		t.Error(err)
	}

	err = pointer.ParseFromString(s)
	if err != nil {
		t.Error(err)
	}

	CheckPointerMembers(t, pointer)
}

func TestLoadingPointer(t *testing.T) {
	ParsePointerTest(t, TEST_CONFIG)
}
