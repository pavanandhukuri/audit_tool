package ports

import (
	"security_audit_tool/domain/entities/core"
)

type RuleEvaluatorPort interface {
	EvaluateRules(rules []core.Rule, data map[string]interface{}) *core.ValidationResult
}
