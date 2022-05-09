.PHONY: help fmt lint test test-cover all

# Basic Makefile for Golang project
# Includes GRPC Gateway, Protocol Buffers
FILES		?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")

default: help

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

tests:
	echo "Executing tests"
	go test ./...

fmt:    ## format the go source files
	echo "Formating tests"
	go fmt ./...
	goimports -w $(FILES)

tools:  ## fetch and install all required tools
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/golang/lint/golint