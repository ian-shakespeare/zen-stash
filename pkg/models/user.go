package models

import "time"

type User struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	PasswordDigest string    `json:"passwordDigest"`
	CreatedAt      time.Time `json:"createdAt"`
}
