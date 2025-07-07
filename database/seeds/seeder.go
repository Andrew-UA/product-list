package seeds

import (
	"context"
	"errors"
	"fmt"
	"github.com/Andrew-UA/product-list/app/dto"
	"github.com/Andrew-UA/product-list/app/models"
	"github.com/Andrew-UA/product-list/app/repositories"
	rErrors "github.com/Andrew-UA/product-list/app/repositories/errors"
	"github.com/Andrew-UA/product-list/internal/config"
	"github.com/Andrew-UA/product-list/internal/helpers"
	"github.com/Andrew-UA/product-list/pkg/auth"
)

type Seeder struct {
	cfg             *config.Config
	passwordManager auth.PasswordManager
	userRepository  repositories.UserRepository
}

func NewSeeder(cfg *config.Config, passwordManager auth.PasswordManager, userRepository repositories.UserRepository) *Seeder {
	return &Seeder{
		cfg:             cfg,
		passwordManager: passwordManager,
		userRepository:  userRepository,
	}
}

func (s *Seeder) Seed() error {
	err := s.CreateAdmin()
	if err != nil {
		return err
	}

	return nil
}

func (s *Seeder) CreateAdmin() error {
	ctx := context.Background()
	adminEmail := s.cfg.AdminEmail

	admin, err := s.userRepository.GetByEmail(ctx, adminEmail)
	if err != nil && !errors.Is(err, rErrors.NotFoundError) {
		return fmt.Errorf("Seeder.CreateAdmin: unable to get admin by email %q: %w", adminEmail, err)
	}
	if admin != nil {
		return nil
	}

	hashedPwd, err := s.passwordManager.HashAndSalt(s.cfg.AdminPassword)
	if err != nil {
		return fmt.Errorf("Seeder.CreateAdmin: unable to hash password: %w", err)
	}

	adminData := dto.UserDTO{
		FirstName:  helpers.Ptr("Admin"),
		SecondName: helpers.Ptr("Admin"),
		Email:      helpers.Ptr(adminEmail),
		Role:       helpers.Ptr(string(models.USER_ADMIN_ROLE)),
		Nickname:   dto.NewNullable("Admin"),
		Password:   dto.NewNullable(hashedPwd),
	}

	admin, err = s.userRepository.Create(ctx, adminData)
	if err != nil {
		return fmt.Errorf("Seeder.CreateAdmin: unable to create admin: %w", err)
	}

	return nil
}
