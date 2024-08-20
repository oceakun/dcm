// package terminalplot

// import (
// 	"dcm/process"
// 	"fmt"
// 	"sort"
// 	"strings"

// 	"github.com/gdamore/tcell/v2"
// 	"github.com/rivo/tview"
// )

// type ProcessInfo struct {
// 	Name       string
// 	CPUPercent float64
// 	Status     string
// 	IsRunning  bool
// 	Username   string
// }

// var (
// 	processesInfo    []ProcessInfo // Currently displayed processes
// 	allProcessesInfo []ProcessInfo // All processes fetched

// 	// Global variables for buttons and sorting states
// 	sortByNameButton *tview.Button
// 	sortByCPUButton  *tview.Button
// 	sortByNameAsc    = true // Initial sort order for Name
// 	sortByCPUAsc     = true // Initial sort order for CPU%
// )

// func CreateTviewProcessTable() *tview.Flex {
//     table := tview.NewTable().SetBorders(true)

//     // Set up header
//     headers := []string{"Name", "CPU%", "Status", "Running", "Username"}
//     for col, header := range headers {
//         cell := tview.NewTableCell(header).
//             SetTextColor(tcell.ColorYellow).
//             SetAlign(tview.AlignCenter).
//             SetSelectable(false)
//         table.SetCell(0, col, cell)
//     }

//     // Populate table with initial process data
//     UpdateTviewProcessTable(table)

//     // Set up search input
//     inputField := tview.NewInputField().
//         SetLabel("Search: ").
//         SetChangedFunc(func(text string) {
//             searchTable(table, text)
//         })

//     // Create "Sort by Name" button with initial label
//     sortByNameButton = tview.NewButton("Name ↑").
//         SetSelectedFunc(func() {
//             sortTableByName()
//             populateTable(table, processesInfo)
//             table.ScrollToBeginning()
//         })

//     // Create "Sort by CPU%" button with initial label
//     sortByCPUButton = tview.NewButton("CPU % ↑").
//         SetLabelColor(tcell.Color100).
//         SetSelectedFunc(func() {
//             sortTableByCPU()
//             populateTable(table, processesInfo)
//             table.ScrollToBeginning()
//         })

//     // Layout the buttons horizontally
//     buttonFlex := tview.NewFlex().
//         SetDirection(tview.FlexColumn).
//         AddItem(sortByNameButton, 0, 1, false).
//         AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorDefault), 1, 0, false).
//         AddItem(sortByCPUButton, 0, 1, false)

//     // Combine search input and buttons in a single row
//     inputAndButtonsFlex := tview.NewFlex().
//         SetDirection(tview.FlexColumn).
//         AddItem(inputField, 0, 3, false).
//         AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorDefault), 1, 0, false).
//         AddItem(buttonFlex, 0, 2, false)

//     // Main layout
//     flex := tview.NewFlex().
//         SetDirection(tview.FlexRow).
//         AddItem(inputAndButtonsFlex, 1, 0, false).
//         AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorDefault), 1, 0, false).
//         AddItem(table, 0, 1, true)

//     flex.SetFullScreen(true)
//     return flex
// }

// // UpdateTviewProcessTable updates the table with the latest process data
// func UpdateTviewProcessTable(table *tview.Table) {
// 	allProcessesInfo = getProcessesInfo()
// 	processesInfo = allProcessesInfo // Initialize displayed processes with all processes
// 	populateTable(table, processesInfo)
// }

// func getProcessesInfo() []ProcessInfo {
// 	rawProcessesInfo := process.GetProcessesInfo()
// 	processesInfo := make([]ProcessInfo, len(rawProcessesInfo))

// 	for i, pInfo := range rawProcessesInfo {

// 		processesInfo[i] = ProcessInfo{
// 			Name:       pInfo.Name,
// 			CPUPercent: pInfo.CPUPercent,
// 			Status:     pInfo.Status,
// 			IsRunning:  pInfo.IsRunning,
// 			Username:   pInfo.Username,
// 		}
// 	}

// 	return processesInfo
// }

// func populateTable(table *tview.Table, data []ProcessInfo) {
// 	table.Clear() // Clear the table before repopulating

// 	// Set up header again after clearing
// 	headers := []string{"Name", "CPU%", "Status", "Running", "Username"}
// 	for col, header := range headers {
// 		cell := tview.NewTableCell(header).
// 			SetTextColor(tcell.ColorYellow).
// 			SetAlign(tview.AlignCenter).
// 			SetSelectable(false) // Headers are not selectable
// 		table.SetCell(0, col, cell)
// 	}

