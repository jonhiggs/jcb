package ui

import (
	"jcb/config"
	"jcb/lib/category"
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"
	"time"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var reportTable *cview.Table

func createReportTable() *cview.Table {
	reportTable = cview.NewTable()
	reportTable.Select(0, 0)
	reportTable.SetBorders(false)
	reportTable.SetFixed(1, 1)
	reportTable.SetSelectable(true, false)
	reportTable.SetSeparator(' ')
	reportTable.SetRect(0, 0, config.MAX_WIDTH, 20)
	reportTable.SetScrollBarVisibility(cview.ScrollBarNever)
	reportTable.SetSelectionChangedFunc(func(r int, c int) { closeStatus() })

	c := cbind.NewConfiguration()
	c.Set("q", handleCloseReport)
	c.Set("j", handleReportSelectNext)
	c.Set("k", handleReportSelectPrev)
	c.Set("F1", handleOpenHelp)
	c.Set("F2", handleOpenTransactions)
	c.Set("F3", handleOpenReport)
	c.Set(":", handleCommand)
	reportTable.SetInputCapture(c.Capture)

	reportTable.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			reportTable.SetSelectable(true, true)
		}
	})

	return reportTable
}

func updateReportTable() {
	var cell *cview.TableCell

	st, _ := transaction.Find(selectionId())
	year := st.Date.Year()

	columns := []string{
		"CATEGORY",
		"JAN",
		"FEB",
		"MAR",
		"APR",
		"MAY",
		"JUN",
		"JUL",
		"AUG",
		"SEP",
		"OCT",
		"NOV",
		"DEC",
		"TOTAL",
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

		reportTable.SetCell(0, i, cell)
	}

	for row, catName := range category.All() {
		for col, _ := range columns {
			if col == 0 {
				cell = cview.NewTableCell(catName)
				cell.SetTextColor(config.COLOR_TITLE_FG)
			} else {
				var startTime time.Time
				var endTime time.Time

				switch col {
				case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
					startTime = time.Date(year, time.Month(col), 1, 0, 0, 0, 0, time.UTC)
					endTime = startTime.AddDate(0, 1, 0)
				case 13:
					startTime = time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
					endTime = startTime.AddDate(1, 0, 0)
				}

				cents := category.Sum(catName, startTime, endTime)
				cell = cview.NewTableCell(" " + stringf.Cents(cents))

				if col == 13 {
					cell.SetAttributes(tcell.AttrBold)
					cell.SetBackgroundColor(config.COLOR_LIGHT_BG)
				}

				cell.SetAlign(cview.AlignRight)
			}
			reportTable.SetCell(row+1, col, cell)
		}
	}

	row := reportTable.GetRowCount()

	// XXX: don't add new total rows for every update
	if reportTable.GetCell(row-1, 0).GetText() == "" {
		row = row - 1
	}

	for col, _ := range columns {
		var startTime time.Time
		var endTime time.Time

		if col == 0 {
			continue
		}

		switch col {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
			startTime = time.Date(year, time.Month(col), 1, 0, 0, 0, 0, time.UTC)
			endTime = startTime.AddDate(0, 1, 0)
		case 13:
			startTime = time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
			endTime = startTime.AddDate(1, 0, 0)
		}

		cents := transaction.Sum(startTime, endTime)
		cell = cview.NewTableCell(" " + stringf.Cents(cents))
		cell.SetSelectable(false)
		cell.SetAlign(cview.AlignRight)
		cell.SetAttributes(tcell.AttrBold)
		cell.SetBackgroundColor(config.COLOR_LIGHT_BG)

		if cents < 0 {
			cell.SetTextColor(config.COLOR_NEGATIVE_FG)
		} else {
			cell.SetTextColor(config.COLOR_POSITIVE_FG)
		}

		reportTable.SetCell(row, col, cell)
	}
}
