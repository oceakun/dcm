package terminalplot

import (
	"dcm/process"
	"fmt"
	"sort"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ProcessInfo struct {
	Name       string
	CPUPercent float64
	Status     string
	IsRunning  bool
	Username   string
}

var (
	processesInfo    []ProcessInfo // Currently displayed processes
	allProcessesInfo []ProcessInfo // All processes fetched

	// Global variables for buttons and sorting states
	sortByNameButton *tview.Button
	sortByCPUButton  *tview.Button
	sortByNameAsc    = true // Initial sort order for Name
	sortByCPUAsc     = true // Initial sort order for CPU%
)

func CreateTviewProcessTable() *tview.Flex {
    table := tview.NewTable().SetBorders(true)
    
    // Set up header
    headers := []string{"Name", "CPU%", "Status", "Running", "Username"}
    for col, header := range headers {
        cell := tview.NewTableCell(header).
            SetTextColor(tcell.ColorYellow).
            SetAlign(tview.AlignCenter).
            SetSelectable(false)
        table.SetCell(0, col, cell)
    }
    
    // Populate table with initial process data
    UpdateTviewProcessTable(table)
    
    // Set up search input
    inputField := tview.NewInputField().
        SetLabel("Search: ").
        SetChangedFunc(func(text string) {
            searchTable(table, text)
        })
    
    // Create "Sort by Name" button with initial label
    sortByNameButton = tview.NewButton("Name ↑").
        SetSelectedFunc(func() {
            sortTableByName()
            populateTable(table, processesInfo)
            table.ScrollToBeginning()
        })
    
    // Create "Sort by CPU%" button with initial label
    sortByCPUButton = tview.NewButton("CPU % ↑").
        SetLabelColor(tcell.Color100).
        SetSelectedFunc(func() {
            sortTableByCPU()
            populateTable(table, processesInfo)
            table.ScrollToBeginning()
        })
    
    // Layout the buttons horizontally
    buttonFlex := tview.NewFlex().
        SetDirection(tview.FlexColumn).
        AddItem(sortByNameButton, 0, 1, false).
        AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorDefault), 1, 0, false).
        AddItem(sortByCPUButton, 0, 1, false)
    
    // Combine search input and buttons in a single row
    inputAndButtonsFlex := tview.NewFlex().
        SetDirection(tview.FlexColumn).
        AddItem(inputField, 0, 3, false).
        AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorDefault), 1, 0, false).
        AddItem(buttonFlex, 0, 2, false)
    
    // Main layout
    flex := tview.NewFlex().
        SetDirection(tview.FlexRow).
        AddItem(inputAndButtonsFlex, 1, 0, false).
        AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorDefault), 1, 0, false).
        AddItem(table, 0, 1, true)
    
    flex.SetFullScreen(true)
    return flex
}


// UpdateTviewProcessTable updates the table with the latest process data
func UpdateTviewProcessTable(table *tview.Table) {
	allProcessesInfo = getProcessesInfo()
	processesInfo = allProcessesInfo // Initialize displayed processes with all processes
	populateTable(table, processesInfo)
}

func getProcessesInfo() []ProcessInfo {
	rawProcessesInfo := process.GetProcessesInfo()
	processesInfo := make([]ProcessInfo, len(rawProcessesInfo))

	for i, pInfo := range rawProcessesInfo {

		processesInfo[i] = ProcessInfo{
			Name:       pInfo.Name,
			CPUPercent: pInfo.CPUPercent,
			Status:     pInfo.Status,
			IsRunning:  pInfo.IsRunning,
			Username:   pInfo.Username,
		}
	}

	return processesInfo
}

func populateTable(table *tview.Table, data []ProcessInfo) {
	table.Clear() // Clear the table before repopulating

	// Set up header again after clearing
	headers := []string{"Name", "CPU%", "Status", "Running", "Username"}
	for col, header := range headers {
		cell := tview.NewTableCell(header).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter).
			SetSelectable(false) // Headers are not selectable
		table.SetCell(0, col, cell)
	}

	// Populate the table with the provided process data
	for row, pInfo := range data {
		table.SetCell(row+1, 0, tview.NewTableCell(pInfo.Name).SetSelectable(true))
		table.SetCell(row+1, 1, tview.NewTableCell(fmt.Sprintf("%.2f", pInfo.CPUPercent)).SetSelectable(true))
		table.SetCell(row+1, 2, tview.NewTableCell(pInfo.Status).SetSelectable(true))
		table.SetCell(row+1, 3, tview.NewTableCell(getRunningStatus(pInfo.IsRunning)).SetSelectable(true))
		table.SetCell(row+1, 4, tview.NewTableCell(pInfo.Username).SetSelectable(true))
	}
}

func sortTableByName() {
	if sortByNameAsc {
		sort.Slice(processesInfo, func(i, j int) bool {
			return strings.ToLower(processesInfo[i].Name) < strings.ToLower(processesInfo[j].Name)
		})
		sortByNameButton.SetLabel("Name ↓")
	} else {
		sort.Slice(processesInfo, func(i, j int) bool {
			return strings.ToLower(processesInfo[i].Name) > strings.ToLower(processesInfo[j].Name)
		})
		sortByNameButton.SetLabel("Name ↑")
	}
	sortByNameAsc = !sortByNameAsc
}

func sortTableByCPU() {
	if sortByCPUAsc {
		sort.Slice(processesInfo, func(i, j int) bool {
			return processesInfo[i].CPUPercent < processesInfo[j].CPUPercent
		})
		sortByCPUButton.SetLabel("CPU % ↓")
	} else {
		sort.Slice(processesInfo, func(i, j int) bool {
			return processesInfo[i].CPUPercent > processesInfo[j].CPUPercent
		})
		sortByCPUButton.SetLabel("CPU % ↑")
	}
	sortByCPUAsc = !sortByCPUAsc
}

func getRunningStatus(isRunning bool) string {
	if isRunning {
		return "Running"
	}
	return "Stopped"
}

// searchTable filters the table based on the query and updates the displayed rows
func searchTable(table *tview.Table, query string) {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		// If query is empty, display all processes and scroll to the top
		processesInfo = allProcessesInfo
		populateTable(table, processesInfo)
		table.ScrollToBeginning() // Scroll to the top of the table
	} else {
		// Filter processesInfo based on the query
		var filtered []ProcessInfo
		for _, pInfo := range allProcessesInfo {
			if strings.Contains(strings.ToLower(pInfo.Name), query) ||
				strings.Contains(strings.ToLower(fmt.Sprintf("%.2f", pInfo.CPUPercent)), query) ||
				strings.Contains(strings.ToLower(pInfo.Status), query) ||
				strings.Contains(strings.ToLower(getRunningStatus(pInfo.IsRunning)), query) ||
				strings.Contains(strings.ToLower(pInfo.Username), query) {
				filtered = append(filtered, pInfo)
			}
		}
		processesInfo = filtered
		populateTable(table, processesInfo)
	}
}
