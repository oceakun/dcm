package terminalplot

import (
	"dcm/temperature"
	"fmt"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func CreateTempTable() *widgets.Table {
	table := widgets.NewTable()
	table.Title = "Temperature Sensors"
	table.Rows = [][]string{{"Sensor", "Temperature"}}
	table.TextStyle = termui.NewStyle(termui.ColorWhite)
	table.RowSeparator = false
	table.BorderStyle = termui.NewStyle(termui.ColorBlue)
	table.RowStyles[0] = termui.NewStyle(termui.ColorGreen, termui.ColorClear, termui.ModifierBold)
	// Set the title color
    table.TitleStyle = termui.NewStyle(termui.ColorYellow, termui.ColorClear, termui.ModifierBold)
	UpdateTempTable(table)
	return table
}

func UpdateTempTable(table *widgets.Table) {
	temperatures := temperature.GetTemperatures()

	newRows := [][]string{{"Sensor", "Temperature"}}

	for sensor, temp := range temperatures {
		newRows = append(newRows, []string{sensor, fmt.Sprintf("%.1f°C", temp)})
	}

	table.Rows = newRows
	table.Title = fmt.Sprintf("Temperature Sensors (%d)", len(temperatures))
}