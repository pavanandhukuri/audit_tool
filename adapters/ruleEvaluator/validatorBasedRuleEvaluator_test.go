package ruleEvaluator

import (
	"security_audit_tool/domain/entities"
	"testing"
)

func TestValidatorBasedRuleEvaluator_EvaluateRules(t *testing.T) {
	t.Run("Test EvaluateRules with Success", func(t *testing.T) {
		//Arrange
		rules := []entities.Rule{
			{
				Field:     "Field1",
				Operation: "eq=true",
				Message:   "Field1 should be true",
			},
			{
				Field:     "Field2",
				Operation: "eq=abc",
				Message:   "Field2 should be abc",
			},
		}

		data := map[string]interface{}{
			"Field1": "true",
			"Field2": "abc",
		}

		validatorBasedRuleEvaluator := NewValidatorBasedRuleEvaluator()
		evaluateRules, err := validatorBasedRuleEvaluator.EvaluateRules(rules, data)
		if err != nil {
			t.Errorf("Error while evaluating rules: %v", err)
		}

		if evaluateRules.Status != entities.Success {
			t.Errorf("Expected status: Success, got: %v", evaluateRules.Status)
		}

		if len(evaluateRules.ValidationErrors) != 0 {
			t.Errorf("Expected 0 validation error, got: %v", len(evaluateRules.ValidationErrors))
		}
	})

	t.Run("Test EvaluateRules with failure", func(t *testing.T) {
		//Arrange
		rules := []entities.Rule{
			{
				Field:     "Field1",
				Operation: "eq=true",
				Message:   "Field1 should be true",
			},
			{
				Field:     "Field2",
				Operation: "eq=abc",
				Message:   "Field2 should be abc",
			},
		}

		data := map[string]interface{}{
			"Field1": "false",
			"Field2": "abc",
		}

		validatorBasedRuleEvaluator := NewValidatorBasedRuleEvaluator()
		evaluateRules, err := validatorBasedRuleEvaluator.EvaluateRules(rules, data)
		if err != nil {
			t.Errorf("Error while evaluating rules: %v", err)
		}

		if evaluateRules.Status != entities.SuccessWithErrors {
			t.Errorf("Expected status: SuccessWithErrors, got: %v", evaluateRules.Status)
		}
		if len(evaluateRules.ValidationErrors) != 1 {
			t.Errorf("Expected 1 validation error, got: %v", len(evaluateRules.ValidationErrors))
		}

		if evaluateRules.ValidationErrors[0].Field != "Field1" {
			t.Errorf("Expected Field1, got: %v", evaluateRules.ValidationErrors[0].Field)
		}

		if evaluateRules.ValidationErrors[0].Message != "Field1 should be true" {
			t.Errorf("Field1 should be true, got: %v", evaluateRules.ValidationErrors[0].Message)
		}
	})

	t.Run("Test Evaluate Rules with nested fields", func(t *testing.T) {
		//Arrange
		rules := []entities.Rule{
			{
				Field: "Field1",
				NestedRules: entities.NestedRule{
					IdentifiedBy: "Id",
					Rules: []entities.Rule{
						{
							Field:     "Field2",
							Operation: "eq=abc",
							Message:   "Field2 should be abc",
						},
					}},
			},
		}

		data := map[string]interface{}{
			"Field1": []interface{}{
				map[string]interface{}{
					"Id":     "1",
					"Field2": "abc",
				},
			},
		}

		validatorBasedRuleEvaluator := NewValidatorBasedRuleEvaluator()
		evaluateRules, err := validatorBasedRuleEvaluator.EvaluateRules(rules, data)
		if err != nil {
			t.Errorf("Error while evaluating rules: %v", err)
		}

		if evaluateRules.Status != entities.Success {
			t.Errorf("Expected status: Success, got: %v", evaluateRules.Status)
		}

		if len(evaluateRules.ValidationErrors) != 0 {
			t.Errorf("Expected 0 validation error, got: %v", len(evaluateRules.ValidationErrors))
		}
	})

	t.Run("Test Evaluate Rules with nested fields with failure", func(t *testing.T) {
		//Arrange
		rules := []entities.Rule{
			{
				Field: "Field1",
				NestedRules: entities.NestedRule{
					IdentifiedBy: "Id",
					Rules: []entities.Rule{
						{
							Field:     "Field2",
							Operation: "eq=abc",
							Message:   "Field2 should be abc",
						},
					}},
			},
		}

		data := map[string]interface{}{
			"Field1": []interface{}{
				map[string]interface{}{
					"Id":     "1",
					"Field2": "abcd",
				},
			},
		}

		validatorBasedRuleEvaluator := NewValidatorBasedRuleEvaluator()
		evaluateRules, err := validatorBasedRuleEvaluator.EvaluateRules(rules, data)
		if err != nil {
			t.Errorf("Error while evaluating rules: %v", err)
		}

		if evaluateRules.Status != entities.SuccessWithErrors {
			t.Errorf("Expected status: SuccessWithErrors, got: %v", evaluateRules.Status)
		}

		if len(evaluateRules.ValidationErrors) != 1 {
			t.Errorf("Expected 1 validation error, got: %v", len(evaluateRules.ValidationErrors))
		}

		if evaluateRules.ValidationErrors[0].Field != "Field2" {
			t.Errorf("Expected Field2, got: %v", evaluateRules.ValidationErrors[0].Field)
		}

		if evaluateRules.ValidationErrors[0].Message != "[1] Field2 should be abc" {
			t.Errorf("Field2 should be abc, got: %v", evaluateRules.ValidationErrors[0].Message)
		}
	})

	t.Run("Test Evaluate Rules with custom operation", func(t *testing.T) {
		//Arrange
		rules := []entities.Rule{
			{
				Field:     "Field1",
				Operation: "isfalse",
				Message:   "Field1 should be false",
			},
		}

		data := map[string]interface{}{
			"Field1": false,
		}

		validatorBasedRuleEvaluator := NewValidatorBasedRuleEvaluator()
		evaluateRules, err := validatorBasedRuleEvaluator.EvaluateRules(rules, data)
		if err != nil {
			t.Errorf("Error while evaluating rules: %v", err)
		}

		if evaluateRules.Status != entities.Success {
			t.Errorf("Expected status: Success, got: %v", evaluateRules.Status)
		}

		if len(evaluateRules.ValidationErrors) != 0 {
			t.Errorf("Expected 0 validation error, got: %v", len(evaluateRules.ValidationErrors))
		}
	})
}
