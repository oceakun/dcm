package terminalplot

import (
	"dcm/internal/network"
	"fmt"
	// "strings"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// CreateNetworkTable initializes and returns a termui Table displaying network interface information.
func CreateNetworkTable() *widgets.Table {
	table := widgets.NewTable()
	table.Title = "Network Interfaces"
	table.TextStyle = termui.NewStyle(termui.ColorWhite)
	table.RowSeparator = false
	table.BorderStyle = termui.NewStyle(termui.ColorBlue)
	table.TitleStyle = termui.NewStyle(termui.ColorYellow, termui.ColorClear, termui.ModifierBold)

	// Set the header row
	headers := []string{"Name", "MTU", "Hardware Address", "Flags", "Interface Addresses"}
	table.Rows = [][]string{headers}

	// Update the table with network interface data
	UpdateNetworkTable(table)

	return table
}

// UpdateNetworkTable refreshes the termui Table with the latest network interface data.
func UpdateNetworkTable(table *widgets.Table) {
	interfaces := network.GetNetworkInterfaces()

	// Populate the table with data starting from the first row (after the header)
	newRows := [][]string{{"Name", "MTU", "Hardware Address", "Flags", "Interface Addresses"}}

	// Variables to store the max width for each column
	maxWidths := []int{len("Name"), len("MTU"), len("Hardware Address"), len("Flags"), len("Interface Addresses")}

	for _, iface := range interfaces {
		// Collect all addresses as a single string
		addresses := fmt.Sprintf("%v", iface.Addresses)
		row := []string{
			iface.Name,
			fmt.Sprintf("%d", iface.MTU),
			iface.HardwareAddr,
			fmt.Sprintf("%v", iface.Flags),
			addresses,
		}

		// Update the max width for each column
		for i, cell := range row {
			if len(cell) > maxWidths[i] {
				maxWidths[i] = len(cell)
			}
		}

		newRows = append(newRows, row)
	}

	table.Rows = newRows
	table.Title = fmt.Sprintf("Network Interfaces (%d)", len(interfaces))
	// Highlight the header row
	table.RowStyles[0] = termui.NewStyle(termui.ColorGreen, termui.ColorClear, termui.ModifierBold)
	// Set the title color
	table.TitleStyle = termui.NewStyle(termui.ColorYellow, termui.ColorClear, termui.ModifierBold)

	// Get the total width of the terminal
	totalWidth, _ := termui.TerminalDimensions()

	// Calculate the total width needed based on content
	totalContentWidth := 0
	for _, w := range maxWidths {
		totalContentWidth += w
	}

	// If there's remaining space, distribute it proportionally
	if totalContentWidth < totalWidth {
		extraSpace := totalWidth - totalContentWidth
		for i := range maxWidths {
			// Distribute extra space proportionally to each column based on its width
			proportionalSpace := (maxWidths[i] * extraSpace) / totalContentWidth
			maxWidths[i] += proportionalSpace
		}
	}

	// Set the column widths
	table.ColumnWidths = maxWidths
}
