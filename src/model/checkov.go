package model

import "fmt"

type Checkov struct {
	Project string        `json:"project"`
	Report  CheckovReport `json:"report"`
}

type CheckovReport struct {
	CheckType string        `json:"check_type"`
	Results   CheckovResult `json:"results"`
	Summary   Summary       `json:"summary"`
}

type CheckovResult struct {
	FailedChecks []FailedCheck `json:"failed_checks"`
}

type FailedCheck struct {
	CheckId     string      `json:"check_id"`
	CheckName   string      `json:"check_name"`
	CheckResult CheckResult `json:"check_result"`
	FilePath    string      `json:"file_path"`
	Resource    string      `json:"resource"`
	Guideline   string      `json:"guideline"`
}

type CheckResult struct {
	Result string `json:"result"`
}

type Summary struct {
	Passed int `json:"passed"`
	Failed int `json:"failed"`
}

func (checkov *Checkov) BuildReport() string {
	var report string

	report += fmt.Sprintf(`

		- Project: %s
		- Report:
			- CheckType: %s
			- Results:
	`, checkov.Project, checkov.Report.CheckType)

	for index, failedCheck := range checkov.Report.Results.FailedChecks {
		report += fmt.Sprintf(`
				- Check: %d
				- Id: %s
				- Name: %s
				- Result: %s
				- FilePath: %s
				- Resource: %s
				- Guideline: %s

		`, index, failedCheck.CheckId, failedCheck.CheckName, failedCheck.CheckResult.Result, failedCheck.FilePath, failedCheck.Resource, failedCheck.Guideline)
	}

	report += fmt.Sprintf(`
			- Summary:
				- Passed: %d
				- Failed: %d

	`, checkov.Report.Summary.Passed, checkov.Report.Summary.Failed)

	return report
}
