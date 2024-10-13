package models

import (
  "database/sql"
  "time"
)

type Book struct {
  ID        int
  Title     string
  Status    string
  Author    string
  Started   time.Time
  Finished  time.Time
}

type BookModel struct {
  DB *sql.DB
}

func (m *BookModel) Insert(title, status, author string, started, finished sql.NullTime) (int, error) {

  stmt := "INSERT INTO books(title, status, author, started, finished) VALUES (?, ?, ?, ?, ?)"

  result, err := m.DB.Exec(stmt, title, status, author, started, finished)
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
  return Book{}, nil
}

