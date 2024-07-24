package ruleEvaluator

import "github.com/go-playground/validator/v10"

func IsFalse(fl validator.FieldLevel) bool {
	return !fl.Field().Bool()
}
