package ui

import (
	"fmt"
	"jcb/db"
	"jcb/lib/format"
	dataf "jcb/lib/formatter/data"
	"jcb/lib/transaction"
	"jcb/lib/validate"
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
	value, err := format.Description(description)
	if err != nil {
		printStatus(fmt.Sprint(err))
		return
	}

	skipped := 0

	for _, id := range ids {
		t, _ := transaction.Find(id)

		if t.Description.String() == value {
			skipped++
		} else {
			t.Description.SetText(value)
			err = t.Save()
			if err != nil {
				panic(err)
			}
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
	err := validate.Date(date)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	startingId := selectionId()
	value, _ := format.Date(date)
	skipped := 0
	lastCommittedDate := db.DateLastCommitted()

	value, err = format.Date(date)
	if err != nil {
		printStatus(fmt.Sprint(err))
		return
	}

	for _, id := range ids {
		t, _ := transaction.Find(id)

		if lastCommittedDate.Unix() > value.Unix() {
			skipped++
			continue
		}

		t.Date.SetValue(value)

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
