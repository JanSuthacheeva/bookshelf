package main

import (
  "net/http"
)

func (app *application) routes() http.Handler {
  mux := http.NewServeMux()

  mux.HandleFunc("/", app.getDashboard)

  mux.HandleFunc("GET /books", app.getBooks)
  mux.HandleFunc("POST /books", app.postBooksCreate)
  mux.HandleFunc("GET /books/create", app.getBooksCreate)
  mux.HandleFunc("GET /books/{id}", app.getBookView)

  return mux
}
