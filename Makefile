SHELL := /bin/bash

.PHONY: build run test vet lint run-dev-appserver \
  run-on-docker mockgen coverage gce-deploy

build:
	go build .

run: 
	go build . && ./Kud

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

gce-deploy:
	./scripts/gce_deploy.sh

moq:
	moq -out ./adapter/repository/github/github_mock.go -pkg github ./entity GitHubRepository
	moq -out ./adapter/repository/badger/slack/slack_mock.go -pkg slack ./entity SlackRepository
