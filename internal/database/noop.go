package database

import (
	"database/sql"
)

type NoOpConnection struct{}

func (_ NoOpConnection) Exec(query string, args ...any) (sql.Result, error) {
	return NoOpResult{}, nil
}

func (_ NoOpConnection) Query(query string, args ...any) (*sql.Rows, error) {
	return &sql.Rows{}, nil
}

func (_ NoOpConnection) QueryRow(query string, args ...any) *sql.Row {
	return &sql.Row{}
}

type NoOpResult struct{}

func (_ NoOpResult) LastInsertId() (int64, error) {
	return -1, nil
}

func (_ NoOpResult) RowsAffected() (int64, error) {
	return -1, nil
}
