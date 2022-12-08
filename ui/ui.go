package ui

import (
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var lowestBalance int64
var lowestBalanceDate time.Time

func Start(year int) {
	app = tview.NewApplication()

	table := TransactionMenu(year)

	box0 := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetText("═══════════════════════════════════════════════════════════════").
		SetTextAlign(tview.AlignCenter)

	balance := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetText("balance").
		SetTextAlign(tview.AlignRight)

	status := tview.NewTextArea().
		SetPlaceholder("status")

	grid := tview.NewGrid().
		SetRows(0, 1, 1).
		SetColumns(40, 20, 0).
		SetBorders(false).
		AddItem(table, 0, 0, 1, 2, 0, 0, true).
		AddItem(box0, 1, 0, 1, 2, 0, 0, true).
		AddItem(status, 2, 0, 1, 1, 0, 0, false).
		AddItem(balance, 2, 1, 1, 1, 0, 0, false)

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

func TransactionMenu(year int) *tview.Table {
	table := tview.NewTable().
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

	table.SetCell(0, 0, tview.NewTableCell("DATE").
		SetTextColor(tcell.ColorYellow).
		SetAttributes(tcell.AttrUnderline|tcell.AttrBold).
		SetSelectable(false).
		SetAlign(tview.AlignLeft))

	table.SetCell(0, 1, tview.NewTableCell("DESCRIPTION").
		SetTextColor(tcell.ColorYellow).
		SetAttributes(tcell.AttrUnderline|tcell.AttrBold).
		SetSelectable(false).
		SetAlign(tview.AlignLeft))

	table.SetCell(0, 2, tview.NewTableCell("AMOUNT").
		SetTextColor(tcell.ColorYellow).
		SetAttributes(tcell.AttrUnderline|tcell.AttrBold).
		SetSelectable(false).
		SetAlign(tview.AlignRight))

	table.SetCell(0, 3, tview.NewTableCell("BALANCE").
		SetTextColor(tcell.ColorYellow).
		SetAttributes(tcell.AttrUnderline|tcell.AttrBold).
		SetSelectable(false).
		SetAlign(tview.AlignRight))

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

		table.SetCell(i+1, 0, tview.NewTableCell(date).
			SetTextColor(color).
			SetAttributes(attributes).
			SetAlign(tview.AlignLeft))

		table.SetCell(i+1, 1, tview.NewTableCell(description).
			SetTextColor(color).
			SetAttributes(attributes).
			SetAlign(tview.AlignLeft))

		table.SetCell(i+1, 2, tview.NewTableCell(cents).
			SetTextColor(color).
			SetAttributes(attributes).
			SetAlign(tview.AlignRight))

		table.SetCell(i+1, 3, tview.NewTableCell(balance).
			SetTextColor(color).
			SetAttributes(attributes).
			SetAlign(tview.AlignRight))

	}

	return table
}

func Modal() *tview.Modal {
	return tview.NewModal().
		SetText("Do you want to quit the application?").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})
}
