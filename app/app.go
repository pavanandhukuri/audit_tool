package app

import (
	"os"
	"security_audit_tool/adapters"
	"security_audit_tool/adapters/ruleEvaluator"
	"security_audit_tool/adapters/vcs/github"
	"security_audit_tool/logger"
	"security_audit_tool/services/versionControl"
)

func Run() {
	logger.LogInfo("Starting...")

	githubAdapter := github.NewGithubVcsAdapter()
	currentWorkingDirectory, _ := os.Getwd()
	ruleRepository := adapters.NewYamlBasedRuleRepository(currentWorkingDirectory + "/resources/version_control_system_rules.yml")
	reportGenerator := adapters.NewTextReportGenerator()
	evaluator := ruleEvaluator.NewValidatorBasedRuleEvaluator()
	versionControlAuditorService := versionControl.NewVersionControlAuditorService(githubAdapter, ruleRepository, reportGenerator, evaluator)
	err := versionControlAuditorService.Audit()
	if err != nil {
		logger.LogError("Error auditing version control system" + err.Error())
		return
	}

	logger.LogInfo("Finished.")
}
