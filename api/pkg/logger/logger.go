package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger wraps zerolog.Logger with additional functionality
type Logger struct {
	*zerolog.Logger
}

// New creates a new logger instance based on environment
func New(environment string) *Logger {
	var log zerolog.Logger

	// Configure based on environment
	if environment == "production" {
		// JSON logging for production
		log = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		// Pretty console logging for development
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		log = zerolog.New(output).With().Timestamp().Logger()
	}

	// Set global log level
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if environment == "development" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return &Logger{&log}
}

// WithFields creates a child logger with additional fields
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	ctx := l.Logger.With()
	for k, v := range fields {
		ctx = ctx.Interface(k, v)
	}
	newLogger := ctx.Logger()
	return &Logger{&newLogger}
}

// WithRequestID creates a child logger with request ID
func (l *Logger) WithRequestID(requestID string) *Logger {
	newLogger := l.Logger.With().Str("request_id", requestID).Logger()
	return &Logger{&newLogger}
}
