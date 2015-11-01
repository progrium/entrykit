NAME=entrykit
ARCH=$(shell uname -m)
ORG=progrium
VERSION=0.3.0

.PHONY: build release test

build:
	glu build darwin,linux ./cmd

docker: build
	docker build -t entrykit .
deps:
	go get github.com/gliderlabs/glu
	go get -u github.com/progrium/basht/...
	go get -d ./cmd

test: docker
	basht tests/*.bash
	docker rmi entrykit:switchtest

release:
	glu release v$(VERSION)
