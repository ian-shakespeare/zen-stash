package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ian-shakespeare/zen-stash/internal/database"
	"github.com/ian-shakespeare/zen-stash/internal/handlers"
	"github.com/ian-shakespeare/zen-stash/pkg/utils"
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

	port := utils.FallbackEnv("PORT", "8080")
	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}
