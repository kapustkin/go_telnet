all: format test lint build

install:
	go build -o bin/telnet

format:
	gofmt -w .

test:
	go test -cover ./...

lint:
	$(shell go env GOPATH)/bin/golangci-lint run --enable-all

build:
	go build -o bin/telnet