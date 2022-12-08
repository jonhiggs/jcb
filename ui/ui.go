package ui

import (
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Start(year int) {
	app := tview.NewApplication()

	table := TransactionMenu(year)

	frame := tview.NewFrame(table).
		SetBorders(0, 0, 0, 0, 0, 0)

	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})
	if err := app.SetRoot(frame, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func TransactionMenu(year int) *tview.Table {
	table := tview.NewTable().
		SetBorders(false)

	table.SetSelectable(true, false)

	committed, _ := transaction.Committed(year)
	uncommitted, _ := transaction.Uncommitted(year)
	all := committed
	for _, t := range uncommitted {
		all = append(all, t)
	}

	table.SetCell(0, 0, tview.NewTableCell("DATE ").
		SetTextColor(tcell.ColorYellow).
		SetSelectable(false).
		SetAlign(tview.AlignLeft))

	table.SetCell(0, 1, tview.NewTableCell("DESCRIPTION ").
		SetTextColor(tcell.ColorYellow).
		SetSelectable(false).
		SetAlign(tview.AlignLeft))

	table.SetCell(0, 2, tview.NewTableCell("AMOUNT ").
		SetTextColor(tcell.ColorYellow).
		SetSelectable(false).
		SetAlign(tview.AlignRight))

	table.SetCell(0, 3, tview.NewTableCell("BALANCE ").
		SetTextColor(tcell.ColorYellow).
		SetSelectable(false).
		SetAlign(tview.AlignRight))

	b := int64(0)
	for i, t := range all {
		b += t.Cents
		date, _ := stringf.Date(t.Date)
		description, _ := stringf.Description(t.Description)
		cents, _ := stringf.Cents(t.Cents)
		balance, _ := stringf.Cents(b)

		table.SetCell(i+1, 0, tview.NewTableCell(date).
			SetTextColor(tcell.Color245).
			SetAlign(tview.AlignLeft))

		table.SetCell(i+1, 1, tview.NewTableCell(description).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

		var centsColor tcell.Color
		if t.Cents < 0 {
			centsColor = tcell.ColorRed
		} else {
			centsColor = tcell.ColorWhite
		}

		table.SetCell(i+1, 2, tview.NewTableCell(cents).
			SetTextColor(centsColor).
			SetAlign(tview.AlignRight))

		var balanceColor tcell.Color
		if b < 0 {
			balanceColor = tcell.ColorRed
		} else {
			balanceColor = tcell.ColorWhite
		}
		table.SetCell(i+1, 3, tview.NewTableCell(balance).
			SetTextColor(balanceColor).
			SetAlign(tview.AlignRight))

	}

	return table
}
