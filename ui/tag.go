package ui

import (
	"fmt"
	"jcb/lib/find"
	"jcb/lib/transaction"
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
	updateInfo()
}

func removeTag(id int64) {
	var newTransactionIds []int64
	for _, i := range taggedTransactionIds {
		if i != id {
			newTransactionIds = append(newTransactionIds, i)
		}
	}
	taggedTransactionIds = newTransactionIds
	updateInfo()
}

func toggleTag(id int64) {
	if isTagged(id) {
		removeTag(id)
	} else {
		applyTag(id)
	}
}

func tagMatches(id int64) {
	matchCount := 0

	for r, i := range transactionIds {
		t, _ := transaction.Find(i)
		if t.IsCommitted() || isTagged(i) {
			continue
		}

		if find.TableRowMatches(transactionsTable, r) {
			matchCount += 1
			applyTag(i)
		}
	}

	selectTransaction(id)
	printStatus(fmt.Sprintf("Tagged %d transactions", matchCount))
	updateTransactionsTable()
}

func untagMatches(id int64) {
	matchCount := 0

	for _, i := range taggedTransactionIds {
		selectTransaction(i)
		r, _ := transactionsTable.GetSelection()

		if find.TableRowMatches(transactionsTable, r) {
			matchCount += 1
			removeTag(i)
		}
	}

	selectTransaction(id)
	printStatus(fmt.Sprintf("Untagged %d transactions", matchCount))
	updateTransactionsTable()
}
