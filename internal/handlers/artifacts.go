package handlers

import (
	"database/sql"
	"net/http"
)

func ArtifactHandlers(db *sql.DB) http.Handler {
	artifacts := http.NewServeMux()

	return artifacts
}
