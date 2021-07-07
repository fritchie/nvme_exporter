package main

// Export nvme smart-log metrics in prometheus format

import (
	"flag"
	"log"
	"net/http"
	"os/exec"
	"os/user"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tidwall/gjson"
)

var labels = []string{"device"}

type nvmeCollector struct {
	nvmeCriticalWarning *prometheus.Desc
	nvmeTemperature *prometheus.Desc
	nvmeAvailSpare *prometheus.Desc
	nvmeSpareThresh *prometheus.Desc
	nvmePercentUsed *prometheus.Desc
	nvmeEnduranceGrpCriticalWarningSummary *prometheus.Desc
	nvmeDataUnitsRead *prometheus.Desc
	nvmeDataUnitsWritten *prometheus.Desc
	nvmeHostReadCommands *prometheus.Desc
	nvmeHostWriteCommands *prometheus.Desc
	nvmeControllerBusyTime *prometheus.Desc
	nvmePowerCycles *prometheus.Desc
	nvmePowerOnHours *prometheus.Desc
	nvmeUnsafeShutdowns *prometheus.Desc
	nvmeMediaErrors *prometheus.Desc
	nvmeNumErrLogEntries *prometheus.Desc
	nvmeWarningTempTime *prometheus.Desc
	nvmeCriticalCompTime *prometheus.Desc
	nvmeThmTemp1TransCount *prometheus.Desc
	nvmeThmTemp2TransCount *prometheus.Desc
	nvmeThmTemp1TotalTime *prometheus.Desc
	nvmeThmTemp2TotalTime *prometheus.Desc
}

// nvme smart-log field descriptions can be found on page 180 of:
// https://nvmexpress.org/wp-content/uploads/NVM-Express-Base-Specification-2_0-2021.06.02-Ratified-5.pdf

func newNvmeCollector() prometheus.Collector {
	return &nvmeCollector{
		nvmeCriticalWarning: prometheus.NewDesc(
			"nvme_critical_warning",
			"Critical warnings for the state of the controller",
			labels,
			nil,
		),
		nvmeTemperature: prometheus.NewDesc(
			"nvme_temperature",
			"Temperature in degrees fahrenheit",
			labels,
			nil,
		),
		nvmeAvailSpare: prometheus.NewDesc(
			"nvme_avail_spare",
			"Normalized percentage of remaining spare capacity available",
			labels,
			nil,
		),
		nvmeSpareThresh: prometheus.NewDesc(
			"nvme_spare_thresh",
			"Async event completion may occur when avail spare < threshold",
			labels,
			nil,
		),
		nvmePercentUsed: prometheus.NewDesc(
			"nvme_percent_used",
			"Vendor specific estimate of the percentage of life used",
			labels,
			nil,
		),
		nvmeEnduranceGrpCriticalWarningSummary: prometheus.NewDesc(
			"nvme_endurance_grp_critical_warning_summary",
			"Critical warnings for the state of endurance groups",
			labels,
			nil,
		),
		nvmeDataUnitsRead: prometheus.NewDesc(
			"nvme_data_units_read",
			"Number of 512 byte data units host has read",
			labels,
			nil,
		),
		nvmeDataUnitsWritten: prometheus.NewDesc(
			"nvme_data_units_written",
			"Number of 512 byte data units the host has written",
			labels,
			nil,
		),
		nvmeHostReadCommands: prometheus.NewDesc(
			"nvme_host_read_commands",
			"Number of read commands completed",
			labels,
			nil,
		),
		nvmeHostWriteCommands: prometheus.NewDesc(
			"nvme_host_write_commands",
			"Number of write commands completed",
			labels,
			nil,
		),
		nvmeControllerBusyTime: prometheus.NewDesc(
			"nvme_controller_busy_time",
			"Amount of time in minutes controller busy with IO commands",
			labels,
			nil,
		),
		nvmePowerCycles: prometheus.NewDesc(
			"nvme_power_cycles",
			"Number of power cycles",
			labels,
			nil,
		),
		nvmePowerOnHours: prometheus.NewDesc(
			"nvme_power_on_hours",
			"Number of power on hours",
			labels,
			nil,
		),
		nvmeUnsafeShutdowns: prometheus.NewDesc(
			"nvme_unsafe_shutdowns",
			"Number of unsafe shutdowns",
			labels,
			nil,
		),
		nvmeMediaErrors: prometheus.NewDesc(
			"nvme_media_errors",
			"Number of unrecovered data integrity errors",
			labels,
			nil,
		),
		nvmeNumErrLogEntries: prometheus.NewDesc(
			"nvme_num_err_log_entries",
			"Lifetime number of error log entries",
			labels,
			nil,
		),
		nvmeWarningTempTime: prometheus.NewDesc(
			"nvme_warning_temp_time",
			"Amount of time in minutes temperature > warning threshold",
			labels,
			nil,
		),
		nvmeCriticalCompTime: prometheus.NewDesc(
			"nvme_critical_comp_time",
			"Amount of time in minutes temperature > critical threshold",
			labels,
			nil,
		),
		nvmeThmTemp1TransCount: prometheus.NewDesc(
			"nvme_thm_temp1_trans_count",
			"Number of times controller transitioned to lower power",
			labels,
			nil,
		),
		nvmeThmTemp2TransCount: prometheus.NewDesc(
			"nvme_thm_temp2_trans_count",
			"Number of times controller transitioned to lower power",
			labels,
			nil,
		),
		nvmeThmTemp1TotalTime: prometheus.NewDesc(
			"nvme_thm_temp1_trans_time",
			"Total number of seconds controller transitioned to lower power",
			labels,
			nil,
		),
		nvmeThmTemp2TotalTime: prometheus.NewDesc(
			"nvme_thm_temp2_trans_time",
			"Total number of seconds controller transitioned to lower power",
			labels,
			nil,
		),
	}
}

