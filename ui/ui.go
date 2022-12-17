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
var status *cview.TextView

func Start(year int) {
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

	status = cview.NewTextView()
	status.SetText("status")

	grid := cview.NewGrid()
	grid.SetRows(0, 1, 1)
	grid.SetColumns(40, 20, 0)
	grid.SetBorders(false)
	grid.AddItem(panels, 0, 0, 1, 2, 0, 0, true)
	grid.AddItem(box0, 1, 0, 1, 2, 0, 0, true)
	grid.AddItem(status, 2, 0, 1, 1, 0, 0, false)
	grid.AddItem(balance, 2, 1, 1, 1, 0, 0, false)

	c := cbind.NewConfiguration()
	handleExit := func(ev *tcell.EventKey) *tcell.EventKey {
		pn, _ := panels.GetFrontPanel()
		if pn == "transactions" {
			app.Stop()
		} else {
			handleCloseInsert()
		}
		return nil
	}

	c.SetRune(tcell.ModCtrl, 'c', handleExit)

	app.SetInputCapture(c.Capture)

	//app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	//	switch event.Key() {
	//	case tcell.KeyCtrlD:
	//		status.SetText("hi")
	//	case tcell.KeyCtrlC:
	//		status.SetText("worked")
	//	}
	//	return event
	//})

	app.SetRoot(grid, true)
	app.Run()
}
