-- +goose Up
CREATE SCHEMA IF NOT EXISTS libraryapp;

CREATE TABLE IF NOT EXISTS books(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(150) DEFAULT 'No title' NOT NULL,
    description TEXT DEFAULT 'No description' NOT NULL,
    author VARCHAR(100) DEFAULT 'No author' NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS books;

DROP SCHEMA IF EXISTS libraryapp;
