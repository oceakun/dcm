package terminalplot

import (
	"fmt"
	"log"
	"math"
	"dcm/memory"
	"github.com/gizak/termui/v3/widgets"
)

func CreateStoragePieChart() *widgets.PieChart {
	pc := widgets.NewPieChart()
	pc.Title = "Storage"
	pc.AngleOffset = -.5 * math.Pi
	UpdateStoragePieChart(pc)
	return pc
}

func UpdateStoragePieChart(pc *widgets.PieChart) {
	storageStats, err := memory.GetStorageStats("/")
	if err != nil {
		log.Printf("Error retrieving storage stats: %v", err)
		return
	}
	pc.Data = []float64{storageStats.UsedGB, storageStats.FreeGB}
	pc.Title = fmt.Sprintf("Storage - Used: %.2f GB, Free: %.2f GB", storageStats.UsedGB, storageStats.FreeGB)
}
