package ui

import (
	"fmt"
	"jcb/config"
	"jcb/lib/transaction"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var transactionsTable *cview.Table
var transactions []*transaction.Transaction

func createTransactionsTable() *cview.Table {
	start, end := transaction.DateRange()
	transactions = transaction.All(start, end)

	transactionsTable = cview.NewTable()
	transactionsTable.Select(0, 0)
	transactionsTable.SetBorders(false)
	transactionsTable.SetFixed(1, 1)
	transactionsTable.SetSelectable(true, false)
	transactionsTable.SetSeparator(' ')
	transactionsTable.SetRect(0, 0, config.MAX_WIDTH, 20)
	transactionsTable.SetScrollBarVisibility(cview.ScrollBarNever)
	transactionsTable.SetSelectionChangedFunc(func(r int, c int) {
		closeStatus()
		updateInfo()
	})

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
	c.Set("]", handleSelectModifiedNext)
	c.Set("[", handleSelectModifiedPrev)
	c.Set("x", handleDeleteTransaction)
	c.Set("r", handleRepeat)
	c.Set("t", handleTagToggle)
	c.Set("C", handleCommitTransaction)
	c.Set("c", handleCommitSingleTransaction)
	c.Set(":", handleCommand)
	c.Set(";", handleTagCommand)
	c.Set("/", handleFindForwards)
	c.Set("?", handleFindBackwards)
	c.Set("T", handleTagMatches)
	c.SetRune(tcell.ModCtrl, 'T', handleUntagMatches)
	c.Set("n", handleSelectMatchNext)
	c.Set("N", handleSelectMatchPrev)
	c.Set(">", handleSelectYearNext)
	c.Set("<", handleSelectYearPrev)
	c.Set("F1", handleOpenHelp)
	c.Set("F2", handleOpenTransactions)
	c.Set("F3", handleOpenReport)
	c.Set("=", handleEditSingleTransaction)
	c.Set("D", handleEditSingleTransaction)
	c.Set("d", handleEditSingleTransaction)
	c.Set("@", handleEditSingleTransaction)
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
	var cell *cview.TableCell

	cell = cview.NewTableCell("")
	cell.SetTextColor(config.COLOR_TITLE_FG)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignRight)
	cell.SetBackgroundColor(config.COLOR_TITLE_BG)
	transactionsTable.SetCell(0, config.ATTR_COLUMN, cell)

	cell = cview.NewTableCell("DATE")
	cell.SetTextColor(config.COLOR_TITLE_FG)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignLeft)
	cell.SetBackgroundColor(config.COLOR_TITLE_BG)
	transactionsTable.SetCell(0, config.DATE_COLUMN, cell)

	cell = cview.NewTableCell("CATEGORY")
	cell.SetTextColor(config.COLOR_TITLE_FG)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignLeft)
	cell.SetBackgroundColor(config.COLOR_TITLE_BG)
	transactionsTable.SetCell(0, config.CATEGORY_COLUMN, cell)

	cell = cview.NewTableCell("DESCRIPTION")
	cell.SetTextColor(config.COLOR_TITLE_FG)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignLeft)
	cell.SetBackgroundColor(config.COLOR_TITLE_BG)
	transactionsTable.SetCell(0, config.DESCRIPTION_COLUMN, cell)

	cell = cview.NewTableCell("AMOUNT")
	cell.SetTextColor(config.COLOR_TITLE_FG)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignRight)
	cell.SetBackgroundColor(config.COLOR_TITLE_BG)
	transactionsTable.SetCell(0, config.AMOUNT_COLUMN, cell)

	cell = cview.NewTableCell("BALANCE")
	cell.SetTextColor(config.COLOR_TITLE_FG)
	cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
	cell.SetSelectable(false)
	cell.SetAlign(cview.AlignRight)
	cell.SetBackgroundColor(config.COLOR_TITLE_BG)
	transactionsTable.SetCell(0, config.BALANCE_COLUMN, cell)

	for i, t := range transactions {

		var colorFg tcell.Color
		var colorBg tcell.Color
		var attributes tcell.AttrMask

		if t.IsCommitted() {
			colorFg = config.COLOR_COMMITTED_FG
			colorBg = config.COLOR_COMMITTED_BG
			attributes = 0
		} else {
			colorFg = config.COLOR_UNCOMMITTED_FG
			colorBg = config.COLOR_UNCOMMITTED_BG
		}

		if !t.IsSaved() {
			colorFg = config.COLOR_MODIFIED_FG
			colorBg = config.COLOR_MODIFIED_BG
		}

		if t.Tagged {
			colorFg = config.COLOR_TAGGED_FG
			colorBg = config.COLOR_TAGGED_BG
		}

		cell = cview.NewTableCell(t.GetAttributeString())
		cell.SetTextColor(colorFg)
		cell.SetBackgroundColor(colorBg)
		cell.SetAttributes(attributes)
		transactionsTable.SetCell(i+1, config.ATTR_COLUMN, cell)

		cell = cview.NewTableCell(fmt.Sprint(&t.Category))
		cell.SetTextColor(colorFg)
		cell.SetBackgroundColor(colorBg)
		cell.SetAttributes(attributes)
		transactionsTable.SetCell(i+1, config.CATEGORY_COLUMN, cell)

		cell = cview.NewTableCell(t.Date.GetText())
		cell.SetTextColor(colorFg)
		cell.SetBackgroundColor(colorBg)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignLeft)
		transactionsTable.SetCell(i+1, config.DATE_COLUMN, cell)

		cell = cview.NewTableCell(fmt.Sprint(&t.Description))
		cell.SetTextColor(colorFg)
		cell.SetBackgroundColor(colorBg)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignLeft)
		transactionsTable.SetCell(i+1, config.DESCRIPTION_COLUMN, cell)

		cell = cview.NewTableCell(fmt.Sprint(&t.Cents))
		if t.Cents.IsDebit() {
			cell.SetTextColor(config.COLOR_NEGATIVE_FG)
		} else {
			cell.SetTextColor(config.COLOR_POSITIVE_FG)
		}
		cell.SetBackgroundColor(colorBg)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignRight)
		transactionsTable.SetCell(i+1, config.AMOUNT_COLUMN, cell)

		balance := t.Balance()
		cell = cview.NewTableCell(fmt.Sprint(balance))
		if balance.IsDebit() {
			cell.SetTextColor(config.COLOR_NEGATIVE_FG)
		} else {
			cell.SetTextColor(config.COLOR_POSITIVE_FG)
		}
		cell.SetBackgroundColor(colorBg)
		cell.SetAttributes(tcell.AttrBold)
		cell.SetAlign(cview.AlignRight)
		transactionsTable.SetCell(i+1, config.BALANCE_COLUMN, cell)
	}
}

// select transaction by id
func selectTransaction(id int) {
	for i, t := range transactions {
		if t.Id == id {
			transactionsTable.Select(i, 0)
		}
	}

}

// get the id of the selection
func selectionId() int {
	r, _ := transactionsTable.GetSelection()
	return transactions[r-1].Id
}

// get Transaction of the selection
func selectionTransaction() *transaction.Transaction {
	return transactions[selectionId()]
}
