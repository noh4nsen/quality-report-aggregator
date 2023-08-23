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
	<ul>
		<li>Project: %s</li>
		<li>Report:</li>
		<ul>
			<li>CheckType: %s</li>
			<li>Results:</li>
			<ul>
	`, checkov.Project, checkov.Report.CheckType)

	for index, failedCheck := range checkov.Report.Results.FailedChecks {
		report += fmt.Sprintf(`
				<li>Check: %d</li>
				<li>Id: %s</li>
				<li>Name: %s</li>
				<li>Result: %s</li>
				<li>FilePath: %s</li>
				<li>Resource: %s</li>
				<li>Guideline: %s</li>
			<br>
		`, index, failedCheck.CheckId, failedCheck.CheckName, failedCheck.CheckResult.Result, failedCheck.FilePath, failedCheck.Resource, failedCheck.Guideline)
	}

	report += fmt.Sprintf(`
			</ul>
			<li>Summary:</li>
			<ul>
				<li>Passed: %d</li>
				<li>Failed: %d</li>
			</ul>
		</ul>
	</ul>
	`, checkov.Report.Summary.Passed, checkov.Report.Summary.Failed)

	return report
}
