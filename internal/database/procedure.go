package database

import "embed"

//go:embed procedures/*.sql
var procedureFiles embed.FS

func LoadProcedures(c Connection) error {
	procedures, err := procedureFiles.ReadDir("procedures")
	if err != nil {
		return err
	}

	for _, p := range procedures {
		procedure, err := procedureFiles.ReadFile("procedures/" + p.Name())
		if err != nil {
			return err
		}

		if _, err = c.Exec(string(procedure)); err != nil {
			return err
		}
	}

	return nil
}
