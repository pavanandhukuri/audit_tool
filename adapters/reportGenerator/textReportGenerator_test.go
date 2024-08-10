package reportGenerator

import (
	"security_audit_tool/domain/entities"
	"testing"
)

func TestTextReportGenerator_Generate(t *testing.T) {

	t.Run("Should generate text report", func(t *testing.T) {
		//Arrange
		reportGenerator := NewTextReportGenerator()
		mockValidationResult := &entities.ValidationResult{
			Status: "Failed",
			ValidationErrors: []entities.ValidationError{
				{
					Field:        "Field1",
					Message:      "Message 1",
					CurrentValue: "Value 1",
				},
			},
		}

		//Act
		fileName, err := reportGenerator.Generate(mockValidationResult)

		//Assert
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}

		if fileName == "" {
			t.Errorf("Expected file name but got empty string")
		}
	})
}
