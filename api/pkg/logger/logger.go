package logger

import "log/slog"

type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

func NewLogger(level slog.Level) Logger {

	logger := slog.Default()

	return &log{
		log: logger,
	}
}

type log struct {
	log *slog.Logger
}

func (l log) Info(args ...interface{}) {
	l.log.Info("DEBUG", args...)
}

func (l log) Warn(args ...interface{}) {
	l.log.Warn("WARN", args...)
}

func (l log) Error(args ...interface{}) {
	l.log.Error("ERROR", args...)
}
