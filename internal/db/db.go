package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type postgresConfig struct {
	DB struct {
		Conn string `mapstructure:"conn_string"`
		Pool struct {
			MaxConns          int32         `mapstructure:"max_conns"`
			MinConns          int32         `mapstructure:"min_conns"`
			HealthCheckPeriod time.Duration `mapstructure:"health_check_period"`
		} `mapstructure:"pool"`
	} `mapstructure:"db"`
}

func loadPgxConfig() (cfg postgresConfig, err error) {
	slog.Info("Starting load database config with viper")
	DB_PATH_CONFIG := os.Getenv("DB_PATH_CONFIG")

	if DB_PATH_CONFIG == "" {
		return postgresConfig{}, fmt.Errorf("Empty path: %s", DB_PATH_CONFIG)
	}

	viper.SetConfigFile(DB_PATH_CONFIG)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("Can't read config file", "err", err, "path", DB_PATH_CONFIG)
		return postgresConfig{}, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		slog.Error("Can't unmarshal yaml file into struct", "err", err)
		return postgresConfig{}, err
	}

	return cfg, nil

}

func InitDB() (pool *pgxpool.Pool, err error) {
	slog.Info("Start to init database")
	cfg, err := loadPgxConfig()
	if err != nil {
		return nil, err
	}

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

	return pool, nil

}
