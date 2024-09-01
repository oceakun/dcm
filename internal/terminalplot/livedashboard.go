package terminalplot

import (
	"log"
	"time"
	ui "github.com/gizak/termui/v3"
)

func CreateLiveDashboard() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	processTable := CreateProcessTable()
	storagePie := CreateStoragePieChart()
	tempTable := CreateTempTable()
	memoryPlot := CreateMemoryPlot()
	networkInterface := CreateNetworkTable()

	grid.Set(
		ui.NewRow(2.0/3,
		ui.NewCol(1.0/2, 
			ui.NewRow(1.0/2, memoryPlot),
			ui.NewRow(1.0/2,ui.NewCol(1, ui.NewCol(1.0/2,tempTable), ui.NewCol(1.0/2,storagePie))),
		),
		ui.NewCol(1.0/2,processTable),
	),
		ui.NewCol(1,networkInterface),
	)

	ui.Render(grid)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

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
			UpdateProcessTable(processTable)
			UpdateStoragePieChart(storagePie)
			UpdateTempTable(tempTable)
			UpdateMemoryPlot(memoryPlot)
			ui.Render(grid)
		}
	}
}