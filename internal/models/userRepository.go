package models

// UserRepository defines the methods for user-related data operations.
type UserRepository interface {
	GetAllUsers() ([]User, error)
	GetUserById(id uint) (User, error)
	CreateUser(user User) (User, error)
	DeleteUserById(id uint) (User, error)
	UpdateUser(user User) (User, error)
	UpdatePassword(id uint, newPassword string) error
}
