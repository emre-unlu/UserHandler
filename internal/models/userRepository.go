package models

type UserRepository interface {
	GetUserList(page int, limit int) ([]User, int64, error)
	GetUserById(id uint) (User, error)
	CreateUser(user User) (User, error)
	DeactivateUserById(user User) (User, error)
	ActivateUserById(user User) (User, error)
	SuspendUserById(user User) (User, error)
	UpdateUser(user User) (User, error)
	UpdatePassword(id uint, newPassword string) error
	CheckUserByEmail(email string) (*User, error)
}
