package factory

import (
	"database/sql"
	"errors"
	"github.com/Andrew-UA/product-list/app/repositories"
	"github.com/Andrew-UA/product-list/app/repositories/mongodb"
	"github.com/Andrew-UA/product-list/app/repositories/mysql"
	"github.com/Andrew-UA/product-list/app/repositories/postgres"
	"github.com/Andrew-UA/product-list/app/repositories/sqlite"
	"github.com/Andrew-UA/product-list/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRepository(dbConnector db.DatabaseConnector) (repositories.Repository, error) {
	conn := dbConnector.Connection()
	switch dbConnector.Type() {
	case db.MySQL:
		dbConn, _ := conn.(*sql.DB)
		return mysql.NewRepository(dbConn), nil
	case db.Postgres:
		dbConn, _ := conn.(*sql.DB)
		return postgres.NewRepository(dbConn), nil
	case db.SQLite:
		dbConn, _ := conn.(*sql.DB)
		return sqlite.NewRepository(dbConn), nil
	case db.Mongo:
		dbConn, _ := conn.(*mongo.Client)
		return mongodb.NewRepository(dbConn), nil
	default:
		return nil, errors.New("invalid db type")
	}
}
