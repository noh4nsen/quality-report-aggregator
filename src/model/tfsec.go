package model

import (
	"fmt"
	"strings"
)

type Tfsec struct {
	Project string      `json:"project"`
	Report  TfsecReport `json:"report"`
}

type TfsecReport struct {
	Results []TfsecResult `json:"results"`
}

type TfsecResult struct {
	RuleId          string   `json:"rule_id"`
	RuleDescription string   `json:"rule_description"`
	RuleProvider    string   `json:"rule_provider"`
	RuleService     string   `json:"rule_service"`
	Impact          string   `json:"impact"`
	Resolution      string   `json:"resolution"`
	Links           []string `json:"links"`
	Severity        string   `json:"severity"`
	Resource        string   `json:"resource"`
	Location        Location `json:"location"`
}

type Location struct {
	Filename  string `json:"filename"`
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
}

func (tfsec *Tfsec) BuildReport() string {
	var report string

	report += fmt.Sprintf(`
	<ul>
		<li>Project: %s</li>
		<li>Report:</li>
		<ul>
	`, tfsec.Project)

	for index, result := range tfsec.Report.Results {
		report += fmt.Sprintf(`
			<li>Issue: %d</li>
			<li>Id: %s</li>
			<li>Severity: %s</li>
			<li>Provider: %s</li>
			<li>Service: %s</li>
			<li>Spec: %s</li>
			<li>Impact: %s</li>
			<li>Resolution: %s</li>
			<li>Links:</li>
			<ul>
		`, index, result.RuleId, result.Severity, result.RuleProvider, result.RuleService, strings.Replace(result.RuleDescription, "'", "", -1), result.Impact, result.Resolution)

		for _, link := range result.Links {
			report += fmt.Sprintf(`
				<li>%s</li>
			`, link)
		}

		report += fmt.Sprintf(`
			</ul>
			<li>Resource: %s</li>
			<li>Location:</li>
			<ul>
				<li>Filename: %s</li>
				<li>StartLine: %d</li>
				<li>EndLine: %d</li>
			</ul>
		<br>
		`, result.Resource, result.Location.Filename, result.Location.StartLine, result.Location.EndLine)
	}

	report += fmt.Sprintf(`
		</ul>
	</ul>
	`)

	return report
}
