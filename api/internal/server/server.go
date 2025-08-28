package server

import (
	"net/http"

	"github.com/Xebec19/jibe/api/internal/lib/environment"
)

type Server interface {
	StartServer() void
	ShutdownServer() void
}

func NewServer() *Server {

	config := environment.GetConfig()

	srv := &http.Server{
		Handler: r,
		Addr:    config.Port,
	}
}
