package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jansuthacheeva/bookshelf/internal/models"
	"github.com/jansuthacheeva/bookshelf/internal/validator"
)

type bookCreateForm struct {
  Title		  string  `form:"title"`
  Author	  string  `form:"author"`
  Started	  string  `form:"started"`
  Finished	  string  `form:"finished"`
  validator.Validator	  `form:"-"`
}

type userRegisterForm struct {
  Name		      string  `form:"name"`
  Email		      string  `form:"email"`
  Password	      string  `form:"password"`
  Password_confirm    string  `form:"password_confirm"`
  validator.Validator	      `form:"-"`
}


func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, http.StatusOK, "login.tmpl.html", "base_guest", templateData{})
}

func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
  data := app.newTemplateData(r)
  data.Form = userRegisterForm{}
  app.render(w, r, http.StatusOK, "register.tmpl.html", "base_guest", data)
}


func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, http.StatusOK, "home.tmpl.html", "base_guest", templateData{})
}

func (app *application) getDashboard(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Dashboard."))
}

func (app *application) getBooks(w http.ResponseWriter, r *http.Request) {
  books, err := app.books.All()
  if err != nil {
    app.serverError(w, r, err)
  }

  data := app.newTemplateData(r)
  data.Books = books

  app.render(w, r, http.StatusOK, "books_index.tmpl.html", "base_auth", data)
}

func (app *application) getBooksCreate(w http.ResponseWriter, r *http.Request) {
  data := app.newTemplateData(r)
  data.Form = bookCreateForm{}
  app.render(w, r, http.StatusOK, "books_create.tmpl.html", "base_auth", data)
}

func (app *application) postBooksCreate(w http.ResponseWriter, r *http.Request) {
  var form bookCreateForm

  err := app.decodePostForm(r, &form)
  if err != nil {
    app.clientError(w, http.StatusBadRequest)
    return
  }

  started, err := app.transformDateStringToSqlNullTime(form.Started)
  if err != nil {
    form.AddFieldError("started", "This field must be a valid date.")
    return
  }

  finished, err := app.transformDateStringToSqlNullTime(form.Finished)
  if err != nil {
    form.AddFieldError("finished", "This field must be a valid date.")
    return
  }


  form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank.")
  form.CheckField(validator.MaxChars(form.Title, 120), "title", "This field cannot be more than 120 characters long.")
  form.CheckField(validator.NotBlank(form.Author), "author", "This field cannot be blank.")
  form.CheckField(validator.MaxChars(form.Author, 120), "author", "This field cannot be more than 120 characters long.")

  if !form.Valid() {
    data := app.newTemplateData(r)
    data.Form = form
    app.render(w, r, http.StatusUnprocessableEntity, "books_create.tmpl.html", "bookCreateForm", data)
    return
  }

  id, err := app.books.Insert(form.Title, form.Author, started, finished)
  if err != nil {
    app.serverError(w, r, err)
    return
  }

  app.sessionManager.Put(r.Context(), "flash", "Book added successfully!")

  w.Header().Set("HX-Redirect", fmt.Sprintf("/books/%d", id));
  w.WriteHeader(http.StatusSeeOther)
}

func (app *application) getBookView(w http.ResponseWriter, r *http.Request) {

  id, err := strconv.Atoi(r.PathValue("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }

  book, err := app.books.Get(id)
  if err != nil {
    if errors.Is(err, models.ErrNoRecord) {
      http.NotFound(w, r)
    } else {
      app.serverError(w, r,   err)
    }
    return
  }

  data := app.newTemplateData(r)
  data.Book = book

  app.render(w, r, http.StatusOK, "books_view.tmpl.html", "base_auth", data)
}

func (app *application) postLogin(w http.ResponseWriter, r *http.Request) {

}

func (app *application) postLogout(w http.ResponseWriter, r *http.Request) {

}

func (app *application) postRegister(w http.ResponseWriter, r *http.Request) {
  var form userRegisterForm

  err := app.decodePostForm(r, &form)
  if err != nil {
    app.clientError(w, http.StatusBadRequest)
    return
  }

  form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank.")
  form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank.")
  form.CheckField(validator.MatchesRegExp(form.Email, validator.EmailRX), "email", "This field must be a valid email address.")
  form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank.")
  form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least eight characters long.")
  form.CheckField(validator.NotBlank(form.Password_confirm), "password_confirm", "This field cannot be blank.")
  form.CheckField(validator.MinChars(form.Password_confirm, 8), "password_confirm", "This field must be at least eight characters long.")
  form.CheckField(validator.Matches(form.Password, form.Password_confirm), "password_confirm", "The password fields must be identical.")

  if !form.Valid() {
    data := app.newTemplateData(r)
    data.Form = form
    app.render(w, r, http.StatusUnprocessableEntity, "register.tmpl.html", "userRegisterForm", data)
    return
  }

  err = app.users.Insert(form.Name, form.Email, form.Password)
  if err != nil {
    if errors.Is(err, models.ErrDuplicateEmail) {
      form.AddFieldError("email", "Email address is already in use.")
      data := app.newTemplateData(r)
      data.Form = form
      app.render(w, r, http.StatusUnprocessableEntity, "register.tmpl.html", "userRegisterForm", data)
      return
    } else {
      app.serverError(w, r, err)
    }
  }

  app.sessionManager.Put(r.Context(), "flash", "Your registration was successful. Please log in.")

  w.Header().Set("HX-Redirect", "/sessions/create")
  w.WriteHeader(http.StatusSeeOther)
}
