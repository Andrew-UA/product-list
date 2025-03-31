package sqlite

import (
	"database/sql"
	"github.com/Andrew-UA/product-list/app/repositories/sqlbase"
	"github.com/Masterminds/squirrel"
)

type UserRepository struct {
	*sqlbase.UserRepository
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		UserRepository: &sqlbase.UserRepository{
			DB:     db,
			Format: squirrel.Question,
			Table:  "users",
		},
	}
}
