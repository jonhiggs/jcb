package ui

import (
	"jcb/domain"
	"log"

	dataf "jcb/lib/formatter/data"
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var editForm *cview.Form
var editInputFieldDate *cview.InputField
var editInputFieldDescription *cview.InputField
var editInputFieldCents *cview.InputField

func handleOpenEdit() {
	if transaction.IsCommitted(selectionId()) {
		status.SetText("Cannot edit a committed transaction")
		return
	}

	panels.ShowPanel("edit")

	t, _ := transaction.Find(selectionId())
	ts := stringf.Transaction(t)

	editInputFieldDate.SetText(ts.Date)
	editInputFieldDescription.SetText(ts.Description)
	editInputFieldCents.SetText(ts.Cents)
}

func handleCloseEdit() {
	panels.HidePanel("edit")
	editForm.SetFocus(0)
}

func readEditForm() domain.Transaction {
	date := dataf.Date(editInputFieldDate.GetText())
	description := dataf.Description(editInputFieldDescription.GetText())
	cents := dataf.Cents(editInputFieldCents.GetText())

	return domain.Transaction{
		selectionId(),
		date,
		description,
		cents,
	}
}

func handleEditTransaction() {
	t := readEditForm()
	err := transaction.Edit(t)
	if err == nil {
		updateTransactionsTable()
		handleCloseEdit()
	} else {
		log.Fatal(err)
	}
}

func createEditForm() *cview.Form {
	editForm = cview.NewForm()
	editForm.SetCancelFunc(handleCloseEdit)

	editInputFieldDate = cview.NewInputField()
	editInputFieldDate.SetLabel("Date:")
	editInputFieldDate.SetFieldWidth(11)
	//editInputFieldDate.SetAcceptanceFunc(dateInputFieldAcceptance)
	//editInputFieldDate.SetChangedFunc(dateInputFieldChanged)

	editInputFieldDescription = cview.NewInputField()
	editInputFieldDescription.SetLabel("Description:")
	editInputFieldDescription.SetFieldWidth(0)

	editInputFieldCents = cview.NewInputField()
	editInputFieldCents.SetLabel("Amount:")
	editInputFieldCents.SetFieldWidth(6)

	editForm.AddFormItem(editInputFieldDate)
	editForm.AddFormItem(editInputFieldDescription)
	editForm.AddFormItem(editInputFieldCents)

	editForm.AddButton("Save", handleEditTransaction)
	editForm.AddButton("Close", handleCloseEdit)
	editForm.SetBorder(true)
	editForm.SetBorderAttributes(tcell.AttrBold)
	editForm.SetRect(6, 4, 45, 11)
	editForm.SetTitleAlign(cview.AlignCenter)
	editForm.SetTitle(" Edit Transaction ")
	editForm.SetWrapAround(true)
	editForm.SetFieldBackgroundColor(tcell.Color242)
	editForm.SetFieldBackgroundColorFocused(tcell.ColorRed)

	return editForm
}
