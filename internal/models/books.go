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

func (m *BookModel) Insert(title, status, author string, started, finished time.Time) (int, error) {
  return 0, nil
}

func (m *BookModel) Get(id int) (Book, error) {
  return Book{}, nil
}

