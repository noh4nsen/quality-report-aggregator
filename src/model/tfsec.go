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
	
		- Project: %s
		- Report:
	`, tfsec.Project)

	for index, result := range tfsec.Report.Results {
		report += fmt.Sprintf(`
			- Issue: %d
			- Id: %s
			- Severity: %s
			- Provider: %s
			- Service: %s
			- Spec: %s
			- Impact: %s
			- Resolution: %s
			- Links:	
		`, index, result.RuleId, result.Severity, result.RuleProvider, result.RuleService, strings.Replace(result.RuleDescription, "'", "", -1), result.Impact, result.Resolution)

		for _, link := range result.Links {
			report += fmt.Sprintf(`
				- %s
			`, link)
		}

		report += fmt.Sprintf(`
			- Resource: %s
			- Location:
				- Filename: %s
				- StartLine: %d
				- EndLine: %d
		`, result.Resource, result.Location.Filename, result.Location.StartLine, result.Location.EndLine)
	}

	return report
}
