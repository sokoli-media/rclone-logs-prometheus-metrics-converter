package converter

import (
	"bufio"
	"encoding/json"
	"os"
)

func LoadRCloneStatsFromFile(path string) (*Stats, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		var entry LogEntry
		if err := json.Unmarshal([]byte(sc.Text()), &entry); err != nil {
			continue
		}
		if entry.Stats == nil {
			continue
		}

		return entry.Stats, nil
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

func StatsToMetrics(rCloneCommand string, shareName string, stats *Stats) []MetricLine {
	var metrics []MetricLine

	context := MetricLineContext{
		RCloneCommand: rCloneCommand,
		ShareName:     shareName,
	}

	if metricLine, err := buildMetricLine(context, "bytes", stats.Bytes); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "checks", stats.Checks); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "deleted_dirs", stats.DeletedDirs); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "deletes", stats.Deletes); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "elapsed_time", stats.ElapsedTime); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "errors", stats.Errors); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "fatal_error", stats.FatalError); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "renames", stats.Renames); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "server_side_copies", stats.ServerSideCopies); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "server_side_copy_bytes", stats.ServerSideCopyBytes); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "server_side_move_bytes", stats.ServerSideMoveBytes); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "server_side_moves", stats.ServerSideMoves); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "server_side_moves", stats.ServerSideMoves); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "total_bytes", stats.TotalBytes); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "total_checks", stats.TotalChecks); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "total_transfers", stats.TotalTransfers); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "transfer_time", stats.TransferTime); err == nil {
		metrics = append(metrics, *metricLine)
	}
	if metricLine, err := buildMetricLine(context, "transfers", stats.Transfers); err == nil {
		metrics = append(metrics, *metricLine)
	}

	return metrics
}
