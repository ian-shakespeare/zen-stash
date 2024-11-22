CMD_DIR := cmd/
SQL_DIR := sql/
BIN_DIR := bin/

up:
	docker compose up

down:
	docker compose down --remove-orphans

build:
	go build -o $(BIN_DIR)zen-stash $(CMD_DIR)main.go

clean:
	rm -rf bin/ .docker-storage/
