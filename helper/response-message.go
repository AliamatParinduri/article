package helper

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type ServiceError struct {
	Message string `json:"message"`
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func CustomeValidateError(payload any) string {

	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	err := validate.Struct(payload)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			return e.Translate(trans)
		}
	}
	return ""
}
