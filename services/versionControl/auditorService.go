package versionControl

import (
	"security_audit_tool/logger"
	"security_audit_tool/ports"
	"security_audit_tool/services/core/validator"
)

type AuditorService struct {
	versionControlSystemAdapter ports.VersionControlSystemPort
	ruleRepository              ports.RuleRepository
	reportGenerator             ports.ReportGenerator
}

func (service *AuditorService) Audit() error {
	versionControlData, err := service.versionControlSystemAdapter.GetInfo()
	if err != nil {
		logger.LogError("Error getting version control data")
		return err
	}

	rules, err := service.ruleRepository.GetRules()
	if err != nil {
		logger.LogError("Error getting rules")
		return err
	}

	validationResult := validator.EvaluateRules(rules, versionControlData)
	err = service.reportGenerator.Generate(validationResult)
	if err != nil {
		return err
	}

	return nil
}

func NewVersionControlAuditorService(versionControlSystemAdapter ports.VersionControlSystemPort, ruleRepository ports.RuleRepository, reportGenerator ports.ReportGenerator) *AuditorService {
	return &AuditorService{versionControlSystemAdapter, ruleRepository, reportGenerator}
}
