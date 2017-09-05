###################################################################
#
# Configuration Library for Go
#
# Copyright (C) Satoshi Konno 2015
#
# This is licensed under BSD-style license, see file COPYING.
#
###################################################################

PRODUCT=config
PRODUCT_DIR=./${PRODUCT}
GITHUB=github.com/cybergarage/go-config

PACKAGES = ${GITHUB}/${PRODUCT}

all: build

setup:
	go get -u ${GITHUB}/${PRODUCT}

format:
	gofmt -w src

build: format
	go build -v ${PACKAGES}

test: build
	go test -v -cover ${PACKAGES}

install: build
	go install ${PACKAGES}

clean:
	rm -rf _obj
	go clean -i ${PACKAGES}
