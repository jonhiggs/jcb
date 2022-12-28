package ui

import (
	"fmt"
	"jcb/config"
	"jcb/lib/transaction"
	"strings"
)

var taggedTransactionIds []int64

func isTagged(id int64) bool {
	for _, i := range taggedTransactionIds {
		if i == id {
			return true
		}
	}

	return false
}

func applyTag(id int64) {
	taggedTransactionIds = append(taggedTransactionIds, id)
}

func removeTag(id int64) {
	var newTransactionIds []int64
	for _, i := range taggedTransactionIds {
		if i != id {
			newTransactionIds = append(newTransactionIds, i)
		}
	}
	taggedTransactionIds = newTransactionIds
}

func tagMatches(id int64) {
	matchCount := 0

	for r, i := range transactionIds {
		if transaction.Attributes(id).Committed || isTagged(i) {
			continue
		}

		if strings.Contains(strings.ToLower(transactionsTable.GetCell(r, config.DESCRIPTION_COLUMN).GetText()), strings.ToLower(findQuery)) {
			matchCount += 1
			applyTag(i)
		}
	}

	selectTransaction(id)
	printStatus(fmt.Sprintf("Tagged %d transactions", matchCount))
	updateTransactionsTable()
}
