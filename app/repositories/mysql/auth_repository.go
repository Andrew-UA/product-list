package mysql

import (
	"database/sql"
	"github.com/Andrew-UA/product-list/app/repositories/sqlbase"
	"github.com/Masterminds/squirrel"
)

type AuthRepository struct {
	*sqlbase.AuthRepository
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		AuthRepository: &sqlbase.AuthRepository{
			DB:        db,
			Format:    squirrel.Question,
			TableName: "jwt_auth",
		},
	}
}
