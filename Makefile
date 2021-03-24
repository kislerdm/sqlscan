# Dmitry Kisler © 2021-present
# www.dkisler.com <admin@dkisler.com>

SHELL=/bin/bash

PROGRAM := sql_scan

VER := `cat VERSION`
# amd64, arm64, 386, arm
ARCH := amd64
# linux, darwin, windows
OS := linux

BIN_NAME := $(PROGRAM)-$(OS)-$(ARCH)

test:
	@go test -tags test -coverprofile="go-cover.tmp" ./...
	@go tool cover -func go-cover.tmp
	@rm go-cover.tmp

build:
	@CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=$(OS) \
		go build -a -gcflags=all="-l -B -C" -ldflags="-w -s" -o $(BIN_NAME) *.go

compress:
	@upx -9 --ultra-brute $(BIN_NAME)

coverage-bump:
	@./coverage_bump.py
