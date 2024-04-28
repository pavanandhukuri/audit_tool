package ruleEvaluator

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"security_audit_tool/domain/entities"
)

type ValidatorBasedRuleEvaluator struct {
	validator *validator.Validate
}

func (r *ValidatorBasedRuleEvaluator) EvaluateRules(rules []entities.Rule, data interface{}) *entities.ValidationResult {

	validationResult := &entities.ValidationResult{}

	for _, rule := range rules {
		evaluateRule(data, rule, r.validator, validationResult)
	}

	if len(validationResult.ValidationErrors) > 0 {
		validationResult.Status = entities.SuccessWithErrors
	} else {
		validationResult.Status = entities.Success
	}

	return validationResult
}

func evaluateRule(data interface{}, rule entities.Rule, validate *validator.Validate, validationResult *entities.ValidationResult) {
	field := getField(data, rule)

	errs := validate.Var(field.Interface(), rule.Operation)

	if errs != nil {
		validationResult.ValidationErrors = append(validationResult.ValidationErrors, entities.ValidationError{
			Field:   rule.Field,
			Message: rule.Message,
		})
	}
}

func getField(data interface{}, rule entities.Rule) reflect.Value {
	v := reflect.ValueOf(data)
	if kind := v.Kind(); kind == reflect.Ptr {
		v = v.Elem()
	}
	field := v.FieldByName(rule.Field)
	return field
}

func NewValidatorBasedRuleEvaluator() *ValidatorBasedRuleEvaluator {
	validatorInstance, err := GetValidator()
	if err != nil {
		panic(err)
	}
	return &ValidatorBasedRuleEvaluator{validatorInstance}
}