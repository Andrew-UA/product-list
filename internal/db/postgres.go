package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Andrew-UA/product-list/internal/config"

	_ "github.com/lib/pq"
)

type PostgresConnector struct {
	cfg *config.Config
	db  *sql.DB
}

// NewPostgresConnector створює новий конектор на основі конфігурації
func NewPostgresConnector(cfg *config.Config) *PostgresConnector {
	return &PostgresConnector{cfg: cfg}
}

// Connect створює з'єднання з PostgresSQL
func (p *PostgresConnector) Connect() (any, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", p.cfg.DbUser, p.cfg.DbPassword, p.cfg.DbHost, p.cfg.DbPort, p.cfg.DbName, p.cfg.DbSslMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("PostgreSQL is not responding: %w", err)
	}
	return db, nil
}

func (p *PostgresConnector) Close(ctx context.Context) error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}
