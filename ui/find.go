package ui

import (
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

	findInputField.SetLabel(string(ev.Rune()))
	c := cbind.NewConfiguration()
	switch ev.Rune() {
	case '/':
		c.SetKey(0, tcell.KeyEnter, handleFindForwards)
	case '?':
		c.SetKey(0, tcell.KeyEnter, handleFindBackwards)
	}
	findInputField.SetInputCapture(c.Capture)

	findInputField.SetText("")
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
