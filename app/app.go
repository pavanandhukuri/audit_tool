package app

import (
	"os"
	"security_audit_tool/adapters/reportGenerator"
	"security_audit_tool/adapters/ruleEvaluator"
	"security_audit_tool/adapters/ruleRepository"
	"security_audit_tool/adapters/vcs/github"
	"security_audit_tool/logger"
	"security_audit_tool/services/core"
)

func Run() {
	logger.LogInfo("Starting...")

	githubAdapter, ruleRepository, reportGenerator, evaluator := createAdapters()

	versionControlAuditorService := core.NewVersionControlAuditorService(githubAdapter, ruleRepository, evaluator, reportGenerator)
	err := versionControlAuditorService.Audit()
	if err != nil {
		logger.LogError("Error auditing version control system" + err.Error())
		panic(err)
	}

	logger.LogInfo("Finished.")
}

func createAdapters() (*github.VcsAdapter, *ruleRepository.YamlBasedRuleRepository, *reportGenerator.TextReportGenerator, *ruleEvaluator.ValidatorBasedRuleEvaluator) {
	currentWorkingDirectory, _ := os.Getwd()
	schemaConfigFilepath := currentWorkingDirectory + "/resources/schema_config.yml"
	vcsRulesFilePath := currentWorkingDirectory + "/resources/version_control_system_rules.yml"

	githubAdapter := github.NewGithubVcsAdapter(schemaConfigFilepath)
	ruleRepository := ruleRepository.NewYamlBasedRuleRepository(vcsRulesFilePath)
	reportGenerator := reportGenerator.NewTextReportGenerator()
	evaluator := ruleEvaluator.NewValidatorBasedRuleEvaluator()

	return githubAdapter, ruleRepository, reportGenerator, evaluator
}
