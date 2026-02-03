.PHONY: build linux

# build: build the mutation-webhook binary
build: linux

linux:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/mutation-webhook cmd/*.go
