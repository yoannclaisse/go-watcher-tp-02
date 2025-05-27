package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"tp2/internal/checker"
)

func ExportResultsToJsonFile(filepath string, results []checker.ReportEntry) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal results to JSON: %w", err)
	}
	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to write results to file: %w", err)
	}
	return nil
}
