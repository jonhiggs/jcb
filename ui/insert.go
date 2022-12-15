package ui

import (
	"jcb/domain"

	dataf "jcb/lib/formatter/data"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var insertForm *cview.Form
var insertInputFieldDate *cview.InputField
var insertInputFieldDescription *cview.InputField
var insertInputFieldCents *cview.InputField
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
	insertInputFieldCents.SetText("")
	insertInputFieldRepeatRule.SetText("0d")
	insertInputFieldRepeatUntil.SetText("2022-12-31")
	insertForm.SetFocus(0)
	return
}

func readInsertForm() domain.Transaction {
	return domain.Transaction{
		0,
		dataf.Date(insertInputFieldDate.GetText()),
		dataf.Description(insertInputFieldDescription.GetText()),
		dataf.Cents(insertInputFieldCents.GetText()),
	}
}

func createInsertForm() *cview.Form {
	insertForm = cview.NewForm()

	insertInputFieldDate = cview.NewInputField()
	insertInputFieldDate.SetLabel("Date:")
	insertInputFieldDate.SetFieldWidth(11)

	insertInputFieldDescription = cview.NewInputField()
	insertInputFieldDescription.SetLabel("Description:")
	insertInputFieldDescription.SetFieldWidth(0)

	insertInputFieldCents = cview.NewInputField()
	insertInputFieldCents.SetLabel("Amount:")
	insertInputFieldCents.SetFieldWidth(6)

	insertInputFieldRepeatRule = cview.NewInputField()
	insertInputFieldRepeatRule.SetLabel("Repeat Every:")
	insertInputFieldRepeatRule.SetFieldWidth(14)
	insertInputFieldRepeatRule.SetText("0d")
	insertInputFieldRepeatRule.SetFieldNote(`<number>(dwm)`)

	insertInputFieldRepeatUntil = cview.NewInputField()
	insertInputFieldRepeatUntil.SetLabel("Repeat Until:")
	insertInputFieldRepeatUntil.SetFieldWidth(11)
	insertInputFieldRepeatUntil.SetText("2022-12-31")

	insertForm.AddFormItem(insertInputFieldDate)
	insertForm.AddFormItem(insertInputFieldDescription)
	insertForm.AddFormItem(insertInputFieldCents)
	insertForm.AddFormItem(insertInputFieldRepeatRule)
	insertForm.AddFormItem(insertInputFieldRepeatUntil)

	insertForm.AddButton("Save", func() {})
	insertForm.AddButton("Quit", handleCloseInsert)
	insertForm.SetBorder(true)
	insertForm.SetBorderAttributes(tcell.AttrBold)
	insertForm.SetRect(4, 4, 45, 18)
	insertForm.SetTitleAlign(cview.AlignCenter)
	insertForm.SetTitle(" New Transaction ")
	insertForm.SetWrapAround(true)

	return insertForm
}
