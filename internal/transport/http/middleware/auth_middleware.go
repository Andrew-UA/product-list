package middleware

import (
	"context"
	"github.com/Andrew-UA/product-list/app/services"
	"github.com/Andrew-UA/product-list/pkg/auth"
	"net/http"
	"strconv"
	"strings"
)

type contextKey string

const AuthUserKey contextKey = "authUser"

type AuthMiddleware struct {
	UserService  services.UsersServiceInterface
	AuthService  services.AuthServiceInterface
	TokenManager auth.TokenManager
}

func NewAuthMiddleware(userService services.UsersServiceInterface, authService services.AuthServiceInterface, tokenManager auth.TokenManager) *AuthMiddleware {
	return &AuthMiddleware{
		UserService:  userService,
		AuthService:  authService,
		TokenManager: tokenManager,
	}
}

func (m *AuthMiddleware) HandleFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		claims, err := m.TokenManager.ParseToken(strings.TrimPrefix(authHeader, "Bearer "))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		token, err := m.AuthService.GetTokenByTokenId(ctx, claims.Id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		authorizedUserID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		user, err := m.UserService.GetUserById(ctx, uint64(authorizedUserID))
		if err != nil || user == nil || token == nil || user.ID != token.UserID {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, AuthUserKey, user)

		next(w, r.WithContext(ctx))
	}
}
