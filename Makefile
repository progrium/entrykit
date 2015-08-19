NAME=entrykit
ARCH=$(shell uname -m)
ORG=progrium
VERSION=0.2.1

.PHONY: build release

build:
	mkdir -p build/Linux  && GOOS=linux  go build -ldflags "-X main.Version $(VERSION)" -o build/Linux/$(NAME) ./cmd
	mkdir -p build/Darwin && GOOS=darwin go build -ldflags "-X main.Version $(VERSION)" -o build/Darwin/$(NAME) ./cmd

deps:
	go get -u github.com/progrium/gh-release/...
	go get -u ./...

release:
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_Linux_$(ARCH).tgz -C build/Linux $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_Darwin_$(ARCH).tgz -C build/Darwin $(NAME)
	gh-release create progrium/$(NAME) $(VERSION) \
		$(shell git rev-parse --abbrev-ref HEAD) v$(VERSION)

circleci:
	rm ~/.gitconfig
	rm -rf /home/ubuntu/.go_workspace/src/github.com/$(ORG)/$(NAME) && cd .. \
		&& mkdir -p /home/ubuntu/.go_workspace/src/github.com/$(ORG) \
		&& mv $(NAME) /home/ubuntu/.go_workspace/src/github.com/$(ORG)/$(NAME) \
		&& ln -s /home/ubuntu/.go_workspace/src/github.com/$(ORG)/$(NAME) $(NAME)


clean:
	rm -rf build release
