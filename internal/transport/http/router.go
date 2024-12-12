package http

import (
	"github.com/Andrew-UA/product-list/internal/config"
	"net/http"
)

type Router struct {
	Mux           *http.ServeMux
	HealthHandler HealthHandlerInterface
}

func NewRouter(config *config.Config, healthHandler HealthHandlerInterface) *Router {
	mux := http.NewServeMux()

	router := &Router{
		Mux:           mux,
		HealthHandler: healthHandler,
	}

	router.registerHealthRoutes()

	return router
}

func (r *Router) registerHealthRoutes() {
	r.Mux.HandleFunc("GET /health", r.HealthHandler.Health)
	r.Mux.HandleFunc("GET /ping", r.HealthHandler.Ping)
}
