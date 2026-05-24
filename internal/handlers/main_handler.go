package handlers

import (
	"library-app/internal/service"
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

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(time.Minute))

	r.Route("/api", func(r chi.Router) {
		r.Get("/books", h.GetBooks)
		r.Post("/books", h.CreateBook) // {title, author, manufacture, description}
		r.Patch("/books/{id}", h.UpdateBook)
	})

	return r
}
