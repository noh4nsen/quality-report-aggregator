package model

type Checkov struct {
	Project string        `json:"project"`
	Report  CheckovReport `json:"report"`
}

type CheckovReport struct {
	CheckType string  `json:"check_type"`
	Results   Result  `json:"results"`
	Summary   Summary `json:"summary"`
}

type Result struct {
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
