package ui

import (
	"fmt"
	"jcb/config"
	"jcb/domain"

	dataf "jcb/lib/formatter/data"
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"
	"jcb/lib/validator"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var editForm *cview.Form
var editInputFieldDate *cview.InputField
var editInputFieldDescription *cview.InputField
var editInputFieldCents *cview.InputField
var editInputFieldNotes *cview.InputField
var editInputFieldCategory *cview.InputField

func handleOpenEdit() {
	if transaction.Attributes(selectionId()).Committed {
		printStatus("Cannot edit a committed transaction")
		return
	}

	panels.ShowPanel("edit")
	panels.SendToFront("edit")

	t, _ := transaction.Find(selectionId())
	ts := stringf.Transaction(t)

	editInputFieldDate.SetText(ts.Date)
	editInputFieldDescription.SetText(ts.Description)
	editInputFieldCents.SetText(ts.Cents)
	editInputFieldNotes.SetText(transaction.Notes(selectionId()))
	editInputFieldCategory.SetText(ts.Category)
}

func handleCloseEdit() {
	panels.HidePanel("edit")
	editForm.SetFocus(0)
}

func readEditForm() domain.Transaction {
	date := dataf.Date(editInputFieldDate.GetText())
	description := dataf.Description(editInputFieldDescription.GetText())
	cents := dataf.Cents(editInputFieldCents.GetText())
	notes := dataf.Notes(editInputFieldNotes.GetText())
	category := dataf.Category(editInputFieldCategory.GetText())

	return domain.Transaction{
		selectionId(),
		date,
		description,
		cents,
		notes,
		category,
	}
}

func handleEditTransaction(ev *tcell.EventKey) *tcell.EventKey {
	if !checkEditForm() {
		return nil
	}

	t := readEditForm()
	err := transaction.Edit(t)
	if err == nil {
		updateTransactionsTable()
		handleCloseEdit()
	} else {
		printStatus(fmt.Sprint(err))
	}
	return nil
}

func checkEditForm() bool {
	var err error
	err = validator.Date(editInputFieldDate.GetText())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return false
	}

	err = validator.Description(editInputFieldDescription.GetText())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return false
	}

	err = validator.Cents(editInputFieldCents.GetText())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return false
	}

	return true
}

func createEditForm() *cview.Form {
	editForm = cview.NewForm()
	editForm.SetCancelFunc(handleCloseEdit)

	editInputFieldDate = cview.NewInputField()
	editInputFieldDate.SetLabel("Date:")
	editInputFieldDate.SetFieldWidth(11)

	editInputFieldCategory = cview.NewInputField()
	editInputFieldCategory.SetLabel("Category:")
	editInputFieldCategory.SetFieldWidth(0)

	editInputFieldDescription = cview.NewInputField()
	editInputFieldDescription.SetLabel("Description:")
	editInputFieldDescription.SetFieldWidth(0)
	editInputFieldDescription.SetAcceptanceFunc(cview.InputFieldMaxLength(config.DESC_MAX_LENGTH))

	editInputFieldCents = cview.NewInputField()
	editInputFieldCents.SetLabel("Amount:")
	editInputFieldCents.SetFieldWidth(10)
	editInputFieldCents.SetAcceptanceFunc(cview.InputFieldFloat)

	editInputFieldNotes = cview.NewInputField()
	editInputFieldNotes.SetLabel("Notes:")
	editInputFieldNotes.SetFieldWidth(0)
	editInputFieldNotes.SetAcceptanceFunc(cview.InputFieldMaxLength(config.NOTES_MAX_LENGTH))

	editForm.AddFormItem(editInputFieldDate)
	editForm.AddFormItem(editInputFieldCategory)
	editForm.AddFormItem(editInputFieldDescription)
	editForm.AddFormItem(editInputFieldCents)
	editForm.AddFormItem(editInputFieldNotes)

	editForm.SetBorder(true)
	editForm.SetBorderAttributes(tcell.AttrBold)
	editForm.SetRect(15, 4, config.MAX_WIDTH-(15*2), 13)
	editForm.SetTitleAlign(cview.AlignCenter)
	editForm.SetTitle(" Edit Transaction ")
	editForm.SetWrapAround(true)
	editForm.SetFieldBackgroundColor(tcell.Color242)
	editForm.SetFieldBackgroundColorFocused(tcell.ColorRed)

	c := cbind.NewConfiguration()
	c.SetKey(tcell.ModNone, tcell.KeyEnter, handleEditTransaction)
	c.SetKey(tcell.ModCtrl, tcell.KeyCtrlD, handleInputFormCustomBindings)
	c.SetKey(tcell.ModCtrl, tcell.KeyCtrlF, handleInputFormCustomBindings)
	c.SetKey(tcell.ModCtrl, tcell.KeyCtrlB, handleInputFormCustomBindings)
	editInputFieldDate.SetInputCapture(c.Capture)
	editInputFieldCategory.SetInputCapture(c.Capture)
	editInputFieldDescription.SetInputCapture(c.Capture)
	editInputFieldCents.SetInputCapture(c.Capture)
	editInputFieldNotes.SetInputCapture(c.Capture)

	return editForm
}
