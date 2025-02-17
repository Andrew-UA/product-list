package db

import (
	"context"
	"database/sql"
	"github.com/Andrew-UA/product-list/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteConnector struct {
	cfg *config.Config
	db  *sql.DB
}

// NewSQLiteConnector створює новий конектор на основі конфігурації
func NewSQLiteConnector(cfg *config.Config) *SQLiteConnector {
	return &SQLiteConnector{cfg: cfg}
}

// Connect створює з'єднання з SQLite
func (s *SQLiteConnector) Connect() (any, error) {
	db, err := sql.Open("sqlite3", s.cfg.DbFilepath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *SQLiteConnector) Close(ctx context.Context) error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
