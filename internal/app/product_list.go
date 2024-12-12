package app

import (
	"context"
	"github.com/Andrew-UA/product-list/internal/config"
	"github.com/Andrew-UA/product-list/internal/server"
	"github.com/Andrew-UA/product-list/internal/transport/http"
	"github.com/Andrew-UA/product-list/internal/transport/http/handlers"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	conf, err := config.InitConfig()

	if err != nil {
		log.Err(err).Msg("failed to read config")
	}
	config.InitLogger(conf)

	healthHandler := handlers.NewHealthHandler(conf)

	router := http.NewRouter(conf, healthHandler)
	srv := server.NewServer(conf, router.Mux)

	go func() {
		if err := srv.Run(); err != nil {
			log.Err(err).Msg("SERVER ERROR")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithCancel(context.Background())
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to stop HTTP server")
	}
}
