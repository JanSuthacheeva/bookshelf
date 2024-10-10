package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func parseTemplates(w http.ResponseWriter, layout string, files *[]string) error {
  tmpl, err := template.ParseFiles(*files...)
  if err != nil {
    fmt.Println(err)
    return err
  }

  err = tmpl.ExecuteTemplate(w, layout, nil)
  if err != nil {
    fmt.Println(err)
    return err
  }
  return nil
}
