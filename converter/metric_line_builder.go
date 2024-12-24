package converter

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type MetricLineContext struct {
	RCloneCommand string
	ShareName     string
}

func dumpMetric(value any) []byte {
	if boolValue, ok := value.(bool); ok {
		if boolValue {
			return []byte("1")
		} else {
			return []byte("0")
		}
	}

	dumpedValue, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	return dumpedValue
}

func buildMetricLine(context MetricLineContext, metricName string, value any) (string, error) {
	matched, err := regexp.MatchString("^[a-z_]+$", metricName)
	if !matched {
		return "", fmt.Errorf("metric name %s includes not allowed characters", metricName)
	}
	if err != nil {
		return "", err
	}

	fullMetricName := fmt.Sprintf("rclone_%s_%s", context.RCloneCommand, metricName)
	dumpedValue := dumpMetric(value)

	metricLines := []string{
		fmt.Sprintf("# HELP %s Metric from rclone.", fullMetricName),
		fmt.Sprintf("# TYPE %s gauge", fullMetricName),
		fmt.Sprintf("%s{share=\"%s\"} %s", fullMetricName, context.ShareName, dumpedValue),
	}

	return strings.Join(metricLines, "\n"), nil
}
