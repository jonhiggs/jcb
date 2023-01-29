package ui

import (
	"fmt"
	"jcb/config"
	"jcb/lib/find"
	"jcb/lib/transaction"
	"jcb/lib/validate"
	"jcb/ui/acceptanceFunction"
	inputBindings "jcb/ui/input-bindings"
	"regexp"
	"strconv"
	"strings"
	"time"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var repeatRuleValue string
var repeatUntilValue time.Time

func handleOpenTransactions(ev *tcell.EventKey) *tcell.EventKey {
	panels.ShowPanel("transactions")
	panels.ShowPanel("info")
	panels.HidePanel("status")
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
	for i, t := range transactions {
		if !t.IsCommitted() {
			selectTransaction(int64(i))
			return nil
		}
	}
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

		if i == len(transactions) {
			i = 0
		}
	}

	return nil
}

func handleSelectMonthPrev(ev *tcell.EventKey) *tcell.EventKey {
	curDate := selectionTransaction().Date.GetText()
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
	curDate := selectionTransaction().Date.GetText()
	curMonth, _ := strconv.Atoi(curDate[5:7])
	curYear, _ := strconv.Atoi(curDate[0:4])

	r, _ := transactionsTable.GetSelection()
	for i := r; i < len(transactions); i++ {
		date := transactionsTable.GetCell(i, config.DATE_COLUMN).GetText()
		month, _ := strconv.Atoi(date[5:7])
		year, _ := strconv.Atoi(date[0:4])

		if month > curMonth || year > curYear {
			transactionsTable.Select(i, 0)
			return nil
		}
	}
	transactionsTable.Select(len(transactions)-1, 0)

	return nil
}

func handleSelectYearPrev(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()
	selectedDate := new(transaction.Date)
	selectedDate.SetText(transactionsTable.GetCell(curRow, config.DATE_COLUMN).GetText())

	for i := curRow; i > 0; i-- {
		curDate := new(transaction.Date)
		curDate.SetText(transactionsTable.GetCell(i, config.DATE_COLUMN).GetText())
		if curDate.Year() != selectedDate.Year() {
			transactionsTable.Select(i, 0)
			return nil
		}
	}

	transactionsTable.Select(1, 0)

	return nil
}

func handleSelectYearNext(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()
	selectedDate := new(transaction.Date)
	selectedDate.SetText(transactionsTable.GetCell(curRow, config.DATE_COLUMN).GetText())

	for i := curRow; i < len(transactions)-1; i++ {
		curDate := new(transaction.Date)
		curDate.SetText(transactionsTable.GetCell(i, config.DATE_COLUMN).GetText())
		if curDate.Year() != selectedDate.Year() {
			transactionsTable.Select(i, 0)
			return nil
		}
	}

	return nil
}

func handleSelectModifiedPrev(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()

	for i := r - 1; i != r; i-- {
		if i == 0 {
			i = len(transactions) - 1
			continue
		}

		t, _ := transaction.Find(selectionId())
		if !t.IsSaved() {
			transactionsTable.Select(i, 0)
			return nil
		}
	}

	return nil
}

func handleSelectModifiedNext(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()

	for i := r + 1; i != r; i++ {
		t, _ := transaction.Find(selectionId())
		if !t.IsSaved() {
			transactionsTable.Select(i, 0)
			return nil
		}

		if i == len(transactions)-1 {
			i = 0
		}
	}

	return nil
}

func handleFindForwards(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt("/", "", func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")

		err := find.SetQuery(promptInputField.GetText())
		if err != nil {
			printStatus(fmt.Sprintf("%s", err))
			return nil
		}

		handleSelectMatchNext(ev)
		return nil
	})

	return nil
}

func handleFindBackwards(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt("?", "", func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")

		err := find.SetQuery(promptInputField.GetText())
		if err != nil {
			printStatus(fmt.Sprintf("%s", err))
			return nil
		}

		handleSelectMatchPrev(ev)
		return nil
	})

	return nil
}

func handleSelectMatchNext(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()

	for i := curRow + 1; i != curRow; i++ {

		if find.TableRowMatches(transactionsTable, i) {
			transactionsTable.Select(i, 0)
			return nil
		}

		if i == len(transactions)-1 {
			i = 0
		}
	}

	printStatus("No matches found")

	return nil
}

func handleSelectMatchPrev(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()

	for i := curRow - 1; i != curRow; i-- {
		if find.TableRowMatches(transactionsTable, i) {
			transactionsTable.Select(i, 0)
			return nil
		}

		if i == 0 {
			i = len(transactions) - 1
		}
	}

	printStatus("No matches found")

	return nil
}

func handleDeleteTransaction(ev *tcell.EventKey) *tcell.EventKey {
	id := selectionId()

	curRow, _ := transactionsTable.GetSelection()
	var r int
	if curRow == len(transactions)-1 {
		r = curRow - 1
	} else {
		r = curRow
	}

	t, _ := transaction.Find(id)
	err := t.Delete()
	if err != nil {
		printStatus(fmt.Sprint(err))
		return nil
	}

	transactionsTable.RemoveRow(curRow)
	removeTag(transactions[curRow].Id)
	updateTransactionsTable()
	transactionsTable.Select(r, 0)

	return nil
}