func (c *nvmeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.nvmeCriticalWarning
	ch <- c.nvmeTemperature
	ch <- c.nvmeAvailSpare
	ch <- c.nvmeSpareThresh
	ch <- c.nvmePercentUsed
	ch <- c.nvmeEnduranceGrpCriticalWarningSummary
	ch <- c.nvmeDataUnitsRead
	ch <- c.nvmeDataUnitsWritten
	ch <- c.nvmeHostReadCommands
	ch <- c.nvmeHostWriteCommands
	ch <- c.nvmeControllerBusyTime
	ch <- c.nvmePowerCycles
	ch <- c.nvmePowerOnHours
	ch <- c.nvmeUnsafeShutdowns
	ch <- c.nvmeMediaErrors
	ch <- c.nvmeNumErrLogEntries
	ch <- c.nvmeWarningTempTime
	ch <- c.nvmeCriticalCompTime
	ch <- c.nvmeThmTemp1TransCount
	ch <- c.nvmeThmTemp2TransCount
	ch <- c.nvmeThmTemp1TotalTime
	ch <- c.nvmeThmTemp2TotalTime
}

func (c *nvmeCollector) Collect(ch chan<- prometheus.Metric) {
	nvmeDeviceCmd, err := exec.Command("nvme", "list", "-o", "json").Output()
	if err != nil {
		log.Fatalf("Error running nvme command: %s\n", err)
	}
	if !gjson.Valid(string(nvmeDeviceCmd)) {
		log.Fatal("nvmeDeviceCmd json is not valid")
	}
	nvmeDeviceList := gjson.Get(string(nvmeDeviceCmd), "Devices.#.DevicePath")
	for _, nvmeDevice := range nvmeDeviceList.Array() {
		nvmeSmartLog, err := exec.Command("nvme", "smart-log", nvmeDevice.String(), "-o", "json").Output()
		if err != nil {
			log.Fatalf("Error running nvme smart-log command for device %s: %s\n", nvmeDevice.String(), err)
		}
		if !gjson.Valid(string(nvmeSmartLog)) {
			log.Fatalf("nvmeSmartLog json is not valid for device: %s: %s\n", nvmeDevice.String(), err)
		}
		nvmeSmartLogMetrics := gjson.GetMany(string(nvmeSmartLog),
                                                     "critical_warning",
                                                     "temperature",
                                                     "avail_spare",
                                                     "spare_thresh",
                                                     "percent_used",
                                                     "endurance_grp_critical_warning_summary",
                                                     "data_units_read",
                                                     "data_units_written",
                                                     "host_read_commands",
                                                     "host_write_commands",
                                                     "controller_busy_time",
                                                     "power_cycles",
                                                     "power_on_hours",
                                                     "unsafe_shutdowns",
                                                     "media_errors",
                                                     "num_err_log_entries",
                                                     "warning_temp_time",
                                                     "critical_comp_time",
                                                     "thm_temp1_trans_count",
                                                     "thm_temp2_trans_count",
                                                     "thm_temp1_total_time",
                                                     "thm_temp2_total_time",)

		ch <- prometheus.MustNewConstMetric(c.nvmeCriticalWarning, prometheus.GaugeValue, nvmeSmartLogMetrics[0].Float(), nvmeDevice.String())
		// convert kelvin to fahrenheit
		ch <- prometheus.MustNewConstMetric(c.nvmeTemperature, prometheus.GaugeValue, (nvmeSmartLogMetrics[1].Float() - 273.15) * 9/5 + 32, nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeAvailSpare, prometheus.GaugeValue, nvmeSmartLogMetrics[2].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeSpareThresh, prometheus.GaugeValue, nvmeSmartLogMetrics[3].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmePercentUsed, prometheus.GaugeValue, nvmeSmartLogMetrics[4].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeEnduranceGrpCriticalWarningSummary, prometheus.GaugeValue, nvmeSmartLogMetrics[5].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeDataUnitsRead, prometheus.CounterValue, nvmeSmartLogMetrics[6].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeDataUnitsWritten, prometheus.CounterValue, nvmeSmartLogMetrics[7].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeHostReadCommands, prometheus.CounterValue, nvmeSmartLogMetrics[8].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeHostWriteCommands, prometheus.CounterValue, nvmeSmartLogMetrics[9].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeControllerBusyTime, prometheus.CounterValue, nvmeSmartLogMetrics[10].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmePowerCycles, prometheus.CounterValue, nvmeSmartLogMetrics[11].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmePowerOnHours, prometheus.CounterValue, nvmeSmartLogMetrics[12].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeUnsafeShutdowns, prometheus.CounterValue, nvmeSmartLogMetrics[13].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeMediaErrors, prometheus.CounterValue, nvmeSmartLogMetrics[14].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeNumErrLogEntries, prometheus.CounterValue, nvmeSmartLogMetrics[15].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeWarningTempTime, prometheus.CounterValue, nvmeSmartLogMetrics[16].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeCriticalCompTime, prometheus.CounterValue, nvmeSmartLogMetrics[17].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeThmTemp1TransCount, prometheus.CounterValue, nvmeSmartLogMetrics[18].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeThmTemp2TransCount, prometheus.CounterValue, nvmeSmartLogMetrics[19].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeThmTemp1TotalTime, prometheus.CounterValue, nvmeSmartLogMetrics[20].Float(), nvmeDevice.String())
		ch <- prometheus.MustNewConstMetric(c.nvmeThmTemp2TotalTime, prometheus.CounterValue, nvmeSmartLogMetrics[21].Float(), nvmeDevice.String())
	}
}

func main() {
	port := flag.String("port", "9998", "port to listen on")
	flag.Parse()
	// check user
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting current user  %s\n", err)
	}
	if currentUser.Username != "root" {
		log.Fatalf("Error: you must be root to use nvme-cli")
	}
	// check for nvme-cli executable
	_, err = exec.LookPath("nvme")
	if err != nil {
		log.Fatalf("Cannot find nvme command in path: %s\n", err)
	}
	prometheus.MustRegister(newNvmeCollector())
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
