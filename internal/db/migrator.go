package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Andrew-UA/product-list/database/migrations"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func RunMigrations(dbConnector DatabaseConnector) error {
	dbType := dbConnector.Type()
	conn := dbConnector.Connection()

	migrationsPath := fmt.Sprintf("%s", dbType)

	sourceDriver, err := iofs.New(migrations.MigrationFiles, migrationsPath)
	if err != nil {
		return fmt.Errorf("RunMigrations: failed to init embedded migration source: %w", err)
	}

	var dbDriver database.Driver

	switch dbType {
	case Postgres:
		db := conn.(*sql.DB)
		dbDriver, err = postgres.WithInstance(db, &postgres.Config{})
	case SQLite:
		db := conn.(*sql.DB)
		dbDriver, err = sqlite3.WithInstance(db, &sqlite3.Config{})
	case MySQL:
		db := conn.(*sql.DB)
		dbDriver, err = mysql.WithInstance(db, &mysql.Config{})
	case Mongo:
		return nil
	default:
		return fmt.Errorf("RunMigrations: unsupported db type: %s", dbType)
	}

	if err != nil {
		return fmt.Errorf("RunMigrations: failed to create db driver: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", sourceDriver, string(dbType), dbDriver)
	if err != nil {
		return fmt.Errorf("RunMigrations: failed to init migrate: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("RunMigrations: migration error: %w", err)
	}

	return nil
}
