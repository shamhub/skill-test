.PHONY: build env test lint run

env:
	go env -w GO111MODULE=on

build: env
	GOOS=windows GOARCH=amd64 go build -o bin/pdfwriter cmd/app/main.go

test: env
	go test -v -race --tags=unit -cover ./...

lint: env
	golangci-lint run ./... --timeout 5m0s

run: env
	go run cmd/app/main.go
