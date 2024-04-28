package adapters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYamlBasedRuleRepository_GetRules(t *testing.T) {
	t.Run("Should return rules when rules file is valid", func(t *testing.T) {
		//Arrange
		ruleRepository := NewYamlBasedRuleRepository("../test/rules.yaml")

		//Act
		rules, err := ruleRepository.GetRules()

		//Assert
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}

		if len(rules) != 2 {
			t.Errorf("Expected 2 rules but got %v", len(rules))
		}

		assert.Equal(t, rules[0].Field, "Field1")
		assert.Equal(t, rules[0].Operation, "required")
		assert.Equal(t, rules[0].Message, "Message 1")

		assert.Equal(t, rules[1].Field, "Field2")
		assert.Equal(t, rules[1].Operation, "isfalse")
		assert.Equal(t, rules[1].Message, "Message 2")
	})

	t.Run("Should return error when rules file is invalid", func(t *testing.T) {
		//Arrange
		ruleRepository := NewYamlBasedRuleRepository("../test/invalidRules.yaml")

		//Act
		_, err := ruleRepository.GetRules()

		//Assert
		if err == nil {
			t.Errorf("Expected error but got nil")
		}
	})
}
