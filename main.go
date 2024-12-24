package main

import (
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

	stats, err := converter.LoadRCloneStatsFromFile(*path)
	if err != nil {
		fmt.Printf("couldn't load stats: %s\n", err)
		syscall.Exit(1)
	}

	for _, metric := range converter.StatsToMetrics(*rCloneCommand, *shareName, stats) {
		fmt.Println(metric)
	}
}
