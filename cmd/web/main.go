package main

import (
  "database/sql"
  "flag"
  "html/template"
  "log/slog"
  "net/http"
  "os"
  "time"

  "github.com/alexedwards/scs/v2"
  "github.com/alexedwards/scs/mysqlstore"
  "github.com/jansuthacheeva/bookshelf/internal/models"
  _ "github.com/go-sql-driver/mysql"
  "github.com/joho/godotenv"
  "github.com/go-playground/form/v4"
)

type application struct {
  books           *models.BookModel
  logger          *slog.Logger
  templateCache   map[string]*template.Template
  users           *models.UserModel
  formDecoder     *form.Decoder
  sessionManager  *scs.SessionManager
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

  formDecoder := form.NewDecoder()

  sessionManager := scs.New()
  sessionManager.Store = mysqlstore.New(db)
  sessionManager.Lifetime = 12 * time.Hour

  app := &application{
    books: &models.BookModel{DB : db},
    logger: logger,
    templateCache: templateCache,
    users: &models.UserModel{DB: db},
    formDecoder: formDecoder,
    sessionManager: sessionManager,
  }

  srv := &http.Server{
    Addr:     *addr,
    Handler:  app.routes(),
  }

  logger.Info("Starting server at localhost", slog.String("addr", *addr))
  err = srv.ListenAndServe()
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
