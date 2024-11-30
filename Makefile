CMD_DIR := cmd/
BIN_DIR := bin/
BUILD_FLAGS := CGO_ENABLED=0

all: up

up:
	docker compose up --build web

down:
	docker compose down --remove-orphans

build:
	$(BUILD_FLAGS) go build -o $(BIN_DIR)zen-stash $(CMD_DIR)main.go

lint:
	golangci-lint run ./...

test-unit:
	go test ./...

test: test-unit

clean:
	rm -rf bin/ .docker-storage/

.PHONY: all up down lint test-unit test clean
