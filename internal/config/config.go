package config

import (
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
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
