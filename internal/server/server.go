package server

import (
	"context"
	"github.com/Andrew-UA/product-list/internal/config"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(config *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           config.AppHost + ":" + config.AppPort,
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1,
		},
	}
}

func (s *Server) Run() error {
	log.Info().Msgf("Starting server on %s", s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	log.Info().Msgf("Stopping server on %s", s.httpServer.Addr)

	return s.httpServer.Shutdown(ctx)
}
