package main

import (
	"html/template"
	"log/slog"
	"net/http"
  "runtime/debug"
)

func (app *application) parseTemplates(w http.ResponseWriter, layout string, files *[]string, data any) error {
  tmpl, err := template.ParseFiles(*files...)
  if err != nil {
    return err
  }

  err = tmpl.ExecuteTemplate(w, layout, data)
  if err != nil {
    return err
  }
  return nil
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
    var (
      method  = r.Method
      uri     = r.URL.RequestURI()
      trace   = string(debug.Stack())
    )

    app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri), slog.String("trace", trace))
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
    http.Error(w, http.StatusText(status), status)
}
