package validator

import "quality-report-aggregator/helper"

func CheckInputArgs(args []string) {
	helper.CheckAndLogError("Invalid number of arguments", "Expected 3.", len(args) == 4)
}
