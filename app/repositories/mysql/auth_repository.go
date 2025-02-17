package mysql

import (
	"context"
	"database/sql"
	"github.com/Andrew-UA/product-list/app/models"
)

type AuthRepository struct {
	conn *sql.DB
}

func NewAuthRepository(conn *sql.DB) *AuthRepository {
	return &AuthRepository{
		conn: conn,
	}
}

func (a AuthRepository) SetPassword(ctx context.Context, user models.User, password string) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthRepository) SetToken(ctx context.Context, user models.User, token string) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthRepository) GetUserToken(ctx context.Context, user models.User) (string, error) {
	//TODO implement me
	panic("implement me")
}
