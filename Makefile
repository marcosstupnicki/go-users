.PHONY: help fmt lint tests all

# Basic Makefile for Golang project
# Includes GRPC Gateway, Protocol Buffers
FILES		?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")

default: help

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

tests:  ## execute the go source tests.
	echo "Executing tests"
	go test ./...

coverage:  ## execute the go source tests with code coverage info.
	echo "Executing tests with coverage"
	go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out

fmt:  ## format the go source files.
	echo "Formatting files"
	go fmt ./...
	goimports -w $(FILES)

tools:  ## fetch and install all required tools.
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/golang/lint/golint