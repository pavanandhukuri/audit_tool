package ruleEvaluator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"security_audit_tool/domain/entities"
	"strings"
)

type ValidatorBasedRuleEvaluator struct {
	validator *validator.Validate
}

func (r *ValidatorBasedRuleEvaluator) EvaluateRules(rules []entities.Rule, data map[string]interface{}) (*entities.ValidationResult, error) {

	validationResult := &entities.ValidationResult{}

	for _, rule := range rules {
		evaluateRule(data, rule, r.validator, validationResult, "")
	}

	if len(validationResult.ValidationErrors) > 0 {
		validationResult.Status = entities.SuccessWithErrors
	} else {
		validationResult.Status = entities.Success
	}

	return validationResult, nil
}

func evaluateRule(data map[string]interface{}, rule entities.Rule, validate *validator.Validate, validationResult *entities.ValidationResult, messagePrefix string) {

	// If there are no nested rules, validate the field
	if rule.NestedRules.Rules == nil {
		field := data[rule.Field]
		// Log field to be validated
		fmt.Println("Validating field: ", rule.Field)
		errs := validate.Var(field, rule.Operation)

		if errs != nil {
			validationResult.ValidationErrors = append(validationResult.ValidationErrors, entities.ValidationError{
				Field:        rule.Field,
				Message:      strings.TrimSpace(messagePrefix + " " + rule.Message),
				CurrentValue: field,
			})
		}
	} else {
		// If there are nested rules, validate the nested rules
		nestedRules := rule.NestedRules.Rules
		identifiedBy := rule.NestedRules.IdentifiedBy

		// Get the field and expect it to be array so cast it
		field := data[rule.Field].([]interface{})

		// Loop through the array and evaluate the nested rules
		for _, item := range field {
			currentItem := item.(map[string]interface{})
			for _, nestedRule := range nestedRules {
				evaluateRule(currentItem, nestedRule, validate, validationResult, "["+currentItem[identifiedBy].(string)+"]")
			}
		}
	}
}

func NewValidatorBasedRuleEvaluator() *ValidatorBasedRuleEvaluator {
	validatorInstance, err := GetValidator()
	if err != nil {
		panic(err)
	}
	return &ValidatorBasedRuleEvaluator{validatorInstance}
}
