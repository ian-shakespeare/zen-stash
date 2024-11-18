package database

import "embed"

//go:embed migrations/*.sql
var migrations embed.FS

func Migrate(version int) error {
	entries, err := migrations.ReadDir("migrations")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entry.Name()
	}

	return nil
}
