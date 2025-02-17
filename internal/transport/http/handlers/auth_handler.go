package handlers

import (
	"github.com/Andrew-UA/product-list/app/services"
	"net/http"
)

type AuthHandler struct {
	UserService services.UsersServiceInterface
	AuthService services.AuthServiceInterface
}

func NewAuthHandler(userService services.UsersServiceInterface, authService services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		UserService: userService,
		AuthService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {

}
