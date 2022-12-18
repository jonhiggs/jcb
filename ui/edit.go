package ui

import (
	"jcb/domain"

	dataf "jcb/lib/formatter/data"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var editForm *cview.Form
var editInputFieldDate *cview.InputField
var editInputFieldDescription *cview.InputField
var editInputFieldCents *cview.InputField

func handleOpenEdit(ev *tcell.EventKey) *tcell.EventKey {
	panels.ShowPanel("edit")
	return nil
}

func handleCloseEdit() {
	panels.HidePanel("edit")
	editForm.SetFocus(0)
	return
}

func readEditForm() domain.Transaction {
	date := dataf.Date(editInputFieldDate.GetText())
	description := dataf.Description(editInputFieldDescription.GetText())
	cents := dataf.Cents(editInputFieldCents.GetText())

	return domain.Transaction{
		-1,
		date,
		description,
		cents,
	}
}

//func handleSaveTransaction() {
//	for _, t := range readInsertForm() {
//		id, err := transaction.Insert(t)
//		if err == nil {
//			updateTransactionsTable()
//			selectTransaction(id)
//		} else {
//			status.SetText(fmt.Sprint(err))
//		}
//	}
//
//	handleCloseInsert()
//}

func createEditForm() *cview.Form {
	editForm = cview.NewForm()
	editForm.SetCancelFunc(handleCloseEdit)

	editInputFieldDate = cview.NewInputField()
	editInputFieldDate.SetLabel("Date:")
	editInputFieldDate.SetText("2022-")
	editInputFieldDate.SetFieldWidth(11)
	editInputFieldDate.SetAcceptanceFunc(dateInputFieldAcceptance)
	editInputFieldDate.SetChangedFunc(dateInputFieldChanged)

	editInputFieldDescription = cview.NewInputField()
	editInputFieldDescription.SetLabel("Description:")
	editInputFieldDescription.SetFieldWidth(0)

	editInputFieldCents = cview.NewInputField()
	editInputFieldCents.SetLabel("Amount:")
	editInputFieldCents.SetFieldWidth(6)

	editForm.AddFormItem(insertInputFieldDate)
	editForm.AddFormItem(insertInputFieldDescription)
	editForm.AddFormItem(insertInputFieldCents)

	editForm.AddButton("Save", handleSaveTransaction)
	editForm.AddButton("Quit", handleCloseInsert)
	editForm.SetBorder(true)
	editForm.SetBorderAttributes(tcell.AttrBold)
	editForm.SetRect(6, 4, 45, 16)
	editForm.SetTitleAlign(cview.AlignCenter)
	editForm.SetTitle(" Edit Transaction ")
	editForm.SetWrapAround(true)
	editForm.SetFieldBackgroundColor(tcell.Color242)
	editForm.SetFieldBackgroundColorFocused(tcell.ColorRed)

	return editForm
}
