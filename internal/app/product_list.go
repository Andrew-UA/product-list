package app

import (
	"context"
	"github.com/Andrew-UA/product-list/app/repositories"
	"github.com/Andrew-UA/product-list/app/services"
	"github.com/Andrew-UA/product-list/internal/config"
	"github.com/Andrew-UA/product-list/internal/db"
	"github.com/Andrew-UA/product-list/internal/server"
	"github.com/Andrew-UA/product-list/internal/transport/http"
	"github.com/Andrew-UA/product-list/internal/transport/http/handlers"
	"github.com/Andrew-UA/product-list/internal/transport/http/middleware"
	"github.com/Andrew-UA/product-list/pkg/auth"
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

	// Init BD connection
	var dbConnector db.DatabaseConnector
	if dbConnector, err = db.GetDataBaseConnector(conf); err != nil {
		log.Fatal().Err(err).Msg("failed get db connector")
	}

	repositoryFactory, err := repositories.GetRepositoryFactory(dbConnector)
	if err != nil {
		log.Fatal().Err(err).Msg("failed get repository factory")
	}

	// Innit Services
	passwordManager := auth.NewBcryptPasswordManager()
	tokenManager := auth.NewJWTTokenManager(conf.AppKey)
	userService := services.NewUserService(repositoryFactory.UserRepository())
	authService := services.NewAuthService(passwordManager, tokenManager, repositoryFactory.AuthRepository())

	// Innit Middleware
	authMiddleware := middleware.NewAuthMiddleware(userService, authService, tokenManager)

	// Innit Handlers
	healthHandler := handlers.NewHealthHandler(conf)
	authHandler := handlers.NewAuthHandler(userService, authService)
	userHandler := handlers.NewUserHandler(userService)

	router := http.NewRouter(conf, authMiddleware.HandleFunc, healthHandler, authHandler, userHandler)
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

	dbConnector.Close(ctx)

	if err := srv.Stop(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to stop HTTP server")
	}
}
