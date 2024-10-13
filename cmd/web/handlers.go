package main

import (
  "net/http"
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
  return
}

func (app *application) getBookView(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Here comes your book."))
}

