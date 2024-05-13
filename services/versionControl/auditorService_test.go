package versionControl

import (
	"github.com/stretchr/testify/assert"
	"security_audit_tool/domain/entities"
	"security_audit_tool/mocks"
	"testing"
)

func TestAuditorService_Audit(t *testing.T) {
	//Arrange
	versionControlSystemMock := mocks.NewVersionControlSystemPort(t)
	ruleRepositoryMock := mocks.NewRuleRepository(t)
	reportGeneratorMock := mocks.NewReportGenerator(t)
	ruleEvaluatorMock := mocks.NewRuleEvaluatorPort(t)

	t.Run("Should return nil when audit is successful", func(t *testing.T) {
		// Mocks
		getInfoMock := versionControlSystemMock.On("GetData").Return(entities.VersionControlSystem{}, nil)
		defer getInfoMock.Unset()
		getRulesMock := ruleRepositoryMock.On("GetRules").Return([]entities.Rule{}, nil)
		defer getRulesMock.Unset()
		evaluateMock := ruleEvaluatorMock.On("EvaluateRules", []entities.Rule{}, entities.VersionControlSystem{}).Return(getMockValidationResult())
		defer evaluateMock.Unset()
		generateMock := reportGeneratorMock.On("Generate", getMockValidationResult()).Return(nil)
		defer generateMock.Unset()

		//Act
		auditorService := NewVersionControlAuditorService(versionControlSystemMock, ruleRepositoryMock, reportGeneratorMock, ruleEvaluatorMock)
		err := auditorService.Audit()

		//Assert
		assert.Nil(t, err)
		versionControlSystemMock.AssertExpectations(t)
	})

	t.Run("Should return error when GetData fails", func(t *testing.T) {
		// Mocks
		getInfoMock := versionControlSystemMock.On("GetData").Return(entities.VersionControlSystem{}, assert.AnError)
		defer getInfoMock.Unset()

		//Act
		auditorService := NewVersionControlAuditorService(versionControlSystemMock, ruleRepositoryMock, reportGeneratorMock, ruleEvaluatorMock)
		err := auditorService.Audit()

		//Assert
		assert.NotNil(t, err)
		versionControlSystemMock.AssertExpectations(t)
	})
	t.Run("Should return error when GetRules fails", func(t *testing.T) {
		// Mocks
		getInfoMock := versionControlSystemMock.On("GetData").Return(entities.VersionControlSystem{}, nil)
		defer getInfoMock.Unset()
		getRulesMock := ruleRepositoryMock.On("GetRules").Return([]entities.Rule{}, assert.AnError)
		defer getRulesMock.Unset()

		//Act
		auditorService := NewVersionControlAuditorService(versionControlSystemMock, ruleRepositoryMock, reportGeneratorMock, ruleEvaluatorMock)
		err := auditorService.Audit()

		//Assert
		assert.NotNil(t, err)
		versionControlSystemMock.AssertExpectations(t)
		ruleRepositoryMock.AssertExpectations(t)
	})

	t.Run("Should return error when Generate fails", func(t *testing.T) {
		// Mocks
		getInfoMock := versionControlSystemMock.On("GetData").Return(entities.VersionControlSystem{}, nil)
		defer getInfoMock.Unset()
		getRulesMock := ruleRepositoryMock.On("GetRules").Return([]entities.Rule{}, nil)
		defer getRulesMock.Unset()
		evaluateMock := ruleEvaluatorMock.On("EvaluateRules", []entities.Rule{}, entities.VersionControlSystem{}).Return(getMockValidationResult())
		defer evaluateMock.Unset()
		generateMock := reportGeneratorMock.On("Generate", getMockValidationResult()).Return(assert.AnError)
		defer generateMock.Unset()

		//Act
		auditorService := NewVersionControlAuditorService(versionControlSystemMock, ruleRepositoryMock, reportGeneratorMock, ruleEvaluatorMock)
		err := auditorService.Audit()

		//Assert
		assert.NotNil(t, err)
	})

}

func getMockValidationResult() *entities.ValidationResult {
	return &entities.ValidationResult{
		Status:           entities.Success,
		ValidationErrors: nil,
	}
}
