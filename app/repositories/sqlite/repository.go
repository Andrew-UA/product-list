package sqlite

import (
	"database/sql"
	"github.com/Andrew-UA/product-list/app/repositories"
)

type Repository struct {
	db       *sql.DB
	authRepo repositories.AuthRepository
	userRepo repositories.UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) AuthRepository() repositories.AuthRepository {
	if r.authRepo == nil {
		r.authRepo = NewAuthRepository(r.db)
	}

	return r.authRepo
}

func (r *Repository) UserRepository() repositories.UserRepository {
	if r.userRepo == nil {
		r.userRepo = NewUserRepository(r.db)
	}

	return r.userRepo
}
