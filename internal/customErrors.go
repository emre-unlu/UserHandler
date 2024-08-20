package internal

import "errors"

var (
	ErrUserAlreadyDisactive = errors.New("user is already disactive")
	ErrUserAlreadyActive    = errors.New("user is already active")
	ErrUserAlreadySuspended = errors.New("user is already suspended")
	ErrUserDeleted          = errors.New("Deleted user cannot be reactivated")
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrHashingError      = errors.New("error hashing password")
)
