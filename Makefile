# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

BINARY = backend

VERSION ?= latest
COMMIT = $(shell git rev-parse HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.BRANCH=${BRANCH} -X main.COMMIT=${COMMIT}"

all: clean test build

build:
	${GOBUILD} ${LDFLAGS} -o ${BINARY}

test:
	$(GOTEST) -v ./...

clean:
	-rm -f ${BINARY}

.PHONY: build test clean
