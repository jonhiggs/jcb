package ui

import (
	"fmt"
	"jcb/config"
	"jcb/db"
	dataf "jcb/lib/formatter/data"
	"jcb/lib/repeater"
	"jcb/lib/transaction"
	acceptanceFunction "jcb/ui/acceptance-functions"
	"log"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

var repeatRuleValue string
var repeatUntilValue time.Time

func handleOpenTransactions(ev *tcell.EventKey) *tcell.EventKey {
	panels.ShowPanel("transactions")
	panels.HidePanel("report")
	panels.SendToFront("transactions")
	return nil
}

func handleSelectNext(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if transactionsTable.GetRowCount() > r+1 {
		transactionsTable.Select(r+1, 0)
	}

	return nil
}

func handleSelectPrev(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	transactionsTable.Select(r-1, 0)
	return nil
}

func handleHalfPageDown(ev *tcell.EventKey) *tcell.EventKey {
	_, h := app.GetScreenSize()
	r, _ := transactionsTable.GetSelection()

	if r+(h/2) < transactionsTable.GetRowCount() {
		transactionsTable.Select(r+(h/2), 0)
	} else {
		transactionsTable.Select(transactionsTable.GetRowCount()-1, 0)
	}

	return nil
}

func handleHalfPageUp(ev *tcell.EventKey) *tcell.EventKey {
	_, h := app.GetScreenSize()
	r, _ := transactionsTable.GetSelection()

	if r-(h/2) > 0 {
		transactionsTable.Select(r-(h/2), 0)
	} else {
		transactionsTable.Select(0, 0)
	}

	return nil
}

func handleSelectFirstUncommitted(ev *tcell.EventKey) *tcell.EventKey {
	uncommitted, _ := transaction.Uncommitted()
	if len(uncommitted) > 0 {
		firstUncommitted := uncommitted[0]

		for i, v := range transactionIds {
			if firstUncommitted.Id == v {
				transactionsTable.Select(i, 0)
				return nil
			}
		}
	}

	transactionsTable.Select(len(transactionIds)-1, 0)
	return nil
}

func handleSelectSimilar(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()
	curDescription := transactionsTable.GetCell(curRow, config.DESCRIPTION_COLUMN).GetText()

	for i := curRow + 1; i != curRow; i++ {
		if transactionsTable.GetCell(i, config.DESCRIPTION_COLUMN).GetText() == curDescription {
			transactionsTable.Select(i, 0)
			break
		}

		if i == len(transactionIds) {
			i = 0
		}
	}

	return nil
}

func handleSelectMonthPrev(ev *tcell.EventKey) *tcell.EventKey {
	curDate := selectedDate()
	curMonth, _ := strconv.Atoi(curDate[5:7])
	curYear, _ := strconv.Atoi(curDate[0:4])

	r, _ := transactionsTable.GetSelection()

	for i := r - 1; i > 0; i-- {
		date := transactionsTable.GetCell(i, config.DATE_COLUMN).GetText()
		month, _ := strconv.Atoi(date[5:7])
		year, _ := strconv.Atoi(date[0:4])

		if month < curMonth || year < curYear {
			transactionsTable.Select(i, 0)
			return nil
		}
		transactionsTable.Select(1, 0)
	}

	return nil
}

func handleSelectMonthNext(ev *tcell.EventKey) *tcell.EventKey {
	curDate := selectedDate()
	curMonth, _ := strconv.Atoi(curDate[5:7])
	curYear, _ := strconv.Atoi(curDate[0:4])

	r, _ := transactionsTable.GetSelection()
	for i := r; i < len(transactionIds); i++ {
		date := transactionsTable.GetCell(i, config.DATE_COLUMN).GetText()
		month, _ := strconv.Atoi(date[5:7])
		year, _ := strconv.Atoi(date[0:4])

		if month > curMonth || year > curYear {
			transactionsTable.Select(i, 0)
			return nil
		}
	}
	transactionsTable.Select(len(transactionIds)-1, 0)

	return nil
}

func handleSelectYearPrev(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()
	curYear := dataf.Date(transactionsTable.GetCell(curRow, config.DATE_COLUMN).GetText()).Year()

	for i := curRow; i > 0; i-- {
		year := dataf.Date(transactionsTable.GetCell(i, config.DATE_COLUMN).GetText()).Year()
		if int(year) != int(curYear) {
			transactionsTable.Select(i, 0)
			return nil
		}
	}

	transactionsTable.Select(1, 0)

	return nil
}

func handleSelectYearNext(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()
	curYear := dataf.Date(transactionsTable.GetCell(curRow, config.DATE_COLUMN).GetText()).Year()

	for i := curRow; i < len(transactionIds)-1; i++ {
		year := dataf.Date(transactionsTable.GetCell(i, config.DATE_COLUMN).GetText()).Year()
		if int(year) != int(curYear) {
			transactionsTable.Select(i, 0)
			return nil
		}
	}

	return nil
}

func handleFindForwards(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt("/", "", func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		findQuery = promptInputField.GetText()
		handleSelectMatchNext(ev)
		return nil
	}, acceptanceFunction.Any)

	return nil
}

