package main

import (
  "net/http"

  "github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
  mux := http.NewServeMux()

  fileServer := http.FileServer(http.Dir("./ui/static"))
  mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

  dynamic := alice.New(app.sessionManager.LoadAndSave)

  mux.Handle("/", dynamic.ThenFunc(app.getHome))
  mux.Handle("GET /login", dynamic.ThenFunc(app.getLogin))
  mux.Handle("GET /register", dynamic.ThenFunc(app.getRegister))
  mux.Handle("GET /books", dynamic.ThenFunc(app.getBooks))
  mux.Handle("POST /books/create", dynamic.ThenFunc(app.postBooksCreate))
  mux.Handle("GET /books/create", dynamic.ThenFunc(app.getBooksCreate))
  mux.Handle("GET /books/{id}", dynamic.ThenFunc(app.getBookView))

  standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

  return standard.Then(mux)
}
