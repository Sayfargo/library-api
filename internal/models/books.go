package models

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmptyFields = errors.New("Title or author are empty")
	ErrNotFound    = errors.New("Book not found")
)

type Book struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Author      string    `json:"author" db:"author"`
	Manufacture uint16    `json:"manufacture" db:"manufacture"`
	Description string    `json:"description" db:"description"`
}

func (b *Book) Validate() error {
	if len(b.Title) == 0 || len(b.Author) == 0 {
		return ErrEmptyFields
	}

	return nil
}
