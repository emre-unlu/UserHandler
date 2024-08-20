package services

import (
	"errors"
	"github.com/emre-unlu/GinTest/internal"
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/models"
	"github.com/emre-unlu/GinTest/internal/utils"
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
	isUserExist, err := s.userRepo.CheckUserByEmail(userDto.Email)

	if isUserExist != (models.User{}) {
		return dtos.UserDto{}, "", errors.New("There is a active or suspended user with this same email")
	}

	generatedPassword, err := passwordgen.GeneratePassword(10)
	if err != nil {
		return dtos.UserDto{}, "", errors.New("Error generating password")
	}

	hashedPassword, err := utils.HashPassword(generatedPassword)
	if err != nil {
		return dtos.UserDto{}, "", errors.New("Error hashing password")
	}

	user, err := userDto.ToUser()
	user.Password = hashedPassword

	generatedUser, err := s.userRepo.CreateUser(user)
	return dtos.ToUserDto(generatedUser), generatedPassword, err
}

func (s *UserService) GetUserById(id uint) (dtos.UserDto, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return dtos.UserDto{}, internal.ErrUserNotFound
	}
	return dtos.ToUserDto(user), nil
}
func (s *UserService) GetUsersWithPagination(page int, limit int) ([]models.User, int64, error) {
	return s.userRepo.GetUsersWithPagination(page, limit)
}
func (s *UserService) DeactivateUserById(id uint) (dtos.UserDto, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return dtos.UserDto{}, internal.ErrUserNotFound
	}
	if user.Status == models.StatusInactive {
		return dtos.UserDto{}, internal.ErrUserAlreadyDisactive
	}

	user.Status = models.StatusInactive
	updatedUser, err := s.userRepo.DeactivateUserById(user)

	if err != nil {
		return dtos.UserDto{}, err
	}
	return dtos.ToUserDto(updatedUser), err
}
func (s *UserService) SuspendUserById(id uint) (dtos.UserDto, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return dtos.UserDto{}, internal.ErrUserNotFound
	}
	if user.Status == models.StatusSuspended {
		return dtos.UserDto{}, internal.ErrUserAlreadySuspended
	}
	if user.Status == models.StatusInactive {
		return dtos.UserDto{}, internal.ErrUserDeleted
	}

	user.Status = models.StatusSuspended
	updatedUser, err := s.userRepo.SuspendUserById(user)

	if err != nil {
		return dtos.UserDto{}, err
	}
	return dtos.ToUserDto(updatedUser), err
}
func (s *UserService) ActivateUserById(id uint) (dtos.UserDto, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return dtos.UserDto{}, internal.ErrUserNotFound
	}
	if user.Status == models.StatusActive {
		return dtos.UserDto{}, internal.ErrUserAlreadyActive
	}
	if user.Status == models.StatusInactive {
		return dtos.UserDto{}, internal.ErrUserDeleted
	}

	user.Status = models.StatusActive
	updatedUser, err := s.userRepo.ActivateUserById(user)

	if err != nil {
		return dtos.UserDto{}, err
	}
	return dtos.ToUserDto(updatedUser), err
}
func (s *UserService) UpdateUser(id uint, updatedUserDto dtos.UserDto) (dtos.UserDto, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return dtos.UserDto{}, internal.ErrUserNotFound
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

	updatedUser, err := s.userRepo.UpdateUser(user)
	if err != nil {
		return dtos.UserDto{}, internal.ErrUserNotFound
	}
	return dtos.ToUserDto(updatedUser), nil

}
func (s *UserService) UpdatePassword(id uint, PasswordUpdateDto dtos.PasswordUpdateDto) error {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return internal.ErrUserNotFound
	}
	if !utils.VerifyPassword(PasswordUpdateDto.OldPassword, user.Password) {

		return internal.ErrIncorrectPassword
	}

	hashedPassword, err := utils.HashPassword(PasswordUpdateDto.NewPassword)
	if err != nil {

		return internal.ErrHashingError
	}
	user.Password = hashedPassword
	_, err = s.userRepo.UpdateUser(user)
	return err
}
