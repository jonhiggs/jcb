package ui

import (
	"fmt"
	"jcb/lib/transaction"
)

func updateCategory(s string, ids []int64) {
	orig := new(transaction.Category)
	err := orig.SetText(s)
	if err != nil {
		printStatus(fmt.Sprint(err))
		return
	}

	skipped := 0

	for _, id := range ids {
		t, _ := transaction.Find(id)

		// TODO: create a function to detect whether the data has changed
		t.Category.SetText(orig.GetValue())
		if true {
			t.Save()
		} else {
			skipped++
		}
	}

	modified := len(ids) - skipped
	printStatus(fmt.Sprintf("Updated category for %d transactions", modified))
	updateTransactionsTable()
}

func updateDescription(s string, ids []int64) {
	orig := new(transaction.Description)
	err := orig.SetText(s)
	if err != nil {
		printStatus(fmt.Sprint(err))
		return
	}

	skipped := 0

	for _, id := range ids {
		t, _ := transaction.Find(id)

		// TODO: create a function to detect whether the data has changed
		t.Description.SetText(orig.GetValue())
		if true {
			t.Save()
		} else {
			skipped++
		}
	}

	modified := len(ids) - skipped
	printStatus(fmt.Sprintf("Updated description for %d transactions", modified))
	updateTransactionsTable()
}

func updateCents(cents string, ids []int64) {
	orig := new(transaction.Cents)
	err := orig.SetText(cents)
	if err != nil {
		printStatus(fmt.Sprint(err))
		return
	}

	skipped := 0

	for _, id := range ids {
		t, _ := transaction.Find(id)

		t.Cents.SetValue(orig.GetValue())

		// TODO: create a function to detect whether the data has changed
		if true {
			t.Save()
		} else {
			skipped++
		}
	}

	modified := len(ids) - skipped
	printStatus(fmt.Sprintf("Updated amount for %d transactions", modified))
	updateTransactionsTable()
}

func updateDate(s string, ids []int64) {
	orig := new(transaction.Date)
	err := orig.SetText(s)
	if err != nil {
		printStatus(fmt.Sprint(err))
		return
	}

	lastCommitted, err := transaction.FindLastCommitted()
	if err != nil {
		panic("You should not make it here. You cannot modify a committed transaction so there must be uncommitted transactions!")
	}

	if lastCommitted.Date.Unix() > orig.Unix() {
		return
	}

	startingId := selectionId()
	skipped := 0

	for _, id := range ids {
		t, _ := transaction.Find(id)

		t.Date.SetValue(orig.GetValue())

		// TODO: create a function to detect whether the data has changed
		if true {
			t.Save()
		} else {
			skipped++
		}
	}

	updateTransactionsTable()
	selectTransaction(startingId)

	modified := len(ids) - skipped
	printStatus(fmt.Sprintf("Updated date for %d transactions", modified))
}
