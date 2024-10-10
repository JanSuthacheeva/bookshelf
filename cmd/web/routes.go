package main

import (
  "net/http"

  "github.com/jansuthacheeva/bookshelf/ui"
)

func (app *application) routes() http.Handler {
  mux := http.NewServeMux()

  mux.Handle("GET /static/", http.FileServerFS(ui.Files))

  mux.HandleFunc("/", app.getHome)

  mux.HandleFunc("GET /login", app.getLogin)
  mux.HandleFunc("GET /register", app.getRegister)

  mux.HandleFunc("GET /books", app.getBooks)
  mux.HandleFunc("POST /books", app.postBooksCreate)
  mux.HandleFunc("GET /books/create", app.getBooksCreate)
  mux.HandleFunc("GET /books/{id}", app.getBookView)

  return mux
}