// 	// Populate the table with the provided process data
// 	for row, pInfo := range data {
// 		table.SetCell(row+1, 0, tview.NewTableCell(pInfo.Name).SetSelectable(true))
// 		table.SetCell(row+1, 1, tview.NewTableCell(fmt.Sprintf("%.2f", pInfo.CPUPercent)).SetSelectable(true))
// 		table.SetCell(row+1, 2, tview.NewTableCell(pInfo.Status).SetSelectable(true))
// 		table.SetCell(row+1, 3, tview.NewTableCell(getRunningStatus(pInfo.IsRunning)).SetSelectable(true))
// 		table.SetCell(row+1, 4, tview.NewTableCell(pInfo.Username).SetSelectable(true))
// 	}
// }

// func sortTableByName() {
// 	if sortByNameAsc {
// 		sort.Slice(processesInfo, func(i, j int) bool {
// 			return strings.ToLower(processesInfo[i].Name) < strings.ToLower(processesInfo[j].Name)
// 		})
// 		sortByNameButton.SetLabel("Name ↓")
// 	} else {
// 		sort.Slice(processesInfo, func(i, j int) bool {
// 			return strings.ToLower(processesInfo[i].Name) > strings.ToLower(processesInfo[j].Name)
// 		})
// 		sortByNameButton.SetLabel("Name ↑")
// 	}
// 	sortByNameAsc = !sortByNameAsc
// }

// func sortTableByCPU() {
// 	if sortByCPUAsc {
// 		sort.Slice(processesInfo, func(i, j int) bool {
// 			return processesInfo[i].CPUPercent < processesInfo[j].CPUPercent
// 		})
// 		sortByCPUButton.SetLabel("CPU % ↓")
// 	} else {
// 		sort.Slice(processesInfo, func(i, j int) bool {
// 			return processesInfo[i].CPUPercent > processesInfo[j].CPUPercent
// 		})
// 		sortByCPUButton.SetLabel("CPU % ↑")
// 	}
// 	sortByCPUAsc = !sortByCPUAsc
// }

// func getRunningStatus(isRunning bool) string {
// 	if isRunning {
// 		return "Running"
// 	}
// 	return "Stopped"
// }

// // searchTable filters the table based on the query and updates the displayed rows
// func searchTable(table *tview.Table, query string) {
// 	query = strings.ToLower(strings.TrimSpace(query))
// 	if query == "" {
// 		// If query is empty, display all processes and scroll to the top
// 		processesInfo = allProcessesInfo
// 		populateTable(table, processesInfo)
// 		table.ScrollToBeginning() // Scroll to the top of the table
// 	} else {
// 		// Filter processesInfo based on the query
// 		var filtered []ProcessInfo
// 		for _, pInfo := range allProcessesInfo {
// 			if strings.Contains(strings.ToLower(pInfo.Name), query) ||
// 				strings.Contains(strings.ToLower(fmt.Sprintf("%.2f", pInfo.CPUPercent)), query) ||
// 				strings.Contains(strings.ToLower(pInfo.Status), query) ||
// 				strings.Contains(strings.ToLower(getRunningStatus(pInfo.IsRunning)), query) ||
// 				strings.Contains(strings.ToLower(pInfo.Username), query) {
// 				filtered = append(filtered, pInfo)
// 			}
// 		}
// 		processesInfo = filtered
// 		populateTable(table, processesInfo)
// 	}
// }




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

	// Global variables for filters
	statusFilter   *tview.DropDown
	runningFilter  *tview.DropDown
	usernameFilter *tview.DropDown
)


