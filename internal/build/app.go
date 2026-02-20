package build

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
)

type Application struct {
	httpServer *http.Server
	logger     *slog.Logger
}

func (a *Application) Start() {
	a.logger.Info("Staring application...")

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic("server error: " + err.Error())
			}
		}
	}()
}

func (a *Application) Stop(ctx context.Context) {
	a.logger.Info("Stopping application...")

	if err := a.httpServer.Shutdown(ctx); err != nil {
		a.logger.Error("Failed to shutdown HTTP server", slog.Any("err", err))
	}
}
