package handlers

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/ian-shakespeare/zen-stash/internal/auth"
	"github.com/ian-shakespeare/zen-stash/internal/database"
	"github.com/ian-shakespeare/zen-stash/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthResponse struct {
	AccessToken string    `json:"accessToken"`
	Expiration  time.Time `json:"expiration"`
}

func AuthHandlers(db database.Connection, a *auth.AuthManager) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
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

		err = database.CreateUser(db, firstName, lastName, email, string(passwordDigest))
		if err != nil {
			_ = NewHandlerError("Could not create user", err).Send(w, http.StatusInternalServerError)
			return
		}

		u, err := database.GetUser(db, email)
		if err != nil {
			_ = NewHandlerError("Could not create user", err).Send(w, http.StatusInternalServerError)
			return
		}

		expires := utils.TwoWeeksFromNow()
		tok, err := a.GenerateToken(&u, expires)
		if err != nil {
			_ = NewHandlerError("Could not sign in", err).Send(w, http.StatusInternalServerError)
			return
		}

		res := AuthResponse{
			AccessToken: string(tok),
			Expiration:  expires,
		}

		b, err := json.Marshal(res)
		if err != nil {
			_ = NewHandlerError("Could not send access token", err).Send(w, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(201)
		_, _ = w.Write(b)
	})

	router.HandleFunc("POST /signin", func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil || r.ParseForm() != nil {
			_ = NewHandlerError("Invalid form", nil).Send(w, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		email := strings.Trim(r.FormValue("email"), " \t\r\n")
		password := strings.Trim(r.FormValue("password"), " \t\r\n")

		u, err := database.GetUser(db, email)
		if err != nil {
			_ = NewHandlerError("User not found", err).Send(w, http.StatusNotFound)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password)); err != nil {
			_ = NewHandlerError("Invalid credentials", err).Send(w, http.StatusNotFound)
			return
		}

		expires := utils.TwoWeeksFromNow()
		tok, err := a.GenerateToken(&u, expires)
		if err != nil {
			_ = NewHandlerError("Could not sign in", err).Send(w, http.StatusInternalServerError)
			return
		}

		res := AuthResponse{
			AccessToken: string(tok),
			Expiration:  expires,
		}

		b, err := json.Marshal(res)
		if err != nil {
			_ = NewHandlerError("Could not send access token", err).Send(w, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		_, _ = w.Write(b)
	})

	return router
}
