package ui

import (
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var insertForm *cview.Form
var insertInputFieldDate *cview.InputField

func handleOpenInsert(ev *tcell.EventKey) *tcell.EventKey {
	panels.ShowPanel("insert")
	return nil
}

func handleCloseInsert() {
	panels.HidePanel("insert")
	insertInputFieldDate.SetText("")
	insertForm.SetFocus(0)
	return
}

func createInsertForm() *cview.Form {
	insertForm = cview.NewForm()

	insertInputFieldDate = cview.NewInputField()
	insertInputFieldDate.SetLabel("Date")
	insertInputFieldDate.SetFieldWidth(11)
	insertInputFieldDate.SetFieldNote("Your complete address")

	insertForm.AddFormItem(insertInputFieldDate)
	insertForm.AddInputField("Description:", "", 0, nil, nil)
	insertForm.AddInputField("Amount:", "", 6, nil, nil)
	insertForm.AddInputField("Repeat Every:", "", 4, nil, nil)
	insertForm.AddInputField("Repeat Until", "2022-12-31", 11, func(t string, c rune) bool { return true }, nil)
	insertForm.AddButton("Save", func() {})
	insertForm.AddButton("Quit", handleCloseInsert)
	insertForm.SetBorder(true)
	insertForm.SetBorderAttributes(tcell.AttrBold)
	insertForm.SetRect(4, 4, 50, 20)
	insertForm.SetTitleAlign(cview.AlignCenter)
	insertForm.SetTitle(" New Transaction ")
	insertForm.SetWrapAround(true)

	return insertForm
}
