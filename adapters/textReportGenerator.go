package adapters

import (
	"bufio"
	"fmt"
	"os"
	"security_audit_tool/domain/entities/core"
	"time"
)

type TextReportGenerator struct {
}

func (r *TextReportGenerator) Generate(result *core.ValidationResult) error {

	// Create a new file with current timestamp suffix in file name
	fileName := fmt.Sprintf("report-%d.txt", time.Now().Unix())

	f, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	fmt.Fprintf(w, "Security Audit Tool Report\n")
	fmt.Fprintf(w, "==========================\n")
	fmt.Fprintf(w, "Status: %s\n", result.Status)

	if len(result.ValidationErrors) > 0 {
		fmt.Fprintf(w, "Validation Errors:\n")
		for _, err := range result.ValidationErrors {
			fmt.Fprintf(w, "  - Field: %s \n    Message: %s\n    CurrentValue: %s\n\n", err.Field, err.Message, err.CurrentValue)
		}
	}

	w.Flush()

	// Print the location of the file
	fmt.Println("Report generated at: report.txt")
	return nil
}

func NewTextReportGenerator() *TextReportGenerator {
	return &TextReportGenerator{}
}
