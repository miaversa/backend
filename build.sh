#!/usr/bin/env bash

go clean
go fmt ./...

go-bindata -prefix templates -o templates/template.go -pkg templates templates/*.html
go-bindata -prefix assets -o assets/assets.go -pkg assets assets/*.css

go build -o backend
