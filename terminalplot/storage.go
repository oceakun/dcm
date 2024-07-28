package terminalplot

import (
	"fmt"
	"log"
	"math"
	"dcm/memory"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func CreateStoragePieChart() *widgets.PieChart {
	pc := widgets.NewPieChart()
	pc.Title = "Storage"
	pc.TitleStyle = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)

	pc.AngleOffset = .5 * math.Pi

	// Set the colors for the pie chart segments
	pc.Colors = []ui.Color{ui.ColorRed, ui.ColorGreen} // Example colors for used and free storage

	// Set the border color to blue
	pc.BorderStyle.Fg = ui.ColorBlue

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
