package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/jansuthacheeva/bookshelf/internal/models"
)

type bookCreateForm struct {
  Title		  string
  Author	  string
  Started	  string
  Finished	  string
  FieldErrors	  map[string]string
}


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
  data := app.newTemplateData(r)
  data.Form = bookCreateForm{}
  app.render(w, r, http.StatusOK, "books_create.tmpl.html", "base_auth", data)
}

func (app *application) postBooksCreate(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    app.clientError(w, http.StatusBadRequest)
    return
  }

  form := bookCreateForm{
    Title:	  r.PostForm.Get("title"),
    Author:   	  r.PostForm.Get("author"),
    Started:  	  r.PostForm.Get("started"),
    Finished: 	  r.PostForm.Get("finished"),
    FieldErrors:  map[string]string{},
  }

  started, err := app.transformDateStringToSqlNullTime(form.Started)
  if err != nil {
    form.FieldErrors["started"] = "This field must be a valid date.";
    return
  }

  finished, err := app.transformDateStringToSqlNullTime(form.Finished)
  if err != nil {
    form.FieldErrors["finished"] = "This field must be a valid date.";
    return
  }


  if strings.TrimSpace(form.Title) == "" {
    form.FieldErrors["title"] = "This field cannot be blank."
  } else if utf8.RuneCountInString(form.Title) > 120 {
    form.FieldErrors["title"] = "This field cannot be more than 120 characters long."
  }

  if strings.TrimSpace(form.Author) == "" {
    form.FieldErrors["author"] = "This field cannot be blank."
  } else if utf8.RuneCountInString(form.Author) > 120 {
    form.FieldErrors["author"] = "This field cannot be more than 120 characters long."
  }

  if len(form.FieldErrors) > 0 {
    data := app.newTemplateData(r)
    data.Form = form
    tmpl, err := template.ParseFiles("./ui/html/bookCreateForm.html")
    if err != nil {
      app.serverError(w, r, err)
    }
    w.WriteHeader(http.StatusUnprocessableEntity)
    err = tmpl.ExecuteTemplate(w, "bookCreateForm", data)
    if err != nil {
      app.serverError(w, r, err)
    }
    return
  }

  id, err := app.books.Insert(form.Title, form.Author, started, finished)
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

