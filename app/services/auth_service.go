package services

import (
	"context"
	models "github.com/Andrew-UA/product-list/app/models"
	"github.com/Andrew-UA/product-list/app/repositories"
)

type AuthService struct {
	AuthRepository repositories.AuthRepository
}

func NewAuthService(repository repositories.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: repository,
	}
}

func (authService *AuthService) Login(ctx context.Context, user *models.User, password string) (string, error) {
	return "", nil
}

func (authService *AuthService) Logout(ctx context.Context, user *models.User) error {
	return nil
}
