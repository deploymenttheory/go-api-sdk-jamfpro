.PHONY: terraform/* build

##### Go
NAME := $(notdir $(PWD))
VERSION := 1.0.0

build: ## go build
	CGO_ENABLED=0 go build -o bin/$(NAME) *.go