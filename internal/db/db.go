package db

import (
	"context"
	"fmt"
	"github.com/Andrew-UA/product-list/internal/config"
)

type DatabaseConnector interface {
	Connect() (any, error)
	Close(ctx context.Context) error
}

func GetDataBaseConnector(cfg *config.Config) (DatabaseConnector, error) {
	switch cfg.DbType {
	case "mysql":
		return NewMySQLConnector(cfg), nil

	case "postgres":
		return NewPostgresConnector(cfg), nil

	case "sqlite":
		return NewSQLiteConnector(cfg), nil

	case "mongo":
		return NewMongoConnector(cfg), nil

	default:
		return nil, fmt.Errorf("unknown database type: %s", cfg.DbType)
	}
}
