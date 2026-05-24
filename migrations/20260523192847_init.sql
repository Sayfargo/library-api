-- +goose Up
CREATE SCHEMA IF NOT EXISTS libraryapp;

CREATE TABLE IF NOT EXISTS libraryapp.books(
    id UUID PRIMARY KEY,
    title VARCHAR(150) DEFAULT 'No title' NOT NULL,
    author VARCHAR(100) DEFAULT 'No author' NOT NULL,
    manufacture SMALLINT CONSTRAINT year_valid CHECK(manufacture > 0) NOT NULL,
    description TEXT DEFAULT 'No description' NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS libraryapp.books;

DROP SCHEMA IF EXISTS libraryapp;
