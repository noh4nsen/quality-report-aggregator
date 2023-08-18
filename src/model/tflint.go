package model

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
