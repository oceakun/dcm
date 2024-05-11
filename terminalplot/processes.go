package terminalplot

import (
	"fmt"
	"log"
	"dcm/process"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func ProcessTable() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	processesInfo := process.GetProcessesInfo()

	table := widgets.NewTable()
	table.Title = "Processes"
	table.Rows = [][]string{
		{"Name", "CPU%", "Status", "Running", "Username"},
	}

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

	table.TextStyle = termui.NewStyle(termui.ColorWhite)
	table.RowSeparator = false
	table.SetRect(0, 0, 100, len(processesInfo)+2) 
	termui.Render(table)

	for e := range termui.PollEvents() {
		if e.Type == termui.KeyboardEvent {
			break
		}
	}
}
