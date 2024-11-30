package handlers

import (
	"net/http"
	"net/mail"
	"strings"

	"github.com/ian-shakespeare/zen-stash/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func UserHandlers(db database.Connection) http.Handler {
	users := http.NewServeMux()

	users.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil || r.ParseForm() != nil {
			_ = NewHandlerError("Invalid form", nil).Send(w, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		firstName := strings.Trim(r.FormValue("firstName"), " \t\r\n")
		if len(firstName) < 2 {
			_ = NewHandlerError("First name must be at least 2 characters in length", nil).Send(w, http.StatusBadRequest)
			return
		} else if len(firstName) > 64 {
			_ = NewHandlerError("First name must be no more than 64 characters in length", nil).Send(w, http.StatusBadRequest)
			return
		}

		lastName := strings.Trim(r.FormValue("lastName"), " \t\r\n")
		if len(lastName) < 2 {
			_ = NewHandlerError("Last name must be at least 2 characters in length", nil).Send(w, http.StatusBadRequest)
			return
		} else if len(lastName) > 64 {
			_ = NewHandlerError("Last name must be no more than 64 characters in length", nil).Send(w, http.StatusBadRequest)
			return
		}

		email := strings.Trim(r.FormValue("email"), " \t\r\n")
		if _, err := mail.ParseAddress(email); err != nil {
			_ = NewHandlerError("Invalid email address", err).Send(w, http.StatusBadRequest)
			return
		}

		password := strings.Trim(r.FormValue("password"), " \t\r\n")
		if len(password) < 8 {
			_ = NewHandlerError("Password must be at least 8 characters in length", nil).Send(w, http.StatusBadRequest)
			return
		} else if len(password) > 72 {
			_ = NewHandlerError("Password must be no more than 72 characters in length", nil).Send(w, http.StatusBadRequest)
			return
		}

		passwordDigest, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		if err != nil {
			_ = NewHandlerError("Invalid password", err).Send(w, http.StatusBadRequest)
			return
		}

		_, err = db.Exec("CALL create_user($1, $2, $3, $4)", firstName, lastName, email, passwordDigest)
		if err != nil {
			_ = NewHandlerError("could not create user", err).Send(w, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(201)
		_, _ = w.Write([]byte("created"))
	})

	return users
}
