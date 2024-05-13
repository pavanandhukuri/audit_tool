package entities

type Branch struct {
	Name string
}
type User struct {
	Email                            string
	IsTwoFactorAuthenticationEnabled bool
}
type Repository struct {
	Name                                     string
	DefaultWorkflowPermission                string
	IsSecretScanningAndPushProtectionEnabled bool
	IsPrivateVulnerabilityReportingEnabled   bool
	IsPrivate                                bool
	InteractionLimit                         string
	Branches                                 []Branch
}
type VersionControlSystem struct {
	CanMembersCreatePublicRepositories bool
	Repositories                       []Repository
}
