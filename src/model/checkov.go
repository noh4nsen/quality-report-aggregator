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
	\n
	\t	- Project: %s \n
	\t	- Report: \n
	\t\t		- CheckType: %s \n
	\t\t		- Results: \n
	`, checkov.Project, checkov.Report.CheckType)

	for index, failedCheck := range checkov.Report.Results.FailedChecks {
		report += fmt.Sprintf(`
		\t\t\t		- Check: \t %d \n
		\t\t\t		- Id: \t %s \n
		\t\t\t		- Name: \t %s \n
		\t\t\t		- Result: \t %s \n
		\t\t\t		- FilePath: \t %s \n
		\t\t\t		- Resource: \t %s \n
		\t\t\t		- Guideline: \t %s \n
		\n
		`, index, failedCheck.CheckId, failedCheck.CheckName, failedCheck.CheckResult.Result, failedCheck.FilePath, failedCheck.Resource, failedCheck.Guideline)
	}

	report += fmt.Sprintf(`
	\t\t		- Summary: \n
	\t\t\t			- Passed: \t %d \n
	\t\t\t			- Failed: \t %d \n
	\n
	`, checkov.Report.Summary.Passed, checkov.Report.Summary.Failed)

	return report
}
