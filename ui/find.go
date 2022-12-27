package ui

import (
	"fmt"
	"jcb/config"
	"strings"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var findForm *cview.Form
var findInputField *cview.InputField
var findQuery string

func handleOpenFind(ev *tcell.EventKey) *tcell.EventKey {
	panels.ShowPanel("find")
	panels.SendToFront("find")

	c := cbind.NewConfiguration()
	switch ev.Rune() {
	case '/':
		findInputField.SetLabel("/")
		c.SetKey(0, tcell.KeyEnter, handleFindForwards)
		findInputField.SetText("")
	case '?':
		findInputField.SetLabel("?")
		c.SetKey(0, tcell.KeyEnter, handleFindBackwards)
		findInputField.SetText("")
	case 'T':
		findInputField.SetLabel("Tag matched transactions:")
		findInputField.SetText(selectedDescription())
		c.SetKey(0, tcell.KeyEnter, handleTagMatches)
	}
	findInputField.SetInputCapture(c.Capture)
	findForm.SetFocus(0)
	return nil
}

func handleCloseFind() {
	panels.HidePanel("find")
}

func handleFindForwards(ev *tcell.EventKey) *tcell.EventKey {
	findQuery = findInputField.GetText()
	handleCloseFind()
	handleSelectNextMatch(ev)

	return nil
}

func handleFindBackwards(ev *tcell.EventKey) *tcell.EventKey {
	findQuery = findInputField.GetText()
	handleCloseFind()
	handleSelectPrevMatch(ev)

	return nil
}

func handleSelectNextMatch(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()

	for i := curRow + 1; i != curRow; i++ {
		if strings.Contains(strings.ToLower(transactionsTable.GetCell(i, config.DESCRIPTION_COLUMN).GetText()), strings.ToLower(findQuery)) {
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

func handleSelectPrevMatch(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()

	for i := curRow - 1; i != curRow; i-- {
		if strings.Contains(strings.ToLower(transactionsTable.GetCell(i, config.DESCRIPTION_COLUMN).GetText()), strings.ToLower(findQuery)) {
			transactionsTable.Select(i, 0)
			break
		}

		if i == 0 {
			i = len(transactionIds) - 1
		}
	}

	return nil
}

func handleTagMatches(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := transactionsTable.GetSelection()
	findQuery = findInputField.GetText()

	matchCount := 0

	for i := 1; i < len(transactionIds); i++ {
		if isCommitted(i) || isTagged(i) {
			continue
		}

		if strings.Contains(strings.ToLower(transactionsTable.GetCell(i, config.DESCRIPTION_COLUMN).GetText()), strings.ToLower(findQuery)) {
			matchCount += 1
			applyTag(i)
		}
	}

	transactionsTable.Select(curRow, 0)

	printStatus(fmt.Sprintf("Tagged %d transactions", matchCount))

	updateTransactionsTable()
	handleCloseFind()

	return nil
}

func createFindForm() *cview.Form {
	findForm = cview.NewForm()
	findForm.SetBorder(false)
	findForm.SetCancelFunc(handleCloseFind)
	findForm.SetItemPadding(0)
	findForm.SetPadding(0, 0, 0, 0)
	findForm.SetLabelColor(tcell.ColorWhite)
	findForm.SetFieldBackgroundColor(tcell.ColorBlack)
	findForm.SetFieldBackgroundColorFocused(tcell.ColorBlack)

	findInputField = cview.NewInputField()
	findInputField.SetFieldWidth(24)

	findForm.AddFormItem(findInputField)

	return findForm
}
