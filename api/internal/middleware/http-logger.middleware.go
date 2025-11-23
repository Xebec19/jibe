package middleware

import (
	"net/http"
	"time"

	"github.com/Xebec19/jibe/api/pkg/logger"
)

// responseWriter wraps http.ResponseWriter to capture status code and bytes written
type responseWriter struct {
	http.ResponseWriter
	status      int
	bytes       int
	wroteHeader bool
}

func HttpLogger(logger logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap response writer to capture status code
			wrapped := &responseWriter{ResponseWriter: w, status: http.StatusOK}

			defer func() {
				logger.Info("HTTP Request: ",
					"method", r.Method,
					"path", r.URL.Path,
					"remote_addr", r.RemoteAddr,
					"status", wrapped.status,
					"bytes", wrapped.bytes,
					"duration_ms", time.Since(start))
			}()

			next.ServeHTTP(wrapped, r)
		})
	}
}
