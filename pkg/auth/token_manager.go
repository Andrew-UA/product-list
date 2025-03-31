package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type TokenManager interface {
	CreateToken(userId string, ttl time.Duration) (string, string, error)
	ParseToken(token string) (*jwt.StandardClaims, error)
}

type JWTTokenManager struct {
	signingKey string
}

func NewJWTTokenManager(signingKey string) *JWTTokenManager {

	return &JWTTokenManager{signingKey: signingKey}
}

func (t *JWTTokenManager) CreateToken(subject string, ttl time.Duration) (tokenString string, tokenId string, err error) {
	tokenId = uuid.New().String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        tokenId,
		ExpiresAt: time.Now().Add(ttl).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   subject,
	})

	tokenString, err = token.SignedString([]byte(t.signingKey))
	if err != nil {
		return "", "", err
	}

	return
}

func (t *JWTTokenManager) ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
