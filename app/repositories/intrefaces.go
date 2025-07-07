package repositories

import (
	"context"
	"github.com/Andrew-UA/product-list/app/dto"
	"github.com/Andrew-UA/product-list/app/models"
)

type AuthRepository interface {
	CreateToken(ctx context.Context, token models.JwtAuth) (*models.JwtAuth, error)
	GetTokenByUserId(ctx context.Context, userId uint64) (*models.JwtAuth, error)
	GetTokenByTokenId(ctx context.Context, tokenId string) (*models.JwtAuth, error)
	DeleteToken(ctx context.Context, token models.JwtAuth) error
}

type UserRepository interface {
	GetList(ctx context.Context) ([]models.User, error)
	GetById(ctx context.Context, userId uint64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, data dto.UserDTO) (*models.User, error)
	Update(ctx context.Context, data dto.UserDTO, user *models.User) (*models.User, error)
	Delete(ctx context.Context, user *models.User) error
}

type Repository interface {
	UserRepository() UserRepository
	AuthRepository() AuthRepository
}
