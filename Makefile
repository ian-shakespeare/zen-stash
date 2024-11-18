CMD_DIR := cmd/
SQL_DIR := sql/
BIN_DIR := bin/

migrate:
	go build -o $(BIN_DIR)run-migrations  $(CMD_DIR)$(SQL_DIR)main.go

clean:
	rm -rf bin/