func handleCommitTransaction(ev *tcell.EventKey) *tcell.EventKey {
	var err error
	t, _ := transaction.Find(selectionId())

	if t.IsCommitted() {
		err = t.Uncommit()
	} else {
		err = t.Commit()
	}

	if err != nil {
		printStatus(fmt.Sprint(err))
		return nil
	}

	updateTransactionsTable()
	return nil
}

func handleCommitSingleTransaction(ev *tcell.EventKey) *tcell.EventKey {
	r, _ := transactionsTable.GetSelection()
	t, _ := transaction.Find(transactions[r].Id)
	err := t.Commit()
	if err != nil {
		printStatus(fmt.Sprint(err))
	}

	updateTransactionsTable()
	return nil
}

func handleEditSingleTransaction(ev *tcell.EventKey) *tcell.EventKey {
	t := selectionTransaction()
	if t.IsCommitted() {
		printStatus("Cannot edit committed transactions")
		return nil
	}

	var function func(s string, ts []*transaction.Transaction) []*transaction.Transaction
	var label string
	var dataType string
	var placeholderText string

	switch ev.Rune() {
	case '@':
		function = transaction.UpdateDate
		label = "Date:"
		dataType = "date"
		placeholderText = t.Date.GetText()
	case '=':
		function = transaction.UpdateCents
		label = "Amount:"
		dataType = "amount"
		placeholderText = t.Cents.GetText()
	case 'd':
		function = transaction.UpdateDescription
		label = "Description:"
		dataType = "description"
		placeholderText = t.Description.GetText()
	case 'D':
		function = transaction.UpdateCategory
		label = "Category:"
		dataType = "category"
		placeholderText = t.Category.GetText()
	}

	openPrompt(label, placeholderText, func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		t := []*transaction.Transaction{selectionTransaction()}
		modifiedTransactions := function(promptInputField.GetText(), t)
		printStatus(fmt.Sprintf("Updated %s for %d transactions", dataType, len(modifiedTransactions)))
		if len(modifiedTransactions) > 0 {
			for _, t := range modifiedTransactions {
				t.Save()
			}

			updateTransactionsTable()
		}
		return nil
	})

	return nil
}

func handleTagToggle(ev *tcell.EventKey) *tcell.EventKey {
	t := selectionTransaction()
	if t.IsCommitted() {
		printStatus("Cannot tag committed transactions")
		return nil
	}

	if isTagged(t.Id) {
		removeTag(t.Id)
	} else {
		applyTag(t.Id)
	}

	updateTransactionsTable()
	handleSelectNext(ev)

	return nil
}

func handleTagMatches(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt("Tag matched transactions:", selectionTransaction().Description.GetText(), func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")

		err := find.SetQuery(promptInputField.GetText())
		if err != nil {
			printStatus(fmt.Sprintf("%s", err))
			return nil
		}

		startingRow, _ := transactionsTable.GetSelection()
		tagMatches(transactions[startingRow].Id)
		return nil
	})

	return nil
}

func handleUntagMatches(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt("Untag matched transactions:", selectionTransaction().Description.GetText(), func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")

		err := find.SetQuery(promptInputField.GetText())
		if err != nil {
			printStatus(fmt.Sprintf("%s", err))
			return nil
		}

		startingRow, _ := transactionsTable.GetSelection()
		untagMatches(transactions[startingRow].Id)
		return nil
	})

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
				selectTransaction(transactions[r].Id)
				handleDeleteTransaction(e)
				startingRow, _ = transactionsTable.GetSelection()
			}
		case 't':
			for _, r := range taggedTransactionIds {
				removeTag(r)
			}
		case 'D':

			openPrompt("Category:", selectionTransaction().Category.GetText(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				modifiedTransactions := transaction.UpdateCategory(promptInputField.GetText(), taggedTransactions())
				if len(modifiedTransactions) > 0 {
					for _, t := range modifiedTransactions {
						t.Save()
					}
					updateTransactionsTable()
				}
				return nil
			})
		case 'd':
			openPrompt("Description:", selectionTransaction().Description.GetText(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				modifiedTransactions := transaction.UpdateDescription(promptInputField.GetText(), taggedTransactions())
				if len(modifiedTransactions) > 0 {
					for _, t := range modifiedTransactions {
						t.Save()
					}
					updateTransactionsTable()
				}
				return nil
			})
		case '=':
			openPrompt("Amount:", selectionTransaction().Cents.GetText(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				modifiedTransactions := transaction.UpdateCents(promptInputField.GetText(), taggedTransactions())
				if len(modifiedTransactions) > 0 {
					for _, t := range modifiedTransactions {
						t.Save()
					}
					updateTransactionsTable()
				}
				return nil
			})
		case '@':
			openPrompt("Date:", selectionTransaction().Date.GetText(), func(ev *tcell.EventKey) *tcell.EventKey {
				panels.HidePanel("prompt")
				modifiedTransactions := transaction.UpdateDate(promptInputField.GetText(), taggedTransactions())
				if len(modifiedTransactions) > 0 {
					for _, t := range modifiedTransactions {
						t.Save()
					}
					updateTransactionsTable()
				}
				return nil
			})
		}

		switch e.Key() {
		case tcell.KeyCtrlT:
			for _, r := range taggedTransactionIds {
				removeTag(r)
			}
		}

		openPrompt("Date:", selectionTransaction().Date.GetText(), func(ev *tcell.EventKey) *tcell.EventKey {
			panels.HidePanel("prompt")
			modifiedTransactions := transaction.UpdateDate(promptInputField.GetText(), taggedTransactions())
			if len(modifiedTransactions) > 0 {
				for _, t := range modifiedTransactions {
					t.Save()
				}
				updateTransactionsTable()
			}
			return nil
		})
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
	panels.HidePanel("info")
	panels.HidePanel("status")
	panels.ShowPanel("report")
	panels.SendToFront("report")
	return nil
}

