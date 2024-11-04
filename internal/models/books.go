package models

import (
	"database/sql"
	"errors"
	"time"
)

type BookModelInterface interface {
	Insert(title, author string, startet, finished sql.NullTime) (int, error)
	Get(id int) (Book, error)
	All() ([]Book, error)
}

type Book struct {
	ID       int
	Title    string
	Author   string
	Started  *time.Time
	Finished *time.Time
	Status   string
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
	stmt := `SELECT id, title, author, started, finished,
  CASE WHEN started IS NULL THEN "Not Started" WHEN finished IS NOT NULL THEN "Finished"
  ELSE "Reading" END AS status FROM books WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	var b Book

	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.Started, &b.Finished, &b.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Book{}, ErrNoRecord
		} else {
			return Book{}, err
		}
	}
	return b, nil
}

func (m *BookModel) All() ([]Book, error) {
	stmt := `SELECT id, title, author, started, finished,
  CASE WHEN started IS NULL THEN "Not Started" WHEN finished IS NOT NULL THEN "Finished"
  ELSE "Reading" END AS status FROM books ORDER BY id DESC`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var b Book

		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Started, &b.Finished, &b.Status)
		if err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
