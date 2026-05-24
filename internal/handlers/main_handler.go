package handlers

import (
	"library-app/internal/service"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // для разработки
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(time.Minute))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./internal/web/index.html")
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/books", h.GetBooks)
		r.Post("/books", h.CreateBook) // {title, author, manufacture, description}
		r.Put("/books/{id}", h.UpdateBook)
		r.Patch("/books/{id}", h.SoftDelete)
	})

	return r
}
