package validator

import "github.com/go-playground/validator/v10"

func GetValidator() (*validator.Validate, error) {
	validate := validator.New()
	err := validate.RegisterValidation("isfalse", isFalse)
	if err != nil {
		return nil, err
	}
	return validate, nil
}

func isFalse(fl validator.FieldLevel) bool {
	return !fl.Field().Bool()
}
