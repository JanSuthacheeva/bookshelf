package main

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/jansuthacheeva/bookshelf/internal/assert"
)

func TestUserCreate(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/users/create")
	validCSRFToken := extractCSRFToken(t, body)

	const (
		validName            = "Bob"
		validPassword        = "pa$$word"
		validPasswordConfirm = "pa$$word"
		validEmail           = "bob@example.com"
		formTag              = `<form id="registerForm">`
	)

	tests := []struct {
		name            string
		userName        string
		password        string
		passwordConfirm string
		email           string
		csrfToken       string
		wantCode        int
		wantFormTag     string
	}{
		{
			name:            "Valid submission",
			userName:        validName,
			password:        validPassword,
			passwordConfirm: validPasswordConfirm,
			email:           validEmail,
			csrfToken:       validCSRFToken,
			wantCode:        http.StatusSeeOther,
		},
		{
			name:            "Invalid CSRF Token",
			userName:        validName,
			password:        validPassword,
			passwordConfirm: validPasswordConfirm,
			email:           validEmail,
			csrfToken:       "wrongToken",
			wantCode:        http.StatusBadRequest,
		},
		{
			name:            "Empty Name",
			userName:        "",
			password:        validPassword,
			passwordConfirm: validPasswordConfirm,
			email:           validEmail,
			csrfToken:       validCSRFToken,
			wantCode:        http.StatusUnprocessableEntity,
			wantFormTag:     formTag,
		},
		{
			name:            "Empty Email",
			userName:        validName,
			password:        validPassword,
			passwordConfirm: validPasswordConfirm,
			email:           "",
			csrfToken:       validCSRFToken,
			wantCode:        http.StatusUnprocessableEntity,
			wantFormTag:     formTag,
		},
		{
			name:            "Empty Password",
			userName:        validName,
			password:        "",
			passwordConfirm: validPasswordConfirm,
			email:           validEmail,
			csrfToken:       validCSRFToken,
			wantCode:        http.StatusUnprocessableEntity,
			wantFormTag:     formTag,
		},
		{
			name:            "Empty Password Confirm",
			userName:        validName,
			password:        validPassword,
			passwordConfirm: "",
			email:           validEmail,
			csrfToken:       validCSRFToken,
			wantCode:        http.StatusUnprocessableEntity,
			wantFormTag:     formTag,
		},
		{
			name:            "Invalid Email",
			userName:        validName,
			password:        validPassword,
			passwordConfirm: validPasswordConfirm,
			email:           "bob@example.",
			csrfToken:       validCSRFToken,
			wantCode:        http.StatusUnprocessableEntity,
			wantFormTag:     formTag,
		},
		{
			name:            "Short Password",
			userName:        validName,
			password:        "dudu",
			passwordConfirm: "dudu",
			email:           validEmail,
			csrfToken:       validCSRFToken,
			wantCode:        http.StatusUnprocessableEntity,
			wantFormTag:     formTag,
		},
		{
			name:            "Not Same Password",
			userName:        validName,
			password:        "dududu31",
			passwordConfirm: "dududu32",
			email:           validEmail,
			csrfToken:       validCSRFToken,
			wantCode:        http.StatusUnprocessableEntity,
			wantFormTag:     formTag,
		},
		{
			name:            "Duplicate Mail",
			userName:        validName,
			password:        validPassword,
			passwordConfirm: validPasswordConfirm,
			email:           "duplicate@example.com",
			csrfToken:       validCSRFToken,
			wantCode:        http.StatusUnprocessableEntity,
			wantFormTag:     formTag,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tt.userName)
			form.Add("email", tt.email)
			form.Add("password", tt.password)
			form.Add("password_confirm", tt.passwordConfirm)
			form.Add("csrf_token", tt.csrfToken)

			code, _, body := ts.postForm(t, "/users/create", form)

			assert.Equal(t, code, tt.wantCode)

			if tt.wantFormTag != "" {
				assert.StringContains(t, body, tt.wantFormTag)
			}
		})
	}
}

func TestBookView(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/books/1",
			wantCode: http.StatusOK,
			wantBody: "A mock book",
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/books/2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlPath:  "/books/-2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal  ID",
			urlPath:  "/books/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/books/foo",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			assert.Equal(t, code, tt.wantCode)

			if tt.wantBody != "" {
				assert.StringContains(t, body, tt.wantBody)
			}
		})
	}
}

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
