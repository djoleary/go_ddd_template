package slog

import (
	"io"
	"log/slog"
)

func NewJsonLogger(w io.Writer, lvl string) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: getSlogLevel(lvl),
	}
	jsonHandler := slog.NewJSONHandler(w, opts)
	return slog.New(jsonHandler)
}

func getSlogLevel(l string) slog.Level {
	switch l {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	}

	return slog.LevelInfo
}
