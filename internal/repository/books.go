package repository

import (
	"context"
	"errors"
	"library-app/internal/models"
	"log/slog"

	"github.com/google/uuid"
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

func (r *BooksRepo) UpdateBook(ctx context.Context, req models.Book) (resp models.Book, err error) {

	if ctx.Err() != nil {
		return models.Book{}, ctx.Err()
	}

	query := `
		UPDATE libraryapp.books 
		SET title = $1, 
		    author = $2, 
		    manufacture = $3, 
		    description = $4 
		WHERE id = $5 
		RETURNING id, title, author, manufacture, description;
	`

	err = r.db.QueryRow(ctx, query,
		req.Title,
		req.Author,
		req.Manufacture,
		req.Description,
		req.ID,
	).Scan(
		&resp.ID,
		&resp.Title,
		&resp.Author,
		&resp.Manufacture,
		&resp.Description,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			slog.Error("There is no book with this id", "err", err, "id", req.ID)

		} else {
			slog.Error("Failed request execution", "err", err, "query", query)
		}
		return models.Book{}, err
	}

	return resp, nil
}

func (r *BooksRepo) CreateBook(ctx context.Context, req models.Book) (resp models.Book, err error) {

	if ctx.Err() != nil {
		return models.Book{}, ctx.Err()
	}

	query := `
				INSERT INTO libraryapp.books 
				(id, title, author, manufacture, description)
				VALUES ($1, $2, $3, $4, $5) 
				RETURNING id, title, author, manufacture, description
			`

	err = r.db.QueryRow(ctx, query,
		uuid.New(),
		req.Title,
		req.Author,
		req.Manufacture,
		req.Description,
	).Scan(
		&resp.ID,
		&resp.Title,
		&resp.Author,
		&resp.Manufacture,
		&resp.Description,
	)

	if err != nil {
		slog.Error("Failed request execution", "err", err, "query", query)
		return models.Book{}, err
	}

	return resp, nil
}

func (r *BooksRepo) GetBooks(ctx context.Context) (resp []models.Book, err error) {

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	query := `SELECT id, title, author, manufacture, description FROM libraryapp.books;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		slog.Error("Failed request execution", "err", err, "query", query)
		return nil, err
	}

	resp, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Book])
	if err != nil {
		slog.Error("Failed to collect rows into struct slice", "err", err, "query", query)
		return nil, err
	}

	return resp, nil

}
