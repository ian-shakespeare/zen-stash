package handlers

import (
	"database/sql"
	"net/http"
)

func New(db *sql.DB) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		_, _ = w.Write([]byte("OK"))
	})

	router.Handle("/users", UserHandlers(db))

	return router
}
