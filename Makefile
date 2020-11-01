.PHONY: test/unit_test test/all  build

gen/go:
	@scripts/gen.sh go

lint:
	@scripts/lint.sh

test/unit_test:
	go test ./... -v -race -covermode=atomic -cover -short

test/all:
	go test ./... -v -race -coverprofile=coverage.txt -covermode=atomic

build:
	go build -o ./bin/simple_ctftime_bot ./cmd/simple_ctftime_bot
