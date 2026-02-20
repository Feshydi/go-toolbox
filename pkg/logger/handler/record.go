package handler

import (
	"context"
	"log/slog"
	"sync"
)

type RecordHandlerMiddlware struct {
	mu      sync.Mutex
	records []slog.Record
}

func NewRecordHandler() *RecordHandlerMiddlware {
	return &RecordHandlerMiddlware{}
}

func (h *RecordHandlerMiddlware) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (h *RecordHandlerMiddlware) Handle(_ context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.records = append(h.records, r)
	return nil
}

func (h *RecordHandlerMiddlware) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *RecordHandlerMiddlware) WithGroup(_ string) slog.Handler {
	return h
}

func (h *RecordHandlerMiddlware) Records() []slog.Record {
	h.mu.Lock()
	defer h.mu.Unlock()
	recordsCopy := make([]slog.Record, len(h.records))
	copy(recordsCopy, h.records)
	return recordsCopy
}
