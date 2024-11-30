package handlers_test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/ian-shakespeare/zen-stash/internal/database"
	"github.com/ian-shakespeare/zen-stash/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestUserHandlers(t *testing.T) {
	t.Parallel()

	h := handlers.UserHandlers(database.NoOpConnection{})

	t.Run("post/invalidForm", func(t *testing.T) {
		t.Parallel()

		w := NewResponseWriter()

		r, err := http.NewRequest("POST", "/", nil)
		assert.NoError(t, err)

		h.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.StatusCode)
		assert.Contains(t, string(w.Body), "Invalid form")
	})

	t.Run("post/invalidFirstName", func(t *testing.T) {
		t.Parallel()

		w := NewResponseWriter()

		tooShort := []string{
			"",
			"   a   ",
			"\t\r\na\t\r\n",
		}
		for _, short := range tooShort {
			form := url.Values{}
			form.Add("firstName", short)

			r, err := http.NewRequest("POST", "/", strings.NewReader(form.Encode()))
			assert.NoError(t, err)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			h.ServeHTTP(w, r)
			assert.Equal(t, http.StatusBadRequest, w.StatusCode)
			assert.Contains(t, string(w.Body), "First name must be at least 2 characters in length")
			w.Reset()
		}

		tooLong := []string{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   ",
			"\t\r\naaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\t\r\n",
		}
		for _, long := range tooLong {
			form := url.Values{}
			form.Add("firstName", long)

			r, err := http.NewRequest("POST", "/", strings.NewReader(form.Encode()))
			assert.NoError(t, err)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			h.ServeHTTP(w, r)
			assert.Equal(t, http.StatusBadRequest, w.StatusCode)
			assert.Contains(t, string(w.Body), "First name must be no more than 64 characters in length")
			w.Reset()
		}
	})

	t.Run("post/invalidLastName", func(t *testing.T) {
		t.Parallel()

		w := NewResponseWriter()

		tooShort := []string{
			"",
			"   a   ",
			"\t\r\na\t\r\n",
		}
		for _, short := range tooShort {
			form := url.Values{}
			form.Add("firstName", "firstName")
			form.Add("lastName", short)

			r, err := http.NewRequest("POST", "/", strings.NewReader(form.Encode()))
			assert.NoError(t, err)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			h.ServeHTTP(w, r)
			assert.Equal(t, http.StatusBadRequest, w.StatusCode)
			assert.Contains(t, string(w.Body), "Last name must be at least 2 characters in length")
			w.Reset()
		}

		tooLong := []string{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   ",
			"\t\r\naaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\t\r\n",
		}
		for _, long := range tooLong {
			form := url.Values{}
			form.Add("firstName", "firstName")
			form.Add("lastName", long)

			r, err := http.NewRequest("POST", "/", strings.NewReader(form.Encode()))
			assert.NoError(t, err)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			h.ServeHTTP(w, r)
			assert.Equal(t, http.StatusBadRequest, w.StatusCode)
			assert.Contains(t, string(w.Body), "Last name must be no more than 64 characters in length")
			w.Reset()
		}
	})

	t.Run("post/invalidEmail", func(t *testing.T) {
		t.Parallel()

		w := NewResponseWriter()

		badEmail := []string{
			"",
			"@",
			"jdoe@",
			"jdoe@.com",
			"@gmail",
			"@gmail.com",
		}
		for _, email := range badEmail {
			form := url.Values{}
			form.Add("firstName", "firstName")
			form.Add("lastName", "lastName")
			form.Add("email", email)

			r, err := http.NewRequest("POST", "/", strings.NewReader(form.Encode()))
			assert.NoError(t, err)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			h.ServeHTTP(w, r)
			assert.Equal(t, http.StatusBadRequest, w.StatusCode)
			assert.Contains(t, string(w.Body), "Invalid email")
			w.Reset()
		}
	})

	t.Run("post/invalidPassword", func(t *testing.T) {
		t.Parallel()

		w := NewResponseWriter()

		tooShort := []string{
			"",
			"aaaaaaa",
			"   aaaaaaa   ",
			"\t\r\naaaaaaa\t\r\n",
		}
		for _, short := range tooShort {
			form := url.Values{}
			form.Add("firstName", "firstName")
			form.Add("lastName", "lastName")
			form.Add("email", "test@test.com")
			form.Add("password", short)

			r, err := http.NewRequest("POST", "/", strings.NewReader(form.Encode()))
			assert.NoError(t, err)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			h.ServeHTTP(w, r)
			assert.Equal(t, http.StatusBadRequest, w.StatusCode)
			assert.Contains(t, string(w.Body), "Password must be at least 8 characters in length")
			w.Reset()
		}

		tooLong := []string{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"   aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa   ",
			"\t\r\naaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\t\r\n",
		}
		for _, long := range tooLong {
			form := url.Values{}
			form.Add("firstName", "firstName")
			form.Add("lastName", "lastName")
			form.Add("email", "test@test.com")
			form.Add("password", long)

			r, err := http.NewRequest("POST", "/", strings.NewReader(form.Encode()))
			assert.NoError(t, err)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			h.ServeHTTP(w, r)
			assert.Equal(t, http.StatusBadRequest, w.StatusCode)
			assert.Contains(t, string(w.Body), "Password must be no more than 72 characters in length")
			w.Reset()
		}
	})
}
