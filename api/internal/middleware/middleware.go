package middleware

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/Xebec19/jibe/api/internal/config"
	"github.com/Xebec19/jibe/api/pkg/logger"
	"github.com/google/uuid"
	"github.com/rs/cors"
)

// contextKey is a type for context keys
type contextKey string

const requestIDKey contextKey = "requestID"

// RequestID middleware adds a unique request ID to each request
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID retrieves the request ID from context
func GetRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		return reqID
	}
	return ""
}

// responseWriter wraps http.ResponseWriter to capture status code and bytes written
type responseWriter struct {
	http.ResponseWriter
	status      int
	bytes       int
	wroteHeader bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.wroteHeader {
		rw.status = code
		rw.ResponseWriter.WriteHeader(code)
		rw.wroteHeader = true
	}
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusOK)
	}
	n, err := rw.ResponseWriter.Write(b)
	rw.bytes += n
	return n, err
}

// Logger middleware logs HTTP requests
func Logger(log *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap response writer to capture status code
			wrapped := &responseWriter{ResponseWriter: w, status: http.StatusOK}

			defer func() {
				log.Info().
					Str("method", r.Method).
					Str("path", r.URL.Path).
					Str("remote_addr", r.RemoteAddr).
					Int("status", wrapped.status).
					Int("bytes", wrapped.bytes).
					Dur("duration_ms", time.Since(start)).
					Str("request_id", GetRequestID(r.Context())).
					Msg("HTTP request")
			}()

			next.ServeHTTP(wrapped, r)
		})
	}
}

// Recoverer middleware recovers from panics and logs them
func Recoverer(log *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rvr := recover(); rvr != nil {
					log.Error().
						Str("request_id", GetRequestID(r.Context())).
						Interface("panic", rvr).
						Bytes("stack", debug.Stack()).
						Msg("Panic recovered")

					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

// CORS middleware configures CORS headers
func CORS(cfg *config.Config) func(next http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowedMethods:   cfg.AllowedMethods,
		AllowedHeaders:   cfg.AllowedHeaders,
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	return func(next http.Handler) http.Handler {
		return c.Handler(next)
	}
}

// ContentType middleware sets the Content-Type header
func ContentType(contentType string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentType)
			next.ServeHTTP(w, r)
		})
	}
}

// RateLimit is a placeholder for rate limiting middleware
// You can implement this using github.com/go-chi/httprate or similar
func RateLimit() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Implement rate limiting
			next.ServeHTTP(w, r)
		})
	}
}

// Auth is a placeholder for authentication middleware
func Auth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Implement authentication
			// Example: Check JWT token, validate, add user to context

			// For now, just pass through
			next.ServeHTTP(w, r)
		})
	}
}

// SecurityHeaders adds security-related headers
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		next.ServeHTTP(w, r)
	})
}

// HealthCheckBypass allows health check endpoints to bypass certain middleware
func HealthCheckBypass(paths ...string) func(next http.Handler) http.Handler {
	pathMap := make(map[string]bool)
	for _, path := range paths {
		pathMap[path] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if pathMap[r.URL.Path] {
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, "OK")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
