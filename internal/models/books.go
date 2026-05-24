package models

import (
	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Author      string    `json:"author" db:"author"`
	Manufacture uint16    `json:"manufacture" db:"manufacture"`
	Description string    `json:"description" db:"description"`
}
