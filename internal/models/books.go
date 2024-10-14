package models

import (
	"database/sql"
	"errors"
	"time"
)

type Book struct {
  ID        int
  Title     string
  Author    string
  Started   *time.Time
  Finished  *time.Time
}

type BookModel struct {
  DB *sql.DB
}

func (m *BookModel) Insert(title, author string, started, finished sql.NullTime) (int, error) {

  stmt := "INSERT INTO books(title, author, started, finished) VALUES (?, ?, ?, ?)"

  result, err := m.DB.Exec(stmt, title, author, started, finished)
  if err != nil {
    return 0, err
  }

  id, err := result.LastInsertId()
  if err != nil {
    return 0, err
  }

  return int(id), nil
}

func (m *BookModel) Get(id int) (Book, error) {
  stmt := "SELECT id, title, author, started, finished FROM books WHERE id = ?"

  row := m.DB.QueryRow(stmt, id)

  var b Book

  err := row.Scan(&b.ID, &b.Title, &b.Author, &b.Started, &b.Finished)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return Book{}, ErrNoRecord
    } else {
      return Book{}, err
    }
  }
  return b, nil
}


