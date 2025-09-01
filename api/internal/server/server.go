package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Xebec19/jibe/api/internal/lib/environment"
	"github.com/Xebec19/jibe/api/internal/pkg/logger"
)

type Server interface {
	StartServer() error
	ShutdownServer()
}

type API struct {
	ctx    context.Context
	cancel context.CancelFunc
	srv    *http.Server
}

func (s *API) StartServer() error {
	logger.Info("Starting Server")

	// Start server in a goroutine
	return s.srv.ListenAndServe()
}

func (s *API) ShutdownServer() {
	logger.Info("Stopping Server")

	// Create a timeout context for shutdown
	ctx, cancel := context.WithTimeout(s.ctx, 30*time.Second)
	defer cancel()

	err := s.srv.Shutdown(ctx)
	if err != nil {
		logger.Error("Graceful shutdown failed, forcing shutdown", err)
		os.Exit(1)
	}

	s.cancel() // Cancel the main context
	logger.Info("Server stopped successfully")
}

func NewServer() Server {
	config := environment.GetConfig()
	r := createRoutes()

	// Create base context
	ctx, cancel := context.WithCancel(context.Background())

	srv := &http.Server{
		Handler:      r,
		Addr:         config.Port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
		// Good practice for keeping our server secure
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	return &API{
		ctx:    ctx,
		cancel: cancel,
		srv:    srv,
	}
}
