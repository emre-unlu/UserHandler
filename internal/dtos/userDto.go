package dtos

import "github.com/emre-unlu/GinTest/internal/models"

type UserDto struct {
	ID        uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string        `json:"name" validate:"required"`
	Surname   string        `json:"surname" validate:"required"`
	Email     string        `json:"email" validate:"required,email"`
	Phone     string        `json:"phone" validate:"required,e164"`
	Birthdate string        `json:"birthdate" validate:"required,datetime=2006-01-02"`
	Status    models.Status `json:"status" `
}

type PasswordUpdateDto struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,password-strength"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Id           uint   `json:"id"`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type UserListRequestDto struct {
	Page  uint `form:"page" validate:"required,gte=0"`
	Limit uint `form:"limit" validate:"required,gt=0"`
}

type UserListDto struct {
	Total uint      `json:"total"`
	Users []UserDto `json:"users"`
}

type ErrorMessageDto struct {
	Message string `json:"message"`
}
type ForgotPasswordDto struct {
	Email string `json:"email" validate:"required,email"`
}
