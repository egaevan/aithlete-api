.PHONY: build run clean test deps fmt lint migrate-up migrate-down migrate-create

BINARY_NAME=aithlete-api
CMD_PATH=./cmd/server
MIGRATIONS_PATH=./migrations
DATABASE_URL ?= postgres://postgres:postgres@localhost:5432/aithlete?sslmode=disable

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

migrate-up:
	@bash -c 'set -a; [ -f .env ] && source .env; \
	migrate -path $(MIGRATIONS_PATH) -database "$${DATABASE_URL}" up'

migrate-down:
	@bash -c 'set -a; [ -f .env ] && source .env; \
	migrate -path $(MIGRATIONS_PATH) -database "$${DATABASE_URL}" down'

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$name
