package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/spf13/cobra"
)

var (
	checkName = "system-profile-metrics"

	rootCmd = &cobra.Command{
		Use:   checkName,
		Short: "Output system metrics (cpu, memory, etc.)",
		Run: func(cmd *cobra.Command, args []string) {
			outputMetrics()
		},
	}
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
}

func main() {
	err := rootCmd.Execute()
	checkErr(err)
}

func outputMetric(name string, value string) {
	hostInfo, err := host.Info()
	checkErr(err)

	fmt.Printf("%s.%s.%s, %s\n",
		hostInfo.Hostname,
		"cpu",
		name,
		value,
	)
}

func outputCPUMetric(cpu string, name string, value string) {
	key := fmt.Sprintf("%s.%s", cpu, name)
	outputMetric(key, value)
}

func outputCPUMetrics(cpu cpu.TimesStat) {
	outputCPUMetric(cpu.CPU, "user", fmt.Sprint(cpu.User))
	outputCPUMetric(cpu.CPU, "system", fmt.Sprint(cpu.System))
	outputCPUMetric(cpu.CPU, "idle", fmt.Sprint(cpu.Idle))
	outputCPUMetric(cpu.CPU, "nice", fmt.Sprint(cpu.Nice))
	outputCPUMetric(cpu.CPU, "iowait", fmt.Sprint(cpu.Iowait))
	outputCPUMetric(cpu.CPU, "irq", fmt.Sprint(cpu.Irq))
	outputCPUMetric(cpu.CPU, "softirq", fmt.Sprint(cpu.Softirq))
	outputCPUMetric(cpu.CPU, "steal", fmt.Sprint(cpu.Steal))
	outputCPUMetric(cpu.CPU, "guest", fmt.Sprint(cpu.Guest))
	outputCPUMetric(cpu.CPU, "guest_nice", fmt.Sprint(cpu.GuestNice))
	outputCPUMetric(cpu.CPU, "stolen", fmt.Sprint(cpu.Stolen))
}

func outputMetrics() {
	perCPUTimes, err := cpu.Times(true)
	checkErr(err)

	for _, cpu := range perCPUTimes {
		outputCPUMetrics(cpu)
	}

	totalCPUTimes, err := cpu.Times(false)
	checkErr(err)

	for _, cpu := range totalCPUTimes {
		outputCPUMetrics(cpu)
	}
}
