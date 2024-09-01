package terminalplot

import (
	"fmt"
	"log"
	"dcm/internal/memory"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func CreateMemoryPlot() *widgets.Plot {
	plot := widgets.NewPlot()
	plot.Title = "Virtual Memory Usage"
	plot.Data = make([][]float64, 1)  // Only one data series for used memory
	plot.Data[0] = make([]float64, 1) // Initialize with one data point
	plot.LineColors[0] = ui.ColorYellow // Used memory line color
	plot.Marker = widgets.MarkerBraille

	// Set the color of the axis labels and markers
	plot.AxesColor = ui.ColorWhite

	// Set border style, including border color
	plot.BorderStyle.Fg = ui.ColorBlue // Set the border color to blue
    plot.TitleStyle = ui.NewStyle(ui.ColorYellow, ui.ColorClear, ui.ModifierBold)

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
	
	// Keep only the last 100 data points
	if len(plot.Data[0]) > 100 {
		plot.Data[0] = plot.Data[0][1:]
	}
	
	// Update the title with the used memory
	plot.Title = fmt.Sprintf("Virtual Memory - Used: %.2f GB, Free: %.2f GB",
		usedGB, freeGB)
	}

func FetchRAMUsage() (float64, float64, error) {
	v, err := memory.GetVirtualMemoryStats()
	if err != nil {
		return 0, 0, err
	}
	return v.UsedGB, v.FreeGB, nil
}
