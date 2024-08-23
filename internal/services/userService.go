package services

import (
	"errors"
	"fmt"
	"github.com/emre-unlu/GinTest/internal"
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/models"
	"github.com/emre-unlu/GinTest/internal/utils"
	"github.com/emre-unlu/GinTest/pkg/InformationSystem"
	"github.com/emre-unlu/go-passwordgen/passwordgen"
	"time"
)

type UserService struct {
	userRepo models.UserRepository
}

func NewUserService(userRepo models.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(userDto dtos.UserDto) (dtos.UserDto, error) {

	isUserExist, err := s.userRepo.CheckUserByEmail(userDto.Email)
	if err != nil {
		return dtos.UserDto{}, err
	}

	if isUserExist != nil {
		return dtos.UserDto{}, internal.ErrThereIsActiveOrSuspendedUser
	}

	generatedPassword, err := passwordgen.GeneratePassword(10)
	if err != nil {
		return dtos.UserDto{}, internal.ErrGeneratingPassword
	}

	hashedPassword, err := utils.HashPassword(generatedPassword)
	if err != nil {
		return dtos.UserDto{}, internal.ErrHashingError
	}

	user, err := userDto.ToUser()
	user.Password = hashedPassword

	generatedUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return dtos.UserDto{}, err
	}

	// Send the email with the password
	subject := "Your Account Has Been Created"
	body := fmt.Sprintf("Dear %s,\n\nYour account has been successfully created.  \nYour password is: ( %s ) \nBest Wishes by Emre", userDto.Name, generatedPassword)
	err = InformationSystem.SendEmail(userDto.Email, subject, body)
	if err != nil {
		return dtos.ToUserDto(generatedUser), fmt.Errorf("user created but failed to send email: \n With id : %d \n With password : %s", userDto.ID, generatedPassword)
	}

	return dtos.ToUserDto(generatedUser), err
}

func (s *UserService) GetUserById(id uint) (dtos.UserDto, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return dtos.UserDto{}, err
	}
	return dtos.ToUserDto(user), nil
}
func (s *UserService) GetUserList(userListRequestDto dtos.UserListRequestDto) (*dtos.UserListDto, error) {
	users, total, err := s.userRepo.GetUserList(userListRequestDto.Page, userListRequestDto.Limit)
	if err != nil {
		return nil, err
	}
	userDtos := dtos.ConvertUsersToDtos(users)

	userListDto := &dtos.UserListDto{
		Total: uint(total),
		Users: userDtos,
	}

	return userListDto, nil
}

func (s *UserService) SuspendUserById(id uint) error {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return internal.ErrUserNotFound
	}
	if user.Status == models.StatusSuspended {
		return internal.ErrUserAlreadySuspended
	}
	if user.Status == models.StatusInactive {
		return internal.ErrUserDeleted
	}

	user.Status = models.StatusSuspended
	err = s.userRepo.SuspendUserById(user)

	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeactivateUserById(id uint) error {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return internal.ErrUserNotFound
	}
	if user.Status == models.StatusInactive {
		return internal.ErrUserAlreadyDisactive
	}

	user.Status = models.StatusInactive
	err = s.userRepo.DeactivateUserById(user)

	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) ActivateUserById(id uint) error {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return internal.ErrUserNotFound
	}
	if user.Status == models.StatusActive {
		return internal.ErrUserAlreadyActive
	}
	if user.Status == models.StatusInactive {
		return internal.ErrUserDeleted
	}

	user.Status = models.StatusActive
	err = s.userRepo.ActivateUserById(user)

	if err != nil {
		return err
	}
	return err
}

func (s *UserService) UpdateUser(id uint, updatedUserDto dtos.UserDto) error {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return err
	}

	isUserFound, err := s.userRepo.CheckUserByEmail(updatedUserDto.Email)
	if err != nil {
		return err
	}

	if isUserFound != nil {
		return internal.ErrThereIsActiveOrSuspendedUser
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
			return errors.New("invalid birthdate")
		}
		user.Birthdate = birthdate
	}

	err = s.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil

}

func (s *UserService) UpdatePassword(id uint, PasswordUpdateDto dtos.PasswordUpdateDto) error {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return err
	}
	if !utils.VerifyPassword(PasswordUpdateDto.OldPassword, user.Password) {

		return internal.ErrIncorrectPassword
	}

	hashedPassword, err := utils.HashPassword(PasswordUpdateDto.NewPassword)
	if err != nil {

		return internal.ErrHashingError
	}
	user.Password = hashedPassword
	err = s.userRepo.UpdateUser(user)

	if err != nil {
		return err
	}

	return nil
}
