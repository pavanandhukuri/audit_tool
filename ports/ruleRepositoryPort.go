package ports

import (
	"security_audit_tool/domain/entities/core"
)

type RuleRepository interface {
	GetRules() ([]core.Rule, error)
}
