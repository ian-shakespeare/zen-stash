package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ian-shakespeare/zen-stash/internal/database"
)

func main() {
	conn, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = database.Migrate(conn)
	if err != nil {
		log.Fatal(err.Error())
	}

	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
