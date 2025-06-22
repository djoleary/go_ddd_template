// Package slog provides a facade for working with loggers
package slog

import (
	"io"
	"log/slog"
)

type Logger struct {
	slog *slog.Logger
}

func NewJsonLogger(w io.Writer, lvl string) *Logger {
	opts := &slog.HandlerOptions{
		Level: getLevel(lvl),
	}
	jsonHandler := slog.NewJSONHandler(w, opts)
	return &Logger{
		slog: slog.New(jsonHandler),
	}
}

// getLevel converts an lowercase english string info a log level
func getLevel(l string) slog.Level {
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

func (l *Logger) Debug(msg string, args ...any) {
	l.slog.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.slog.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.slog.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.slog.Error(msg, args...)
}
