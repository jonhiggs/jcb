package ui

import (
	"fmt"
	dataf "jcb/lib/formatter/data"
	"jcb/lib/transaction"
	"jcb/lib/validator"
)

func updateCategory(category string, rows []int) {
	err := validator.Category(category)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	value := dataf.Category(category)
	skipped := 0

	for _, r := range rows {
		t, _ := transaction.Find(transactionIds[r])
		if t.Category == value {
			skipped += 1
			continue
		}

		t.Category = value
		transaction.Edit(t)
	}

	modified := len(rows) - skipped
	printStatus(fmt.Sprintf("Updated category for %d transactions", modified))
	updateTransactionsTable()
}

func updateDescription(description string, rows []int) {
	err := validator.Description(description)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	value := dataf.Description(description)
	skipped := 0

	for _, r := range rows {
		t, _ := transaction.Find(transactionIds[r])

		if t.Description == value {
			skipped += 1
			continue
		}

		t.Description = value
		transaction.Edit(t)
	}

	modified := len(rows) - skipped
	printStatus(fmt.Sprintf("Updated description for %d transactions", modified))
	updateTransactionsTable()
}

func updateCents(cents string, rows []int) {
	err := validator.Cents(cents)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	value := dataf.Cents(cents)
	skipped := 0

	for _, r := range rows {
		t, _ := transaction.Find(transactionIds[r])

		if t.Cents == value {
			skipped += 1
			continue
		}

		t.Cents = value
		err = transaction.Edit(t)
		if err != nil {
			printStatus(fmt.Sprint(err))
		}
	}

	modified := len(rows) - skipped
	printStatus(fmt.Sprintf("Updated amount for %d transactions", modified))
	updateTransactionsTable()
}

func updateDate(date string, rows []int) {
	err := validator.Date(date)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	id := selectionId()
	value := dataf.Date(date)
	skipped := 0

	for _, r := range rows {
		t, _ := transaction.Find(transactionIds[r])

		if t.Date == value {
			skipped += 1
			continue
		}

		t.Date = value
		err = transaction.Edit(t)
		if err != nil {
			printStatus(fmt.Sprint(err))
		}
	}

	updateTransactionsTable()
	selectTransaction(id)

	modified := len(rows) - skipped
	printStatus(fmt.Sprintf("Updated date for %d transactions", modified))
}
