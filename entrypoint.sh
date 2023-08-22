#!/bin/sh

set -eou pipefail

tflint_report=$1
tfsec_report=$2
checkov_report=$3
report=$(quality-report-aggregator "$tflint_report" "$tfsec_report" "$checkov_report")

echo "--- Executing quality-report-aggregator ---"
echo "report='$(echo $report)'" >> $GITHUB_OUTPUT