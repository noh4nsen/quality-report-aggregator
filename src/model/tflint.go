package model

type Tflint struct {
	Project string
	Report  TflintReport
}

type TflintReport struct {
	Errors []string
	Issues []Issue
}

type Issue struct {
	Message string
	Rule    Rule
}

type Rule struct {
	Link     string
	Name     string
	Severity string
}
