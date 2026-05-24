package handlers

import (
	"context"
	"errors"
	"library-app/internal/jsonutil"
	"library-app/internal/service"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	bookService *service.BookService
}

func NewHandler(bookService *service.BookService) *Handler {
	return &Handler{
		bookService: bookService,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(time.Minute))

	r.Route("/api", func(r chi.Router) {
		r.Get("/books", h.GetBooks)
	})

	return r
}

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {

	books, err := h.bookService.GetBooks(r.Context())
	if err != nil {
		if errors.Is(err, context.Canceled) {
			slog.Info("Connection refused by user")
		} else {
			jsonutil.EncodeError(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	jsonutil.EncodeSuccess(w, books, http.StatusOK)
}
