package user_responses

import (
	"github.com/Andrew-UA/product-list/app/models"
	"time"
)

type UserResponseData struct {
	ID         uint64     `json:"id"`
	FirstName  string     `json:"first_name"`
	SecondName string     `json:"second_name"`
	Email      string     `json:"email"`
	Role       string     `json:"role"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

func NewUserResponse(user *models.User) *UserResponseData {
	return &UserResponseData{
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

type UserListResponseData struct {
	Users []*UserResponseData `json:"users"`
}

func NewUserListResponse(users []models.User) *UserListResponseData {
	usersData := make([]*UserResponseData, len(users))
	for i, user := range users {
		usersData[i] = NewUserResponse(&user)
	}

	return &UserListResponseData{
		Users: usersData,
	}
}
