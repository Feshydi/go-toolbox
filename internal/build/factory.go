package build

import (
	"go-toolbox/internal/config"
	"go-toolbox/pkg/logger"
)

func Create(cfg *config.Config) *Application {
	logger := logger.New(cfg.Logger.Level)

	return &Application{
		logger: logger,
	}
}
