package terminalplot

import (
	"fmt"
	"log"
	"math"

	"dcm/memory"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var run = true  

func PieChart() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	storageStats, err := memory.GetStorageStats("/") 
	if err != nil {
		log.Fatalf("Error retrieving storage stats: %v", err)
	}

	usedMemory := storageStats.UsedGB  
	freeMemory := storageStats.FreeGB 
	pc := widgets.NewPieChart()
	pc.Title = "Storage"
	pc.Border = false
	pc.SetRect(0, 0, 20, 20) 
	pc.Data = []float64{usedMemory, freeMemory}
	pc.AngleOffset = -.5 * math.Pi
	labels := widgets.NewParagraph()
	labels.Text = fmt.Sprintf("[Used: %.2f GB](fg:red)\n[Free: %.2f GB](fg:green)", usedMemory, freeMemory)

	labels.Border = false
	labels.SetRect(0, 0, 100, 6) 
	ui.Render(pc, labels)

	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		}
	}
}
