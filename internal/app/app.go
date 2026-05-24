package app

import (
	"context"
	"errors"
	"library-app/internal/config"
	"library-app/internal/db"
	"library-app/internal/handlers"
	"library-app/internal/repository"
	"library-app/internal/service"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	httpServer *http.Server
	db         *pgxpool.Pool
}

func NewApp(cfg *config.Config) (app *App, err error) {

	db, err := db.InitDB(cfg)
	if err != nil {
		return nil, err
	}

	bookRepo := repository.NewBooksRepo(db)
	bookService := service.NewBookService(bookRepo)
	handler := handlers.NewHandler(bookService)

	server := &http.Server{
		Addr:         cfg.Server.Port,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
		Handler:      handler.InitRoutes(),
	}

	return &App{
		httpServer: server,
		db:         db,
	}, nil

}

func (a *App) Run() {
	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Failed to start server", "err", err)
	}
}

func (a *App) Shutdown(ctx context.Context) error {
	slog.Info("Starting shutdown server")

	if err := a.httpServer.Shutdown(ctx); err != nil {
		slog.Error("Failed shutdown server", "err", err)
		return err
	}

	a.db.Close()

	slog.Info("Server and database are successfull shutdown")

	return nil
}
