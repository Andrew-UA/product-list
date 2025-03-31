package mongodb

import (
	"context"
	"github.com/Andrew-UA/product-list/app/dto"
	"github.com/Andrew-UA/product-list/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{client: client}
}

func (u UserRepository) GetList(ctx context.Context) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) GetById(ctx context.Context, userId uint64) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Create(ctx context.Context, data dto.UserDTO) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Update(ctx context.Context, data dto.UserDTO, user *models.User) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(ctx context.Context, user *models.User) error {
	//TODO implement me
	panic("implement me")
}
