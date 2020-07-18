
.PHONY: test
test:
	go test ./... -v -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: build
build:
	go build -o ./bin/app ./cmd/simple_ctftime_bot
