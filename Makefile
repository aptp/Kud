SHELL := /bin/bash

.PHONY: build test vet lint run-dev-appserver \
  run-on-docker mockgen coverage

build:
	go build .

test: 
	go test -race -v $(shell go list ./...)

vet:
	go vet $(shell go list ./...)

lint:
	go lint $(shell go list ./...)

run-dev_appserver:
	dev_appserver.py app.yaml

coverage:
	go tool cover -html=coverage.out
