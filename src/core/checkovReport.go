package core

import "quality-report-aggregator/model"

func buildCheckovReport(reports []model.Checkov) string {
	var report string

	report = `--- CHECKOV REPORT ---`

	for _, checkovReport := range reports {
		report += checkovReport.BuildReport()
	}

	return report
}
