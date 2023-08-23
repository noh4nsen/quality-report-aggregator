package model

import "fmt"

type Tflint struct {
	Project string       `json:"project"`
	Report  TflintReport `json:"report"`
}

type TflintReport struct {
	Errors []string `json:"errors"`
	Issues []Issue  `json:"issues"`
}

type Issue struct {
	Message string `json:"message"`
	Rule    Rule   `json:"rule"`
}

type Rule struct {
	Link     string `json:"link"`
	Name     string `json:"name"`
	Severity string `json:"severity"`
}

func (tflint *Tflint) BuildReport() string {
	var report string

	report += fmt.Sprintf(`

		- Project: %s
		- Report:
	`, tflint.Project)

	for index, issue := range tflint.Report.Issues {
		report += fmt.Sprintf(`
				- Issue: %d
				- Message: %s
				- Rule:
					- Link: %s
					- Name: %s
					- Severity: %s

		`, index, issue.Message, issue.Rule.Link, issue.Rule.Name, issue.Rule.Severity)
	}

	report += fmt.Sprintf(`
			- Errors:
	`)

	for _, error := range tflint.Report.Errors {
		report += fmt.Sprintf(`
				- %s
		`, error)
	}

	return report
}
