package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/Xebec19/jibe/api/internal/db"
	"github.com/Xebec19/jibe/api/internal/layers/container"
	"github.com/Xebec19/jibe/api/internal/routes"
	"github.com/Xebec19/jibe/api/pkg/config"
	"github.com/Xebec19/jibe/api/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	Container container.Container
	Srv       *http.Server
}

func NewServer(ctx context.Context, cfg *config.Config) (*Server, error) {

	logger := logger.NewLogger(slog.LevelInfo)

	pool, err := pgxpool.New(context.TODO(), cfg.DbConn)
	if err != nil {
		logger.Error("DB Pool creation failed!", "error", err)
	}

	q := db.New(pool)

	c := container.NewContainer(ctx, cfg, logger, pool, q)

	c.SetupRepositories()
	c.SetupServices()

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
		Container: c,
		Srv:       srv,
	}, nil
}

func (s *Server) Run() error {

	s.Container.Logger.Info("Server started", "PORT", s.Container.Cfg.Port)

	return s.Srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {

	s.Container.Dbpool.Close()

	s.Srv.Close()

	return nil
}
