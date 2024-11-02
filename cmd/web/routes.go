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
  mux.Handle("/home", dynamic.ThenFunc(app.getDashboard))
  mux.Handle("GET /sessions/create", dynamic.ThenFunc(app.getLogin))
  mux.Handle("POST /sessions/create",  dynamic.ThenFunc(app.postLogin))
  mux.Handle("GET /users/create", dynamic.ThenFunc(app.getRegister))
  mux.Handle("POST /users/create", dynamic.ThenFunc(app.postRegister))
  mux.Handle("POST /sessions/delete", dynamic.ThenFunc(app.postLogout))
  mux.Handle("GET /books", dynamic.ThenFunc(app.getBooks))
  mux.Handle("POST /books/create", dynamic.ThenFunc(app.postBooksCreate))
  mux.Handle("GET /books/create", dynamic.ThenFunc(app.getBooksCreate))
  mux.Handle("GET /books/{id}", dynamic.ThenFunc(app.getBookView))

  mux.Handle("GET /buttons/eye-open", dynamic.ThenFunc(app.getButtonEyeOpen))
  mux.Handle("GET /buttons/eye-closed", dynamic.ThenFunc(app.getButtonEyeClosed))

  standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

  return standard.Then(mux)
}
