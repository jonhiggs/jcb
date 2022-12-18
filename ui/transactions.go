package ui

import (
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var table *cview.Table
var transactionIds []int64
var initialBalance int64

func handleSelectNext(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := table.GetSelection()
	if table.GetRowCount() > r+1 {
		table.Select(r+1, 0)
	}

	return nil
}

func handleSelectPrev(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := table.GetSelection()
	table.Select(r-1, 0)
	return nil
}

func handleHalfPageDown(ev *tcell.EventKey) *tcell.EventKey {
	_, h := app.GetScreenSize()
	r, _ := table.GetSelection()

	if r+(h/2) < table.GetRowCount() {
		table.Select(r+(h/2), 0)
	} else {
		table.Select(table.GetRowCount()-1, 0)
	}

	return nil
}

func handleHalfPageUp(ev *tcell.EventKey) *tcell.EventKey {
	_, h := app.GetScreenSize()
	r, _ := table.GetSelection()

	if r-(h/2) > 0 {
		table.Select(r-(h/2), 0)
	} else {
		table.Select(0, 0)
	}

	return nil
}

func handleDeleteTransaction(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := table.GetSelection()
	transaction.Delete(transactionIds[r])
	updateTransactionsTable()
	table.Select(r, 0)
	return nil
}

func handleCommitTransaction(ev *tcell.EventKey) *tcell.EventKey {
	year := 2022
	r, _ := table.GetSelection()
	id := transactionIds[r]
	if transaction.IsCommitted(id, year) {
		transaction.Uncommit(id)
	} else {
		transaction.Commit(id, initialBalance, year)
		updateTransactionsTable()
	}
	return nil
}

func createTransactionsTable() *cview.Table {
	initialBalance = 0
	table = cview.NewTable()
	table.Select(0, 0)
	table.SetBorders(false)
	table.SetFixed(1, 1)
	table.SetSelectable(true, false)
	table.SetSeparator(' ')
	table.SetRect(0, 0, 72, 20)
	table.SetScrollBarVisibility(cview.ScrollBarNever)

	c := cbind.NewConfiguration()
	c.Set("i", handleOpenInsert)
	c.Set("j", handleSelectNext)
	c.Set("k", handleSelectPrev)
	c.SetRune(tcell.ModCtrl, 'd', handleHalfPageDown)
	c.SetRune(tcell.ModCtrl, 'u', handleHalfPageUp)
	c.Set("x", handleDeleteTransaction)
	c.Set("C", handleCommitTransaction)
	table.SetInputCapture(c.Capture)

	updateTransactionsTable()

	table.SetDoneFunc(func(key tcell.Key) {
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

func updateTransactionsTable() {
	year := 2022
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

	b := initialBalance
	transactionIds = make([]int64, len(all)+1)
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

		transactionIds[i+1] = t.Id
	}
}

func selectTransaction(id int64) {
	for i, v := range transactionIds {
		if v == id {
			table.Select(i, 0)
		}
	}

}
