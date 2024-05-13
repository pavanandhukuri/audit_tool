package versionControl

import (
	"security_audit_tool/logger"
	"security_audit_tool/ports"
)

type AuditorService struct {
	versionControlSystemAdapter ports.VersionControlSystemPort
	ruleRepository              ports.RuleRepository
	reportGenerator             ports.ReportGenerator
	ruleEvaluator               ports.RuleEvaluatorPort
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

	validationResult := service.ruleEvaluator.EvaluateRules(rules, versionControlData)
	err = service.reportGenerator.Generate(validationResult)
	if err != nil {
		return err
	}

	return nil
}

func NewVersionControlAuditorService(versionControlSystemAdapter ports.VersionControlSystemPort, ruleRepository ports.RuleRepository, reportGenerator ports.ReportGenerator, ruleEvaluator ports.RuleEvaluatorPort) *AuditorService {
	return &AuditorService{versionControlSystemAdapter, ruleRepository, reportGenerator, ruleEvaluator}
}
