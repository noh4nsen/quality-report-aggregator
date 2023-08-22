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
	\n
	\t	- Project: %s \n
	\t	- Report: \n
	`, tflint.Project)

	for index, issue := range tflint.Report.Issues {
		report += fmt.Sprintf(`
		\t\t		- Issue: \t %d \n
		\t\t		- Message: \t %s \n
		\t\t		- Rule: \n
		\t\t\t			- Link: \t %s \n
		\t\t\t			- Name: \t %s \n
		\t\t\t			- Severity: \t %s \n
		\n
		`, index, issue.Message, issue.Rule.Link, issue.Rule.Name, issue.Rule.Severity)
	}

	report += fmt.Sprintf(`
	\t\t		- Errors: \n
	`)

	for _, error := range tflint.Report.Errors {
		report += fmt.Sprintf(`
				- %s \n
		`, error)
	}

	return report
}
