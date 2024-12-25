# rclone-logs-prometheus-metrics-converter

Basic project to enable exporting stats from rclone logs to prometheus.

### "Installation"

For linux/amd64:

```shell
wget -O rclone-logs-prometheus-metrics-converter https://github.com/sokoli-media/rclone-logs-prometheus-metrics-converter/releases/download/nightly/rclone-logs-prometheus-metrics-converter-linux-amd64
chmod +x rclone-logs-prometheus-metrics-converter
```

Remember to move it to some place where you store your executables and adjust all the examples below to use your
`rclone-logs-prometheus-metrics-converter` path.

### Example usage

```shell
export LOG_PATH=/mnt/logs/very-important-data-$(date +"%Y-%m-%d-%H-%M-%S").log
rclone copy /mnt/very-important-data/ google-cloud:very-important-backup --log-level INFO --use-json-log --log-file $LOG_PATH || true

rclone-logs-prometheus-metrics-converter -logs-path $LOG_PATH -rclone-command copy -share-name very-important-data >> /mnt/metrics/very-important-data.prom
```

`|| true` is added here in order to process logs even if copying data failed

### Telegraf config to expose metrics to prometheus

```
[[inputs.tail]]
  files = ["/mnt/metrics/*.prom"]
  data_format = "json"
  json_name_key = "name"
  tag_keys = ["share"]
  from_beginning = false
  watch_method = "inotify"
  pipe = false

[[outputs.prometheus_client]]
  listen = ":9273"
  metric_version = 2
```

### But it can be done in telegraf config / sth else!

I guess, but I'm not a telegraf ninja, so this way was much easier for me.

### Example metrics

```
rclone_copy_bytes_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 1.3503398e+07
rclone_copy_checks_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 406
rclone_copy_deleted_dirs_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 0
rclone_copy_deletes_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 11
rclone_copy_elapsed_time_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 13.407117061
rclone_copy_errors_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 0
rclone_copy_fatal_error_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 0
rclone_copy_renames_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 0
rclone_copy_server_side_copies_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 0
rclone_copy_server_side_copy_bytes_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 0
rclone_copy_server_side_move_bytes_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 0
rclone_copy_server_side_moves_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 0
rclone_copy_total_bytes_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 1.3503398e+07
rclone_copy_total_checks_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 406
rclone_copy_total_transfers_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 304
rclone_copy_transfer_time_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 13.121614488
rclone_copy_transfers_value{host="my-unraid-hostname",path="/mnt/metrics/very-important-data.prom",share="very-important-data"} 304
```
