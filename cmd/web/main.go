package main

import (
  "database/sql"
  "net/http"
  "flag"
  "fmt"
  "os"

  _ "github.com/go-sql-driver/mysql"
)

func main () {
  addr := flag.String("addr", ":4444", "HTTP network address.")
  dsn := flag.String("dsn", "web:pass@/bookshelf?parseTime=true", "MySQL data source name.")
  server := &http.Server{
    Addr: *addr,
    Handler: routes(),
  }

  db, err := openDB(*dsn)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
  // defer the close function so that it's always closing the connection before the main function is
  // finished.
  defer db.Close()

  fmt.Printf("Starting server at localhost%s\n", *addr)
  err = server.ListenAndServe()
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
