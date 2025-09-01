package logger

import (
	"log/slog"
	"os"
)

var log *slog.Logger

func init() {
	// isDev := environment.IsDevEnvironment()

	// if isDev {
	Init(slog.LevelDebug, "text")
	// } else {
	// 	Init(slog.LevelError, "json")
	// }
}

// Init sets up the default logger for the app.
func Init(level slog.Level, format string) {
	opts := &slog.HandlerOptions{Level: level}

	var h slog.Handler
	if format == "text" {
		h = slog.NewTextHandler(os.Stdout, opts)
	} else {
		h = slog.NewJSONHandler(os.Stdout, opts)
	}

	log = slog.New(h)
	slog.SetDefault(log) // optional: replace global slog
}

// Exposed helpers
func Info(msg string, args ...any)  { log.Info(msg, args...) }
func Warn(msg string, args ...any)  { log.Warn(msg, args...) }
func Error(msg string, args ...any) { log.Error(msg, args...) }
func Debug(msg string, args ...any) { log.Debug(msg, args...) }
