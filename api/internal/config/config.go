package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Environment     string
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration

	// CORS settings
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string

	// Database (example - adjust based on your needs)
	DatabaseURL string

	// JWT/Auth (example)
	JWTSecret string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists (ignore error if file doesn't exist)
	godotenv.Load()

	cfg := &Config{
		Environment:     getEnv("ENVIRONMENT", "development"),
		Port:            getEnv("PORT", "8080"),
		ReadTimeout:     getDurationEnv("READ_TIMEOUT", 15*time.Second),
		WriteTimeout:    getDurationEnv("WRITE_TIMEOUT", 15*time.Second),
		IdleTimeout:     getDurationEnv("IDLE_TIMEOUT", 60*time.Second),
		ShutdownTimeout: getDurationEnv("SHUTDOWN_TIMEOUT", 20*time.Second),

		AllowedOrigins: getSliceEnv("ALLOWED_ORIGINS", []string{"*"}),
		AllowedMethods: getSliceEnv("ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		AllowedHeaders: getSliceEnv("ALLOWED_HEADERS", []string{"Accept", "Authorization", "Content-Type"}),

		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret:   getEnv("JWT_SECRET", ""),
	}

	// Validate required fields for production
	if cfg.Environment == "production" {
		if err := cfg.Validate(); err != nil {
			return nil, fmt.Errorf("configuration validation failed: %w", err)
		}
	}

	return cfg, nil
}

// Validate checks if required configuration values are set
func (c *Config) Validate() error {
	if c.Environment == "production" {
		if c.JWTSecret == "" {
			return fmt.Errorf("JWT_SECRET is required in production")
		}
		if c.DatabaseURL == "" {
			return fmt.Errorf("DATABASE_URL is required in production")
		}
	}
	return nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getDurationEnv retrieves a duration environment variable or returns a default value
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	// Try parsing as seconds (integer)
	if seconds, err := strconv.Atoi(valueStr); err == nil {
		return time.Duration(seconds) * time.Second
	}

	// Try parsing as duration string (e.g., "15s", "1m")
	if duration, err := time.ParseDuration(valueStr); err == nil {
		return duration
	}

	return defaultValue
}

// getSliceEnv retrieves a comma-separated environment variable as a slice
func getSliceEnv(key string, defaultValue []string) []string {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	// Simple split by comma - enhance if needed
	result := []string{}
	for i, s := 0, 0; i <= len(valueStr); i++ {
		if i == len(valueStr) || valueStr[i] == ',' {
			if i > s {
				result = append(result, valueStr[s:i])
			}
			s = i + 1
		}
	}

	if len(result) == 0 {
		return defaultValue
	}
	return result
}
