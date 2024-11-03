package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)
	public := dynamic.Append(app.alreadyAuthenticated)
	protected := dynamic.Append(app.requireAuth)

	mux.Handle("/", public.ThenFunc(app.getHome))
	mux.Handle("GET /users/create", public.ThenFunc(app.getRegister))
	mux.Handle("POST /users/create", public.ThenFunc(app.postRegister))
	mux.Handle("GET /sessions/create", public.ThenFunc(app.getSessionCreate))
	mux.Handle("POST /sessions/create", public.ThenFunc(app.postSessionCreate))

	mux.Handle("/home", protected.ThenFunc(app.getDashboard))
	mux.Handle("POST /sessions/delete", protected.ThenFunc(app.postSessionDelete))
	mux.Handle("GET /books", protected.ThenFunc(app.getBooks))
	mux.Handle("POST /books/create", protected.ThenFunc(app.postBooksCreate))
	mux.Handle("GET /books/create", protected.ThenFunc(app.getBooksCreate))
	mux.Handle("GET /books/{id}", protected.ThenFunc(app.getBookView))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
