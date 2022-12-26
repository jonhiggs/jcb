package ui

import (
	"fmt"
	"jcb/config"
	"jcb/domain"
	dataf "jcb/lib/formatter/data"
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var table *cview.Table
var transactionIds []int64
var transactionAttributes []domain.Attributes
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

func handleSelectFirstUncommitted(ev *tcell.EventKey) *tcell.EventKey {
	uncommitted, _ := transaction.Uncommitted()
	if len(uncommitted) > 0 {
		firstUncommitted := uncommitted[0]

		for i, v := range transactionIds {
			if firstUncommitted.Id == v {
				table.Select(i, 0)
				return nil
			}
		}
	}

	table.Select(len(transactionIds)-1, 0)
	return nil
}

func handleSelectSimilar(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := table.GetSelection()
	curDescription := table.GetCell(curRow, 2).GetText()

	for i := curRow + 1; i != curRow; i++ {
		if table.GetCell(i, 2).GetText() == curDescription {
			table.Select(i, 0)
			break
		}

		if i == len(transactionIds) {
			i = 0
		}
	}

	return nil
}

func handleSelectMonthNext(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := table.GetSelection()
	curMonth := dataf.Date(table.GetCell(curRow, 1).GetText()).Month()

	for i := curRow + 1; i < len(transactionIds); i++ {
		month := dataf.Date(table.GetCell(i, 1).GetText()).Month()
		if int(month) > int(curMonth) {
			table.Select(i, 0)
			return nil
		}
	}

	table.Select(len(transactionIds)-1, 0)

	return nil
}

func handleSelectMonthPrev(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := table.GetSelection()
	curMonth := dataf.Date(table.GetCell(curRow, 1).GetText()).Month()

	for i := curRow + 1; i > 0; i-- {
		month := dataf.Date(table.GetCell(i, 1).GetText()).Month()
		if int(month) < int(curMonth) {
			table.Select(i, 0)
			return nil
		}
	}

	table.Select(1, 0)

	return nil
}

func handleSelectYear(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := table.GetSelection()
	curYear := dataf.Date(table.GetCell(curRow, 1).GetText()).Year()

	if ev.Rune() == '<' {
		for i := curRow; i > 0; i-- {
			year := dataf.Date(table.GetCell(i, 1).GetText()).Year()
			if int(year) != int(curYear) {
				table.Select(i, 0)
				return nil
			}
		}

		table.Select(1, 0)
	} else {
		for i := curRow; i < len(transactionIds)-1; i++ {
			year := dataf.Date(table.GetCell(i, 1).GetText()).Year()
			if int(year) != int(curYear) {
				table.Select(i, 0)
				return nil
			}
		}

		table.Select(len(transactionIds)-1, 0)
	}

	return nil
}

func handleDeleteTransaction(ev *tcell.EventKey) *tcell.EventKey {
	id := selectionId()

	curRow, _ := table.GetSelection()
	var r int
	if curRow == len(transactionIds)-1 {
		r = curRow - 1
	} else {
		r = curRow
	}

	err := transaction.Delete(id)
	if err != nil {
		printStatus(fmt.Sprint(err))
		return nil
	}

	table.RemoveRow(curRow)
	updateTransactionsTable()
	table.Select(r, 0)

	return nil
}

func handleCommitTransaction(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := table.GetSelection()
	id := transactionIds[r]

	if transaction.Attributes(id).Committed {
		transaction.Uncommit(id)
	} else {
		transaction.Commit(id, initialBalance)
	}
	updateTransactionsTable()
	return nil
}

func handleCommitSingleTransaction(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := table.GetSelection()
	id := transactionIds[r]

	var err error
	if transaction.Attributes(id).Committed {
		err = transaction.UncommitSingle(id)
	} else {
		err = transaction.CommitSingle(id)
	}

	if err != nil {
		printStatus(fmt.Sprint(err))
		return nil
	}

	updateTransactionsTable()
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
	table.SetSelectionChangedFunc(func(r int, c int) { handleCloseStatus() })

	c := cbind.NewConfiguration()
	c.Set("i", handleOpenInsert)
	c.Set("j", handleSelectNext)
	c.Set("k", handleSelectPrev)
	c.SetRune(tcell.ModCtrl, 'd', handleHalfPageDown)
	c.SetRune(tcell.ModCtrl, 'u', handleHalfPageUp)
	c.Set("0", handleSelectFirstUncommitted)
	c.Set("*", handleSelectSimilar)
	c.Set("}", handleSelectMonthNext)
	c.Set("{", handleSelectMonthPrev)
	c.Set("x", handleDeleteTransaction)
	c.Set("r", handleOpenRepeat)
	c.Set("C", handleCommitTransaction)
	c.Set("c", handleCommitSingleTransaction)
	c.Set(":", handleOpenCommand)
	c.Set("/", handleOpenFind)
	c.Set("?", handleOpenFind)
	c.Set("n", handleSelectNextMatch)
	c.Set("N", handleSelectPrevMatch)
	c.Set(">", handleSelectYear)
	c.Set("<", handleSelectYear)
	table.SetInputCapture(c.Capture)

	updateTransactionsTable()

	table.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	})

	table.SetSelectedFunc(func(row int, column int) {
		handleOpenEdit()
	})

	return table
}

func updateTransactionsTable() {
	committed, _ := transaction.Committed()
	uncommitted, _ := transaction.Uncommitted()
	all := committed
	for _, t := range uncommitted {
		all = append(all, t)
	}

	cell := cview.NewTableCell("")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignRight)
	table.SetCell(0, 0, cell)

	cell = cview.NewTableCell("DATE")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignLeft)
	table.SetCell(0, 1, cell)

	cell = cview.NewTableCell("DESCRIPTION")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignLeft)
	table.SetCell(0, 2, cell)

	cell = cview.NewTableCell("AMOUNT")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignRight)
	table.SetCell(0, 3, cell)

	cell = cview.NewTableCell("BALANCE")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignRight)
	table.SetCell(0, 4, cell)

	b := initialBalance
	transactionIds = make([]int64, len(all)+1)
	transactionAttributes = make([]domain.Attributes, len(all)+1)
	for i, t := range all {
		b += t.Cents
		date := stringf.Date(t.Date)
		description := stringf.Description(t.Description)
		cents := stringf.Cents(t.Cents)
		balance := stringf.Cents(b)
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

		transactionIds[i+1] = t.Id
		transactionAttributes[i+1] = transaction.Attributes(t.Id)

		cell = cview.NewTableCell(stringf.Attributes(transactionAttributes[i+1]))
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignRight)
		table.SetCell(i+1, 0, cell)

		cell = cview.NewTableCell(date)
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignLeft)
		table.SetCell(i+1, 1, cell)

		if len(description) > config.DESC_MAX_LENGTH {
			description = description[0:config.DESC_MAX_LENGTH]
		}
		cell = cview.NewTableCell(fmt.Sprintf("%-*s", config.DESC_MAX_LENGTH, description))
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignLeft)
		table.SetCell(i+1, 2, cell)

		cell = cview.NewTableCell(fmt.Sprintf("%10s", cents))
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignRight)
		table.SetCell(i+1, 3, cell)

		cell = cview.NewTableCell(fmt.Sprintf("%10s", balance))
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignRight)
		table.SetCell(i+1, 4, cell)
	}
}

func selectTransaction(id int64) {
	for i, v := range transactionIds {
		if v == id {
			table.Select(i, 0)
		}
	}

}

func selectionId() int64 {
	r, _ := table.GetSelection()
	return transactionIds[r]
}
