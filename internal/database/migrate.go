package database

import (
	"embed"
	"errors"
	"io/fs"
	"sort"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

const MIGRATIONS_TABLE = `
CREATE TABLE IF NOT EXISTS sql_migrations (
  version SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT now()
)
`

func Migrate(c Connection) error {
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
	// downs := []fs.DirEntry{}

	for i := 0; i < len(entries); i += 2 {
		// downs = append(downs, entries[i])
		ups = append(ups, entries[i+1])
	}

	dbVersion := getDatabaseVersion(c)
	for i := dbVersion; i < len(ups); i += 1 {
		migration, err := migrationFiles.ReadFile("migrations/" + ups[i].Name())
		if err != nil {
			return err
		}

		if _, err = c.Exec(string(migration)); err != nil {
			return err
		}
	}

	return nil
}

func getDatabaseVersion(c Connection) int {
	row := c.QueryRow("SELECT MAX(version) AS version FROM sql_migrations")

	var version int
	err := row.Scan(&version)
	if err != nil {
		return 0
	}

	return version
}
