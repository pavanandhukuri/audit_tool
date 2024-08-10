package core

import (
	"security_audit_tool/logger"
	"security_audit_tool/ports"
)

type AuditorService struct {
	versionControlSystemAdapter ports.VersionControlSystemPort
	ruleRepository              ports.RuleRepository
	ruleEvaluator               ports.RuleEvaluatorPort
	reportGenerator             ports.ReportGenerator
}

func (service *AuditorService) Audit() error {
	versionControlData, err := service.versionControlSystemAdapter.GetData()
	if err != nil {
		logger.LogError("Error getting version control data")
		return err
	}

	rules, err := service.ruleRepository.GetRules()
	if err != nil {
		logger.LogError("Error getting rules")
		return err
	}

	validationResult, err := service.ruleEvaluator.EvaluateRules(rules, versionControlData)
	if err != nil {
		logger.LogError("Unable to evaluate rules")
		return err
	}

	fileName, err := service.reportGenerator.Generate(validationResult)
	if err != nil {
		logger.LogError("Unable to generate report")
		return err
	}

	logger.LogInfo("Report generated successfully: " + fileName)

	return nil
}

func NewVersionControlAuditorService(versionControlSystemAdapter ports.VersionControlSystemPort, ruleRepository ports.RuleRepository, ruleEvaluator ports.RuleEvaluatorPort, reportGenerator ports.ReportGenerator) *AuditorService {
	return &AuditorService{versionControlSystemAdapter, ruleRepository, ruleEvaluator, reportGenerator}
}
