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
	c.Set("1", handleOpenHelp)
	c.Set("2", handleOpenTransactions)
	c.Set("3", handleOpenBudget)
	c.Set("4", handleOpenReport)
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
	committed := transaction.Committed()
	all := transaction.All()

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

		var colorFg tcell.Color
		var colorBg tcell.Color
		var attributes tcell.AttrMask

		transactionIds[i+1] = t.Id
		transactionAttributes[i+1] = transaction.Attributes(t.Id)

		if isCommitted {
			colorFg = config.COLOR_COMMITTED_FG
			colorBg = config.COLOR_COMMITTED_BG
			attributes = 0
		} else {
			colorFg = config.COLOR_UNCOMMITTED_FG
			colorBg = config.COLOR_UNCOMMITTED_BG
		}

		if !transactionAttributes[i+1].Saved {
			colorFg = config.COLOR_MODIFIED_FG
			colorBg = config.COLOR_MODIFIED_BG
		}

		if isTagged(t.Id) {
			colorFg = config.COLOR_TAGGED_FG
			colorBg = config.COLOR_TAGGED_BG
		}

		cell = cview.NewTableCell(stringf.Attributes(transactionAttributes[i+1]))
		cell.SetTextColor(colorFg)
		cell.SetBackgroundColor(colorBg)
		cell.SetAttributes(attributes)
		transactionsTable.SetCell(i+1, config.ATTR_COLUMN, cell)

		cell = cview.NewTableCell(fmt.Sprintf("%-10s", stringf.Category(t.Category)))
		cell.SetTextColor(colorFg)
		cell.SetBackgroundColor(colorBg)
		cell.SetAttributes(attributes)
		transactionsTable.SetCell(i+1, config.CATEGORY_COLUMN, cell)

		cell = cview.NewTableCell(date)
		cell.SetTextColor(colorFg)
		cell.SetBackgroundColor(colorBg)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignLeft)
		transactionsTable.SetCell(i+1, config.DATE_COLUMN, cell)

		if len(description) > config.DESCRIPTION_MAX_LENGTH {
			description = description[0:config.DESCRIPTION_MAX_LENGTH]
		}
		cell = cview.NewTableCell(fmt.Sprintf("%-*s", config.DESCRIPTION_MAX_LENGTH, description))
		cell.SetTextColor(colorFg)
		cell.SetBackgroundColor(colorBg)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignLeft)
		transactionsTable.SetCell(i+1, config.DESCRIPTION_COLUMN, cell)

		cell = cview.NewTableCell(fmt.Sprintf("%10s", cents))
		if dataf.Cents(cents) < 0 {
			cell.SetTextColor(config.COLOR_NEGATIVE_FG)
		} else {
			cell.SetTextColor(config.COLOR_POSITIVE_FG)
		}
		cell.SetBackgroundColor(colorBg)
		cell.SetAttributes(attributes)
		cell.SetAlign(cview.AlignRight)
		transactionsTable.SetCell(i+1, config.AMOUNT_COLUMN, cell)

		cell = cview.NewTableCell(fmt.Sprintf("%10s", balance))
		if dataf.Cents(balance) < 0 {
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
