package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Xebec19/jibe/api/internal/config"
	"github.com/Xebec19/jibe/api/internal/container"
	"github.com/Xebec19/jibe/api/internal/handlers"
	"github.com/Xebec19/jibe/api/internal/middleware"
	"github.com/Xebec19/jibe/api/pkg/logger"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/Xebec19/jibe/api/docs"
)

// @title           Jibe API
// @version         1.0
// @description     This is a Jibe API server with user management capabilities.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @schemes http https

func main() {
	// Initialize configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	log := logger.New(cfg.Environment)

	log.Info().
		Str("environment", cfg.Environment).
		Str("port", cfg.Port).
		Msg("Starting server")

	// Initialize dependency injection container
	c := container.New(cfg, log)
	defer c.Shutdown()

	// Initialize router with container
	router := setupRouter(c)

	// Create HTTP server
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Start server in a goroutine
	serverErrors := make(chan error, 1)
	go func() {
		log.Info().
			Str("addr", srv.Addr).
			Msg("Server listening")
		serverErrors <- srv.ListenAndServe()
	}()

	// Setup graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Block until we receive a signal or server error
	select {
	case err := <-serverErrors:
		log.Fatal().Err(err).Msg("Server error")

	case sig := <-shutdown:
		log.Info().
			Str("signal", sig.String()).
			Msg("Starting graceful shutdown")

		// Give outstanding requests a deadline for completion
		ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		// Attempt graceful shutdown
		if err := srv.Shutdown(ctx); err != nil {
			log.Error().Err(err).Msg("Graceful shutdown failed")
			// Force close
			srv.Close()
		}

		log.Info().Msg("Server stopped")
	}
}

func setupRouter(c *container.Container) *mux.Router {
	r := mux.NewRouter()

	// Middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger(c.Logger))
	r.Use(middleware.Recoverer(c.Logger))
	r.Use(middleware.CORS(c.Config))

	// Health check endpoint
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/", handlers.Root).Methods("GET")

	// Initialize handlers with dependency injection
	userHandler := handlers.NewUserHandler(c.UserService, c.Logger)

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// User routes
	api.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
