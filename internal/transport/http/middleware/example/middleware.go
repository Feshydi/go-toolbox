package example

import (
	http_logger "go-toolbox/internal/transport/http/logger"
	"log/slog"
	"net/http"
)

func Middleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const op = "http.middleware.example"

			logCtx := http_logger.WithOpRequest(logger, op, r)

			logCtx.Info("Middleware worked!")
			next.ServeHTTP(w, r)
		})
	}
}
