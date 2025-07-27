package mongodb

import (
	"github.com/Andrew-UA/product-list/app/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	client   *mongo.Client
	authRepo repositories.AuthRepository
	userRepo repositories.UserRepository
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) AuthRepository() repositories.AuthRepository {
	if r.authRepo == nil {
		r.authRepo = NewAuthRepository(r.client)
	}

	return r.authRepo
}

func (r *Repository) UserRepository() repositories.UserRepository {
	if r.userRepo == nil {
		r.userRepo = NewUserRepository(r.client)
	}

	return r.userRepo
}
