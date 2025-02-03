package repositories

import (
	"context"
	models "github.com/Andrew-UA/product-list/app/ models"
)

type AuthRepository interface {
	SetPassword(ctx context.Context, user models.User, password string) error
	SetToken(ctx context.Context, user models.User, token string) error
	GetUserPassword(ctx context.Context, user models.User) (string, error)
	GetUserToken(ctx context.Context, user models.User) (string, error)
}

type UserRepository interface {
	GetList(ctx context.Context) ([]models.User, error)
	GetById(ctx context.Context, userId int64) (*models.User, error)
	GetByLogin(ctx context.Context, login string) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, user *models.User) error
}
