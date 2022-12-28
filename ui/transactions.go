package ui

import (
	"fmt"
	"jcb/config"
	"jcb/domain"
	dataf "jcb/lib/formatter/data"
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"
	"strings"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var transactionsTable *cview.Table
var transactionIds []int64
var transactionAttributes []domain.Attributes
var initialBalance int64

func handleOpenTransactions(ev *tcell.EventKey) *tcell.EventKey {
	panels.ShowPanel("transactions")
	panels.HidePanel("report")
	panels.SendToFront("transactions")
	return nil
}

func handleSelectNext(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if transactionsTable.GetRowCount() > r+1 {
		transactionsTable.Select(r+1, 0)
	}

	return nil
}

func handleSelectPrev(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	transactionsTable.Select(r-1, 0)
	return nil
}

func handleHalfPageDown(ev *tcell.EventKey) *tcell.EventKey {
	_, h := app.GetScreenSize()
	r, _ := transactionsTable.GetSelection()

	if r+(h/2) < transactionsTable.GetRowCount() {
		transactionsTable.Select(r+(h/2), 0)
	} else {
		transactionsTable.Select(transactionsTable.GetRowCount()-1, 0)
	}

	return nil
}

func handleHalfPageUp(ev *tcell.EventKey) *tcell.EventKey {
	_, h := app.GetScreenSize()
	r, _ := transactionsTable.GetSelection()

	if r-(h/2) > 0 {
		transactionsTable.Select(r-(h/2), 0)
	} else {
		transactionsTable.Select(0, 0)
	}

	return nil
}

func handleSelectFirstUncommitted(ev *tcell.EventKey) *tcell.EventKey {
	uncommitted, _ := transaction.Uncommitted()
	if len(uncommitted) > 0 {
		firstUncommitted := uncommitted[0]

		for i, v := range transactionIds {
			if firstUncommitted.Id == v {
				transactionsTable.Select(i, 0)
				return nil
			}
		}
	}

	transactionsTable.Select(len(transactionIds)-1, 0)
	return nil
}

func handleSelectSimilar(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()
	curDescription := transactionsTable.GetCell(curRow, config.DESCRIPTION_COLUMN).GetText()

	for i := curRow + 1; i != curRow; i++ {
		if transactionsTable.GetCell(i, config.DESCRIPTION_COLUMN).GetText() == curDescription {
			transactionsTable.Select(i, 0)
			break
		}

		if i == len(transactionIds) {
			i = 0
		}
	}

	return nil
}

func handleSelectMonthNext(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()
	curMonth := dataf.Date(transactionsTable.GetCell(curRow, config.DATE_COLUMN).GetText()).Month()

	for i := curRow + 1; i < len(transactionIds); i++ {
		month := dataf.Date(transactionsTable.GetCell(i, config.DATE_COLUMN).GetText()).Month()
		if int(month) > int(curMonth) {
			transactionsTable.Select(i, 0)
			return nil
		}
	}

	transactionsTable.Select(len(transactionIds)-1, 0)

	return nil
}

func handleSelectMonthPrev(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()
	curMonth := dataf.Date(transactionsTable.GetCell(curRow, config.DATE_COLUMN).GetText()).Month()

	for i := curRow + 1; i > 0; i-- {
		month := dataf.Date(transactionsTable.GetCell(i, config.DATE_COLUMN).GetText()).Month()
		if int(month) < int(curMonth) {
			transactionsTable.Select(i, 0)
			return nil
		}
	}

	transactionsTable.Select(1, 0)

	return nil
}

func handleSelectYear(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()
	curYear := dataf.Date(transactionsTable.GetCell(curRow, config.DATE_COLUMN).GetText()).Year()

	if ev.Rune() == '<' {
		for i := curRow; i > 0; i-- {
			year := dataf.Date(transactionsTable.GetCell(i, config.DATE_COLUMN).GetText()).Year()
			if int(year) != int(curYear) {
				transactionsTable.Select(i, 0)
				return nil
			}
		}

		transactionsTable.Select(1, 0)
	} else {
		for i := curRow; i < len(transactionIds)-1; i++ {
			year := dataf.Date(transactionsTable.GetCell(i, config.DATE_COLUMN).GetText()).Year()
			if int(year) != int(curYear) {
				transactionsTable.Select(i, 0)
				return nil
			}
		}

		transactionsTable.Select(len(transactionIds)-1, 0)
	}

	return nil
}

func handleDeleteTransaction(ev *tcell.EventKey) *tcell.EventKey {
	id := selectionId()

	curRow, _ := transactionsTable.GetSelection()
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

	transactionsTable.RemoveRow(curRow)
	removeTag(curRow)
	updateTransactionsTable()
	transactionsTable.Select(r, 0)

	return nil
}

func handleCommitTransaction(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
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
	r, _ := transactionsTable.GetSelection()
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

func handleEditCents(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if isCommitted(r) {
		printStatus("Cannot edit committed transactions")
		return nil
	}

	openPrompt("Amount:", selectedAmount(), func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		r, _ := transactionsTable.GetSelection()
		updateCents(promptInputField.GetText(), []int{r})
		return nil
	})

	return nil
}

func handleEditCategory(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if isCommitted(r) {
		printStatus("Cannot edit committed transactions")
		return nil
	}

	openPrompt("Category:", selectedCategory(), func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		r, _ := transactionsTable.GetSelection()
		updateCategory(promptInputField.GetText(), []int{r})
		return nil
	})

	return nil
}

