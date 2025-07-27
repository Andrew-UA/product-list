package db

import (
	"context"
	"fmt"
	"github.com/Andrew-UA/product-list/internal/config"
)

type DatabaseType string

const (
	MySQL    DatabaseType = "mysql"
	SQLite   DatabaseType = "sqlite"
	Postgres DatabaseType = "postgres"
	Mongo    DatabaseType = "mongo"
)

type DatabaseConnector interface {
	Connect() error
	Close(ctx context.Context) error
	Type() DatabaseType
	Connection() any
}

func GetDataBaseConnector(cfg *config.Config) (DatabaseConnector, error) {

	switch DatabaseType(cfg.DbType) {
	case MySQL:
		return NewMySQLConnector(cfg), nil

	case Postgres:
		return NewPostgresConnector(cfg), nil

	case SQLite:
		return NewSQLiteConnector(cfg), nil

	case Mongo:
		return NewMongoConnector(cfg), nil

	default:
		return nil, fmt.Errorf("unknown database type: %s", cfg.DbType)
	}
}