// CreateTviewProcessTable creates the main layout with the process table, filters, and sort buttons.
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

	// Create "Sort by Name" button
	sortByNameButton = tview.NewButton("Name ↑").
		SetSelectedFunc(func() {
			sortTableByName()
			populateTable(table, processesInfo)
			table.ScrollToBeginning()
		})

	// Create "Sort by CPU%" button
	sortByCPUButton = tview.NewButton("CPU % ↑").
		SetLabelColor(tcell.Color100).
		SetSelectedFunc(func() {
			sortTableByCPU()
			populateTable(table, processesInfo)
			table.ScrollToBeginning()
		})

	// Create filter dropdowns
	statusFilter = tview.NewDropDown().
		SetLabel("Status: ").
		SetOptions(append([]string{"All"}, getUniqueValues(allProcessesInfo, "Status")...), func(option string, index int) {
			applyFilters(table)
		})

	runningFilter = tview.NewDropDown().
		SetLabel("Running: ").
		SetOptions([]string{"All", "Running", "Stopped"}, func(option string, index int) {
			applyFilters(table)
		})

	usernameFilter = tview.NewDropDown().
		SetLabel("Username: ").
		SetOptions(append([]string{"All"}, getUniqueValues(allProcessesInfo, "Username")...), func(option string, index int) {
			applyFilters(table)
		})

	// Combine sort buttons and filters vertically
	sortFilterColumnFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("").SetTextAlign(tview.AlignCenter), 1, 0, false).
		AddItem(tview.NewTextView().SetText("Sort by").SetTextAlign(tview.AlignCenter), 1, 0, false).
		AddItem(tview.NewTextView().SetText("").SetTextAlign(tview.AlignCenter), 1, 0, false).
		AddItem(sortByNameButton, 1, 0, false).
		AddItem(tview.NewTextView().SetText("").SetTextAlign(tview.AlignCenter), 1, 0, false).
		AddItem(sortByCPUButton, 1, 0, false).
		AddItem(tview.NewTextView().SetText("").SetTextAlign(tview.AlignCenter), 1, 0, false).
		AddItem(tview.NewTextView().SetText("").SetTextAlign(tview.AlignCenter), 1, 0, false).
		AddItem(tview.NewTextView().SetText("Filters").SetTextAlign(tview.AlignCenter), 1, 0, false).
		AddItem(tview.NewTextView().SetText("").SetTextAlign(tview.AlignCenter), 1, 0, false).
		AddItem(statusFilter, 1, 0, false).
		AddItem(tview.NewTextView().SetText("").SetTextAlign(tview.AlignCenter), 1, 0, false).
		AddItem(runningFilter, 1, 0, false).
		AddItem(tview.NewTextView().SetText("").SetTextAlign(tview.AlignCenter), 1, 0, false).
		AddItem(usernameFilter, 1, 0, false).
		AddItem(tview.NewTextView().SetText("").SetTextAlign(tview.AlignCenter), 1, 0, false)

	// Combine table and the right column (sort + filters) horizontally
	mainFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(table, 0, 3, true).
		AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorDefault), 1, 0, false).
		AddItem(sortFilterColumnFlex, 0, 1, false)

	// Main layout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(inputField, 1, 0, false).
		AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorDefault), 1, 0, false).
		AddItem(mainFlex, 0, 1, true)

	flex.SetFullScreen(true)

	return flex
}

// applyFilters applies the selected filters and updates the displayed processes
func applyFilters(table *tview.Table) {
    filtered := allProcessesInfo

    // Get the selected status option (index and label)
    selectedStatusIdx, _ := statusFilter.GetCurrentOption()
    selectedStatusStr := getStatusOptionLabel(selectedStatusIdx)

    // Filter by status
    if selectedStatusStr != "All" {
        var temp []ProcessInfo
        for _, pInfo := range filtered {
            if pInfo.Status == selectedStatusStr {
                temp = append(temp, pInfo)
            }
        }
        filtered = temp
    }

    // Get the selected running state option (index and label)
    selectedRunningIdx, _ := runningFilter.GetCurrentOption()
    selectedRunningStr := getRunningOptionLabel(selectedRunningIdx)

    // Filter by running state
    if selectedRunningStr != "All" {
        var temp []ProcessInfo
        for _, pInfo := range filtered {
            if (selectedRunningStr == "Running" && pInfo.IsRunning) ||
                (selectedRunningStr == "Stopped" && !pInfo.IsRunning) {
                temp = append(temp, pInfo)
            }
        }
        filtered = temp
    }

    // Get the selected username option (index and label)
    selectedUsernameIdx, _ := usernameFilter.GetCurrentOption()
    selectedUsernameStr := getUsernameOptionLabel(selectedUsernameIdx)

    // Filter by username
    if selectedUsernameStr != "All" {
        var temp []ProcessInfo
        for _, pInfo := range filtered {
            if pInfo.Username == selectedUsernameStr {
                temp = append(temp, pInfo)
            }
        }
        filtered = temp
    }

    // Update the displayed processes
    processesInfo = filtered
    populateTable(table, processesInfo)
}

// Helper functions to get the label from the index
func getStatusOptionLabel(index int) string {
    options := append([]string{"All"}, getUniqueValues(allProcessesInfo, "Status")...)
    if index >= 0 && index < len(options) {
        return options[index]
    }
    return "All"
}

func getRunningOptionLabel(index int) string {
    options := []string{"All", "Running", "Stopped"}
    if index >= 0 && index < len(options) {
        return options[index]
    }
    return "All"
}

func getUsernameOptionLabel(index int) string {
    options := append([]string{"All"}, getUniqueValues(allProcessesInfo, "Username")...)
    if index >= 0 && index < len(options) {
        return options[index]
    }
    return "All"
}


// getUniqueValues returns a list of unique values from the specified column of processesInfo.
func getUniqueValues(data []ProcessInfo, column string) []string {
	valueSet := make(map[string]struct{})

	for _, pInfo := range data {
		var value string
		switch column {
		case "Status":
			value = pInfo.Status
		case "Username":
			value = pInfo.Username
		}
		if _, exists := valueSet[value]; !exists {
			valueSet[value] = struct{}{}
		}
	}

	uniqueValues := make([]string, 0, len(valueSet))
	for value := range valueSet {
		uniqueValues = append(uniqueValues, value)
	}

	return uniqueValues
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
