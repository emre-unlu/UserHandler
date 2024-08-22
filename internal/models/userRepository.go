package models

type UserRepository interface {
	GetUserList(page int, limit int) ([]User, int64, error)
	GetUserById(id uint) (*User, error)
	CreateUser(user User) (*User, error)
	SuspendUserById(user *User) error
	DeactivateUserById(user *User) error
	ActivateUserById(user *User) error
	UpdateUser(user *User) error
	UpdatePassword(id uint, newPassword string) error
	CheckUserByEmail(email string) (*User, error)
}
