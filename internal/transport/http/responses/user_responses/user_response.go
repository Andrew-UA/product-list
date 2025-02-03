package user_responses

import (
	models "github.com/Andrew-UA/product-list/app/ models"
	"github.com/Andrew-UA/product-list/internal/transport/http/responses"
	"time"
)

type UserResponse struct {
	responses.Response

	ID         uint       `json:"id"`
	FirstName  string     `json:"first_name"`
	SecondName string     `json:"second_name"`
	Email      string     `json:"email"`
	Role       string     `json:"role"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		Email:      user.Email,
		Role:       string(user.Role),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeletedAt,
	}
}
