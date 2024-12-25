package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"rclone-logs-prometheus-metrics-converter/converter"
	"syscall"
)

func main() {
	rCloneCommand := flag.String("rclone-command", "", "command used in rclone")
	shareName := flag.String("share-name", "", "share that is synchronized")
	path := flag.String("logs-path", "", "path to the file with rclone logs")

	flag.Parse()

	if *rCloneCommand == "" || *shareName == "" || *path == "" {
		flag.PrintDefaults()
		return
	}

	stats, err := converter.LoadRCloneStatsFromFile(*path)
	if err != nil {
		fmt.Printf("couldn't load stats: %s\n", err)
		syscall.Exit(1)
	}

	for _, metric := range converter.StatsToMetrics(*rCloneCommand, *shareName, stats) {
		dumpedMetric, err := json.Marshal(metric)
		if err == nil {
			fmt.Println(string(dumpedMetric))
		}
	}
}
