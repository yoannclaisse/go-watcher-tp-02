package checker

import (
	"net/http"
	"time"
)

type CheckResult struct {
	Target string
	Status string
	Err    error
}

func CheckUrl(url string) CheckResult {
	// Timeout pour pas bloquer longtime
	client := http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := client.Get(url)
	if err != nil {
		return CheckResult{
			Target: url,
			//avant : Err:    fmt.Errorf("failed to fecth URL: %w", err)}
			Err: &UnreachableURLError{
				URL: url,
				Err: err,
			},
		}
	}

	defer resp.Body.Close()
	return CheckResult{
		Target: url,
		Status: resp.Status,
	}
}
