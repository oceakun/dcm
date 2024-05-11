package main

import (
	"fmt"
	"log"
	"math"
	"time"
	"dcm/memory"
	"dcm/process"
	// "dcm/terminalplot"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func FetchRAMUsage() (float64, float64, error) {
	v, err := memory.GetVirtualMemoryStats()
	if err != nil {
		return 0, 0, err
	}

	usedGB := v.UsedGB
	freeGB := v.FreeGB

	return usedGB, freeGB, nil
}

func Dashboard() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Virtual Memory Usage Line Chart
	lc := widgets.NewPlot()
	lc.Title = "Virtual Memory Usage"
	// lc.Data = [][]float64{{0}, {0}}
	lc.Data = make([][]float64, 2)
	lc.Data[0] = make([]float64, 0, 60) 
	lc.Data[1] = make([]float64, 0, 60) 
	lc.LineColors[0] = ui.ColorYellow
	lc.LineColors[1] = ui.ColorBlue
	lc.Marker = widgets.MarkerBraille

	// Storage Pie Chart
	pc := widgets.NewPieChart()
	pc.Title = "Storage"
	pc.Data = []float64{0, 0}
	pc.AngleOffset = -.5 * math.Pi

	// Processes Table
	table := widgets.NewTable()
	table.Title = "Processes"
	table.Rows = [][]string{{"Name", "CPU%", "Status", "Running", "Username"}}

	// Layout
	termWidth, termHeight := ui.TerminalDimensions()
	lc.SetRect(0, 0, termWidth/2, termHeight/2)
	pc.SetRect(termWidth/2, 0, termWidth, termHeight/2)
	table.SetRect(0, termHeight/2, termWidth, termHeight)

	ui.Render(lc, pc, table)

	ticker := time.NewTicker(2 * time.Second).C
	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			if e.Type == ui.KeyboardEvent {
				return
			}
		case <-ticker:
			// Update Virtual Memory Usage Line Chart
			usedGB, freeGB, err := FetchRAMUsage()
			if err != nil {
				fmt.Printf("Error fetching RAM usage: %v\n", err)
				continue
			}
			lc.Data[0] = append(lc.Data[0], usedGB)
			lc.Data[1] = append(lc.Data[1], freeGB)

			// Update Storage Pie Chart
			storageStats, err := memory.GetStorageStats("/")
			if err != nil {
				log.Fatalf("Error retrieving storage stats: %v", err)
			}
			pc.Data = []float64{storageStats.UsedGB, storageStats.FreeGB}

			// Update Processes Table
			processesInfo := process.GetProcessesInfo()
			table.Rows = [][]string{{"Name", "CPU%", "Status", "Running", "Username"}}
			for _, pInfo := range processesInfo {
				runningStatus := "Stopped"
				if pInfo.IsRunning {
					runningStatus = "Running"
				}
				table.Rows = append(table.Rows, []string{
					pInfo.Name,
					fmt.Sprintf("%.5f", pInfo.CPUPercent),
					pInfo.Status,
					runningStatus,
					pInfo.Username,
				})
			}

			ui.Render(lc, pc, table)
		}
	}
}
