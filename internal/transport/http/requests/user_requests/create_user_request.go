package user_requests

import (
	"github.com/Andrew-UA/product-list/app/dto"
	"github.com/Andrew-UA/product-list/internal/transport/http/requests"
)

type CreateUserRequest struct {
	requests.BaseRequest
	FirstName  string  `json:"first_name" validate:"required,min=2,max=50"`
	SecondName string  `json:"second_name" validate:"required,min=2,max=50"`
	Email      string  `json:"email" validate:"required,email"`
	Nickname   *string `json:"nickname" validate:"omitempty,min=2,max=50"`
}

func (r *CreateUserRequest) ToDTO() dto.UserDTO {
	userDTO := dto.UserDTO{
		FirstName:  &r.FirstName,
		SecondName: &r.SecondName,
		Email:      &r.Email,
		Nickname:   dto.Nullable[string]{},
	}

	if r.Nickname != nil {
		userDTO.Nickname.SetValue(*r.Nickname)
	}

	return userDTO
}
