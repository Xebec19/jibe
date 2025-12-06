package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Xebec19/jibe/api/internal/server"
	"github.com/Xebec19/jibe/api/pkg/config"
)

func main() {

	envPath := flag.String("env", "../../.env", "path to env file")
	flag.Parse()

	cfg, err := config.NewConfig(*envPath)
	if err != nil {
		slog.Error("Config is invalid", "error", err)
		os.Exit(1)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	srv, err := server.NewServer(ctx, cfg)
	if err != nil {
		cancelFunc()
		os.Exit(1)
	}

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- srv.Run()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		cancelFunc()
		slog.Error("Server threw an error", "error", err)

	case <-ctx.Done():
		slog.Info("Context cancelled")

	case <-shutdown:
		slog.Error("Program cancelled")

		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("Server graceful shutdown failed!", "error", err)
			srv.Srv.Close() // Force shutdown
		}

		slog.Info("Server Stopped")
	}
}
