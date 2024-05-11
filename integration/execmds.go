package main

import (
	"fmt"
	"os/exec"
	"dcm/tviewplay"
	"github.com/rivo/tview"
)

func main(){
	Execmds()
}

func Execmds() {
    // Commands successfully executed: 
	// ("ls")
	// ("ps aux")
	// ("sh", "-c", "cat /sys/class/thermal/thermal_zone*/temp")
	// ("sh", "-c", "cat /sys/class/thermal/thermal_zone*/type")

	// tempValsCmd := exec.Command("sh", "-c", "cat /sys/class/thermal/thermal_zone*/temp")
    // tempValsOutput, err := tempValsCmd.Output()
    // if err != nil {
    //     fmt.Println("Error executing command:", err)
    //     return
    // }
	// tempValResult := string(tempValsOutput)
    // tviewplay.TextView(tempValResult)

	tempSrcCmd := exec.Command("sh", "-c", "cat /sys/class/thermal/thermal_zone*/type")
    tempSrcOutput, err := tempSrcCmd.Output()
    if err != nil {
        fmt.Println("Error executing command:", err)
        return
    }
	tempSrcResult := string(tempSrcOutput)
    tviewplay.TextView(tempSrcResult)
}


func Grid() {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	temperature := newPrimitive("Hardware Temperature")
	graph := newPrimitive("Memory Usage")

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("dcm"), 0, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(temperature, 0, 0, 0, 0, 0, 0, false).
		AddItem(graph, 0, 0, 0, 0, 0, 0, false)

	// func (*tview.Grid).AddItem(p tview.Primitive, 
	// row int, 
	// column int, 
	// rowSpan int, 
	// colSpan int, 
	// minGridHeight int, 
	// minGridWidth int, 
	// focus bool
	// ) *tview.Grid

	// Layout for screens wider than 100 cells.
	grid.AddItem(temperature, 1, 0, 1, 1, 0, 100, false).
		AddItem(graph, 1, 1, 1, 1, 0, 100, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
