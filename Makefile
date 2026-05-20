.PHONY: build run clean test deps fmt lint

BINARY_NAME=aithlete-api
CMD_PATH=./cmd/server

deps:
	go mod tidy

fmt:
	go fmt ./...

lint:
	golangci-lint run

build: deps
	go build -o $(BINARY_NAME) $(CMD_PATH)

run: deps
	go run $(CMD_PATH)

clean:
	rm -f $(BINARY_NAME)
	go clean

test:
	go test -v ./...
