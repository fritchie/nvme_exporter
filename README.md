# nvme_exporter
Prometheus exporter for nvme smart-log metrics

## Building and running

### Build

```
go build .
```

A sample Dockerfile and docker-compose.yaml are provided.

### Running

Running the exporter requires the nvme-cli package to be installed on the host.

```
./nvme_exporter <flags>
```

#### Flags

| Name | Description |
|----|-------------------------------------------------|
port | Listen port number. Type: String. Default: 9998 |

### Sample Output

Golang and process metrics have been removed from the sample.

```
# HELP nvme_avail_spare Normalized percentage of remaining spare capacity available
# TYPE nvme_avail_spare gauge
nvme_avail_spare{device="/dev/nvme0n1"} 100
nvme_avail_spare{device="/dev/nvme1n1"} 100
nvme_avail_spare{device="/dev/nvme2n1"} 100
# HELP nvme_controller_busy_time Amount of time in minutes controller busy with IO commands
# TYPE nvme_controller_busy_time counter
nvme_controller_busy_time{device="/dev/nvme0n1"} 26476
nvme_controller_busy_time{device="/dev/nvme1n1"} 2344
nvme_controller_busy_time{device="/dev/nvme2n1"} 426
# HELP nvme_critical_comp_time Amount of time in minutes temperature > critical threshold
# TYPE nvme_critical_comp_time counter
nvme_critical_comp_time{device="/dev/nvme0n1"} 0
nvme_critical_comp_time{device="/dev/nvme1n1"} 0
nvme_critical_comp_time{device="/dev/nvme2n1"} 0
# HELP nvme_critical_warning Critical warnings for the state of the controller
# TYPE nvme_critical_warning gauge
nvme_critical_warning{device="/dev/nvme0n1"} 0
nvme_critical_warning{device="/dev/nvme1n1"} 0
nvme_critical_warning{device="/dev/nvme2n1"} 0
# HELP nvme_data_units_read Number of 512 byte data units host has read
# TYPE nvme_data_units_read counter
nvme_data_units_read{device="/dev/nvme0n1"} 7.24388547e+08
nvme_data_units_read{device="/dev/nvme1n1"} 2.171078e+06
nvme_data_units_read{device="/dev/nvme2n1"} 4.370719e+06
# HELP nvme_data_units_written Number of 512 byte data units the host has written
# TYPE nvme_data_units_written counter
nvme_data_units_written{device="/dev/nvme0n1"} 1.01395942e+08
nvme_data_units_written{device="/dev/nvme1n1"} 3.0735598e+07
nvme_data_units_written{device="/dev/nvme2n1"} 2.960926e+06
# HELP nvme_endurance_grp_critical_warning_summary Critical warnings for the state of endurance groups
# TYPE nvme_endurance_grp_critical_warning_summary gauge
nvme_endurance_grp_critical_warning_summary{device="/dev/nvme0n1"} 0
nvme_endurance_grp_critical_warning_summary{device="/dev/nvme1n1"} 0
nvme_endurance_grp_critical_warning_summary{device="/dev/nvme2n1"} 0
# HELP nvme_host_read_commands Number of read commands completed
# TYPE nvme_host_read_commands counter
nvme_host_read_commands{device="/dev/nvme0n1"} 5.028009993e+09
nvme_host_read_commands{device="/dev/nvme1n1"} 1.34732619e+08
nvme_host_read_commands{device="/dev/nvme2n1"} 2.78362886e+08
# HELP nvme_host_write_commands Number of write commands completed
# TYPE nvme_host_write_commands counter
nvme_host_write_commands{device="/dev/nvme0n1"} 2.517983855e+09
nvme_host_write_commands{device="/dev/nvme1n1"} 9.13277657e+08
nvme_host_write_commands{device="/dev/nvme2n1"} 2.17255509e+08
# HELP nvme_media_errors Number of unrecovered data integrity errors
# TYPE nvme_media_errors counter
nvme_media_errors{device="/dev/nvme0n1"} 0
nvme_media_errors{device="/dev/nvme1n1"} 0
nvme_media_errors{device="/dev/nvme2n1"} 0
# HELP nvme_num_err_log_entries Lifetime number of error log entries
# TYPE nvme_num_err_log_entries counter
nvme_num_err_log_entries{device="/dev/nvme0n1"} 0
nvme_num_err_log_entries{device="/dev/nvme1n1"} 94
nvme_num_err_log_entries{device="/dev/nvme2n1"} 88
# HELP nvme_percent_used Vendor specific estimate of the percentage of life used
# TYPE nvme_percent_used gauge
nvme_percent_used{device="/dev/nvme0n1"} 11
nvme_percent_used{device="/dev/nvme1n1"} 0
nvme_percent_used{device="/dev/nvme2n1"} 1
# HELP nvme_power_cycles Number of power cycles
# TYPE nvme_power_cycles counter
nvme_power_cycles{device="/dev/nvme0n1"} 66
nvme_power_cycles{device="/dev/nvme1n1"} 72
nvme_power_cycles{device="/dev/nvme2n1"} 66
# HELP nvme_power_on_hours Number of power on hours
# TYPE nvme_power_on_hours counter
nvme_power_on_hours{device="/dev/nvme0n1"} 16410
nvme_power_on_hours{device="/dev/nvme1n1"} 3825
nvme_power_on_hours{device="/dev/nvme2n1"} 16342
# HELP nvme_spare_thresh Async event completion may occur when avail spare < threshold
# TYPE nvme_spare_thresh gauge
nvme_spare_thresh{device="/dev/nvme0n1"} 10
nvme_spare_thresh{device="/dev/nvme1n1"} 10
nvme_spare_thresh{device="/dev/nvme2n1"} 5
# HELP nvme_temperature Temperature in degrees fahrenheit
# TYPE nvme_temperature gauge
nvme_temperature{device="/dev/nvme0n1"} 103.73000000000005
nvme_temperature{device="/dev/nvme1n1"} 105.53000000000004
nvme_temperature{device="/dev/nvme2n1"} 91.13000000000004
# HELP nvme_thm_temp1_trans_count Number of times controller transitioned to lower power
# TYPE nvme_thm_temp1_trans_count counter
nvme_thm_temp1_trans_count{device="/dev/nvme0n1"} 0
nvme_thm_temp1_trans_count{device="/dev/nvme1n1"} 0
nvme_thm_temp1_trans_count{device="/dev/nvme2n1"} 0
# HELP nvme_thm_temp1_trans_time Total number of seconds controller transitioned to lower power
# TYPE nvme_thm_temp1_trans_time counter
nvme_thm_temp1_trans_time{device="/dev/nvme0n1"} 0
nvme_thm_temp1_trans_time{device="/dev/nvme1n1"} 0
nvme_thm_temp1_trans_time{device="/dev/nvme2n1"} 0
# HELP nvme_thm_temp2_trans_count Number of times controller transitioned to lower power
# TYPE nvme_thm_temp2_trans_count counter
nvme_thm_temp2_trans_count{device="/dev/nvme0n1"} 0
nvme_thm_temp2_trans_count{device="/dev/nvme1n1"} 0
nvme_thm_temp2_trans_count{device="/dev/nvme2n1"} 0
# HELP nvme_thm_temp2_trans_time Total number of seconds controller transitioned to lower power
# TYPE nvme_thm_temp2_trans_time counter
nvme_thm_temp2_trans_time{device="/dev/nvme0n1"} 0
nvme_thm_temp2_trans_time{device="/dev/nvme1n1"} 0
nvme_thm_temp2_trans_time{device="/dev/nvme2n1"} 0
# HELP nvme_unsafe_shutdowns Number of unsafe shutdowns
# TYPE nvme_unsafe_shutdowns counter
nvme_unsafe_shutdowns{device="/dev/nvme0n1"} 44
nvme_unsafe_shutdowns{device="/dev/nvme1n1"} 49
nvme_unsafe_shutdowns{device="/dev/nvme2n1"} 48
# HELP nvme_warning_temp_time Amount of time in minutes temperature > warning threshold
# TYPE nvme_warning_temp_time counter
nvme_warning_temp_time{device="/dev/nvme0n1"} 0
nvme_warning_temp_time{device="/dev/nvme1n1"} 0
nvme_warning_temp_time{device="/dev/nvme2n1"} 2
```

### Dashboard

A sample Grafana dashboard is available:

[https://grafana.com/grafana/dashboards/14706](https://grafana.com/grafana/dashboards/14706)
