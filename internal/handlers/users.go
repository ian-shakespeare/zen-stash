package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func UserHandlers(db *sql.DB) http.Handler {
	users := http.NewServeMux()

	users.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		query := `
    INSERT INTO users (user_id)
    VALUES (DEFAULT);
    `

		_, err := db.Exec(query)
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte("bad insert"))
			return
		}

		w.WriteHeader(201)
		_, _ = w.Write([]byte("created"))
	})

	users.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		query := `
    SELECT user_id
    FROM users
    `

		rows, err := db.Query(query)
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte("bad select"))
			return
		}

		var users []string
		for rows.Next() {
			var userId string
			err = rows.Scan(&userId)
			if err != nil {
				w.WriteHeader(500)
				_, _ = w.Write([]byte("bad scan"))
				return
			}
			users = append(users, userId)
		}

		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte("bad marshal"))
			return
		}

		w.WriteHeader(200)
		_, _ = w.Write(b)
	})

	return users
}
