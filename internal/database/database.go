package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	username, exists := os.LookupEnv("DB_USER")
	if !exists {
		log.Fatal("missing env DB_USER")
	}

	password, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		log.Fatal("missing env DB_PASSWORD")
	}

	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		log.Fatal("missing env DB_HOST")
	}

	name, exists := os.LookupEnv("DB_NAME")
	if !exists {
		log.Fatal("missing env DB_NAME")
	}

	// TODO: Probably should enable ssl in prod.
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, 5432, username, password, name)

	return sql.Open("postgres", connStr)
}
