package app

import (
	"context"
	"database/sql"
	"github.com/Andrew-UA/product-list/app/repositories/mysql"
	"github.com/Andrew-UA/product-list/app/services"
	"github.com/Andrew-UA/product-list/internal/config"
	"github.com/Andrew-UA/product-list/internal/db"
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

	// Init BD connection
	var dbConnector db.DatabaseConnector
	if dbConnector, err = db.GetDataBaseConnector(conf); err != nil {
		log.Fatal().Err(err).Msg("failed get db connector")
	}

	conn, err := dbConnector.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	// Init Mysql Repositories
	mysqlDB, ok := conn.(*sql.DB)
	if !ok {
		log.Fatal().Msg("expected *sql.DB, got a different type")
	}

	userRepo := mysql.NewUserRepository(mysqlDB)
	authRepo := mysql.NewAuthRepository(mysqlDB)

	// Innit Services
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(authRepo)

	// Innit Handlers
	healthHandler := handlers.NewHealthHandler(conf)
	authHandler := handlers.NewAuthHandler(userService, authService)
	userHandler := handlers.NewUserHandler(userService)

	router := http.NewRouter(conf, healthHandler, authHandler, userHandler)
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
