package main

import (
	"os"
	"os/signal"

	"github.com/Xebec19/jibe/api/internal/pkg/logger"
	"github.com/Xebec19/jibe/api/internal/server"
)

func main() {

	srv := server.NewServer()

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
