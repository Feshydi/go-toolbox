package config

import (
	"log/slog"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port         int           `env:"SERVER_PORT,required"`
		ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT,required"`
		WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT,required"`
		IdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT,required"`
	}
	Logger struct {
		Level slog.Level `env:"LOGGER_LEVEL,required"`
	}
}

func MustLoad() *Config {
	_ = godotenv.Load() // TODO: remove in production

	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		panic("failed to parse environment variables: " + err.Error())
	}

	return &cfg
}
