package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port         string        `mapstructure:"port"`
		Host         string        `mapstructure:"host"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
		IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
	} `mapstructure:"server"`

	DB struct {
		Conn string `mapstructure:"conn_string"`
		Pool struct {
			MaxConns          int32         `mapstructure:"max_conns"`
			MinConns          int32         `mapstructure:"min_conns"`
			HealthCheckPeriod time.Duration `mapstructure:"health_check_period"`
		} `mapstructure:"pool"`
	} `mapstructure:"db"`
}

func LoadConfig() (cfg *Config, err error) {
	slog.Info("Start loading config with viper")

	DB_PATH_CONFIG := os.Getenv("DB_PATH_CONFIG")
	SERVER_PATH_CONFIG := os.Getenv("SERVER_PATH_CONFIG")

	v := viper.New()

	v.SetConfigFile(SERVER_PATH_CONFIG)
	if err := v.ReadInConfig(); err != nil {
		slog.Error("Failed to read server.yaml with path", "err", err, "path", SERVER_PATH_CONFIG)
		return nil, err
	}

	v.SetConfigFile(DB_PATH_CONFIG)
	if err := v.MergeInConfig(); err != nil {
		slog.Error("Failed to merge in config with viper", "err", err, "path", DB_PATH_CONFIG)
		return nil, err
	}

	if err := v.BindEnv("db.conn_string", "DB_CONN_STRING"); err != nil {
		slog.Error("Failed to bind env DB_CONN_STRING", "err", err)
		return nil, err
	}

	v.AutomaticEnv()

	if err := v.Unmarshal(&cfg); err != nil {
		slog.Error("Can't unmarshal to strcut with viper", "err", err)
		return nil, err
	}

	slog.Info("Config loaded successfully", "port", cfg.Server.Port, "db", cfg.DB.Conn)
	return cfg, nil
}
