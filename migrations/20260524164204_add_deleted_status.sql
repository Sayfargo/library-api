-- +goose Up
ALTER TABLE libraryapp.books
ADD COLUMN is_deleted BOOLEAN DEFAULT false NOT NULL;

-- +goose Down
ALTER TABLE libraryapp.books
DROP COLUMN IF EXISTS is_deleted;
