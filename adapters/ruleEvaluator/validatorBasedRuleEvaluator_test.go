package ruleEvaluator

import (
	"github.com/stretchr/testify/assert"
	"security_audit_tool/domain/entities"
	"testing"
)

func TestValidatorBasedRuleEvaluator_EvaluateRules(t *testing.T) {
	t.Run("Should return success when there are no validation errors", func(t *testing.T) {
		//Arrange
		ruleEvaluator := NewValidatorBasedRuleEvaluator()
		rules := []entities.Rule{
			{
				Field:     "CanMembersCreatePublicRepositories",
				Operation: "isfalse",
				Message:   "CanMembersCreatePublicRepositories should be false",
			},
		}
		data := struct {
			CanMembersCreatePublicRepositories bool
		}{
			CanMembersCreatePublicRepositories: false,
		}

		//Act
		result := ruleEvaluator.EvaluateRules(rules, data)

		//Assert
		assert.Equal(t, entities.Success, result.Status)
	})

	t.Run("Should return success with errors when there are validation errors", func(t *testing.T) {
		//Arrange
		ruleEvaluator := NewValidatorBasedRuleEvaluator()
		rules := []entities.Rule{
			{
				Field:     "CanMembersCreatePublicRepositories",
				Operation: "isfalse",
				Message:   "CanMembersCreatePublicRepositories should be false",
			},
		}
		data := struct {
			CanMembersCreatePublicRepositories bool
		}{
			CanMembersCreatePublicRepositories: true,
		}

		//Act
		result := ruleEvaluator.EvaluateRules(rules, data)

		//Assert
		assert.Equal(t, entities.SuccessWithErrors, result.Status)
		assert.Equal(t, 1, len(result.ValidationErrors))
		assert.Equal(t, "CanMembersCreatePublicRepositories should be false", result.ValidationErrors[0].Message)
	})

	t.Run("Should return success for default validator tests", func(t *testing.T) {
		//Arrange
		ruleEvaluator := NewValidatorBasedRuleEvaluator()
		rules := []entities.Rule{
			{
				Field:     "RequiredField",
				Operation: "required",
				Message:   "CanMembersCreatePublicRepositories should be false",
			},
		}
		data := struct {
			RequiredField string
		}{
			RequiredField: "Some Value",
		}

		//Act
		result := ruleEvaluator.EvaluateRules(rules, data)

		//Assert
		assert.Equal(t, entities.Success, result.Status)
	})

	t.Run("Should return success with errors for default validator tests", func(t *testing.T) {
		//Arrange
		ruleEvaluator := NewValidatorBasedRuleEvaluator()
		rules := []entities.Rule{
			{
				Field:     "RequiredField",
				Operation: "required",
				Message:   "RequiredField is required",
			},
		}
		data := struct {
			RequiredField string
		}{
			RequiredField: "",
		}

		//Act
		result := ruleEvaluator.EvaluateRules(rules, data)

		//Assert
		assert.Equal(t, entities.SuccessWithErrors, result.Status)
		assert.Equal(t, 1, len(result.ValidationErrors))
		assert.Equal(t, "RequiredField is required", result.ValidationErrors[0].Message)
	})
}
