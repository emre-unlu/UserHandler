package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Birthdate time.Time `json:"birthdate"`
	Password  string    `json:"password"`
	Status    Status    `gorm:"default:'active'" json:"status"`
}
