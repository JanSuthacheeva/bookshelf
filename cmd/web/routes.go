package main

import (
  "net/http"
)

func (app *application) routes() http.Handler {
  mux := http.NewServeMux()

  fileServer := http.FileServer(http.Dir("./ui/static"))
  mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

  mux.HandleFunc("/", app.getHome)

  mux.HandleFunc("GET /login", app.getLogin)
  mux.HandleFunc("GET /register", app.getRegister)

  mux.HandleFunc("GET /books", app.getBooks)
  mux.HandleFunc("POST /books", app.postBooksCreate)
  mux.HandleFunc("GET /books/create", app.getBooksCreate)
  mux.HandleFunc("GET /books/{id}", app.getBookView)

  return app.logRequest(commonHeaders(mux))
}
