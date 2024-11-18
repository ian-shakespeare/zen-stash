package database

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"sort"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func Migrate(version int) error {
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

	for _, up := range ups {
		b, err := migrationFiles.ReadFile("migrations/" + up.Name())
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	}

	return nil
}
