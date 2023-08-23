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
	<ul>
		<li>Project: %s</li>
		<li>Report:</li>
		<ul>
	`, tflint.Project)

	for index, issue := range tflint.Report.Issues {
		report += fmt.Sprintf(`
				<li>Issue: %d</li>
				<li>Message: %s</li>
				<li>Rule:</li>
				<ul>
					<li>Link: %s</li>
					<li>Name: %s</li>
					<li>Severity: %s</li>
				</ul>
			<br>
		`, index, issue.Message, issue.Rule.Link, issue.Rule.Name, issue.Rule.Severity)
	}

	report += fmt.Sprintf(`
		</ul>
			<li>Errors:</li>
			<ul>
	`)

	for _, error := range tflint.Report.Errors {
		report += fmt.Sprintf(`
				<li>- %s</li>
		`, error)
	}

	report += fmt.Sprintf(`
			</ul>
	</ul>
	`)

	return report
}
