package core

import "quality-report-aggregator/model"

func buildCheckovReport(reports []model.Checkov) string {
	var report string

	report = `\n\n--- CHECKOV REPORT ---\n\n`

	for _, checkovReport := range reports {
		report += checkovReport.BuildReport()
	}

	return report
}
