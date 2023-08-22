package validator

import (
	"log"
	"quality-report-aggregator/helper"
)

func CheckInputArgs(args []string) {
	log.Println(len(args))
	helper.CheckAndLogError("Invalid number of arguments", "Expected 3.", len(args) == 4)
}
