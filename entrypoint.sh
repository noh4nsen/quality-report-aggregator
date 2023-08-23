#!/bin/sh

set -eou pipefail

encoded_tflint_report=$1
encoded_tfsec_report=$2
encoded_checkov_report=$3

tflint_report=$(echo -n $encoded_tflint_report | base64 -d)
tfsec_report=$(echo -n $encoded_tfsec_report | base64 -d)
checkov_report=$(echo -n $encoded_checkov_report | base64 -d)

report=$(quality-report-aggregator "$tflint_report" "$tfsec_report" "$checkov_report")

echo "--- Executing quality-report-aggregator ---"
echo "report=$(echo $report)" >> $GITHUB_OUTPUT