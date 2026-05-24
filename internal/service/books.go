package service

import (
	"context"
	"library-app/internal/models"
)

type BookStorage interface {
	GetBooks(ctx context.Context) (books []models.Book, err error)
}

type BookService struct {
	repo BookStorage
}

func NewBookService(repo BookStorage) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (b *BookService) GetBooks(ctx context.Context) (books []models.Book, err error) {
	if ctx.Err() != nil {
		return []models.Book{}, ctx.Err()
	}

	books, err = b.repo.GetBooks(ctx)
	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}
