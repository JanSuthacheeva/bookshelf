package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
  files := []string{
    "./ui/html/guest_base.tmpl.html",
    "./ui/html/pages/home.tmpl.html",
  }
  tmpl, err := template.ParseFiles(files...)
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

  err = tmpl.ExecuteTemplate(w, "base_guest", nil)
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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

