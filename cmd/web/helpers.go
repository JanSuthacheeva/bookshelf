package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page, tmplType string, data templateData) {
  tmpl, ok := app.templateCache[page]
  if !ok {
    err := fmt.Errorf("The template %s does not exist.", page)
    app.serverError(w, r, err)
    return
  }
  buf := new(bytes.Buffer)
  err := tmpl.ExecuteTemplate(buf, tmplType, data)
  if err != nil {
    app.serverError(w, r, err)
  }

  w.WriteHeader(status)

  buf.WriteTo(w)
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

func (app *application) transformDateStringToSqlNullTime(dateString string) (sql.NullTime, error) {

  var nullTime sql.NullTime

  if dateString == "" {
    nullTime = sql.NullTime{
      Valid: false,
    }
  } else {
    date, err := time.Parse("2006-02-02", dateString)
    if err != nil {
      app.logger.Info(err.Error())
      return nullTime, err
    }
    nullTime = sql.NullTime{
      Time: date,
      Valid: true,
    }
  }

  return nullTime, nil
}
