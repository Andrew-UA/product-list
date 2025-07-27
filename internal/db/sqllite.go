package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Andrew-UA/product-list/internal/config"
	"os"
	"path/filepath"

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
func (s *SQLiteConnector) Connect() error {
	if _, err := os.Stat(s.cfg.DbFilepath); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(s.cfg.DbFilepath), 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to stat database file: %w", err)
	}

	db, err := sql.Open("sqlite3", s.cfg.DbFilepath)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	s.db = db

	return nil
}

func (s *SQLiteConnector) Close(ctx context.Context) error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

func (s *SQLiteConnector) Type() DatabaseType {
	return SQLite
}

func (m *SQLiteConnector) Connection() any {
	return m.db
}
