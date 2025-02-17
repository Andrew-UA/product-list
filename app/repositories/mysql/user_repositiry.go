package mysql

import (
	"context"
	"database/sql"
	"github.com/Andrew-UA/product-list/app/dto"
	"github.com/Andrew-UA/product-list/app/models"
)

type UserRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (u UserRepository) GetList(ctx context.Context) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) GetById(ctx context.Context, userId int64) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) GetByLogin(ctx context.Context, login string) (*models.User, error) {
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
