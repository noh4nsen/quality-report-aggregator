package core

import (
	"quality-report-aggregator/helper"
	"quality-report-aggregator/model"
	"quality-report-aggregator/validator"
)

func BuildReport(args []string) string {
	var report string
	var tflintReports []model.Tflint
	var tfsecReports []model.Tfsec
	var checkovReports []model.Checkov

	validator.CheckInputArgs(args)

	helper.ToTflint(args[1], &tflintReports)
	helper.ToTfsec(args[2], &tfsecReports)
	helper.ToCheckov(args[3], &checkovReports)

	report += buildTflintReport(tflintReports)
	report += buildTfsecReport(tfsecReports)
	report += buildCheckovReport(checkovReports)

	return report
}
