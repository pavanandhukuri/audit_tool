package ruleEvaluator

import (
	"github.com/go-playground/validator/v10"
)

func GetValidator() (*validator.Validate, error) {
	validate := validator.New()
	err := validate.RegisterValidation("isfalse", IsFalse)
	if err != nil {
		return nil, err
	}
	return validate, nil
}
