package main

import (
  "database/sql"
  "flag"
  "html/template"
  "log/slog"
  "net/http"
  "os"

  "github.com/jansuthacheeva/bookshelf/internal/models"
  _ "github.com/go-sql-driver/mysql"
  "github.com/joho/godotenv"
)

type application struct {
  books           *models.BookModel
  logger          *slog.Logger
  templateCache   map[string]*template.Template
  users           *models.UserModel
}

func main () {
  godotenv.Load()
  addr := flag.String("addr", os.Getenv("APPLICATION_PORT"), "HTTP network address.")
  dsn := flag.String("dsn", os.Getenv("GOOSE_DBSTRING"), "MySQL data source name.")

  logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
    AddSource: true,
  }))

  db, err := openDB(*dsn)
  if err != nil {
    logger.Error(err.Error())
    os.Exit(1)
  }
  // defer the close function so that it's always closing the connection before the main function is
  // finished.
  defer db.Close()

  templateCache, err := newTemplateCache()
  if err != nil {
    logger.Error(err.Error())
    os.Exit(1)
  }

  app := &application{
    books: &models.BookModel{DB : db},
    logger: logger,
    templateCache: templateCache,
    users: &models.UserModel{DB: db},
  }

  logger.Info("Starting server at localhost", slog.String("addr", *addr))
  err = http.ListenAndServe(*addr, app.routes())
  if err != nil {
    logger.Error(err.Error())
    os.Exit(1)
  }
}


func openDB(dsn string) (*sql.DB, error) {
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }

  err = db.Ping()
  if err != nil {
    db.Close()
    return nil, err
  }

  return db, nil
}
