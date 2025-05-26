package sqlite

import (
	"database/sql"
	"github.com/Andrew-UA/product-list/app/repositories/interfaces"
	"github.com/Andrew-UA/product-list/internal/config"
)

type Repository struct {
	db       *sql.DB
	authRepo interfaces.AuthRepository
	userRepo interfaces.UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Seed(cfg config.Config) error {
	return nil
}

func (r *Repository) AuthRepository() interfaces.AuthRepository {
	if r.authRepo == nil {
		r.authRepo = NewAuthRepository(r.db)
	}

	return r.authRepo
}

func (r *Repository) UserRepository() interfaces.UserRepository {
	if r.userRepo == nil {
		r.userRepo = NewUserRepository(r.db)
	}

	return r.userRepo
}
