package auth_responses

import "github.com/Andrew-UA/product-list/internal/transport/http/responses"

type LoginResponse struct {
	responses.ResponseJson
	Token string `json:"token"`
}

func NewLoginResponse(token string) *LoginResponse {
	response := &LoginResponse{
		Token: token,
	}
	response.SetStatusCode(200)

	return response
}
