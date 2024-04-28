package app

import (
	"os"
	"security_audit_tool/adapters"
	"security_audit_tool/adapters/ruleEvaluator"
	"security_audit_tool/logger"
	"security_audit_tool/services/config"
	"security_audit_tool/services/versionControl"
)

func Run() {
	logger.LogInfo("Starting...")

	configuration, _ := config.Load()

	githubAdapter := adapters.NewGithubVersionControlSystemAdapter(*configuration)
	currentWorkingDirectory, _ := os.Getwd()
	ruleRepository := adapters.NewYamlBasedRuleRepository(currentWorkingDirectory + "/resources/version_control_system_rules.yml")
	reportGenerator := adapters.NewTextReportGenerator()
	ruleEvaluator := ruleEvaluator.NewValidatorBasedRuleEvaluator()
	versionControlAuditorService := versionControl.NewVersionControlAuditorService(githubAdapter, ruleRepository, reportGenerator, ruleEvaluator)
	err := versionControlAuditorService.Audit()
	if err != nil {
		logger.LogError("Error auditing version control system" + err.Error())
		return
	}

	logger.LogInfo("Finished.")
}
