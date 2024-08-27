package process

import (
	"github.com/shirou/gopsutil/v3/process"
	"strings"
)

// ProcessInfo is the struct that holds the details for each process
type ProcessInfo struct {
	Name       string
	CPUPercent float64
	Status     string
	IsRunning  bool
	Username   string
}

// GetProcessesInfo fetches the list of processes and their details
func GetProcessesInfo() []ProcessInfo {
	var processesInfo []ProcessInfo

	processes, _ := process.Processes()
	for _, proc := range processes {
		name, _ := proc.Name()
		isRunning, _ := proc.IsRunning()
		cpu, _ := proc.CPUPercent()
		cpuRounded := float64(int(cpu*100000)) / 100000
		status, _ := proc.Status()
		statusStr := strings.Join(status, ", ") 
		username, _ := proc.Username()

		processesInfo = append(processesInfo, ProcessInfo{
			Name:       name,
			CPUPercent: cpuRounded,
			Status:     statusStr,
			IsRunning:  isRunning,
			Username:   username,
		})
	}

	return processesInfo
}