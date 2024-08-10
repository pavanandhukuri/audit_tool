package reportGenerator

import (
	"bufio"
	"fmt"
	"os"
	"security_audit_tool/domain/entities"
	"time"
)

type TextReportGenerator struct {
}

func (r *TextReportGenerator) Generate(result *entities.ValidationResult) (string, error) {

	// Create a new file with current timestamp suffix in file name
	fileName := fmt.Sprintf("report-%d.txt", time.Now().Unix())

	f, err := os.Create(fileName)

	if err != nil {
		return "", err
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = fmt.Fprintf(w, "Security Audit Tool Report\n")
	if err != nil {
		return "", err
	}
	_, err = fmt.Fprintf(w, "==========================\n")
	if err != nil {
		return "", err
	}
	_, err = fmt.Fprintf(w, "Status: %s\n", result.Status)
	if err != nil {
		return "", err
	}

	if len(result.ValidationErrors) > 0 {
		_, err := fmt.Fprintf(w, "Validation Errors:\n")
		if err != nil {
			return "", err
		}
		for _, err := range result.ValidationErrors {
			_, err := fmt.Fprintf(w, "  - Field: %s \n    Message: %s\n    CurrentValue: %s\n\n", err.Field, err.Message, err.CurrentValue)
			if err != nil {
				return "", err
			}
		}
	}

	err = w.Flush()
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func NewTextReportGenerator() *TextReportGenerator {
	return &TextReportGenerator{}
}
