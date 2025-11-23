package logger

import "log/slog"

type Logger interface {
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
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

func (l log) Info(msg string, args ...interface{}) {
	l.log.Info(msg, args...)
}

func (l log) Warn(msg string, args ...interface{}) {
	l.log.Warn(msg, args...)
}

func (l log) Error(msg string, args ...interface{}) {
	l.log.Error(msg, args...)
}
