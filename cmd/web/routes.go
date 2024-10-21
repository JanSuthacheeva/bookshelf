package main

import (
  "net/http"

  "github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
  mux := http.NewServeMux()

  fileServer := http.FileServer(http.Dir("./ui/static"))
  mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

  mux.HandleFunc("/", app.getHome)

  mux.HandleFunc("GET /login", app.getLogin)
  mux.HandleFunc("GET /register", app.getRegister)

  mux.HandleFunc("GET /books", app.getBooks)
  mux.HandleFunc("POST /books/create", app.postBooksCreate)
  mux.HandleFunc("GET /books/create", app.getBooksCreate)
  mux.HandleFunc("GET /books/{id}", app.getBookView)

  standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

  return standard.Then(mux)
}
