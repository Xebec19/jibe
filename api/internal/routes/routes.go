package routes

import (
	"github.com/Xebec19/jibe/api/internal/layers/container"
	"github.com/Xebec19/jibe/api/internal/middleware"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, c container.Container) {

	r.Use(middleware.HttpLogger(c.Logger))

	registerHealthRoutes(r, c)
	registerAuthRoutes(r, c)
}
