package customValidator

import (
	"github.com/emre-unlu/GinTest/internal"
	"github.com/emre-unlu/GinTest/pkg/postgresql"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

type CustomValidator struct {
	Validator  *validator.Validate
	Translator ut.Translator
	UserRepo   *postgresql.PGUserRepository
}

// NewValidator initializes the custom customValidator with the repository and translations
func NewValidator(userRepo *postgresql.PGUserRepository) *CustomValidator {
	v := validator.New()

	// Set up the English translator
	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")

	// Register the English translations
	en_translations.RegisterDefaultTranslations(v, trans)
	RegisterCustomTranslations(v, trans)

	cv := &CustomValidator{Validator: v, Translator: trans, UserRepo: userRepo}

	// Register custom validation function
	v.RegisterValidation("password-strength", ValidatePasswordStrength)

	return cv
}

func ValidatePasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		if strings.Contains(internal.Uppercase, string(char)) {
			hasUpper = true
		} else if strings.Contains(internal.Lowercase, string(char)) {
			hasLower = true
		} else if strings.Contains(internal.Numbers, string(char)) {
			hasNumber = true
		} else if strings.Contains(internal.SpecialChars, string(char)) {
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial

}
