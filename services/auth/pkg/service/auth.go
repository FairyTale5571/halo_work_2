package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fairytale5571/halo_work_2/services/auth/pkg/errs"
)

type AuthService struct{}

const userName = "user-name"

type tokenClaims struct {
	jwt.StandardClaims
	Id string
}

func NewAuth() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Authorize(username string) bool {
	if username == userName {
		return true
	}
	return false
}

func (s *AuthService) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		username,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.ErrorInvalidSignMethod
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errs.ErrorCantClaimToken
	}

	return claims.Id, nil
}
