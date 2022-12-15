package ui

import (
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var insertForm *cview.Form
var insertInputFieldDate *cview.InputField
var insertInputFieldDescription *cview.InputField
var insertInputFieldAmount *cview.InputField
var insertInputFieldRepeatRule *cview.InputField
var insertInputFieldRepeatUntil *cview.InputField

func handleOpenInsert(ev *tcell.EventKey) *tcell.EventKey {
	panels.ShowPanel("insert")
	return nil
}

func handleCloseInsert() {
	panels.HidePanel("insert")
	insertInputFieldDate.SetText("")
	insertInputFieldDescription.SetText("")
	insertInputFieldAmount.SetText("")
	insertInputFieldRepeatRule.SetText("0d")
	insertInputFieldRepeatUntil.SetText("2022-12-31")
	insertForm.SetFocus(0)
	return
}

func createInsertForm() *cview.Form {
	insertForm = cview.NewForm()

	insertInputFieldDate = cview.NewInputField()
	insertInputFieldDate.SetLabel("Date:")
	insertInputFieldDate.SetFieldWidth(11)
	insertInputFieldDate.SetFieldNote("Your complete address")

	insertInputFieldDescription = cview.NewInputField()
	insertInputFieldDescription.SetLabel("Description:")
	insertInputFieldDescription.SetFieldWidth(6)

	insertInputFieldAmount = cview.NewInputField()
	insertInputFieldAmount.SetLabel("Amount:")
	insertInputFieldAmount.SetFieldWidth(6)

	insertInputFieldRepeatRule = cview.NewInputField()
	insertInputFieldRepeatRule.SetLabel("Repeat Every:")
	insertInputFieldRepeatRule.SetFieldWidth(4)
	insertInputFieldRepeatRule.SetText("0d")

	insertInputFieldRepeatUntil = cview.NewInputField()
	insertInputFieldRepeatUntil.SetLabel("Repeat Until:")
	insertInputFieldRepeatUntil.SetFieldWidth(11)
	insertInputFieldRepeatUntil.SetText("2022-12-31")

	insertForm.AddFormItem(insertInputFieldDate)
	insertForm.AddFormItem(insertInputFieldDescription)
	insertForm.AddFormItem(insertInputFieldAmount)
	insertForm.AddFormItem(insertInputFieldRepeatRule)
	insertForm.AddFormItem(insertInputFieldRepeatUntil)

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
