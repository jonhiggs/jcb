package ui

import (
	"fmt"
	dataf "jcb/lib/formatter/data"
	"jcb/lib/transaction"
	"jcb/lib/validator"

	"github.com/gdamore/tcell/v2"
)

var taggedTransactions []int

func handleTag(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if isTagged(r) {
		removeTag(r)
	} else {
		applyTag(r)
	}

	handleSelectNext(ev)
	updateTransactionsTable()

	return nil
}

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

func handleTagCommand(ev *tcell.EventKey) *tcell.EventKey {
	screen := app.GetScreen()
	cmdEv := screen.PollEvent()

	startingTransaction, _ := transactionsTable.GetSelection()

	switch e := cmdEv.(type) {
	case *tcell.EventKey:
		switch e.Rune() {
		case 'x':
			for _, r := range taggedTransactions {
				selectTransaction(transactionIds[r])
				handleDeleteTransaction(e)
				startingTransaction, _ = transactionsTable.GetSelection()
			}
		case 't':
			for _, r := range taggedTransactions {
				removeTag(r)
			}
		case 'D':
			openPrompt("Category:", selectedCategory(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				category := dataf.Category(promptInputField.GetText())
				updateCategory(category, taggedTransactions)
				return nil
			})
		case 'd':
			openPrompt("Description:", selectedDescription(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				updateDescription(promptInputField.GetText(), taggedTransactions)
				return nil
			})
		case '=':
			openPrompt("Amount:", selectedAmount(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				cents := promptInputField.GetText()
				updateCents(cents, taggedTransactions)
				return nil
			})
		case '@':
			openPrompt("Date:", selectedDate(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				date := promptInputField.GetText()
				updateDate(date, taggedTransactions)
				return nil
			})
		}
	}

	transactionsTable.Select(startingTransaction, 0)
	updateTransactionsTable()
	return nil
}

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
