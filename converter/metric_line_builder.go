package converter

import (
	"fmt"
	"reflect"
	"regexp"
)

type MetricLine struct {
	Name  string  `json:"name"`
	Share string  `json:"share"`
	Value float64 `json:"value"`
}

type MetricLineContext struct {
	RCloneCommand string
	ShareName     string
}

func dumpMetric(value any) float64 {
	switch v := value.(type) {
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	case bool:
		if v {
			return 1.0
		}
		return 0.0
	default:
		panic(fmt.Sprintf("unknown type: %s", reflect.TypeOf(value)))
	}
}

func buildMetricLine(context MetricLineContext, metricName string, value any) (*MetricLine, error) {
	matched, err := regexp.MatchString("^[a-z_]+$", metricName)
	if !matched {
		return nil, fmt.Errorf("metric name %s includes not allowed characters", metricName)
	}
	if err != nil {
		return nil, err
	}

	fullMetricName := fmt.Sprintf("rclone_%s_%s", context.RCloneCommand, metricName)
	dumpedValue := dumpMetric(value)

	return &MetricLine{
		Name:  fullMetricName,
		Share: context.ShareName,
		Value: dumpedValue,
	}, nil
}
