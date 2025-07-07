package services

import (
	"context"
	"errors"
	models "github.com/Andrew-UA/product-list/app/models"
	"github.com/Andrew-UA/product-list/app/repositories"
	"github.com/Andrew-UA/product-list/pkg/auth"
	"strconv"
	"time"
)

var ttl = time.Hour * 24

type AuthService struct {
	passwordManager auth.PasswordManager
	tokenManager    auth.TokenManager
	authRepository  repositories.AuthRepository
}

func NewAuthService(passwordManager auth.PasswordManager, tokenManager auth.TokenManager, repository repositories.AuthRepository) *AuthService {
	return &AuthService{
		passwordManager: passwordManager,
		tokenManager:    tokenManager,
		authRepository:  repository,
	}
}

func (a *AuthService) Login(ctx context.Context, user *models.User, password string) (string, error) {

	isValidPassword := a.passwordManager.CheckPassword(user.PasswordHash, password)
	if !isValidPassword {
		return "", errors.New("invalid password")
	}

	tokenString, tokenId, err := a.tokenManager.CreateToken(strconv.FormatUint(user.ID, 10), ttl)
	if err != nil {
		return "", err
	}

	token := models.JwtAuth{
		UserID:    user.ID,
		TokenID:   tokenId,
		ExpiresAt: time.Now().Add(ttl),
	}

	_, err = a.authRepository.CreateToken(ctx, token)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *AuthService) Logout(ctx context.Context, user *models.User) error {
	token, err := a.authRepository.GetTokenByUserId(ctx, user.ID)
	if err != nil {
		return err
	}

	if token == nil {
		return nil
	}

	return a.authRepository.DeleteToken(ctx, *token)
}

func (a *AuthService) GetTokenByTokenId(ctx context.Context, tokenId string) (*models.JwtAuth, error) {
	return a.authRepository.GetTokenByTokenId(ctx, tokenId)
}
