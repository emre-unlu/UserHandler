package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:name`
	Age  int    `jason:age`
}

type UserRepository interface {
	GetAllUsers() ([]User, error)
	GetUserById(id uint) (User, error)
	CreateUser(user User) (User, error)
	DeleteUserById(id uint) (User, error)
}

type PGUserRepository struct {
	DB *gorm.DB
}

func NewPGUserRepository(db *gorm.DB) *PGUserRepository {
	return &PGUserRepository{DB: db}
}

func (r *PGUserRepository) GetAllUsers() ([]User, error) {
	var users []User
	result := r.DB.Find(&users)
	return users, result.Error
}
func (r *PGUserRepository) GetUserById(id uint) (User, error) {
	var user User
	result := r.DB.First(&user, id)
	return user, result.Error
}
func (r *PGUserRepository) CreateUser(user User) (User, error) {
	result := r.DB.Create(&user)
	return user, result.Error
}
func (r *PGUserRepository) DeleteUserById(id uint) (User, error) {
	result := r.DB.Delete(&User{}, id)
	return User{}, result.Error
}
