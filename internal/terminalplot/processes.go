package terminalplot

import (
	"fmt"
	"dcm/internal/process"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func CreateProcessTable() *widgets.Table {
    table := widgets.NewTable()
    table.Title = "Processes"
    table.Rows = [][]string{
        {"Name", "CPU%", "Status", "Running", "Username"},
    }
    table.TextStyle = termui.NewStyle(termui.ColorWhite)
    table.RowSeparator = false
    table.BorderStyle = termui.NewStyle(termui.ColorBlue) // Set the border color

    // Highlight the header row
    table.RowStyles[0] = termui.NewStyle(termui.ColorGreen, termui.ColorClear, termui.ModifierBold)
	// Set the title color
    table.TitleStyle = termui.NewStyle(termui.ColorYellow, termui.ColorClear, termui.ModifierBold)


    UpdateProcessTable(table)
    return table
}
func UpdateProcessTable(table *widgets.Table) {
	processesInfo := process.GetProcessesInfo()
	table.Rows = table.Rows[:1] // Keep only the header
	for _, pInfo := range processesInfo {
		runningStatus := "Stopped"
		if pInfo.IsRunning {
			runningStatus = "Running"
		}
		table.Rows = append(table.Rows, []string{
			pInfo.Name,
			fmt.Sprintf("%.2f", pInfo.CPUPercent),
			pInfo.Status,
			runningStatus,
			pInfo.Username,
		})
	}
}