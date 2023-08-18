package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"quality-report-aggregator/model"
	"quality-report-aggregator/validator"
)

func main() {
	validator.CheckInputArgs(os.Args)

	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	var tflintReports []model.Tflint
	var checkovReports []model.Checkov
	var tfsecReports []model.Tfsec

	fmt.Println("TFLINT REPORTS:")
	if err := json.Unmarshal([]byte(fmt.Sprintf(`%s`, os.Args[1])), &tflintReports); err != nil {
		log.Fatal(err)
	}

	for _, obj := range tflintReports {
		fmt.Println(obj)
		fmt.Println()
	}

	fmt.Println("TFSEC REPORTS:")
	if err := json.Unmarshal([]byte(fmt.Sprintf(`%s`, os.Args[2])), &tfsecReports); err != nil {
		log.Fatal(err)
	}

	for _, obj := range tfsecReports {
		fmt.Println(obj)
		fmt.Println()
	}

	fmt.Println("CHECKOV REPORTS:")
	if err := json.Unmarshal([]byte(fmt.Sprintf(`%s`, os.Args[3])), &checkovReports); err != nil {
		log.Fatal(err)
	}

	for _, obj := range checkovReports {
		fmt.Println(obj)
		fmt.Println()
	}
}
