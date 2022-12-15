package ui

import (
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"
	"log"
	"time"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var app *cview.Application
var panels *cview.Panels
var lowestBalance int64
var lowestBalanceDate time.Time

func Start(year int) {
	app = cview.NewApplication()

	table := TransactionMenu(year)
	table.SetRect(0, 0, 72, 20)

	box0 := cview.NewTextView()
	box0.SetDynamicColors(true)
	box0.SetRegions(true)
	box0.SetWordWrap(true)
	box0.SetText("════════════════════════════════════════════════════════════")
	box0.SetTextAlign(cview.AlignCenter)

	balance := cview.NewTextView()
	balance.SetDynamicColors(true)
	balance.SetRegions(true)
	balance.SetWordWrap(true)
	balance.SetText("balance")
	balance.SetTextAlign(cview.AlignRight)

	panels = cview.NewPanels()
	panels.AddPanel("balance", table, false, true)
	panels.AddPanel("insert", createInsertForm(), false, false)

	status := cview.NewTextView()
	status.SetText("status")

	grid := cview.NewGrid()
	grid.SetRows(0, 1, 1)
	grid.SetColumns(40, 20, 0)
	grid.SetBorders(false)
	grid.AddItem(panels, 0, 0, 1, 2, 0, 0, true)
	grid.AddItem(box0, 1, 0, 1, 2, 0, 0, true)
	grid.AddItem(status, 2, 0, 1, 1, 0, 0, false)
	grid.AddItem(balance, 2, 1, 1, 1, 0, 0, false)

	table.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}

		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	})
	table.SetSelectedFunc(func(row int, column int) {
		panels.ShowPanel("insert")
		//table.GetCell(row, column).SetTextColor(tcell.ColorRed.TrueColor())
		//table.SetSelectable(false, false)
	})

	c := cbind.NewConfiguration()

	if err := c.Set("i", handleOpenInsert); err != nil {
		log.Fatalf("failed to set keybind: %s", err)
	}

	app.SetInputCapture(c.Capture)
	app.SetRoot(grid, true)
	app.Run()
}

func TransactionMenu(year int) *cview.Table {
	table := cview.NewTable()
	table.Select(0, 0)
	table.SetBorders(false)
	table.SetFixed(1, 1)
	table.SetSelectable(true, false)
	table.SetSeparator(' ')

	committed, _ := transaction.Committed(year)
	uncommitted, _ := transaction.Uncommitted(year)
	all := committed
	for _, t := range uncommitted {
		all = append(all, t)
	}

	cell := cview.NewTableCell("DATE")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignLeft)
	table.SetCell(0, 0, cell)

	cell = cview.NewTableCell("DESCRIPTION")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignLeft)
	table.SetCell(0, 1, cell)

	cell = cview.NewTableCell("AMOUNT")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignRight)
	table.SetCell(0, 2, cell)

	cell = cview.NewTableCell("BALANCE")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignRight)
	table.SetCell(0, 3, cell)

	b := int64(0)
	for i, t := range all {
		b += t.Cents
		date, _ := stringf.Date(t.Date)
		description, _ := stringf.Description(t.Description)
		cents, _ := stringf.Cents(t.Cents)
		balance, _ := stringf.Cents(b)
		isCommitted := false

		for _, ct := range committed {
			if ct.Id == t.Id {
				isCommitted = true
			}
		}

		var color tcell.Color
		var attributes tcell.AttrMask
		if b < 0 {
			color = tcell.ColorRed
		} else if isCommitted {
			color = tcell.ColorWhite
			attributes = 0
		} else {
			color = tcell.ColorBlue
			attributes = tcell.AttrBold
		}

		cell = cview.NewTableCell(date)
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignLeft)
		table.SetCell(i+1, 0, cell)

		cell = cview.NewTableCell(description)
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignLeft)
		table.SetCell(i+1, 1, cell)

		cell = cview.NewTableCell(cents)
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignRight)
		table.SetCell(i+1, 2, cell)

		cell = cview.NewTableCell(balance)
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignRight)
		table.SetCell(i+1, 3, cell)
	}

	return table
}

func Window() *cview.WindowManager {
	list := cview.NewList()
	list.ShowSecondaryText(false)
	list.AddItem(cview.NewListItem("Item #1"))
	list.AddItem(cview.NewListItem("Item #2"))
	list.AddItem(cview.NewListItem("Item #3"))
	list.AddItem(cview.NewListItem("Item #4"))
	list.AddItem(cview.NewListItem("Item #5"))
	list.AddItem(cview.NewListItem("Item #6"))
	list.AddItem(cview.NewListItem("Item #7"))

	loremIpsum := cview.NewTextView()
	loremIpsum.SetText("balh blah blha")

	wm := cview.NewWindowManager()
	w1 := cview.NewWindow(list)
	w1.SetRect(2, 2, 10, 7)
	w2 := cview.NewWindow(loremIpsum)
	w2.SetRect(7, 4, 12, 12)
	w1.SetTitle("List")
	w2.SetTitle("Lorem Ipsum")
	wm.Add(w1, w2)

	return wm
}

func Box() *cview.Box {
	box := cview.NewBox()
	box.SetBorder(true)
	box.SetBorderAttributes(tcell.AttrBold)
	box.SetRect(10, 10, 20, 20)
	box.SetTitle("New Transaction")

	return box
}
