package ui

import (
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

func createTransactionsTable() *cview.Table {
	year := 2022
	table := cview.NewTable()
	table.Select(0, 0)
	table.SetBorders(false)
	table.SetFixed(1, 1)
	table.SetSelectable(true, false)
	table.SetSeparator(' ')
	table.SetRect(0, 0, 72, 20)

	c := cbind.NewConfiguration()
	c.Set("i", handleOpenInsert)
	table.SetInputCapture(c.Capture)

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

	return table
}
