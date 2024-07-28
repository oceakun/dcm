package main

import (
	"dcm/terminalplot"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Create a grid
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	// Create widgets 
	processTable := terminalplot.CreateProcessTable()
	storagePie := terminalplot.CreateStoragePieChart()
	tempTable := terminalplot.CreateTempTable()
	memoryPlot := terminalplot.CreateMemoryPlot()

	
	// Set up grid layout
	grid.Set(
		ui.NewCol(1.0/2, 
			ui.NewRow(1.0/3, memoryPlot),
			ui.NewRow(1.0/3,tempTable),
			ui.NewRow(1.0/3,storagePie),
		),
		ui.NewCol(1.0/2,processTable),
	)

	// Render the grid
	ui.Render(grid)

	// Set up ticker for updates
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// Event loop
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		case <-ticker.C:
			terminalplot.UpdateProcessTable(processTable)
			terminalplot.UpdateStoragePieChart(storagePie)
			terminalplot.UpdateTempTable(tempTable)
			terminalplot.UpdateMemoryPlot(memoryPlot)
			ui.Render(grid)
		}
	}
}