NAME=entrykit
ARCH=$(shell uname -m)
ORG=progrium
VERSION=0.4.1

.PHONY: build release dep

build: dep
	go build ./...
	go test ./...
	mkdir -p build/Darwin
	go build -a -installsuffix cgo -ldflags "-X main.Version=$(VERSION)" -o build/Darwin/entrykit ./cmd
	mkdir -p build/Linux
	go build -a -installsuffix cgo -ldflags "-X main.Version=$(VERSION)" -o build/Linux/entrykit ./cmd

dep:
	GO111MODULE=on go mod tidy

release:
	git tag v$(VERSION)
	git push origin v$(VERSION)

