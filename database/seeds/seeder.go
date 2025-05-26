package seeds

//
//import (
//	"context"
//	"github.com/Andrew-UA/product-list/app/dto"
//	"github.com/Andrew-UA/product-list/app/models"
//	"github.com/Andrew-UA/product-list/app/repositories/interfaces"
//	"github.com/Andrew-UA/product-list/internal/config"
//	"time"
//)
//
//type Seeder struct {
//	cfg            config.Config
//	userRepository interfaces.UserRepository
//}
//
//func NewSeeder(cfg config.Config, userRepository interfaces.UserRepository) *Seeder {
//	return &Seeder{
//		cfg:            cfg,
//		userRepository: userRepository,
//	}
//}
//
//func (s *Seeder) Seed() error {
//	return s.createAdmin()
//}
//
//
//
//func (s *Seeder) createAdmin()  error {
//	user := dto.UserDTO{
//		FirstName:  "Admin",
//		SecondName: "Admin",
//		Email:      s.cfg.AdminEmail,
//		Nickname:   nil,
//		Role:       models.USER_ADMIN_ROLE,
//		Password:   s.cfg.AdminPassword,
//		CreatedAt:  time.Time{},
//		UpdatedAt:  time.Time{},
//		DeletedAt:  nil,
//	}
//
//	s.userRepository.Create(context.Background(), user)
//}
