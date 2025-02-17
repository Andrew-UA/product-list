package services

import (
	"context"
	"github.com/Andrew-UA/product-list/app/dto"
	"github.com/Andrew-UA/product-list/app/models"
	"github.com/Andrew-UA/product-list/app/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (u UserService) GetUsers(ctx context.Context) ([]models.User, error) {
	return u.userRepository.GetList(ctx)
}

func (u UserService) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return u.userRepository.GetById(ctx, id)
}

func (u UserService) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	return u.userRepository.GetByLogin(ctx, login)
}

func (u UserService) CreateUser(ctx context.Context, data dto.UserDTO) (*models.User, error) {
	user := &models.User{}
	user.Role = models.USER_DEFAULT_ROLE

	return u.userRepository.Create(ctx, user)
}

func (u UserService) UpdateUser(ctx context.Context, data dto.UserDTO, user *models.User) (*models.User, error) {

	return u.userRepository.Update(ctx, user)
}

func (u UserService) DeleteUser(ctx context.Context, user *models.User) error {
	return u.userRepository.Delete(ctx, user)
}

func (u UserService) mapToUser(_ context.Context, data dto.UserDTO, user *models.User) *models.User {
	//TODO:

	return user
}
