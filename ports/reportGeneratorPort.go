package ports

import (
	"security_audit_tool/domain/entities/core"
)

type ReportGenerator interface {
	Generate(result *core.ValidationResult) error
}
