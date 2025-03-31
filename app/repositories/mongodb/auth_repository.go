package mongodb

import (
	"context"
	"github.com/Andrew-UA/product-list/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	client *mongo.Client
}

func NewAuthRepository(client *mongo.Client) *AuthRepository {
	return &AuthRepository{client: client}
}

func (a AuthRepository) CreateToken(ctx context.Context, token models.JwtAuth) (*models.JwtAuth, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthRepository) GetTokenByUserId(ctx context.Context, userId uint64) (*models.JwtAuth, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthRepository) GetTokenByTokenId(ctx context.Context, tokenId string) (*models.JwtAuth, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthRepository) DeleteToken(ctx context.Context, token models.JwtAuth) error {
	//TODO implement me
	panic("implement me")
}
