package internal

import "errors"

type ErrResponse struct {
	Error ServiceError `json:"error"`
}

// ServiceError has fields for Service errors. All fields with no data will
// be omitted
type ServiceError struct {
	Kind    string `json:"kind,omitempty"`
	Code    string `json:"code,omitempty"`
	Param   string `json:"param,omitempty"`
	Message string `json:"message,omitempty"`
}

var (
	ErrUserAlreadyDisactive         = errors.New("user is already disactive")
	ErrUserAlreadyActive            = errors.New("user is already active")
	ErrUserAlreadySuspended         = errors.New("user is already suspended")
	ErrUserDeleted                  = errors.New("Deleted user cannot be reactivated")
	ErrThereIsActiveOrSuspendedUser = errors.New("There is a active or suspended user with this same email")
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrIncorrectPassword  = errors.New("incorrect password")
	ErrHashingError       = errors.New("error hashing password")
	ErrGeneratingPassword = errors.New("error generating password")
)