func handleFindBackwards(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt("?", "", func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		findQuery = promptInputField.GetText()
		handleSelectMatchPrev(ev)
		return nil
	}, acceptanceFunction.Any)

	return nil
}

func handleSelectMatchNext(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()

	for i := curRow + 1; i != curRow; i++ {

		if matchesQuery(i) {
			transactionsTable.Select(i, 0)
			return nil
		}

		if i == len(transactionIds) {
			i = 0
		}
	}

	printStatus("No matches found")

	return nil
}

func handleSelectMatchPrev(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()

	for i := curRow - 1; i != curRow; i-- {
		if matchesQuery(i) {
			transactionsTable.Select(i, 0)
			break
		}

		if i == 0 {
			i = len(transactionIds) - 1
		}
	}

	printStatus("No matches found")

	return nil
}

func handleDeleteTransaction(ev *tcell.EventKey) *tcell.EventKey {
	id := selectionId()

	curRow, _ := transactionsTable.GetSelection()
	var r int
	if curRow == len(transactionIds)-1 {
		r = curRow - 1
	} else {
		r = curRow
	}

	err := transaction.Delete(id)
	if err != nil {
		printStatus(fmt.Sprint(err))
		return nil
	}

	transactionsTable.RemoveRow(curRow)
	removeTag(transactionIds[curRow])
	updateTransactionsTable()
	transactionsTable.Select(r, 0)

	return nil
}

func handleCommitTransaction(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	id := transactionIds[r]

	if transaction.Attributes(id).Committed {
		transaction.Uncommit(id)
	} else {
		transaction.Commit(id, initialBalance)
	}
	updateTransactionsTable()
	return nil
}

func handleCommitSingleTransaction(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	id := transactionIds[r]

	var err error
	if transaction.Attributes(id).Committed {
		err = transaction.UncommitSingle(id)
	} else {
		err = transaction.CommitSingle(id)
	}

	if err != nil {
		printStatus(fmt.Sprint(err))
		return nil
	}

	updateTransactionsTable()
	return nil
}

func handleEditCents(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if isCommitted(r) {
		printStatus("Cannot edit committed transactions")
		return nil
	}

	openPrompt("Amount:", selectedAmount(), func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		r, _ := transactionsTable.GetSelection()
		updateCents(promptInputField.GetText(), []int64{transactionIds[r]})
		return nil
	}, acceptanceFunction.Cents)

	return nil
}

func handleEditCategory(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if isCommitted(r) {
		printStatus("Cannot edit committed transactions")
		return nil
	}

	openPrompt("Category:", selectedCategory(), func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		r, _ := transactionsTable.GetSelection()
		updateCategory(promptInputField.GetText(), []int64{transactionIds[r]})
		return nil
	}, acceptanceFunction.Category)

	return nil
}

func handleEditDescription(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if isCommitted(r) {
		printStatus("Cannot edit committed transactions")
		return nil
	}

	openPrompt("Description:", selectedDescription(), func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		r, _ := transactionsTable.GetSelection()
		updateDescription(promptInputField.GetText(), []int64{transactionIds[r]})
		return nil
	}, acceptanceFunction.Description)

	return nil
}

func handleEditDate(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if isCommitted(r) {
		printStatus("Cannot edit committed transactions")
		return nil
	}

	openPrompt("Date:", selectedDate(), func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		r, _ := transactionsTable.GetSelection()
		updateDate(promptInputField.GetText(), []int64{transactionIds[r]})
		return nil
	}, acceptanceFunction.Date)

	return nil
}

