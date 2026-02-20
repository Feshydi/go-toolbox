package build

import (
	"context"
	"log/slog"
)

type Application struct {
	logger *slog.Logger
}

func (a *Application) Start() {
	a.logger.Info("Staring application...")
}

func (a *Application) Stop(ctx context.Context) {
	a.logger.Info("Stopping application...")
}
