package ui

import (
	"fmt"
	"jcb/lib/find"
	"jcb/lib/transaction"
)

func taggedTransactions() []*transaction.Transaction {
	var ts []*transaction.Transaction
	for _, t := range transactions {
		if t.Tagged {
			ts = append(ts, t)
		}
	}
	return ts
}

func tagMatches(id int) {
	matchCount := 0

	for r, t := range transactions {
		if t.Committed || t.Tagged {
			continue
		}

		if find.TableRowMatches(transactionsTable, r) {
			matchCount += 1
			t.Tag()
		}
	}

	selectTransaction(id)
	printStatus(fmt.Sprintf("Tagged %d transactions", matchCount))
	updateTransactionsTable()
}

func untagMatches(id int) {
	matchCount := 0

	for _, t := range taggedTransactions() {
		selectTransaction(t.Id)
		r, _ := transactionsTable.GetSelection()

		if find.TableRowMatches(transactionsTable, r) {
			matchCount += 1
			t.Untag()
		}
	}

	selectTransaction(id)
	printStatus(fmt.Sprintf("Untagged %d transactions", matchCount))
	updateTransactionsTable()
}
