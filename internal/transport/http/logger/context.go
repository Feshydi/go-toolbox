package logger

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func WithOpRequest(logger *slog.Logger, op string, r *http.Request) *slog.Logger {
	return logger.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
}
