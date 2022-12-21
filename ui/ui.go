package ui

import (
	"time"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var app *cview.Application
var panels *cview.Panels
var lowestBalance int64
var lowestBalanceDate time.Time
var find *cview.TextView

func Start() {
	app = cview.NewApplication()

	box0 := cview.NewTextView()
	box0.SetDynamicColors(true)
	box0.SetRegions(true)
	box0.SetWordWrap(true)
	box0.SetText("")
	box0.SetTextAlign(cview.AlignCenter)

	balance := cview.NewTextView()
	balance.SetDynamicColors(true)
	balance.SetRegions(true)
	balance.SetWordWrap(true)
	balance.SetText("balance")
	balance.SetTextAlign(cview.AlignRight)

	panels = cview.NewPanels()
	panels.AddPanel("transactions", createTransactionsTable(), false, true)
	panels.AddPanel("insert", createInsertForm(), false, false)
	panels.AddPanel("edit", createEditForm(), false, false)
	panels.AddPanel("find", createFindForm(), false, false)
	panels.AddPanel("command", createCommandForm(), false, false)
	panels.AddPanel("status", createStatusTextView(), false, false)
	panels.AddPanel("help", createHelp(), false, false)

	grid := cview.NewGrid()
	grid.SetRows(0)
	grid.SetColumns(72, 0)
	grid.SetBorders(false)
	grid.AddItem(panels, 0, 0, 1, 1, 0, 0, true)

	c := cbind.NewConfiguration()
	handleExit := func(ev *tcell.EventKey) *tcell.EventKey {
		pn, _ := panels.GetFrontPanel()
		if pn == "transactions" {
			printStatus("To quit, use the command ':q'.")
		} else {
			handleCloseInsert()
			handleCloseEdit()
			handleCloseFind()
			handleCloseCommand()
			handleCloseHelp()
		}
		return nil
	}

	c.SetRune(tcell.ModCtrl, 'c', handleExit)

	app.SetInputCapture(c.Capture)

	app.SetAfterResizeFunc(func(w int, h int) {
		table.SetRect(0, 0, 72, h-1)
		status.SetRect(0, h-1, 72, h)
		helpTextView.SetRect(0, 0, 72, h-1)
		findForm.SetRect(0, h-1, 72, h)
		commandForm.SetRect(0, h-1, 72, h)
	})

	app.SetRoot(grid, true)
	app.Run()
}
