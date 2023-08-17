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

	var tflintReports []model.Tflint
	var checkovReports []model.Checkov

	fmt.Println("TFLINT REPORTS:")
	if err := json.Unmarshal([]byte(os.Args[1]), &tflintReports); err != nil {
		log.Fatal(err)
	}

	for _, obj := range tflintReports {
		fmt.Println(obj)
	}

	fmt.Println("CHECKOV REPORTS:")
	if err := json.Unmarshal([]byte(os.Args[3]), &checkovReports); err != nil {
		log.Fatal(err)
	}

	for _, obj := range checkovReports {
		fmt.Println(obj)
	}
}
