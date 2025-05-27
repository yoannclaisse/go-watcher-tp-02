package checker

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"tp2/internal/config"
)

type ReportEntry struct {
	Name   string
	URL    string
	Owner  string
	Status string
	ErrMsg string
}
type CheckResult struct {
	InputTarget config.InputTarget
	Status      string
	Err         error
}

func CheckUrl(target config.InputTarget) CheckResult {
	// Timeout pour pas bloquer longtime
	client := http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := client.Get(target.URL)
	if err != nil {
		return CheckResult{
			InputTarget: target,
			//avant : Err:    fmt.Errorf("failed to fecth URL: %w", err)}
			Err: &UnreachableURLError{
				URL: target.URL,
				Err: err,
			},
		}
	}

	defer resp.Body.Close()
	return CheckResult{
		InputTarget: target,
		Status:      resp.Status,
	}
}

func ConvertToReportEntry(res CheckResult) ReportEntry {
	report := ReportEntry{
		Name:   res.InputTarget.Name,
		URL:    res.InputTarget.URL,
		Owner:  res.InputTarget.Owner,
		Status: res.Status,
	}

	if res.Err != nil {
		var unreachableURL *UnreachableURLError
		if errors.As(res.Err, &unreachableURL) {
			report.Status = "unreachable"
			report.ErrMsg = fmt.Sprintf("failed to fetch URL %s: %v", unreachableURL.URL, unreachableURL.Err)
		} else {
			report.Status = "Error"
			report.ErrMsg = fmt.Sprintf("an error occurred: %v", res.Err)
		}
	}
	return report
}
