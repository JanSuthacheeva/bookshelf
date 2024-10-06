package main

import (
  "net/http"
)

func (app *application) routes() http.Handler {
  mux := http.NewServeMux()

  mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("Hello World!"))
  })

  mux.HandleFunc("GET /books", func (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Here are the books."))
  })

  mux.HandleFunc("POST /books", func (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("New books here."))
  })

  mux.HandleFunc("GET /books/{id}", func (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Here comes your book."))
  })

  return mux
}
