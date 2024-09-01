package memory

import (
	"math"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

const bytesInGigabyte = 1 << 30 // 2^30 bytes in a gigabyte

// VirtualMemoryStats holds statistics about virtual memory.
type VirtualMemoryStats struct {
	TotalGB       float64
	FreeGB        float64
	UsedGB        float64
	UsedPercent   float64
}

// StorageStats holds statistics about storage.
type StorageStats struct {
	TotalGB       float64
	FreeGB        float64
	UsedGB        float64
	UsedPercent   float64
}

// Convert bytes to gigabytes (GB).
func bytesToGB(bytes uint64) float64 {
	return math.Round(float64(bytes)/float64(bytesInGigabyte)*1000) / 1000
}

// GetVirtualMemoryStats retrieves the current virtual memory stats and converts them to GB.
func GetVirtualMemoryStats() (*VirtualMemoryStats, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return &VirtualMemoryStats{
		TotalGB:     bytesToGB(v.Total),
		FreeGB:      bytesToGB(v.Free),
		UsedGB:      bytesToGB(v.Used),
		UsedPercent: v.UsedPercent,
	}, nil
}

// GetStorageStats retrieves the current storage stats for a given path and converts them to GB.
func GetStorageStats(path string) (*StorageStats, error) {
	usageStat, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}
	return &StorageStats{
		TotalGB:     bytesToGB(usageStat.Total),
		FreeGB:      bytesToGB(usageStat.Free),
		UsedGB:      bytesToGB(usageStat.Used),
		UsedPercent: usageStat.UsedPercent,
	}, nil
}
