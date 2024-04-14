package adapters

import (
	"gopkg.in/yaml.v3"
	"os"
	"security_audit_tool/domain/entities"
	"security_audit_tool/logger"
)

type YamlBasedRule struct {
	Field     string `yaml:"field"`
	Operation string `yaml:"operation"`
	Message   string `yaml:"message"`
}

type YamlBasedRuleRepository struct {
	ruleFilePath string
}

func (r *YamlBasedRuleRepository) GetRules() ([]entities.Rule, error) {
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

	var rules []entities.Rule

	for _, yamlBasedRule := range yamlBasedRules {
		rules = append(rules, entities.Rule(yamlBasedRule))
	}

	return rules, nil
}

func NewYamlBasedRuleRepository(ruleFilePath string) *YamlBasedRuleRepository {
	return &YamlBasedRuleRepository{ruleFilePath}
}
