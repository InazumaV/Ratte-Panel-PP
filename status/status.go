package status

import "github.com/shirou/gopsutil/v4/cpu"

type SystemStatus struct {
	CPUUsage    float64
	MemoryUsage float64
	DiskUsage   float64
}

func GetSystemStatus() (SystemStatus, error) {
	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		return SystemStatus{}, err
	}
	memoryUsage, err := cpu.Percent(0, false)
	if err != nil {
		return SystemStatus{}, err
	}
	diskUsage, err := cpu.Percent(0, false)
	if err != nil {
		return SystemStatus{}, err
	}
	return SystemStatus{
		CPUUsage:    cpuUsage[0],
		MemoryUsage: memoryUsage[0],
		DiskUsage:   diskUsage[0],
	}, nil
}
