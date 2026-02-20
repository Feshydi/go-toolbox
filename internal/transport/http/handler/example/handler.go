package example

import (
	"go-toolbox/internal/transport/http/logger"
	"log/slog"
	"net/http"
)

type service interface {
	Handle() []byte
}

type Handler struct {
	service service
	logger  *slog.Logger
}

func New(service service, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) HandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.handler.example.HandlerFunc"

		logCtx := logger.WithOpRequest(h.logger, op, r)

		data := h.service.Handle()

		logCtx.Info("Handler worked!")
		w.Write(data)
	}
}
