package user_requests

import "github.com/Andrew-UA/product-list/internal/transport/http/requests"

type CreateUserRequest struct {
	requests.BaseRequest
	FirstName  string `json:"first_name" validate:"required,min=2,max=50"`
	SecondName string `json:"second_name" validate:"required,min=2,max=50"`
	Email      string `json:"login" validate:"required,min=3,max=20,email"`
}
