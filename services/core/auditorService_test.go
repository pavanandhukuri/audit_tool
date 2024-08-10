package core

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"security_audit_tool/domain/entities"
	"security_audit_tool/mocks"
	"testing"
)

func TestAuditorService_Audit(t *testing.T) {

	var versionControlSystemMock *mocks.VersionControlSystemPort
	var ruleRepositoryMock *mocks.RuleRepository
	var ruleEvaluatorMock *mocks.RuleEvaluatorPort
	var reportGeneratorMock *mocks.ReportGenerator
	var auditorService *AuditorService

	setup := func() {
		versionControlSystemMock = mocks.NewVersionControlSystemPort(t)
		ruleRepositoryMock = mocks.NewRuleRepository(t)
		ruleEvaluatorMock = mocks.NewRuleEvaluatorPort(t)
		reportGeneratorMock = mocks.NewReportGenerator(t)

		auditorService = NewVersionControlAuditorService(versionControlSystemMock, ruleRepositoryMock, ruleEvaluatorMock, reportGeneratorMock)

	}
	mockVersionControlData := make(map[string]interface{})
	var mockRules []entities.Rule
	mockValidationResult := &entities.ValidationResult{}

	t.Run("Should return error if versionControlSystem.GetData() returns an error", func(t *testing.T) {
		// Arrange
		setup()
		versionControlSystemMock.On("GetData").Return(nil, errors.New("some error"))

		// Act
		err := auditorService.Audit()

		// Assert
		if err == nil {
			t.Error("Expected error, got nil")
		}

		assert.Equal(t, "some error", err.Error())

		versionControlSystemMock.AssertExpectations(t)
	})

	t.Run("Should return error if ruleRepository.GetRules() returns an error", func(t *testing.T) {
		// Arrange
		setup()
		versionControlSystemMock.On("GetData").Return(mockVersionControlData, nil)
		ruleRepositoryMock.On("GetRules").Return(nil, errors.New("some error"))

		// Act
		err := auditorService.Audit()

		// Assert
		if err == nil {
			t.Error("Expected error, got nil")
		}

		assert.Equal(t, "some error", err.Error())

		versionControlSystemMock.AssertExpectations(t)
		ruleRepositoryMock.AssertExpectations(t)
	})

	t.Run("Should return error if ruleEvaluator.Evaluate() returns an error", func(t *testing.T) {
		// Arrange
		setup()
		versionControlSystemMock.On("GetData").Return(mockVersionControlData, nil)
		ruleRepositoryMock.On("GetRules").Return(mockRules, nil)
		ruleEvaluatorMock.On("EvaluateRules", mockRules, mockVersionControlData).Return(nil, errors.New("some error"))

		// Act
		err := auditorService.Audit()

		// Assert
		if err == nil {
			t.Error("Expected error, got nil")
		}

		assert.Equal(t, "some error", err.Error())

		versionControlSystemMock.AssertExpectations(t)
		ruleRepositoryMock.AssertExpectations(t)
		ruleEvaluatorMock.AssertExpectations(t)

	})

	t.Run("Should return error if reportGenerator.Generate() returns an error", func(t *testing.T) {
		// Arrange
		setup()
		versionControlSystemMock.On("GetData").Return(mockVersionControlData, nil)
		ruleRepositoryMock.On("GetRules").Return(mockRules, nil)
		ruleEvaluatorMock.On("EvaluateRules", mockRules, mockVersionControlData).Return(mockValidationResult, nil)
		reportGeneratorMock.On("Generate", mockValidationResult).Return("", errors.New("some error"))

		// Act
		err := auditorService.Audit()

		// Assert
		if err == nil {
			t.Error("Expected error, got nil")
		}

		assert.Equal(t, "some error", err.Error())

		versionControlSystemMock.AssertExpectations(t)
		ruleRepositoryMock.AssertExpectations(t)
		ruleEvaluatorMock.AssertExpectations(t)
		reportGeneratorMock.AssertExpectations(t)
	})

	t.Run("Should return nil if all dependencies return nil", func(t *testing.T) {
		// Arrange
		setup()
		versionControlSystemMock.On("GetData").Return(mockVersionControlData, nil)
		ruleRepositoryMock.On("GetRules").Return(mockRules, nil)
		ruleEvaluatorMock.On("EvaluateRules", mockRules, mockVersionControlData).Return(mockValidationResult, nil)
		reportGeneratorMock.On("Generate", mockValidationResult).Return("fileName", nil)

		// Act
		err := auditorService.Audit()

		// Assert
		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		versionControlSystemMock.AssertExpectations(t)
		ruleRepositoryMock.AssertExpectations(t)
		ruleEvaluatorMock.AssertExpectations(t)
		reportGeneratorMock.AssertExpectations(t)
	})
}
