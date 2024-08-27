package main

import (
	"log"
	"time"
	"dcm/internal/terminalplot"
	"github.com/charmbracelet/huh"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/rivo/tview"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	var selectedOption string
	huh.NewSelect[string]().
		Title("View : ").
		Options(
			huh.NewOption("Temperature Table", "tt"),
			// huh.NewOption("Process Table", "pt"),
			huh.NewOption("Storage Pie", "st"),
			huh.NewOption("Virtual Memory", "vm"),
			huh.NewOption("Interactive Process Table", "tpt"),
			huh.NewOption("Live Dashboard", "ldb"),
		).
		Value(&selectedOption).Run()

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	var activeWidget ui.Drawable

	switch selectedOption {
	case "tt":
		activeWidget = terminalplot.CreateTempTable()
	// case "pt":
	// 	activeWidget = terminalplot.CreateInteractiveProcessTable()
	case "st":
		activeWidget = terminalplot.CreateStoragePieChart()
	case "vm":
		activeWidget = terminalplot.CreateMemoryPlot()
	case "ldb":
		terminalplot.CreateLiveDashboard()
		return
	case "tpt":
		// Handle the tview process table option
		app := tview.NewApplication()
		flex := terminalplot.CreateTviewProcessTable()

		go func() {
			if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
				log.Fatalf("Error running tview application: %v", err)
			}
		}()
		select {} // Block the main goroutine, as tview will take over.
	}

	grid.Set(
		ui.NewRow(1.0, activeWidget),
	)

	ui.Clear()
	ui.Render(grid)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

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
			switch selectedOption {
			case "tt":
				terminalplot.UpdateTempTable(activeWidget.(*widgets.Table))
			// case "pt":
			// 	terminalplot.UpdateInteractiveProcessTable(activeWidget.(*widgets.Table))
			case "st":
				terminalplot.UpdateStoragePieChart(activeWidget.(*widgets.PieChart))
			case "vm":
				terminalplot.UpdateMemoryPlot(activeWidget.(*widgets.Plot))
			}
			ui.Render(grid)
		}
	}
}

