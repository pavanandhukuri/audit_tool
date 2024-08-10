package ports

import (
	"security_audit_tool/domain/entities"
)

type ReportGenerator interface {
	Generate(result *entities.ValidationResult) (string, error)
}
