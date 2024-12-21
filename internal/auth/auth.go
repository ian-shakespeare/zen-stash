package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/ian-shakespeare/zen-stash/pkg/models"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type AuthManager struct {
	signingKey []byte
}

func New(signingKey string) *AuthManager {
	return &AuthManager{[]byte(signingKey)}
}

func (a *AuthManager) GenerateToken(user *models.User, expiration time.Time) ([]byte, error) {
	tok, err := jwt.NewBuilder().
		Claim("id", user.ID).
		Claim("firstName", user.FirstName).
		Claim("lastName", user.LastName).
		Claim("email", user.Email).
		Claim("createdAt", user.CreatedAt).
		Issuer("zen-stash").
		Audience([]string{"users"}).
		Expiration(expiration).
		Build()
	if err != nil {
		return nil, err
	}

	return jwt.Sign(tok, jwt.WithKey(jwa.HS256(), a.signingKey))
}

func (a *AuthManager) ParseToken(tok []byte) (string, error) {
	parsed, err := jwt.Parse(tok, jwt.WithKey(jwa.HS256(), a.signingKey))
	if err != nil {
		return "", err
	}

	return a.getUserId(parsed)
}

func (a *AuthManager) ParseHeader(header http.Header) (string, error) {
	parsed, err := jwt.ParseHeader(header, "authorization")
	if err != nil {
		return "", err
	}

	return a.getUserId(parsed)
}

func (a *AuthManager) getUserId(tok jwt.Token) (string, error) {
	var userId interface{}
	err := tok.Get("id", &userId)
	if err != nil {
		return "", err
	}

	userIdStr, ok := userId.(string)
	if !ok {
		return "", errors.New("invalid user id")
	}

	return userIdStr, nil
}
