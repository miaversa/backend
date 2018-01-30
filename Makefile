
BINARY = backend

VERSION ?= latest
COMMIT = $(shell git rev-parse HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.BRANCH=${BRANCH} -X main.COMMIT=${COMMIT}"

all: clean deps gen fmt test build

clean:
	go clean
	-rm -rf vendor
	-rm templates/*.go

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

fmt :
	go fmt ./...

gen:
	go-bindata -o templates/template.go -pkg templates templates/*.html

test:
	go test -v -race ./...

build:
	go build ${LDFLAGS} -o ${BINARY}

push:
	git push origin master

.PHONY: clean deps fmt gen test build push
