package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/Xebec19/jibe/api/internal/lib/environment"
	"github.com/Xebec19/jibe/api/internal/pkg/logger"
	"github.com/Xebec19/jibe/api/internal/server"
)

func main() {

	srv := server.NewServer()

	if environment.IsDevEnvironment() || true {
		config := environment.GetConfig()
		logger.Info("Configuration", slog.Any("config", config))
	}

	go func() {
		err := srv.StartServer()
		if err != nil {
			logger.Error("Server could not start", err)
		}
	}()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	<-ch
	srv.ShutdownServer()
}