func handleCommand(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt(":", "", func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		runCommand(promptInputField.GetText())
		return nil
	})

	return nil
}

func handleRepeat(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt("Repeat pattern:", "1m", func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		repeatRuleValue = promptInputField.GetText()
		err := validate.RepeatRule(repeatRuleValue)
		if err != nil {
			printStatus(fmt.Sprint(err))
			return nil
		}

		lastUncommitted, err := transaction.FindLastUncommitted()
		if err != nil {
			panic("You cannot make it here. You cannot repeat a committed transaction so there must be uncommitted transactions!")
		}

		text := fmt.Sprintf("%d-12-31", lastUncommitted.Date.Year())
		openPrompt("Repeat until:", text, func(ev *tcell.EventKey) *tcell.EventKey {
			panels.HidePanel("prompt")
			repeatUntilString := promptInputField.GetText()

			if !transaction.ValidDateString(repeatUntilString) {
				printStatus(fmt.Sprint(err))
				return nil
			}

			repeatUntil := new(transaction.Date)
			repeatUntil.SetText(repeatUntilString)

			transactionSlice, _ := selectionTransaction().Repeat(repeatRuleValue, repeatUntil.GetValue())
			for _, tr := range transactionSlice {
				tr.Save()
			}

			updateTransactionsTable()
			return nil
		})

		return nil
	})
	return nil
}

func handleOpenHelp(ev *tcell.EventKey) *tcell.EventKey {
	panels.HidePanel("info")
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

func HandleInputFormCustomBindings(ev *tcell.EventKey) *tcell.EventKey {
	pn, _ := panels.GetFrontPanel()
	var field *cview.InputField
	switch pn {
	case "edit":
		fieldId, _ := editForm.GetFocusedItemIndex()
		field = editForm.GetFormItem(fieldId).(*cview.InputField)
	case "insert":
		fieldId, _ := insertForm.GetFocusedItemIndex()
		field = insertForm.GetFormItem(fieldId).(*cview.InputField)
	case "prompt":
		fieldId, _ := promptForm.GetFocusedItemIndex()
		field = promptForm.GetFormItem(fieldId).(*cview.InputField)
	}

	acceptanceFunc := acceptanceFunction.FieldFunc(field)
	originalText := field.GetText()

	switch ev.Key() {
	case tcell.KeyCtrlD:
		inputBindings.DeleteChar(field)
	case tcell.KeyCtrlF:
		inputBindings.ForwardChar(field)
	case tcell.KeyCtrlB:
		inputBindings.BackwardChar(field)
	case tcell.KeyCtrlK:
		inputBindings.KillLine(field)
	case tcell.KeyCtrlW:
		inputBindings.UnixWordRubout(field)
	case tcell.KeyCtrlY:
		inputBindings.Yank(field)
	}

	isChar, _ := regexp.MatchString(`[abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*(),\./<>?;':\"\[\]\{\}\-+]`, string(ev.Rune()))
	if ev.Modifiers() == tcell.ModAlt {
		switch ev.Key() {
		case tcell.KeyBackspace2:
			inputBindings.OtherUnixWordRubout(field)
		}

		switch ev.Rune() {
		case 'd':
			inputBindings.DeleteWord(field)
		case 'f':
			inputBindings.ForwardWord(field)
		case 'b':
			inputBindings.BackwardWord(field)
		}
	} else if isChar {
		// this is to workaround some bugs in cview that prevents a dash editing
		// inputs at or near symbols.
		var text string
		pos := field.GetCursorPosition()

		textSlice := strings.Split(field.GetText(), "")
		for i, c := range textSlice {
			if i == pos {
				text = text + string(ev.Rune())
			}
			text = text + c
		}

		if pos == len(text) {
			text = text + string(ev.Rune())
		}

		field.SetText(text)
		if pos < len(text) {
			field.SetCursorPosition(pos + 1)
		}
	}

	if !acceptanceFunc(field) {
		field.SetText(originalText)
	}

	return nil
}
