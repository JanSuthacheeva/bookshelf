package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jansuthacheeva/bookshelf/internal/models"
)


func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
  files := []string{
    "./ui/html/guest_base.tmpl.html",
    "./ui/html/pages/login.tmpl.html",
  }

  err := app.parseTemplates(w, "base_guest", &files)
  if err != nil {
    app.serverError(w, r, err)
    return
  }
}

func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
  files := []string{
    "./ui/html/guest_base.tmpl.html",
    "./ui/html/pages/register.tmpl.html",
  }

  err := app.parseTemplates(w, "base_guest", &files)
  if err != nil {
    app.serverError(w, r, err)
    return
  }
}


func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
  files := []string{
    "./ui/html/guest_base.tmpl.html",
    "./ui/html/pages/home.tmpl.html",
  }

  err := app.parseTemplates(w, "base_guest", &files)
  if err != nil {
    app.serverError(w, r, err)
    return
  }
}

func (app *application) getDashboard(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Dashboard."))
}

func (app *application) getBooks(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Here are the books."))
}

func (app *application) getBooksCreate(w http.ResponseWriter, r *http.Request) {
  return
}

func (app *application) postBooksCreate(w http.ResponseWriter, r *http.Request) {
  title := "Let's Go"
  author := "Alex Edwards"
  started := sql.NullTime{
    Valid: false,
  }
  finished := sql.NullTime{
    Valid: false,
  }

  id, err := app.books.Insert(title, author, started, finished)
  if err != nil {
    app.serverError(w, r, err)
    return
  }

  http.Redirect(w, r, fmt.Sprintf("/books/%d", id), http.StatusSeeOther)
}

func (app *application) getBookView(w http.ResponseWriter, r *http.Request) {
  files := []string{
    "./ui/html/authenticated_base.tmpl.html",
    "./ui/html/partials/nav.tmpl.html",
    "./ui/html/pages/books/view.tmpl.html",
  }
  err := app.parseTemplates(w, "base_auth", &files)
  if err != nil {
    app.serverError(w, r, err)
    return
  }
  id, err := strconv.Atoi(r.PathValue("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }

  book, err := app.books.Get(id)
  if err != nil {
    if errors.Is(err, models.ErrNoRecord) {
      http.NotFound(w, r)
    } else {
      app.serverError(w, r,   err)
    }
    return
  }

  fmt.Fprintf(w, "%v", book)
}

