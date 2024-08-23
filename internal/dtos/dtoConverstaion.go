package dtos

import (
	"fmt"
	"github.com/emre-unlu/GinTest/internal/models"
	"time"
)

const BirthdateFormat = "2006-01-02"

func ToUserDto(user *models.User) UserDto {
	return UserDto{
		ID:        user.ID,
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		Phone:     user.Phone,
		Birthdate: user.Birthdate.Format(BirthdateFormat),
		Status:    user.Status,
	}
}

func (u *UserDto) ToUser() (*models.User, error) {
	birthdate, err := time.Parse(BirthdateFormat, u.Birthdate)
	if err != nil {
		return nil, fmt.Errorf("failed to convert User to user dto: %w", err)
	}
	return &models.User{
		ID:        u.ID,
		Name:      u.Name,
		Surname:   u.Surname,
		Email:     u.Email,
		Phone:     u.Phone,
		Birthdate: birthdate,
		Status:    u.Status,
	}, err
}
func ConvertUsersToDtos(users []models.User) []UserDto {
	dtos := make([]UserDto, len(users))
	for i, user := range users {
		dtos[i] = ToUserDto(&user)
	}
	return dtos
}
