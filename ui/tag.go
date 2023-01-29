package ui

import (
	"fmt"
	"jcb/lib/find"
	"jcb/lib/transaction"
)

var taggedTransactionIds []int

func isTagged(id int) bool {
	for _, i := range taggedTransactionIds {
		if i == id {
			return true
		}
	}

	return false
}

func taggedTransactions() []*transaction.Transaction {
	var ts []*transaction.Transaction
	for _, id := range taggedTransactionIds {
		t, _ := transaction.Find(id)
		ts = append(ts, t)
	}
	return ts
}

func applyTag(id int) {
	taggedTransactionIds = append(taggedTransactionIds, id)
	updateInfo()
}

func removeTag(id int) {
	var newTransactionIds []int
	for _, i := range taggedTransactionIds {
		if i != id {
			newTransactionIds = append(newTransactionIds, i)
		}
	}
	taggedTransactionIds = newTransactionIds
	updateInfo()
}

func toggleTag(id int) {
	if isTagged(id) {
		removeTag(id)
	} else {
		applyTag(id)
	}
}

func tagMatches(id int) {
	matchCount := 0

	for r, t := range transactions {
		if t.IsCommitted() || isTagged(t.Id) {
			continue
		}

		if find.TableRowMatches(transactionsTable, r) {
			matchCount += 1
			applyTag(t.Id)
		}
	}

	selectTransaction(id)
	printStatus(fmt.Sprintf("Tagged %d transactions", matchCount))
	updateTransactionsTable()
}

func untagMatches(id int) {
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
