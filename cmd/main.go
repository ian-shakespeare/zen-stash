package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ian-shakespeare/zen-stash/internal/database"
	"github.com/ian-shakespeare/zen-stash/internal/handlers"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	err = database.Migrate(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	server := handlers.New(db)

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
