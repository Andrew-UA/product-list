package app

import (
	"context"
	"github.com/Andrew-UA/product-list/app/repositories/factory"
	"github.com/Andrew-UA/product-list/app/services"
	"github.com/Andrew-UA/product-list/database/seeds"
	"github.com/Andrew-UA/product-list/internal/config"
	"github.com/Andrew-UA/product-list/internal/db"
	"github.com/Andrew-UA/product-list/internal/server"
	"github.com/Andrew-UA/product-list/internal/transport/http"
	"github.com/Andrew-UA/product-list/internal/transport/http/handlers"
	"github.com/Andrew-UA/product-list/internal/transport/http/middleware"
	"github.com/Andrew-UA/product-list/internal/validation"
	"github.com/Andrew-UA/product-list/pkg/auth"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	var err error
	var cfg *config.Config
	cfg, err = config.InitConfig()

	if err != nil {
		log.Err(err).Msg("failed to read config")
	}
	config.InitLogger(cfg)

	// Init BD connection
	var dbConnector db.DatabaseConnector
	if dbConnector, err = db.GetDataBaseConnector(cfg); err != nil {
		log.Fatal().Err(err).Msg("failed get db connector")
	}

	// Connect to db
	err = dbConnector.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	//Run migrations
	err = db.RunMigrations(dbConnector)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run migrations")
	}

	repository, err := factory.CreateRepository(dbConnector)
	if err != nil {
		log.Fatal().Err(err).Msg("failed get repository factory")
	}

	// Innit Services
	validator := validation.NewValidator()
	passwordManager := auth.NewBcryptPasswordManager()
	tokenManager := auth.NewJWTTokenManager(cfg.AppKey)
	userService := services.NewUserService(repository.UserRepository())
	authService := services.NewAuthService(passwordManager, tokenManager, repository.AuthRepository())

	//Seeder
	seeder := seeds.NewSeeder(cfg, passwordManager, repository.UserRepository())
	err = seeder.Seed()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to seed")
	}

	// Innit Middleware
	authMiddleware := middleware.NewAuthMiddleware(userService, authService, tokenManager)

	// Innit Handlers
	healthHandler := handlers.NewHealthHandler(cfg)
	authHandler := handlers.NewAuthHandler(validator, userService, authService)
	userHandler := handlers.NewUserHandler(validator, userService)

	router := http.NewRouter(cfg, authMiddleware.HandleFunc, healthHandler, authHandler, userHandler)
	srv := server.NewServer(cfg, router.Mux)

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
