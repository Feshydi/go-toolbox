package handler

import (
	"context"
	"log/slog"
)

type DiscardHandlerMiddlware struct{}

func NewDiscardHandler() *DiscardHandlerMiddlware {
	return &DiscardHandlerMiddlware{}
}

func (h *DiscardHandlerMiddlware) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

func (h *DiscardHandlerMiddlware) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *DiscardHandlerMiddlware) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *DiscardHandlerMiddlware) WithGroup(_ string) slog.Handler {
	return h
}
