package repository

import (
	"context"
	"library-app/internal/models"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BooksRepo struct {
	db *pgxpool.Pool
}

func NewBooksRepo(db *pgxpool.Pool) *BooksRepo {
	return &BooksRepo{
		db: db,
	}
}

func (r *BooksRepo) GetBooks(ctx context.Context) (books []models.Book, err error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	query := `SELECT id, title, author, manufacture, description FROM libraryapp.books;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		slog.Error("Failed request with query", "err", err, "query", query)
		return nil, err
	}

	books, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Book])
	if err != nil {
		slog.Error("Failed to collect rows into struct slice", "err", err, "query", query)
		return nil, err
	}

	return books, nil

}
