package main

import (
	"html/template"
	"path/filepath"

	"github.com/jansuthacheeva/bookshelf/internal/models"
)

type templateData struct{
  Book models.Book
}

func newTemplateCache() (map[string]*template.Template, error) {

  cache := map[string]*template.Template{}

  guestPages, err := filepath.Glob("./ui/html/pages/guestPages/*.tmpl.html")
  if err != nil {
    return nil, err
  }
  authPages, err := filepath.Glob("./ui/html/pages/authPages/*.tmpl.html")
  if err != nil {
    return nil, err
  }

  cache, err = loopOverPages(cache, guestPages, "guest")
  if err != nil {
    return nil, err
  }
  cache, err = loopOverPages(cache, authPages, "auth")
  if err != nil {
    return nil, err
  }

  return cache, nil
}

func loopOverPages(cache map[string]*template.Template, pages []string, templateType string) (map[string]*template.Template, error) {

  var baseTemplate = []string{}
  if templateType == "auth" {
    baseTemplate = append([]string{"./ui/html/authenticated_base.tmpl.html"},"./ui/html/partials/nav.tmpl.html")
  } else {
    baseTemplate = append(baseTemplate, "./ui/html/guest_base.tmpl.html")
  }
  for _, page := range pages {
    name := filepath.Base(page)


    files := append(baseTemplate, page)

    ts, err := template.ParseFiles(files...)
    if err != nil {
      return nil, err
    }

    cache[name] = ts
  }

  return cache, nil
}