func handleEditDescription(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if isCommitted(r) {
		printStatus("Cannot edit committed transactions")
		return nil
	}

	openPrompt("Description:", selectedDescription(), func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		r, _ := transactionsTable.GetSelection()
		updateDescription(promptInputField.GetText(), []int{r})
		return nil
	})

	return nil
}

func handleEditDate(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if isCommitted(r) {
		printStatus("Cannot edit committed transactions")
		return nil
	}

	openPrompt("Date:", selectedDate(), func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		r, _ := transactionsTable.GetSelection()
		updateDate(promptInputField.GetText(), []int{r})
		return nil
	})

	return nil
}

func createTransactionsTable() *cview.Table {
	initialBalance = 0
	transactionsTable = cview.NewTable()
	transactionsTable.Select(0, 0)
	transactionsTable.SetBorders(false)
	transactionsTable.SetFixed(1, 1)
	transactionsTable.SetSelectable(true, false)
	transactionsTable.SetSeparator(' ')
	transactionsTable.SetRect(0, 0, config.MAX_WIDTH, 20)
	transactionsTable.SetScrollBarVisibility(cview.ScrollBarNever)
	transactionsTable.SetSelectionChangedFunc(func(r int, c int) { handleCloseStatus() })

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
	c.Set("t", handleTag)
	c.Set("C", handleCommitTransaction)
	c.Set("c", handleCommitSingleTransaction)
	c.Set(":", handleCommand)
	c.Set(";", handleTagCommand)
	c.Set("/", handleOpenFind)
	c.Set("?", handleOpenFind)
	c.Set("T", handleOpenFind)
	c.Set("n", handleSelectNextMatch)
	c.Set("N", handleSelectPrevMatch)
	c.Set(">", handleSelectYear)
	c.Set("<", handleSelectYear)
	c.Set("F1", handleOpenHelp)
	c.Set("F2", handleOpenTransactions)
	c.Set("F3", handleOpenReport)
	c.Set("=", handleEditCents)
	c.Set("D", handleEditCategory)
	c.Set("d", handleEditDescription)
	c.Set("@", handleEditDate)
	transactionsTable.SetInputCapture(c.Capture)

	updateTransactionsTable()

	transactionsTable.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			transactionsTable.SetSelectable(true, true)
		}
	})

	transactionsTable.SetSelectedFunc(func(row int, column int) {
		handleOpenEdit()
	})

	return transactionsTable
}

