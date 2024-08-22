package services

import (
	"github.com/emre-unlu/GinTest/internal"
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/models"
	"github.com/emre-unlu/GinTest/internal/utils"
)

type AuthService struct {
	userRepo models.UserRepository
}

func NewAuthService(userRepo models.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(email string, password string) (loginDto dtos.LoginDto, err error) {
	user, err := s.userRepo.CheckUserByEmail(email)
	if user == nil {
		return dtos.LoginDto{}, internal.ErrUserNotFound
	}

	if !utils.VerifyPassword(password, user.Password) {
		return dtos.LoginDto{}, internal.ErrIncorrectPassword
	}

	loginDto = dtos.ConvertUserToLoginDto(user)

	return loginDto, nil
}
