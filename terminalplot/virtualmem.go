package terminalplot

import (
	"fmt"
	"time"

	"log"

	"dcm/memory"

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

func LinePlot() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()
	lc := widgets.NewPlot()
	lc.Title = "Virtual Memory Usage"
	lc.Data = [][]float64{{0},{0}}
	
	lc.LineColors[0] = ui.ColorYellow
	lc.LineColors[1] = ui.ColorBlue
	lc.Marker = widgets.MarkerBraille

	termWidth, termHeight := ui.TerminalDimensions()
	lc.SetRect(0, 0, termWidth, termHeight/2)

  uiEvents := ui.PollEvents()
    ticker := time.NewTicker(2 * time.Second).C

    for {
        select {
        case e := <-uiEvents:
            if e.Type == ui.KeyboardEvent {
                return
            }
        case <-ticker:
            usedGB, freeGB, err := FetchRAMUsage()
            if err != nil {
                fmt.Printf("Error fetching RAM usage: %v\n", err)
                continue
            }
            lc.Data[0] = append(lc.Data[0], usedGB)
            lc.Data[1] = append(lc.Data[1], freeGB)
            ui.Render(lc)
        }
    }
	
}
