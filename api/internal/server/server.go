package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/Xebec19/jibe/api/internal/container"
	"github.com/Xebec19/jibe/api/internal/container/config"
	"github.com/Xebec19/jibe/api/internal/routes"
	"github.com/Xebec19/jibe/api/pkg/logger"
	"github.com/gorilla/mux"
)

type Server struct {
	Srv *http.Server
}

func NewServer() (*Server, error) {

	logger := logger.NewLogger(slog.LevelInfo)

	cfg, err := config.NewConfig("../../.env")
	if err != nil {
		logger.Error("configuration setup failed")
		return nil, err
	}

	c := container.NewContainer(cfg, logger)

	// todo init services and repositories in container

	r := mux.NewRouter()

	routes.RegisterRoutes(r, c)

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  time.Duration(60 * time.Second),
		WriteTimeout: time.Duration(60 * time.Second),
		IdleTimeout:  time.Duration(60 * time.Second),
	}

	return &Server{
		Srv: srv,
	}, nil
}

func (s *Server) Run() error {

	return s.Srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	// todo cleanup
	s.Srv.Close()
	return nil
}
