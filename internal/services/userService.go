package services

import (
	"errors"
	"github.com/emre-unlu/GinTest/internal/models"
)

type UserService struct {
	userRepo models.UserRepository
}

func NewUserService(userRepo models.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	if user.Name == "" {
		return models.User{}, errors.New("name is required")
	}
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUserById(id uint) (models.User, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}
func (s *UserService) GetAllusers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}
func (s *UserService) DeleteUserById(id uint) (models.User, error) {
	return s.userRepo.DeleteUserById(id)
}
