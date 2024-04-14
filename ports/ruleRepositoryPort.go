package ports

import "security_audit_tool/domain/entities"

type RuleRepository interface {
	GetRules() ([]entities.Rule, error)
}
