package services

import (
	"errors"
	"github.com/emre-unlu/GinTest/internal/models"
	"github.com/emre-unlu/GinTest/pkg/utils"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type AuthService struct {
	userRepo models.UserRepository
}

func NewAuthService(userRepo models.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(email, password string) (accessToken, refreshToken string, err error) {
	user, err := s.userRepo.CheckUserByEmail(email)
	if err != nil {
		return "", "", ErrInvalidCredentials
	}

	if !utils.VerifyPassword(password, user.Password) {
		return "", "", ErrInvalidCredentials
	}

	accessToken, refreshToken, err = utils.GenerateJWT(email)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
