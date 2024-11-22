package database

import (
	"database/sql"
	"embed"
	"errors"
	"io/fs"
	"sort"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

const MIGRATIONS_TABLE = `
CREATE TABLE IF NOT EXISTS migrations (
  version SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT now()
)
`

func Migrate(conn *sql.DB) error {
	entries, err := migrationFiles.ReadDir("migrations")
	if err != nil {
		return err
	}

	if len(entries)%2 == 1 {
		return errors.New("received an odd number of migrations")
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	ups := []fs.DirEntry{}
	downs := []fs.DirEntry{}

	for i := 0; i < len(entries); i += 2 {
		downs = append(downs, entries[i])
		ups = append(ups, entries[i+1])
	}

	dbVersion := getDatabaseVersion(conn)
	for i := dbVersion; i < len(ups); i += 1 {
		migration, err := migrationFiles.ReadFile("migrations/" + ups[i].Name())
		if err != nil {
			return err
		}

		if _, err = conn.Exec(string(migration)); err != nil {
			return err
		}
	}

	return nil
}

func getDatabaseVersion(conn *sql.DB) int {
	row := conn.QueryRow("SELECT MAX(version) AS version FROM migrations")

	var version int
	err := row.Scan(&version)
	if err != nil {
		return 0
	}

	return version
}
