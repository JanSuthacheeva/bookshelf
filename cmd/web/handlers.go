package main

import (
  "fmt"
  "net/http"
  "html/template"
)


func (app *application) getDashboard(w http.ResponseWriter, r *http.Request) {
  tmpl, err := template.ParseFiles("./ui/html/pages/dashboard.tmpl.html")
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

  err = tmpl.Execute(w, nil)
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
  }
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

