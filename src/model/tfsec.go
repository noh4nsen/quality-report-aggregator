package model

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
