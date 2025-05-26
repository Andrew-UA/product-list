package handlers

import (
	"errors"
	"github.com/Andrew-UA/product-list/app/services"
	"github.com/Andrew-UA/product-list/internal/transport/http/requests"
	"github.com/Andrew-UA/product-list/internal/transport/http/requests/auth_requests"
	"github.com/Andrew-UA/product-list/internal/transport/http/responses"
	"github.com/Andrew-UA/product-list/internal/transport/http/responses/auth_responses"
	"github.com/Andrew-UA/product-list/internal/validation"
	"net/http"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type AuthHandler struct {
	Validator   validation.ValidatorInterface
	UserService services.UsersServiceInterface
	AuthService services.AuthServiceInterface
}

func NewAuthHandler(validator validation.ValidatorInterface, userService services.UsersServiceInterface, authService services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		Validator:   validator,
		UserService: userService,
		AuthService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request auth_requests.LoginRequest

	ctx := r.Context()

	err := requests.ReadAndCLose(r, &request)
	if err != nil {
		responses.NewErrorResponse(err).SetStatusCode(http.StatusBadRequest).Write(w)
		return
	}

	err = h.Validator.ValidateStruct(&request)
	if err != nil {
		responses.NewErrorResponse(err).SetStatusCode(http.StatusBadRequest).Write(w)
	}

	user, err := h.UserService.GetUserByEmail(ctx, request.Email)
	if err != nil || user == nil {
		responses.NewErrorResponse(ErrInvalidCredentials).SetStatusCode(http.StatusBadRequest).Write(w)
		return
	}

	token, err := h.AuthService.Login(ctx, user, request.Password)
	if err != nil {
		responses.NewErrorResponse(ErrInvalidCredentials).SetStatusCode(http.StatusBadRequest).Write(w)
		return
	}

	auth_responses.NewLoginResponse(token).SetStatusCode(200).Write(w)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {

}
