package auth_requests

import (
	"github.com/Andrew-UA/product-list/internal/transport/http/requests"
)

type LoginRequest struct {
	requests.BaseRequest
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
