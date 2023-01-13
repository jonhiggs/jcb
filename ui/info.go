package ui

import (
	"fmt"
	"jcb/config"
	"jcb/lib/transaction"

	"code.rocketnine.space/tslocum/cview"
)

var info *cview.TextView

func updateInfo() {
	row, _ := transactionsTable.GetSelection()
	if row == 0 {
		row = 1
	}

	modifedCount := 0

	start, end := transaction.DateRange()

	for i, t := range transaction.All(start, end) {
		if i == 0 {
			continue
		}

		if !t.IsSaved() {
			modifedCount += 1
		}
	}

	info.SetText(fmt.Sprintf("[%d:%d] [%d] [%d]", row, len(transactionIds)-1, modifedCount, len(taggedTransactionIds)))
}

func createInfoTextView() *cview.TextView {
	info = cview.NewTextView()
	info.SetTextAlign(cview.AlignRight)
	info.SetTextColor(config.COLOR_INFO_FG)
	return info
}
