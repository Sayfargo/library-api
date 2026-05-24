package service

import (
	"context"
	"errors"
	"library-app/internal/models"

	"github.com/jackc/pgx/v5"
)

type BookStorage interface {
	GetBooks(ctx context.Context) (resp []models.Book, err error)
	CreateBook(ctx context.Context, req models.Book) (resp models.Book, err error)
	UpdateBook(ctx context.Context, req models.Book) (resp models.Book, err error)
}

type BookService struct {
	repo BookStorage
}

func NewBookService(repo BookStorage) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (b *BookService) UpdateBook(ctx context.Context, req models.Book) (resp models.Book, err error) {
	if ctx.Err() != nil {
		return models.Book{}, ctx.Err()
	}

	if err = req.Validate(); err != nil {
		return models.Book{}, err
	}

	resp, err = b.repo.UpdateBook(ctx, req)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Book{}, models.ErrNotFound
		}
		return models.Book{}, err
	}
	return resp, nil
}

func (b *BookService) CreateBook(ctx context.Context, req models.Book) (resp models.Book, err error) {
	if ctx.Err() != nil {
		return models.Book{}, ctx.Err()
	}

	if err := req.Validate(); err != nil {
		return models.Book{}, err
	}

	return b.repo.CreateBook(ctx, req)
}

func (b *BookService) GetBooks(ctx context.Context) (resp []models.Book, err error) {
	if ctx.Err() != nil {
		return []models.Book{}, ctx.Err()
	}

	resp, err = b.repo.GetBooks(ctx)
	if err != nil {
		return []models.Book{}, err
	}

	return resp, nil
}
