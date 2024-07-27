package terminalplot

import (
	"fmt"
	"log"
	"dcm/memory"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func CreateMemoryPlot() *widgets.Plot {
	plot := widgets.NewPlot()
	plot.Title = "Virtual Memory Usage"
	plot.Data = make([][]float64, 2)
	plot.Data[0] = make([]float64, 1) // Initialize with one data point
	plot.Data[1] = make([]float64, 1) // Initialize with one data point
	plot.LineColors[0] = ui.ColorYellow // Used memory
	plot.LineColors[1] = ui.ColorBlue   // Free memory
	plot.Marker = widgets.MarkerBraille
	plot.AxesColor = ui.ColorWhite
	// plot.LineStyles[0] = ui.NewStyle(ui.ColorYellow)
	// plot.LineStyles[1] = ui.NewStyle(ui.ColorBlue)
	UpdateMemoryPlot(plot) // Initial update to populate with real data
	return plot
}

func UpdateMemoryPlot(plot *widgets.Plot) {
	usedGB, freeGB, err := FetchRAMUsage()
	if err != nil {
		log.Printf("Error fetching RAM usage: %v", err)
		return
	}
	
	// Append new data points
	plot.Data[0] = append(plot.Data[0], usedGB)
	plot.Data[1] = append(plot.Data[1], freeGB)
	
	// Keep only the last 100 data points
	if len(plot.Data[0]) > 100 {
		plot.Data[0] = plot.Data[0][1:]
		plot.Data[1] = plot.Data[1][1:]
	}
	
	// Ensure we always have at least one data point
	if len(plot.Data[0]) == 0 {
		plot.Data[0] = append(plot.Data[0], usedGB)
		plot.Data[1] = append(plot.Data[1], freeGB)
	}
	
	plot.Title = fmt.Sprintf("Virtual Memory - Used: %.2f GB, Free: %.2f GB", usedGB, freeGB)
}

func FetchRAMUsage() (float64, float64, error) {
	v, err := memory.GetVirtualMemoryStats()
	if err != nil {
		return 0, 0, err
	}
	return v.UsedGB, v.FreeGB, nil
}