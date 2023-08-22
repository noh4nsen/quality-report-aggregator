package core

import "quality-report-aggregator/model"

func buildTfsecReport(reports []model.Tfsec) string {
	var report string

	report = `\n\n--- TFSEC REPORT ---\n\n`

	for _, tfsecReport := range reports {
		report += tfsecReport.BuildReport()
	}

	return report
}
