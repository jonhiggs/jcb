package ui

import (
	"jcb/config"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var budgetTable *cview.Table

func closeBudget() {
	panels.HidePanel("budget")
	closeStatus()
}

func createBudgetTable() *cview.Table {
	budgetTable = cview.NewTable()
	budgetTable.Select(0, 0)
	budgetTable.SetBorders(false)
	budgetTable.SetFixed(1, 1)
	budgetTable.SetSelectable(true, false)
	budgetTable.SetSeparator(' ')
	budgetTable.SetRect(0, 0, config.MAX_WIDTH, 20)
	budgetTable.SetScrollBarVisibility(cview.ScrollBarNever)
	budgetTable.SetSelectionChangedFunc(func(r int, c int) { closeStatus() })

	c := cbind.NewConfiguration()
	c.Set("j", handleBudgetSelectNext)
	c.Set("k", handleBudgetSelectPrev)
	c.Set("1", handleOpenHelp)
	c.Set("2", handleOpenTransactions)
	c.Set("3", handleOpenBudget)
	c.Set("4", handleOpenReport)
	c.Set(":", handleCommand)
	c.Set("i", handleOpenInsertBudget)
	budgetTable.SetInputCapture(c.Capture)

	budgetTable.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			budgetTable.SetSelectable(true, true)
		}
	})

	return budgetTable
}

func updateBudgetTable() {
	var cell *cview.TableCell

	//st, _ := transaction.Find(selectionId())
	//year := st.Date.Year()

	columns := []string{
		"",
		"DATE",
		"CATEGORY",
		"BUDGET",
		"REMAINING",
	}

	for i, c := range columns {
		cell = cview.NewTableCell(c)
		cell.SetTextColor(config.COLOR_TITLE_FG)
		cell.SetAttributes(tcell.AttrUnderline | tcell.AttrBold)
		cell.SetSelectable(false)
		cell.SetBackgroundColor(config.COLOR_TITLE_BG)
		if i == 0 {
			cell.SetAlign(cview.AlignLeft)
		} else {
			cell.SetAlign(cview.AlignRight)
		}

		budgetTable.SetCell(0, i, cell)
	}

	cell = cview.NewTableCell("Ccn+")
	budgetTable.SetCell(1, 0, cell)

	cell = cview.NewTableCell("2023-01-02")
	budgetTable.SetCell(1, 1, cell)

	cell = cview.NewTableCell("shopping")
	budgetTable.SetCell(1, 2, cell)

	cell = cview.NewTableCell("20.33")
	budgetTable.SetCell(1, 3, cell)

	cell = cview.NewTableCell("20.33")
	budgetTable.SetCell(1, 4, cell)

	cell = cview.NewTableCell("Ccn+")
	budgetTable.SetCell(2, 0, cell)

	cell = cview.NewTableCell("2023-01-02")
	budgetTable.SetCell(2, 1, cell)

	cell = cview.NewTableCell("health")
	budgetTable.SetCell(2, 2, cell)

	cell = cview.NewTableCell("340.00")
	budgetTable.SetCell(2, 3, cell)

	cell = cview.NewTableCell("340.00")
	budgetTable.SetCell(2, 4, cell)

	//for row, catName := range category.All() {
	//	for col, _ := range columns {
	//		if col == 0 {
	//			cell = cview.NewTableCell(catName)
	//			cell.SetTextColor(config.COLOR_TITLE_FG)
	//		} else {
	//			var startTime time.Time
	//			var endTime time.Time

	//			switch col {
	//			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	//				startTime = time.Date(year, time.Month(col), 1, 0, 0, 0, 0, time.UTC)
	//				endTime = startTime.AddDate(0, 1, 0)
	//			case 13:
	//				startTime = time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	//				endTime = startTime.AddDate(1, 0, 0)
	//			}

	//			cents := category.Sum(catName, startTime, endTime)
	//			cell = cview.NewTableCell(" " + stringf.Cents(cents))

	//			if col == 13 {
	//				cell.SetAttributes(tcell.AttrBold)
	//				cell.SetBackgroundColor(config.COLOR_LIGHT_BG)
	//			}

	//			cell.SetAlign(cview.AlignRight)
	//		}
	//		budgetTable.SetCell(row+1, col, cell)
	//	}
	//}

	//row := budgetTable.GetRowCount()

	//// XXX: don't add new total rows for every update
	//if budgetTable.GetCell(row-1, 0).GetText() == "" {
	//	row = row - 1
	//}

	//for col, _ := range columns {
	//	var startTime time.Time
	//	var endTime time.Time

	//	if col == 0 {
	//		continue
	//	}

	//	switch col {
	//	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	//		startTime = time.Date(year, time.Month(col), 1, 0, 0, 0, 0, time.UTC)
	//		endTime = startTime.AddDate(0, 1, 0)
	//	case 13:
	//		startTime = time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	//		endTime = startTime.AddDate(1, 0, 0)
	//	}

	//	cents := transaction.Sum(startTime, endTime)
	//	cell = cview.NewTableCell(" " + stringf.Cents(cents))
	//	cell.SetSelectable(false)
	//	cell.SetAlign(cview.AlignRight)
	//	cell.SetAttributes(tcell.AttrBold)
	//	cell.SetBackgroundColor(config.COLOR_LIGHT_BG)

	//	if cents < 0 {
	//		cell.SetTextColor(config.COLOR_NEGATIVE_FG)
	//	} else {
	//		cell.SetTextColor(config.COLOR_POSITIVE_FG)
	//	}

	//	budgetTable.SetCell(row, col, cell)
	//}
}
