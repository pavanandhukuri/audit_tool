package ruleEvaluator

import (
	"github.com/go-playground/validator/v10"
	"security_audit_tool/domain/entities/core"
)

type ValidatorBasedRuleEvaluator struct {
	validator *validator.Validate
}

func (r *ValidatorBasedRuleEvaluator) EvaluateRules(rules []core.Rule, data map[string]interface{}) *core.ValidationResult {

	validationResult := &core.ValidationResult{}

	for _, rule := range rules {
		evaluateRule(data, rule, r.validator, validationResult, "")
	}

	if len(validationResult.ValidationErrors) > 0 {
		validationResult.Status = core.SuccessWithErrors
	} else {
		validationResult.Status = core.Success
	}

	return validationResult
}

func evaluateRule(data map[string]interface{}, rule core.Rule, validate *validator.Validate, validationResult *core.ValidationResult, messagePrefix string) {

	// If there are no nested rules, validate the field
	if rule.NestedRules.Rules == nil {
		field := data[rule.Field]
		errs := validate.Var(field, rule.Operation)

		if errs != nil {
			validationResult.ValidationErrors = append(validationResult.ValidationErrors, core.ValidationError{
				Field:        rule.Field,
				Message:      messagePrefix + " " + rule.Message,
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
