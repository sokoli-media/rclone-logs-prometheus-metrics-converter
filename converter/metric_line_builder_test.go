package converter

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_BuildMetricLine(t *testing.T) {
	context := MetricLineContext{
		RCloneCommand: "copy",
		ShareName:     "important-data",
	}

	metricLine, err := buildMetricLine(context, "bytes", 123.4)
	require.NoError(t, err)
	require.True(t, strings.Contains(metricLine, "rclone_copy_bytes{share=\"important-data\"} 123.4"))

	metricLine, err = buildMetricLine(context, "number", 123)
	require.NoError(t, err)
	require.True(t, strings.Contains(metricLine, "rclone_copy_number{share=\"important-data\"} 123"))

	metricLine, err = buildMetricLine(context, "number", int64(123))
	require.NoError(t, err)
	require.True(t, strings.Contains(metricLine, "rclone_copy_number{share=\"important-data\"} 123"))

	metricLine, err = buildMetricLine(context, "bool", true)
	require.NoError(t, err)
	require.True(t, strings.Contains(metricLine, "rclone_copy_bool{share=\"important-data\"} 1"))

	metricLine, err = buildMetricLine(context, "bool", false)
	require.NoError(t, err)
	require.True(t, strings.Contains(metricLine, "rclone_copy_bool{share=\"important-data\"} 0"))
}
