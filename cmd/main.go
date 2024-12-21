package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ian-shakespeare/zen-stash/internal/auth"
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

	err = database.LoadProcedures(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	signingKey, exists := os.LookupEnv("SIGNING_KEY")
	if !exists {
		log.Fatal("missing env `SIGNING_KEY`")
	}

	a := auth.New(signingKey)

	server := handlers.New(db, a)

	port := utils.FallbackEnv("PORT", "8080")
	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}
