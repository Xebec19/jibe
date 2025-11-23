package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Xebec19/jibe/api/internal/server"
)

func main() {

	srv, err := server.NewServer()
	if err != nil {
		panic(err)
	}

	serverErrors := make(chan error, 1)

	go func() {
		slog.Info("Starting Server")
		serverErrors <- srv.Run()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		slog.Error("Server threw an error", "error", err)

	case <-shutdown:
		slog.Error("Program cancelled")

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("Server graceful shutdown failed!", "error", err)
			srv.Srv.Close() // Force shutdown
		}

		slog.Info("Server Stopped")
	}
}
