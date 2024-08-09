package services

import (
	"errors"
	"github.com/emre-unlu/GinTest/internal"
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/models"
	"github.com/emre-unlu/go-passwordgen/passwordgen"
	"time"
)

type UserService struct {
	userRepo models.UserRepository
}

func NewUserService(userRepo models.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(userDto dtos.UserDto) (dtos.UserDto, string, error) {
	if userDto.Name == "" {
		return dtos.UserDto{}, "", errors.New("name is required")
	}
	generatedPassword, err := passwordgen.GeneratePassword(10)

	user, err := userDto.ToUser()
	user.Password = generatedPassword

	generatedUser, err := s.userRepo.CreateUser(user)
	return dtos.ToUserDto(generatedUser), generatedPassword, err
}

func (s *UserService) GetUserById(id uint) (dtos.UserDto, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return dtos.UserDto{}, errors.New("user not found")
	}
	return dtos.ToUserDto(user), nil
}
func (s *UserService) GetAllUsers() ([]dtos.UserDto, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var userDTOs []dtos.UserDto
	for _, user := range users {
		userDTOs = append(userDTOs, dtos.ToUserDto(user))
	}
	return userDTOs, nil
}
func (s *UserService) DeleteUserById(id uint) (dtos.UserDto, error) {
	user, err := s.userRepo.DeleteUserById(id)
	if err != nil {
		return dtos.UserDto{}, errors.New("unexpected error while deleting user")
	}
	return dtos.ToUserDto(user), err
}
func (s *UserService) UpdateUser(id uint, updatedUserDto dtos.UserDto, newPassword string) (dtos.UserDto, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return dtos.UserDto{}, errors.New("user not found")
	}
	if updatedUserDto.Name != "" {
		user.Name = updatedUserDto.Name
	}
	if updatedUserDto.Email != "" {
		user.Email = updatedUserDto.Email
	}
	if updatedUserDto.Surname != "" {
		user.Surname = updatedUserDto.Surname
	}
	if updatedUserDto.Phone != "" {
		user.Phone = updatedUserDto.Phone
	}
	if updatedUserDto.Birthdate != "" {
		birthdate, err := time.Parse(internal.BirthdayFormat, updatedUserDto.Birthdate)
		if err != nil {
			return dtos.UserDto{}, errors.New("invalid birthdate")
		}
		user.Birthdate = birthdate
	}
	if newPassword != "" {
		user.Password = newPassword
	}

	updatedUser, err := s.userRepo.UpdateUser(user)
	if err != nil {
		return dtos.UserDto{}, errors.New("user not found")
	}
	return dtos.ToUserDto(updatedUser), nil

}
