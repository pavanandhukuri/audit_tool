package ports

import (
	"security_audit_tool/domain/entities"
)

type RuleEvaluatorPort interface {
	EvaluateRules(rules []entities.Rule, data map[string]interface{}) (*entities.ValidationResult, error)
}
