package ui

import (
	"fmt"
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"
	"time"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var app *cview.Application
var lowestBalance int64
var lowestBalanceDate time.Time

func Start(year int) {
	app = cview.NewApplication()

	table := TransactionMenu(year)

	box0 := cview.NewTextView()
	box0.SetDynamicColors(true)
	box0.SetRegions(true)
	box0.SetWordWrap(true)
	box0.SetText("═══════════════════════════════════════════════════════════════")
	box0.SetTextAlign(cview.AlignCenter)

	balance := cview.NewTextView()
	balance.SetDynamicColors(true)
	balance.SetRegions(true)
	balance.SetWordWrap(true)
	balance.SetText("balance")
	balance.SetTextAlign(cview.AlignRight)

	status := cview.NewTextView()
	fmt.Printf(status, "status")

	grid := cview.NewGrid()
	grid.SetRows(0, 1, 1)
	grid.SetColumns(40, 20, 0)
	grid.SetBorders(false)
	grid.AddItem(table, 0, 0, 1, 2, 0, 0, true)
	grid.AddItem(box0, 1, 0, 1, 2, 0, 0, true)
	grid.AddItem(status, 2, 0, 1, 1, 0, 0, false)
	grid.AddItem(balance, 2, 1, 1, 1, 0, 0, false)

	table.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			Modal().Focus()
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})

	err := app.SetRoot(grid, true).
		EnableMouse(true).
		Run()

	if err != nil {
		panic(err)
	}
}

func TransactionMenu(year int) *cview.Table {
	table := cview.NewTable().
		Select(0, 0).
		SetBorders(false).
		SetFixed(1, 1).
		SetSelectable(true, false).
		SetSeparator(' ')

	committed, _ := transaction.Committed(year)
	uncommitted, _ := transaction.Uncommitted(year)
	all := committed
	for _, t := range uncommitted {
		all = append(all, t)
	}

	table.SetCell(0, 0, cview.NewTableCell("DATE").
		SetTextColor(tcell.ColorYellow).
		SetAttributes(tcell.AttrUnderline|tcell.AttrBold).
		SetSelectable(false).
		SetAlign(cview.AlignLeft))

	table.SetCell(0, 1, cview.NewTableCell("DESCRIPTION").
		SetTextColor(tcell.ColorYellow).
		SetAttributes(tcell.AttrUnderline|tcell.AttrBold).
		SetSelectable(false).
		SetAlign(cview.AlignLeft))

	table.SetCell(0, 2, cview.NewTableCell("AMOUNT").
		SetTextColor(tcell.ColorYellow).
		SetAttributes(tcell.AttrUnderline|tcell.AttrBold).
		SetSelectable(false).
		SetAlign(cview.AlignRight))

	table.SetCell(0, 3, cview.NewTableCell("BALANCE").
		SetTextColor(tcell.ColorYellow).
		SetAttributes(tcell.AttrUnderline|tcell.AttrBold).
		SetSelectable(false).
		SetAlign(cview.AlignRight))

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

		table.SetCell(i+1, 0, cview.NewTableCell(date).
			SetTextColor(color).
			SetAttributes(attributes).
			SetAlign(cview.AlignLeft))

		table.SetCell(i+1, 1, cview.NewTableCell(description).
			SetTextColor(color).
			SetAttributes(attributes).
			SetAlign(cview.AlignLeft))

		table.SetCell(i+1, 2, cview.NewTableCell(cents).
			SetTextColor(color).
			SetAttributes(attributes).
			SetAlign(cview.AlignRight))

		table.SetCell(i+1, 3, cview.NewTableCell(balance).
			SetTextColor(color).
			SetAttributes(attributes).
			SetAlign(cview.AlignRight))

	}

	return table
}

func Modal() *cview.Modal {
	return cview.NewModal().
		SetText("Do you want to quit the application?").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})
}
