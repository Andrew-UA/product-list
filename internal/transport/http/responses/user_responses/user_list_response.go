package user_responses

import (
	models "github.com/Andrew-UA/product-list/app/ models"
	"github.com/Andrew-UA/product-list/internal/transport/http/responses"
)

type UserListResponse struct {
	responses.Response

	Users []UserResponse
}

func NewUserListResponse(users []models.User) *UserListResponse {
	userResponses := make([]UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = NewUserResponse(&user)
	}

	return &UserListResponse{
		Users: userResponses,
	}
}
