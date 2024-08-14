package reportGenerator

import (
	"os"
	"security_audit_tool/domain/entities"
	"testing"
)

func TestTextReportGenerator_Generate(t *testing.T) {

	t.Run("Should generate text report for Failed Status", func(t *testing.T) {
		//Arrange
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
		expectedContent := "Security Audit Tool Report\n==========================\nStatus: Failed\nValidation Errors:\n  - Field: Field1 \n    Message: Message 1\n    CurrentValue: Value 1\n\n"

		generateFileAndCheckContent(t, mockValidationResult, expectedContent)
	})

	t.Run("Should generate text report for Passed Status", func(t *testing.T) {
		//Arrange
		mockValidationResult := &entities.ValidationResult{
			Status:           "Passed",
			ValidationErrors: []entities.ValidationError{},
		}
		expectedContent := "Security Audit Tool Report\n==========================\nStatus: Passed\n"

		generateFileAndCheckContent(t, mockValidationResult, expectedContent)
	})

	t.Run("Should generate text report for Failed Status with multiple errors", func(t *testing.T) {
		//Arrange
		mockValidationResult := &entities.ValidationResult{
			Status: "Failed",
			ValidationErrors: []entities.ValidationError{
				{
					Field:        "Field1",
					Message:      "Message 1",
					CurrentValue: "Value 1",
				},
				{
					Field:        "Field2",
					Message:      "Message 2",
					CurrentValue: "Value 2",
				},
			},
		}
		expectedContent := "Security Audit Tool Report\n==========================\nStatus: Failed\nValidation Errors:\n  - Field: Field1 \n    Message: Message 1\n    CurrentValue: Value 1\n\n  - Field: Field2 \n    Message: Message 2\n    CurrentValue: Value 2\n\n"

		generateFileAndCheckContent(t, mockValidationResult, expectedContent)
	})

	t.Run("Should generate text report for Failed Status with no errors", func(t *testing.T) {
		//Arrange
		mockValidationResult := &entities.ValidationResult{
			Status:           "Failed",
			ValidationErrors: []entities.ValidationError{},
		}
		expectedContent := "Security Audit Tool Report\n==========================\nStatus: Failed\n"

		generateFileAndCheckContent(t, mockValidationResult, expectedContent)
	})

	t.Run("Should generate text report for SuccessWithErrors Status", func(t *testing.T) {
		//Arrange
		mockValidationResult := &entities.ValidationResult{
			Status: "SuccessWithErrors",
			ValidationErrors: []entities.ValidationError{
				{
					Field:        "Field1",
					Message:      "Message 1",
					CurrentValue: "Value 1",
				},
			},
		}
		expectedContent := "Security Audit Tool Report\n==========================\nStatus: SuccessWithErrors\nValidation Errors:\n  - Field: Field1 \n    Message: Message 1\n    CurrentValue: Value 1\n\n"

		generateFileAndCheckContent(t, mockValidationResult, expectedContent)
	})
}

func generateFileAndCheckContent(t *testing.T, mockValidationResult *entities.ValidationResult, expectedContent string) {
	reportGenerator := NewTextReportGenerator()

	//Act
	fileName, err := reportGenerator.Generate(mockValidationResult)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("Error while deleting file %v", name)
		}
	}(fileName)

	//Assert
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	if fileName == "" {
		t.Errorf("Expected file name but got empty string")
	}

	//	Read the file and check the content
	file, err := os.Open(fileName)
	if err != nil {
		t.Errorf("Error while opening file %v", fileName)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		t.Errorf("Error while getting file info %v", fileName)
	}

	fileSize := fileInfo.Size()
	if fileSize == 0 {
		t.Errorf("Expected file size > 0 but got %v", fileSize)
	}

	// Read the file content
	// Read the file content
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		t.Errorf("Error while reading file %v", fileName)
	}

	// Check the content

	if string(fileContent) != expectedContent {
		t.Errorf("Expected content %v but got %v", expectedContent, string(fileContent))
	}
}
