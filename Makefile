CMD_DIR := cmd/
BIN_DIR := bin/

up:
	docker compose up --build web

down:
	docker compose down --remove-orphans

build:
	go build -o $(BIN_DIR)zen-stash $(CMD_DIR)main.go

clean:
	rm -rf bin/ .docker-storage/
