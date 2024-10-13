package main

import (
  "database/sql"
  "net/http"
  "flag"
  "fmt"
  "os"

  "github.com/jansuthacheeva/bookshelf/internal/models"
  _ "github.com/go-sql-driver/mysql"
)

type application struct {
  books     *models.BookModel
  users     *models.UserModel
}

func main () {
  addr := flag.String("addr", ":4444", "HTTP network address.")
  dsn := flag.String("dsn", "web:pass@/bookshelf?parseTime=true", "MySQL data source name.")

  db, err := openDB(*dsn)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
  // defer the close function so that it's always closing the connection before the main function is
  // finished.
  defer db.Close()

  app := &application{
    books: &models.BookModel{DB : db},
  }

  fmt.Printf("Starting server at localhost%s\n", *addr)
  err = http.ListenAndServe(*addr, app.routes())
  if err != nil {
    fmt.Println(err)
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
