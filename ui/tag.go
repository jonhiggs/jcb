package ui

import (
	"fmt"
	"jcb/config"
	"strings"
)

var taggedTransactions []int

func isTagged(row int) bool {
	for _, r := range taggedTransactions {
		if r == row {
			return true
		}
	}

	return false
}

func applyTag(row int) {
	if isCommitted(row) {
		printStatus("Cannot tag committed transactions")
	} else {
		taggedTransactions = append(taggedTransactions, row)
	}
}

func removeTag(row int) {
	var t []int
	for _, r := range taggedTransactions {
		if r != row {
			t = append(t, r)
		}
	}
	taggedTransactions = t
}

func tagMatches() {
	curRow, _ := transactionsTable.GetSelection()

	matchCount := 0

	for i := 1; i < len(transactionIds); i++ {
		if isCommitted(i) || isTagged(i) {
			continue
		}

		if strings.Contains(strings.ToLower(transactionsTable.GetCell(i, config.DESCRIPTION_COLUMN).GetText()), strings.ToLower(findQuery)) {
			matchCount += 1
			applyTag(i)
		}
	}

	transactionsTable.Select(curRow, 0)

	printStatus(fmt.Sprintf("Tagged %d transactions", matchCount))
	updateTransactionsTable()
}
