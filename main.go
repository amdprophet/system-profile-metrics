package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

// static const char *cpu_state_names[] = {
// "user",    "system", "wait",  "nice",       "swap", "interrupt",
// "softirq", "steal", "guest", "guest_nice", "idle", "active"};

func checkErr(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
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

func outputCpuMetric(cpu string, name string, value string) {
	key := fmt.Sprintf("%s.%s", cpu, name)
	outputMetric(key, value)
}

func outputCpuMetrics(cpu cpu.TimesStat) {
	outputCpuMetric(cpu.CPU, "user", fmt.Sprint(cpu.User))
	outputCpuMetric(cpu.CPU, "system", fmt.Sprint(cpu.System))
	outputCpuMetric(cpu.CPU, "idle", fmt.Sprint(cpu.Idle))
	outputCpuMetric(cpu.CPU, "nice", fmt.Sprint(cpu.Nice))
	outputCpuMetric(cpu.CPU, "iowait", fmt.Sprint(cpu.Iowait))
	outputCpuMetric(cpu.CPU, "irq", fmt.Sprint(cpu.Irq))
	outputCpuMetric(cpu.CPU, "softirq", fmt.Sprint(cpu.Softirq))
	outputCpuMetric(cpu.CPU, "steal", fmt.Sprint(cpu.Steal))
	outputCpuMetric(cpu.CPU, "guest", fmt.Sprint(cpu.Guest))
	outputCpuMetric(cpu.CPU, "guest_nice", fmt.Sprint(cpu.GuestNice))
	outputCpuMetric(cpu.CPU, "stolen", fmt.Sprint(cpu.Stolen))
}

func main() {
	perCpuTimes, err := cpu.Times(true)
	checkErr(err)

	for _, cpu := range perCpuTimes {
		outputCpuMetrics(cpu)
	}

	totalCpuTimes, err := cpu.Times(false)
	checkErr(err)

	for _, cpu := range totalCpuTimes {
		outputCpuMetrics(cpu)
	}
}
