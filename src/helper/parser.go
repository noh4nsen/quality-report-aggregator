package helper

import (
	"encoding/json"
	"fmt"
	"quality-report-aggregator/model"
	"reflect"
)

func ToTflint(arg string, obj *[]model.Tflint) {
	err := json.Unmarshal([]byte(arg), &obj)
	CheckAndLogError("Could not parse report: ", fmt.Sprintf("object: %s, error: %s", reflect.TypeOf(obj), err), err == nil)
}

func ToTfsec(arg string, obj *[]model.Tfsec) {
	err := json.Unmarshal([]byte(arg), &obj)
	CheckAndLogError("Could not parse report: ", fmt.Sprintf("object: %s, error: %s", reflect.TypeOf(obj), err), err == nil)
}

func ToCheckov(arg string, obj *[]model.Checkov) {
	err := json.Unmarshal([]byte(arg), &obj)
	CheckAndLogError("Could not parse report: ", fmt.Sprintf("object: %s, error: %s", reflect.TypeOf(obj), err), err == nil)
}
