package core

import "quality-report-aggregator/model"

func buildTflintReport(reports []model.Tflint) string {
	var report string

	report = `<h1>TFLINT</h1>`

	for _, tflintReport := range reports {
		report += tflintReport.BuildReport()
	}

	return report
}
