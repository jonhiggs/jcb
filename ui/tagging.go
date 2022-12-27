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

	categoryValue := dataf.Category(category)

	for _, r := range rows {
		t, _ := transaction.Find(transactionIds[r])
		t.Category = categoryValue
		transaction.Edit(t)
	}

	printStatus(fmt.Sprintf("Updated category for %d transactions", len(rows)))
	updateTransactionsTable()
}

func updateDescription(description string, rows []int) {
	err := validator.Description(description)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	descriptionValue := dataf.Description(description)

	for _, r := range rows {
		t, _ := transaction.Find(transactionIds[r])
		t.Description = descriptionValue
		transaction.Edit(t)
	}

	printStatus(fmt.Sprintf("Updated description for %d transactions", len(rows)))
	updateTransactionsTable()
}

func updateCents(cents string, rows []int) {
	err := validator.Cents(cents)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	centsValue := dataf.Cents(cents)
	for _, r := range rows {
		t, _ := transaction.Find(transactionIds[r])
		t.Cents = centsValue
		err = transaction.Edit(t)
		if err != nil {
			printStatus(fmt.Sprint(err))
		}
	}

	printStatus(fmt.Sprintf("Updated amount for %d transactions", len(rows)))
	updateTransactionsTable()
}

func updateDate(date string, rows []int) {
	err := validator.Date(date)
	if err != nil {
		printStatus(fmt.Sprintf("%s", err))
		return
	}

	id := selectionId()

	dateValue := dataf.Date(date)
	for _, r := range rows {
		t, _ := transaction.Find(transactionIds[r])
		t.Date = dateValue
		err = transaction.Edit(t)
		if err != nil {
			printStatus(fmt.Sprint(err))
		}
	}

	updateTransactionsTable()
	selectTransaction(id)
	printStatus(fmt.Sprintf("Updated date for %d transactions", len(rows)))
}
