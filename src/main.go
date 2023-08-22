package main

import (
	"fmt"
	"os"
	"quality-report-aggregator/core"
)

func main() {
	report := core.BuildReport(os.Args)

	fmt.Println(report)
}
