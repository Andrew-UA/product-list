package services

import (
	"context"
	models "github.com/Andrew-UA/product-list/app/models"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, user *models.User, password string) (string, error)
	Logout(ctx context.Context, user *models.User) error
}

type UsersServiceInterface interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserById(ctx context.Context, userId int64) (*models.User, error)
	GetUserByLogin(ctx context.Context, login string) (*models.User, error)
	CreateUser(ctx context.Context, data map[string]any) (*models.User, error)
	UpdateUser(ctx context.Context, data map[string]any, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, user *models.User) error
}
