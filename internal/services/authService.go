package services

import (
	"errors"
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/models"
	"github.com/emre-unlu/GinTest/internal/utils"
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

func (s *AuthService) Login(email string, password string) (loginDto dtos.LoginDto, err error) {
	user, err := s.userRepo.CheckUserByEmail(email)
	if err != nil {
		return dtos.LoginDto{}, ErrInvalidCredentials
	}

	if !utils.VerifyPassword(password, user.Password) {
		return dtos.LoginDto{}, ErrInvalidCredentials
	}

	loginDto = dtos.LoginDto{
		Email:    user.Email,
		Password: user.Password,
		Id:       user.ID,
	}

	return loginDto, nil
}
