package user_requests

import (
	"github.com/Andrew-UA/product-list/app/dto"
	"github.com/Andrew-UA/product-list/internal/transport/http/requests"
)

type UpdateUserRequest struct {
	requests.BaseRequest
	FirstName  *string `json:"first_name,omitempty" validate:"omitempty,min=2,max=50"`
	SecondName *string `json:"second_name,omitempty" validate:"omitempty,min=2,max=50"`
	Email      *string `json:"email,omitempty" validate:"omitempty,email"`
	Nickname   *string `json:"nickname,omitempty" validate:"omitempty,min=2,max=50"`
}

func (r *UpdateUserRequest) ToDTO() dto.UserDTO {
	var exist, isNull bool
	userDTO := dto.UserDTO{
		FirstName:  r.FirstName,
		SecondName: r.SecondName,
		Email:      r.Email,
		Nickname:   dto.Nullable[string]{},
	}

	exist, isNull = r.HasField("nickname")
	if exist && isNull {
		userDTO.Nickname.SetNull()
	}

	return userDTO
}
