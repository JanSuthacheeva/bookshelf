package main

import (
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
  app.render(w, r, http.StatusOK, "books_create.tmpl.html", "base_auth", templateData{})
}

func (app *application) postBooksCreate(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    app.clientError(w, http.StatusBadRequest)
    return
  }
  title := r.PostForm.Get("title")
  author := r.PostForm.Get("author")
  startedReq := r.PostForm.Get("started")
  finishedReq := r.PostForm.Get("finished")

  // var started sql.NullTime
  // if startedReq == "" {
  //   started = sql.NullTime{
  //     Valid: false,
  //   }
  // } else {
  //   parsedDate, err := time.Parse("2006-02-02", startedReq)
  //   if err != nil {
  //     app.logger.Info(err.Error())
  //     app.clientError(w, http.StatusBadRequest)
  //     return
  //   }
  //   started = sql.NullTime{
  //     Time: parsedDate,
  //     Valid: true,
  //   }
  // }
  started, err := app.transformDateStringToSqlNullTime(startedReq)
  if err != nil {
    app.clientError(w, http.StatusBadRequest)
    return
  }

  finished, err := app.transformDateStringToSqlNullTime(finishedReq)
  if err != nil {
    app.clientError(w, http.StatusBadRequest)
    return
  }

  if title == "" {
    app.render(w, r, http.StatusUnprocessableEntity, "books_create.tmpl.html", "createBookForm", templateData{})
  }
  if author == "" {
    app.render(w, r, http.StatusUnprocessableEntity, "books_create.tmpl.html", "createBookForm", templateData{})
  }

  id, err := app.books.Insert(title, author, started, finished)
  if err != nil {
    app.serverError(w, r, err)
    return
  }
  w.Header().Set("HX-Redirect", fmt.Sprintf("/books/%d", id));
  w.WriteHeader(http.StatusSeeOther)

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

