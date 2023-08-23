package core

import "quality-report-aggregator/model"

func buildTfsecReport(reports []model.Tfsec) string {
	var report string

	report = `--- TFSEC REPORT ---`

	for _, tfsecReport := range reports {
		report += tfsecReport.BuildReport()
	}

	return report
}
