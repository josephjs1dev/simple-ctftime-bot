
.PHONY: test/all test/unit_test build

test/unit_test:
	go test ./... -v -race -covermode=atomic -cover -short

test/all:
	go test ./... -v -race -coverprofile=coverage.txt -covermode=atomic

build:
	go build -o ./bin/app ./cmd/simple_ctftime_bot
