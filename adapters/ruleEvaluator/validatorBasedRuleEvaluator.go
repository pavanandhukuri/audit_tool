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

	dataAsValue := getValueIfPointer(data)

	for _, rule := range rules {
		evaluateRule(dataAsValue, rule, r.validator, validationResult, "")
	}

	if len(validationResult.ValidationErrors) > 0 {
		validationResult.Status = entities.SuccessWithErrors
	} else {
		validationResult.Status = entities.Success
	}

	return validationResult
}

func evaluateRule(fieldValue reflect.Value, rule entities.Rule, validate *validator.Validate, validationResult *entities.ValidationResult, messagePrefix string) {

	// If there are no nested rules, validate the field
	if rule.NestedRules.Rules == nil {
		field := getFieldByFieldName(fieldValue, rule.Field)
		errs := validate.Var(field.Interface(), rule.Operation)

		if errs != nil {
			validationResult.ValidationErrors = append(validationResult.ValidationErrors, entities.ValidationError{
				Field:   rule.Field,
				Message: messagePrefix + " " + rule.Message,
			})
		}
	} else {
		// If there are nested rules, validate the nested rules
		nestedRules := rule.NestedRules.Rules
		field := getFieldByFieldName(fieldValue, rule.Field)

		for i := 0; i < field.Len(); i++ {
			subField := field.Index(i)
			identifiedBy := getFieldByFieldName(subField, rule.NestedRules.IdentifiedBy)
			evaluateRule(subField, nestedRules[0], validate, validationResult, rule.Field+"["+identifiedBy.String()+"]")
		}

	}

}

func getFieldByFieldName(fieldValue reflect.Value, fieldName string) reflect.Value {
	field := fieldValue.FieldByName(fieldName)
	return field
}

func getValueIfPointer(data interface{}) reflect.Value {
	v := reflect.ValueOf(data)
	if kind := v.Kind(); kind == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func NewValidatorBasedRuleEvaluator() *ValidatorBasedRuleEvaluator {
	validatorInstance, err := GetValidator()
	if err != nil {
		panic(err)
	}
	return &ValidatorBasedRuleEvaluator{validatorInstance}
}
