package service

import (
	"customer/internal/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login() (string, error)
}

type authService struct {
	server *types.Server
}

func NewAuthService(server *types.Server) AuthService {
	return &authService{
		server: server,
	}
}

func (s authService) Login() (string, error) {
	claims := jwt.MapClaims{
		"user": "nurhakiki",
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	_, tokenString, err := s.server.AuthToken.Encode(claims)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
