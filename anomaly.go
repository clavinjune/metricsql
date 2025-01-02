package metricsql

import "strings"

var anomalyFuncs = map[string]bool{
	"avg_daily":                 true,
	"median_daily":              true,
	"avg_weekly":                true,
	"median_weekly":             true,
	"median_weekly_with_trends": true,
}

func IsAnomalyFunc(s string) bool {
	return anomalyFuncs[strings.ToLower(s)]
}
