SHELL=/bin/bash
GO_GET=go get -u -v all

build: src/*.go
	cd src && $(GO_GET) && go build -v -o ../aws-mfa

test: src/*.go
	cd src && $(GO_GET) && go test ./... -cover