package entities

type NestedRule struct {
	IdentifiedBy string
	Rules        []Rule
}

type Rule struct {
	Field       string
	Operation   string
	Message     string
	NestedRules NestedRule
}
