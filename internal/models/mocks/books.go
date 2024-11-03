package mocks

import (
	"database/sql"

	"github.com/jansuthacheeva/bookshelf/internal/models"
)

var mockBook = models.Book{
	ID:       1,
	Title:    "A mock book",
	Author:   "Mocky Mockerson",
	Started:  nil,
	Finished: nil,
	Status:   "Not Started",
}

type BookModel struct{}

func (m *BookModel) Insert(title, author string, started, finished sql.NullTime) (int, error) {
	return 2, nil
}

func (m *BookModel) Get(id int) (models.Book, error) {
	switch id {
	case 1:
		return mockBook, nil
	default:
		return models.Book{}, models.ErrNoRecord
	}
}

func (m *BookModel) All() ([]models.Book, error) {
	return []models.Book{mockBook}, nil
}
