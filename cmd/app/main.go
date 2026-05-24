package main

import (
	"context"
	"library-app/internal/app"
	"library-app/internal/config"
	"log/slog"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config", "err", err)
		return
	}

	app, err := app.NewApp(cfg)
	if err != nil {
		return
	}

	sigCtx, sigCancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer sigCancel()

	go app.Run()

	<-sigCtx.Done()
	sigCancel()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	if err := app.Shutdown(shutdownCtx); err != nil {
		return
	}

}
