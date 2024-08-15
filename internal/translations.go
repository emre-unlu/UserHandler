package internal

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() *CustomValidator {
	validate := validator.New()
	validate.RegisterValidation("password-strength", ValidatePasswordStrength)
	return &CustomValidator{Validator: validate}
}

func ValidatePasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		if strings.Contains(Uppercase, string(char)) {
			hasUpper = true
		} else if strings.Contains(Lowercase, string(char)) {
			hasLower = true
		} else if strings.Contains(Numbers, string(char)) {
			hasNumber = true
		} else if strings.Contains(SpecialChars, string(char)) {
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial

}
