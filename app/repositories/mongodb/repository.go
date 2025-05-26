package mongodb

import (
	"github.com/Andrew-UA/product-list/app/repositories/interfaces"
	"github.com/Andrew-UA/product-list/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	client   *mongo.Client
	authRepo interfaces.AuthRepository
	userRepo interfaces.UserRepository
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) Seed(cfg config.Config) error {
	return nil
}

func (r *Repository) AuthRepository() interfaces.AuthRepository {
	if r.authRepo == nil {
		r.authRepo = NewAuthRepository(r.client)
	}

	return r.authRepo
}

func (r *Repository) UserRepository() interfaces.UserRepository {
	if r.userRepo == nil {
		r.userRepo = NewUserRepository(r.client)
	}

	return r.userRepo
}
