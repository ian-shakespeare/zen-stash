package database

import (
	"database/sql"

	"github.com/ian-shakespeare/zen-stash/pkg/models"
)

func CreateUser(conn Connection, firstName, lastName, email, passwordDigest string) error {
	_, err := conn.Exec("CALL create_user($1, $2, $3, $4)", firstName, lastName, email, passwordDigest)
	return err
}

func GetUser(conn Connection, email string) (models.User, error) {
	query := `
  SELECT
    user_id,
    first_name,
    last_name,
    email,
    password_digest,
    created_at
  FROM users
  WHERE email = $1
  `

	row := conn.QueryRow(query, email)
	if row == nil {
		return models.User{}, sql.ErrNoRows
	}

	var u models.User
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.PasswordDigest, &u.CreatedAt)

	return u, err
}