func handleTagToggle(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	if isCommitted(r) {
		printStatus("Cannot tag committed transactions")
		return nil
	}

	if isTagged(transactionIds[r]) {
		removeTag(transactionIds[r])
	} else {
		applyTag(transactionIds[r])
	}

	updateTransactionsTable()
	handleSelectNext(ev)

	return nil
}

func handleTagMatches(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt("Tag matched transactions:", selectedDescription(), func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		findQuery = promptInputField.GetText()
		startingRow, _ := transactionsTable.GetSelection()
		tagMatches(transactionIds[startingRow])
		return nil
	}, acceptanceFunction.Any)

	return nil
}

func handleTagCommand(ev *tcell.EventKey) *tcell.EventKey {
	screen := app.GetScreen()
	cmdEv := screen.PollEvent()

	startingRow, _ := transactionsTable.GetSelection()

	switch e := cmdEv.(type) {
	case *tcell.EventKey:
		switch e.Rune() {
		case 'x':
			for _, r := range taggedTransactionIds {
				selectTransaction(transactionIds[r])
				handleDeleteTransaction(e)
				startingRow, _ = transactionsTable.GetSelection()
			}
		case 't':
			for _, r := range taggedTransactionIds {
				removeTag(r)
			}
		case 'D':
			openPrompt("Category:", selectedCategory(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				category := dataf.Category(promptInputField.GetText())
				updateCategory(category, taggedTransactionIds)
				return nil
			}, acceptanceFunction.Category)
		case 'd':
			openPrompt("Description:", selectedDescription(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				updateDescription(promptInputField.GetText(), taggedTransactionIds)
				return nil
			}, acceptanceFunction.Description)
		case '=':
			openPrompt("Amount:", selectedAmount(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				cents := promptInputField.GetText()
				updateCents(cents, taggedTransactionIds)
				return nil
			}, acceptanceFunction.Description)
		case '@':
			openPrompt("Date:", selectedDate(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				date := promptInputField.GetText()
				updateDate(date, taggedTransactionIds)
				return nil
			}, acceptanceFunction.Cents)
		}
	}

	transactionsTable.Select(startingRow, 0)
	updateTransactionsTable()
	return nil
}

func handleCloseReport(ev *tcell.EventKey) *tcell.EventKey {
	panels.HidePanel("report")
	return nil
}

func handleOpenReport(ev *tcell.EventKey) *tcell.EventKey {
	updateReportTable()
	panels.ShowPanel("report")
	panels.SendToFront("report")
	return nil
}

func handleCommand(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt(":", "", func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		runCommand(promptInputField.GetText())
		return nil
	}, acceptanceFunction.Any)

	return nil
}

func handleRepeat(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt("Repeat pattern:", "1m", func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		repeatRuleValue = promptInputField.GetText()

		text := fmt.Sprintf("%d-12-31", db.DateLastUncommitted().Year())
		openPrompt("Repeat until:", text, func(ev *tcell.EventKey) *tcell.EventKey {
			panels.HidePanel("prompt")
			repeatUntilValue = dataf.Date(promptInputField.GetText())

			err := repeater.Insert(selectionId(), repeatRuleValue, repeatUntilValue)
			if err != nil {
				log.Fatal(err)
			}

			updateTransactionsTable()
			return nil
		}, acceptanceFunction.Date)

		return nil
	}, acceptanceFunction.Any)
	return nil
}

func handleOpenHelp(ev *tcell.EventKey) *tcell.EventKey {
	openHelp()
	return nil
}

func handleHelpScroll(ev *tcell.EventKey) *tcell.EventKey {
	_, _, _, h := helpTextView.GetInnerRect()
	offset, _ := helpTextView.GetScrollOffset()
	switch ev.Rune() {
	case ' ', 'd':
		pos := offset + (h / 2)
		helpTextView.ScrollTo(pos, 0)
	case 'u':
		pos := offset - (h / 2)
		helpTextView.ScrollTo(pos, 0)
	}

	return nil
}

func handleReportSelectNext(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := reportTable.GetSelection()
	if r < reportTable.GetRowCount()-2 {
		reportTable.Select(r+1, 0)
	}

	return nil
}

func handleReportSelectPrev(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := reportTable.GetSelection()
	reportTable.Select(r-1, 0)
	return nil
}
