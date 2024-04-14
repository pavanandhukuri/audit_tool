package entities

type Status string

const (
	Success           Status = "Success"
	Failure           Status = "Failure"
	SuccessWithErrors Status = "SuccessWithErrors"
)

type ValidationError struct {
	Field   string
	Message string
}
type ValidationResult struct {
	Status           Status
	ValidationErrors []ValidationError
}
