package customValidator

import (
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomTranslations(v *validator.Validate, trans ut.Translator) {
	// Custom error message for unique email validation
	v.RegisterTranslation("password-strength", trans, func(ut ut.Translator) error {
		return ut.Add("password-strength", "The Password is not include one or more of the following characters : Uppercase , LowerCase , Special , Number ", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("password-strength", fe.Field())
		return t
	})
}
