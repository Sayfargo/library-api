package db

import (
	"context"
	"library-app/internal/config"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(cfg *config.Config) (pool *pgxpool.Pool, err error) {
	slog.Info("Start initing database with config", "conn", cfg.DB.Conn)

	pgxCfg, err := pgxpool.ParseConfig(cfg.DB.Conn)
	if err != nil {
		slog.Error("Can't parse db conn string", "err", err, "conn_string", cfg.DB.Conn)
		return nil, err
	}

	pgxCfg.MaxConns = cfg.DB.Pool.MaxConns
	pgxCfg.MinConns = cfg.DB.Pool.MinConns
	pgxCfg.HealthCheckPeriod = cfg.DB.Pool.HealthCheckPeriod

	pool, err = pgxpool.NewWithConfig(context.Background(), pgxCfg)
	if err != nil {
		slog.Error("Error with pgxpool", "err", err)
		return nil, err
	}

	slog.Info("Try to ping DB")
	if err := pool.Ping(context.Background()); err != nil {
		slog.Error("Can't ping DB, no answer", "err", err)
		return nil, err
	}

	slog.Info("Database started successfully")
	return pool, nil

}
