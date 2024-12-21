package handlers

import (
	"net/http"

	"github.com/ian-shakespeare/zen-stash/internal/auth"
	"github.com/ian-shakespeare/zen-stash/internal/database"
)

func New(db database.Connection, a *auth.AuthManager) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		_, _ = w.Write([]byte("OK"))
	})

	router.Handle("/users", UserHandlers(db, a))

	return router
}
