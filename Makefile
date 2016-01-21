NAME=entrykit
ARCH=$(shell uname -m)
ORG=progrium
VERSION=0.4.0

.PHONY: build release

build:
	glu build darwin,linux ./cmd

deps:
	go get github.com/gliderlabs/glu
	go get -d ./cmd

release:
	glu release v$(VERSION)
