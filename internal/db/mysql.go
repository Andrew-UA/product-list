package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Andrew-UA/product-list/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnector struct {
	cfg *config.Config
	db  *sql.DB
}

// NewMySQLConnector створює новий конектор на основі конфігурації
func NewMySQLConnector(cfg *config.Config) *MySQLConnector {
	return &MySQLConnector{cfg: cfg}
}

// Connect створює з'єднання з MySQL
func (m *MySQLConnector) Connect() (any, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", m.cfg.DbUser, m.cfg.DbPassword, m.cfg.DbHost, m.cfg.DbPort, m.cfg.DbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("MySQL is not responding: %w", err)
	}
	return db, nil
}
func (m *MySQLConnector) Close(ctx context.Context) error {
	if m.db != nil {
		return m.db.Close()
	}
	return nil
}
