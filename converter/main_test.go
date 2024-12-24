package converter

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_LoadMetrics(t *testing.T) {
	f, err := os.CreateTemp("", "")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	content := map[string]any{
		"level":  "info",
		"source": "accounting/stats.go:528",
		"stats": map[string]any{
			"bytes":               13503398,
			"checks":              406,
			"deletedDirs":         0,
			"deletes":             11,
			"elapsedTime":         13.407117061,
			"errors":              0,
			"eta":                 0,
			"fatalError":          false,
			"renames":             0,
			"retryError":          false,
			"serverSideCopies":    0,
			"serverSideCopyBytes": 0,
			"serverSideMoveBytes": 0,
			"serverSideMoves":     0,
			"speed":               1038724.0540969963,
			"totalBytes":          13503398,
			"totalChecks":         406,
			"totalTransfers":      304,
			"transferTime":        13.121614488,
			"transfers":           304,
		},
		"time": "2024-12-24T03:13:40.064695+01:00",
	}
	dumpedContent, err := json.Marshal(content)
	require.NoError(t, err)
	_, err = f.Write(dumpedContent)
	require.NoError(t, err)

	stats, err := LoadRCloneStatsFromFile(f.Name())
	require.NoError(t, err)

	expectedStats := Stats{
		Bytes:               13503398,
		Checks:              406,
		DeletedDirs:         0,
		Deletes:             11,
		ElapsedTime:         13.407117061,
		Errors:              0,
		ETA:                 0,
		FatalError:          false,
		Renames:             0,
		RetryError:          false,
		ServerSideCopies:    0,
		ServerSideCopyBytes: 0,
		ServerSideMoveBytes: 0,
		ServerSideMoves:     0,
		Speed:               1038724.0540969963,
		TotalBytes:          13503398,
		TotalChecks:         406,
		TotalTransfers:      304,
		TransferTime:        13.121614488,
		Transfers:           304,
	}

	require.Equal(t, expectedStats, *stats)
}

func Test_IgnoreNonMetricLines(t *testing.T) {
	f, err := os.CreateTemp("", "")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	content := map[string]any{
		"level":  "info",
		"source": "accounting/stats.go:528",
		"time":   "2024-12-24T03:13:40.064695+01:00",
	}
	dumpedContent, err := json.Marshal(content)
	require.NoError(t, err)
	_, err = f.Write(dumpedContent)
	require.NoError(t, err)

	stats, err := LoadRCloneStatsFromFile(f.Name())
	require.NoError(t, err)
	require.Nil(t, stats)
}

func Test_IgnoreEmptyLines(t *testing.T) {
	f, err := os.CreateTemp("", "")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	_, err = f.Write([]byte("\n"))
	require.NoError(t, err)

	content := map[string]any{
		"level":  "info",
		"source": "accounting/stats.go:528",
		"time":   "2024-12-24T03:13:40.064695+01:00",
	}
	dumpedContent, err := json.Marshal(content)
	require.NoError(t, err)
	_, err = f.Write(dumpedContent)
	require.NoError(t, err)

	stats, err := LoadRCloneStatsFromFile(f.Name())
	require.NoError(t, err)
	require.Nil(t, stats)
}

func Test_IgnoreNonJsonLines(t *testing.T) {
	f, err := os.CreateTemp("", "")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	_, err = f.Write([]byte("Hello, I am not a json."))
	require.NoError(t, err)

	content := map[string]any{
		"level":  "info",
		"source": "accounting/stats.go:528",
		"time":   "2024-12-24T03:13:40.064695+01:00",
	}
	dumpedContent, err := json.Marshal(content)
	require.NoError(t, err)
	_, err = f.Write(dumpedContent)
	require.NoError(t, err)

	stats, err := LoadRCloneStatsFromFile(f.Name())
	require.NoError(t, err)
	require.Nil(t, stats)
}

func Test_ScanForAMetricLine(t *testing.T) {
	f, err := os.CreateTemp("", "")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	nonStatsLine := map[string]any{
		"level":  "info",
		"source": "accounting/stats.go:528",
		"time":   "2024-12-24T03:13:40.064695+01:00",
	}
	dumpedContent, err := json.Marshal(nonStatsLine)
	require.NoError(t, err)
	_, err = f.Write(dumpedContent)
	require.NoError(t, err)

	_, err = f.Write([]byte("\n"))
	require.NoError(t, err)

	statsLine := map[string]any{
		"level":  "info",
		"source": "accounting/stats.go:528",
		"stats": map[string]any{
			"bytes": 13503398,
		},
		"time": "2024-12-24T03:13:40.064695+01:00",
	}
	dumpedContent, err = json.Marshal(statsLine)
	require.NoError(t, err)
	_, err = f.Write(dumpedContent)
	require.NoError(t, err)

	stats, err := LoadRCloneStatsFromFile(f.Name())
	require.NoError(t, err)
	require.NotNil(t, stats)
	require.Equal(t, int64(13503398), stats.Bytes)
}
