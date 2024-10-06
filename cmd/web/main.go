package main

import (
  "net/http"
  "flag"
  "fmt"
  "os"
)

func main () {
  addr := flag.String("addr", ":4444", "HTTP network address.")
  server := &http.Server{
    Addr: *addr,
    Handler: routes(),
  }


  fmt.Printf("Starting server at localhost%s\n", *addr)
  err := server.ListenAndServe()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
