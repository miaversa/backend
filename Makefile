# Go parameters
GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

BINARY = backend

VERSION ?= latest
COMMIT = $(shell git rev-parse HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.BRANCH=${BRANCH} -X main.COMMIT=${COMMIT}"

all: build

clean:
	-rm -f ${BINARY}
	-rm -rf vendor

deps: clean
	${GOGET} -u github.com/golang/dep/cmd/dep
	dep ensure

fmt : deps
	go fmt ./...

test: fmt
	$(GOTEST) -v -race ./...

build: test
	${GOBUILD} ${LDFLAGS} -o ${BINARY}

push:
	git push origin master

.PHONY: clean deps fmt test build push
