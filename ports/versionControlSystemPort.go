package ports

type VersionControlSystemPort interface {
	GetData() (map[string]interface{}, error)
}
