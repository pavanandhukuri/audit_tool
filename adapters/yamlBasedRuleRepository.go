package adapters

import (
	"gopkg.in/yaml.v3"
	"os"
	"security_audit_tool/domain/entities/core"
	"security_audit_tool/logger"
)

type YamlBasedNestedRule struct {
	IdentifiedBy string          `yaml:"identifiedBy"`
	Rules        []YamlBasedRule `yaml:"rules"`
}

type YamlBasedRule struct {
	Field       string              `yaml:"field"`
	Operation   string              `yaml:"operation"`
	Message     string              `yaml:"message"`
	NestedRules YamlBasedNestedRule `yaml:"nestedRules"`
}

type YamlBasedRuleRepository struct {
	ruleFilePath string
}

func (r *YamlBasedRuleRepository) GetRules() ([]core.Rule, error) {
	rulesYaml, err := os.ReadFile(r.ruleFilePath)
	if err != nil {
		logger.LogError("Error reading rules file")
		return nil, err
	}

	var yamlBasedRules []YamlBasedRule
	err = yaml.Unmarshal(rulesYaml, &yamlBasedRules)
	if err != nil {
		logger.LogError("Error unmarshalling rules yaml")
		return nil, err
	}

	var rules []core.Rule

	for _, yamlBasedRule := range yamlBasedRules {
		rules = append(rules, yamlBasedRule.toRule())
	}

	return rules, nil
}

func (yamlBasedRule *YamlBasedRule) toRule() core.Rule {

	rule := core.Rule{
		Field:     yamlBasedRule.Field,
		Operation: yamlBasedRule.Operation,
		Message:   yamlBasedRule.Message,
	}

	// If there are nested rules, convert them to entities.Rule
	if yamlBasedRule.NestedRules.Rules != nil {
		var nestedRules []core.Rule

		for _, yamlBasedRule := range yamlBasedRule.NestedRules.Rules {
			nestedRules = append(nestedRules, yamlBasedRule.toRule())
		}
		rule.NestedRules = core.NestedRule{
			IdentifiedBy: yamlBasedRule.NestedRules.IdentifiedBy,
			Rules:        nestedRules,
		}
	}

	return rule
}

func NewYamlBasedRuleRepository(ruleFilePath string) *YamlBasedRuleRepository {
	return &YamlBasedRuleRepository{ruleFilePath}
}
