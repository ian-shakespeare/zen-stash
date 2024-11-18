package main

import (
	"log"

	"github.com/ian-shakespeare/zen-stash/internal/database"
)

func main() {
	if err := database.Migrate(0); err != nil {
		log.Fatal(err.Error())
	}
}
