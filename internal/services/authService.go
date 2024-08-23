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

func (s *AuthService) Login(email string, password string) (loginResponseDto *dtos.LoginResponseDto, err error) {
	user, err := s.userRepo.CheckUserByEmail(email)
	if user == nil {
		return nil, internal.ErrUserNotFound
	}

	if !utils.VerifyPassword(password, user.Password) {
		return nil, internal.ErrIncorrectPassword
	}

	response := dtos.LoginResponseDto{
		Id: user.ID,
	}

	return &response, nil
}
