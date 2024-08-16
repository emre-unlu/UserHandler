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
