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
  app.render(w, r, http.StatusOK, "login.tmpl.html", "base_guest", templateData{})
}

func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, http.StatusOK, "register.tmpl.html", "base_guest", templateData{})
}


func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, http.StatusOK, "home.tmpl.html", "base_guest", templateData{})
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
  
  data := app.newTemplateData(r)
  data.Book = book

  app.render(w, r, http.StatusOK, "books_view.tmpl.html", "base_auth", data)
}

