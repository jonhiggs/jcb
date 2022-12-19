package ui

import (
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
	findInputField.SetText("")
	findForm.SetFocus(0)
	return nil
}

func handleCloseFind() {
	panels.HidePanel("find")
}

func handleFind(ev *tcell.EventKey) *tcell.EventKey {
	findQuery = findInputField.GetText()
	handleCloseFind()
	handleSelectNextMatch(ev)

	return nil
}

func handleSelectNextMatch(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := table.GetSelection()

	for i := curRow + 1; i != curRow; i++ {
		if strings.Contains(table.GetCell(i, 1).GetText(), findQuery) {
			table.Select(i, 0)
			return nil
		}

		if i == len(transactionIds) {
			i = 0
		}
	}

	status.SetText("No matches found")
	panels.ShowPanel("status")
	panels.SendToFront("transactions")

	return nil
}

func handleSelectPrevMatch(ev *tcell.EventKey) *tcell.EventKey {
	curRow, _ := table.GetSelection()

	for i := curRow - 1; i != curRow; i-- {
		if strings.Contains(table.GetCell(i, 1).GetText(), findQuery) {
			table.Select(i, 0)
			break
		}

		if i == 0 {
			i = len(transactionIds) - 1
		}
	}

	return nil
}

func createFindForm(direction string) *cview.Form {
	findForm = cview.NewForm()
	findForm.SetBorder(false)
	findForm.SetCancelFunc(handleCloseFind)
	findForm.SetItemPadding(0)
	findForm.SetPadding(0, 0, 0, 0)
	findForm.SetLabelColor(tcell.ColorWhite)
	findForm.SetFieldBackgroundColor(tcell.ColorBlack)
	findForm.SetFieldBackgroundColorFocused(tcell.ColorBlack)

	findInputField = cview.NewInputField()
	findInputField.SetLabel("/")
	findInputField.SetFieldWidth(24)

	findForm.AddFormItem(findInputField)

	c := cbind.NewConfiguration()
	c.SetKey(0, tcell.KeyEnter, handleFind)
	findInputField.SetInputCapture(c.Capture)

	return findForm
}
