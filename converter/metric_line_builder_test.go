package converter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_BuildMetricLine_FloatMetric(t *testing.T) {
	context := MetricLineContext{
		RCloneCommand: "copy",
		ShareName:     "important-data",
	}

	metricLine, err := buildMetricLine(context, "bytes", 123.4)
	require.NoError(t, err)
	expectedMetricLine := MetricLine{
		Name:  "rclone_copy_bytes",
		Share: "important-data",
		Value: 123.4,
	}
	require.Equal(t, expectedMetricLine, *metricLine)
}

func Test_BuildMetricLine_IntegerMetric(t *testing.T) {
	context := MetricLineContext{
		RCloneCommand: "copy",
		ShareName:     "important-data",
	}

	metricLine, err := buildMetricLine(context, "number", 123)
	require.NoError(t, err)
	expectedMetricLine := MetricLine{
		Name:  "rclone_copy_number",
		Share: "important-data",
		Value: 123,
	}
	require.Equal(t, expectedMetricLine, *metricLine)
}

func Test_BuildMetricLine_int64Value(t *testing.T) {
	context := MetricLineContext{
		RCloneCommand: "copy",
		ShareName:     "important-data",
	}

	metricLine, err := buildMetricLine(context, "number", int64(123))
	require.NoError(t, err)
	expectedMetricLine := MetricLine{
		Name:  "rclone_copy_number",
		Share: "important-data",
		Value: 123,
	}
	require.Equal(t, expectedMetricLine, *metricLine)
}

func Test_BuildMetricLine_boolValues(t *testing.T) {
	context := MetricLineContext{
		RCloneCommand: "copy",
		ShareName:     "important-data",
	}

	metricLine, err := buildMetricLine(context, "bool", true)
	require.NoError(t, err)
	expectedMetricLine := MetricLine{
		Name:  "rclone_copy_bool",
		Share: "important-data",
		Value: float64(1),
	}
	require.Equal(t, expectedMetricLine, *metricLine)

	metricLine, err = buildMetricLine(context, "bool", false)
	require.NoError(t, err)
	expectedMetricLine = MetricLine{
		Name:  "rclone_copy_bool",
		Share: "important-data",
		Value: float64(0),
	}
	require.Equal(t, expectedMetricLine, *metricLine)
}
