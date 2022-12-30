package ui

import (
	"fmt"
	"jcb/config"

	"code.rocketnine.space/tslocum/cview"
)

var info *cview.TextView

func updateInfo() {
	row, _ := transactionsTable.GetSelection()
	if row == 0 {
		row = 1
	}
	info.SetText(fmt.Sprintf("[%d:%d] [%d]", row, len(transactionIds)-1, len(taggedTransactionIds)))
}

func createInfoTextView() *cview.TextView {
	info = cview.NewTextView()
	info.SetTextAlign(cview.AlignRight)
	info.SetTextColor(config.COLOR_INFO_FG)
	return info
}
