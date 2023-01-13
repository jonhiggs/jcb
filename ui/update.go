package ui

import (
	"fmt"
	"jcb/db"
	dataf "jcb/lib/formatter/data"
	"jcb/lib/transaction"
	"jcb/lib/validator"
)

func updateCategory(category string, ids []int64) {
	err := validator.Category(category)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	value := dataf.Category(category)
	skipped := 0

	for _, id := range ids {
		t, _ := transaction.Find(id)

		if t.SetCategory(value) {
			t.Save()
		} else {
			skipped++
		}
	}

	modified := len(ids) - skipped
	printStatus(fmt.Sprintf("Updated category for %d transactions", modified))
	updateTransactionsTable()
}

func updateDescription(description string, ids []int64) {
	err := validator.Description(description)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	value := dataf.Description(description)
	skipped := 0

	for _, id := range ids {
		t, _ := transaction.Find(id)

		if t.SetDescription(value) {
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
	err := validator.Cents(cents)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	value := dataf.Cents(cents)
	skipped := 0

	for _, id := range ids {
		t, _ := transaction.Find(id)

		if t.SetCents(value) {
			t.Save()
		} else {
			skipped++
		}
	}

	modified := len(ids) - skipped
	printStatus(fmt.Sprintf("Updated amount for %d transactions", modified))
	updateTransactionsTable()
}

func updateDate(date string, ids []int64) {
	err := validator.Date(date)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	startingId := selectionId()
	value := dataf.Date(date)
	skipped := 0
	lastCommittedDate := db.DateLastCommitted()

	for _, id := range ids {
		t, _ := transaction.Find(id)

		if lastCommittedDate.Unix() > value.Unix() {
			skipped++
			continue
		}

		if t.SetDate(value) {
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
