package user_requests

import "github.com/Andrew-UA/product-list/internal/transport/http/requests"

type UpdateUserRequest struct {
	requests.BaseRequest
	FirstName  *string `json:"first_name,omitempty" validate:"omitempty,min=2,max=50"`
	SecondName *string `json:"second_name,omitempty" validate:"omitempty,min=2,max=50"`
	Email      *string `json:"email,omitempty" validate:"omitempty,email"`
}
