package ports

import "security_audit_tool/domain/entities"

type VersionControlSystemPort interface {
	GetData() (entities.VersionControlSystem, error)
}
