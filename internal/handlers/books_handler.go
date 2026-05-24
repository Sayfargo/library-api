package handlers

import (
	"context"
	"errors"
	"library-app/internal/jsonutil"
	"library-app/internal/models"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		jsonutil.EncodeError(w, "Missing book id", http.StatusBadRequest)
		return
	}

	var req models.Book

	if err := jsonutil.DecodeJSON(w, r, &req); err != nil {
		return
	}

	req.ID = uuid

	resp, err := h.bookService.UpdateBook(r.Context(), req)
	makeAndSend(w, resp, err, http.StatusOK)

}

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {

	resp, err := h.bookService.GetBooks(r.Context())
	makeAndSend(w, resp, err, http.StatusOK)
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {

	var req models.Book

	if err := jsonutil.DecodeJSON(w, r, &req); err != nil {
		return
	}

	resp, err := h.bookService.CreateBook(r.Context(), req)
	makeAndSend(w, resp, err, http.StatusCreated)
}

func makeAndSend(w http.ResponseWriter, resp any, err error, status int) {
	if err != nil {
		if errors.Is(err, context.Canceled) {
			slog.Info("Connection refused by user")
		} else if errors.Is(err, models.ErrEmptyFields) {
			jsonutil.EncodeError(w, err.Error(), http.StatusBadRequest)
		} else if errors.Is(err, models.ErrNotFound) {
			jsonutil.EncodeError(w, err.Error(), http.StatusNotFound)
		} else {
			jsonutil.EncodeError(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	jsonutil.EncodeSuccess(w, resp, status)
}
