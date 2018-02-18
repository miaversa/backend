#!/usr/bin/env bash

go clean
go fmt ./...

rm assets/assets.go

go-bindata -prefix templates -o templates/template.go -pkg templates templates/*.html
go-bindata -prefix assets -o assets/assets.go -pkg assets assets/*

go build -o backend