func updateTransactionsTable() {
	committed, _ := transaction.Committed()
	uncommitted, _ := transaction.Uncommitted()
	all := committed
	for _, t := range uncommitted {
		all = append(all, t)
	}

	var cell *cview.TableCell

	cell = cview.NewTableCell("")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignRight)
	transactionsTable.SetCell(0, config.ATTR_COLUMN, cell)

	cell = cview.NewTableCell("DATE")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignLeft)
	transactionsTable.SetCell(0, config.DATE_COLUMN, cell)

	cell = cview.NewTableCell("CATEGORY")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignLeft)
	transactionsTable.SetCell(0, config.CATEGORY_COLUMN, cell)

	cell = cview.NewTableCell("DESCRIPTION")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignLeft)
	transactionsTable.SetCell(0, config.DESCRIPTION_COLUMN, cell)

	cell = cview.NewTableCell("AMOUNT")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignRight)
	transactionsTable.SetCell(0, config.AMOUNT_COLUMN, cell)

	cell = cview.NewTableCell("BALANCE")
	cell.SetTextColor(tcell.ColorYellow)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignRight)
	transactionsTable.SetCell(0, config.BALANCE_COLUMN, cell)

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

		if isTagged(i + 1) {
			color = tcell.ColorGreen
		} else if isCommitted {
			color = tcell.ColorWhite
			attributes = 0
		} else if b < 0 {
			color = tcell.ColorRed
		} else {
			color = tcell.ColorBlue
			attributes = tcell.AttrBold
		}

		transactionIds[i+1] = t.Id
		transactionAttributes[i+1] = transaction.Attributes(t.Id)

		cell = cview.NewTableCell(stringf.Attributes(transactionAttributes[i+1]))
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		transactionsTable.SetCell(i+1, config.ATTR_COLUMN, cell)

		cell = cview.NewTableCell(fmt.Sprintf("%-10s", stringf.Category(t.Category)))
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		transactionsTable.SetCell(i+1, config.CATEGORY_COLUMN, cell)

		cell = cview.NewTableCell(date)
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignLeft)
		transactionsTable.SetCell(i+1, config.DATE_COLUMN, cell)

		if len(description) > config.DESC_MAX_LENGTH {
			description = description[0:config.DESC_MAX_LENGTH]
		}
		cell = cview.NewTableCell(fmt.Sprintf("%-*s", config.DESC_MAX_LENGTH, description))
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignLeft)
		transactionsTable.SetCell(i+1, config.DESCRIPTION_COLUMN, cell)

		cell = cview.NewTableCell(fmt.Sprintf("%10s", cents))
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignRight)
		transactionsTable.SetCell(i+1, config.AMOUNT_COLUMN, cell)

		cell = cview.NewTableCell(fmt.Sprintf("%10s", balance))
		cell.SetTextColor(color)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignRight)
		transactionsTable.SetCell(i+1, config.BALANCE_COLUMN, cell)
	}
}

// select transaction by id
func selectTransaction(id int64) {
	for i, v := range transactionIds {
		if v == id {
			transactionsTable.Select(i, 0)
		}
	}

}

// get the id of the selection
func selectionId() int64 {
	r, _ := transactionsTable.GetSelection()
	return transactionIds[r]
}

func isCommitted(r int) bool {
	if transactionsTable.GetCell(r, config.ATTR_COLUMN).GetText()[0:1] == "C" {
		return true
	} else {
		return false
	}
}

func selectedAmount() string {
	r, _ := transactionsTable.GetSelection()
	return strings.Trim(transactionsTable.GetCell(r, config.AMOUNT_COLUMN).GetText(), " ")
}

func selectedCategory() string {
	r, _ := transactionsTable.GetSelection()
	return strings.Trim(transactionsTable.GetCell(r, config.CATEGORY_COLUMN).GetText(), " ")
}

func selectedDescription() string {
	r, _ := transactionsTable.GetSelection()
	return strings.Trim(transactionsTable.GetCell(r, config.DESCRIPTION_COLUMN).GetText(), " ")
}

func selectedDate() string {
	r, _ := transactionsTable.GetSelection()
	return strings.Trim(transactionsTable.GetCell(r, config.DATE_COLUMN).GetText(), " ")
}
