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
	\n
	\t	- Project: %s \n
	\t 	- Report: \n
	`, tfsec.Project)

	for index, result := range tfsec.Report.Results {
		report += fmt.Sprintf(`
		\t\t		- Issue: %d \n
		\t\t\t		- Id: \t %s \n
		\t\t\t		- Severity: \t %s \n
		\t\t\t		- Provider: \t %s \n
		\t\t\t 		- Service: \t %s \n
		\t\t\t		- Spec: \t %s \n
		\t\t\t 		- Impact: \t %s \n
		\t\t\t 		- Resolution: \t %s \n
		\t\t\t 		- Links: \n	
		`, index, result.RuleId, result.Severity, result.RuleProvider, result.RuleService, strings.Replace(result.RuleDescription, "'", "", -1), result.Impact, result.Resolution)

		for _, link := range result.Links {
			report += fmt.Sprintf(`
			\t\t\t\t	- %s \n
			`, link)
		}

		report += fmt.Sprintf(`
		\t\t\t 		- Resource: \t %s \n
		\t\t\t		- Location: \n
		\t\t\t\t		- Filename: \t %s \n
		\t\t\t\t		- StartLine: \t %d \n
		\t\t\t\t		- EndLine: \t %d \n
		`, result.Resource, result.Location.Filename, result.Location.StartLine, result.Location.EndLine)
	}

	return report
}
