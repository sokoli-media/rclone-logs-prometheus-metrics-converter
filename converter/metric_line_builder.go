package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
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
		return "", errors.New(fmt.Sprintf("metric name %s includes not allowed characters", metricName))
	}
	if err != nil {
		return "", err
	}

	dumpedValue := dumpMetric(value)
	return fmt.Sprintf("rclone_%s_%s{share=\"%s\"} %s", context.RCloneCommand, metricName, context.ShareName, dumpedValue), nil
}
