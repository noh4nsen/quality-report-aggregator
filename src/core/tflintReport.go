package core

import "quality-report-aggregator/model"

func buildTflintReport(reports []model.Tflint) string {
	var report string

	report = `\n\n--- TFLINT REPORT ---\n\n`

	for _, tflintReport := range reports {
		report += tflintReport.BuildReport()
	}

	return report
}
