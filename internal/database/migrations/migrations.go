package migrations

import "embed"

//go:embed *.sql
var migrations embed.FS

func Run(version int) {
}
