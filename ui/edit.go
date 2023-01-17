package ui

import (
	"fmt"
	"jcb/config"

	dataf "jcb/lib/formatter/data"
	"jcb/lib/transaction"
	"jcb/lib/validate"
	"jcb/lib/validator"
	inputBindings "jcb/ui/input-bindings"

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
	t, _ := transaction.Find(selectionId())
	if t.IsCommitted() {
		printStatus("Cannot edit a committed transaction")
		return
	}

	panels.ShowPanel("edit")
	panels.SendToFront("edit")

	editForm.SetTitle(fmt.Sprintf(" Edit Transaction (%d) ", t.GetID()))
	editInputFieldDate.SetText(t.Date.GetText())
	editInputFieldDescription.SetText(t.Description.GetText())
	editInputFieldCents.SetText(t.GetAmount(false))
	editInputFieldNotes.SetText(t.GetNotes())
	editInputFieldCategory.SetText(t.GetCategory(false))
}

func closeEdit() {
	panels.HidePanel("edit")
	editForm.SetFocus(0)
}

func readEditForm() *transaction.Transaction {
	t := new(transaction.Transaction)
	t.Date.SetText(editInputFieldDate.GetText())
	t.Description.SetText(editInputFieldDescription.GetText())
	t.SetCents(dataf.Cents(editInputFieldCents.GetText()))
	t.SetNotes(dataf.Notes(editInputFieldNotes.GetText()))
	t.SetCategory(dataf.Category(editInputFieldCategory.GetText()))

	return t
}

func handleEditTransaction(ev *tcell.EventKey) *tcell.EventKey {
	if !checkEditForm() {
		return nil
	}

	t := readEditForm()
	err := t.Save()
	if err == nil {
		updateTransactionsTable()
		selectTransaction(t.GetID())
		closeEdit()
	} else {
		printStatus(fmt.Sprint(err))
	}
	return nil
}

func checkEditForm() bool {
	var err error
	err = validate.Date(editInputFieldDate.GetText())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return false
	}

	err = validate.Description(editInputFieldDescription.GetText())
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
	editForm.SetCancelFunc(closeEdit)

	editInputFieldDate = cview.NewInputField()
	editInputFieldDate.SetLabel("Date:")
	editInputFieldDate.SetFieldWidth(11)

	editInputFieldCategory = cview.NewInputField()
	editInputFieldCategory.SetLabel("Category:")
	editInputFieldCategory.SetFieldWidth(0)

	editInputFieldDescription = cview.NewInputField()
	editInputFieldDescription.SetLabel("Description:")
	editInputFieldDescription.SetFieldWidth(0)

	editInputFieldCents = cview.NewInputField()
	editInputFieldCents.SetLabel("Amount:")
	editInputFieldCents.SetFieldWidth(10)

	editInputFieldNotes = cview.NewInputField()
	editInputFieldNotes.SetLabel("Notes:")
	editInputFieldNotes.SetFieldWidth(0)

	editForm.AddFormItem(editInputFieldDate)
	editForm.AddFormItem(editInputFieldCategory)
	editForm.AddFormItem(editInputFieldDescription)
	editForm.AddFormItem(editInputFieldCents)
	editForm.AddFormItem(editInputFieldNotes)

	editForm.SetBorder(true)
	editForm.SetBorderAttributes(tcell.AttrBold)
	editForm.SetRect(15, 4, config.MAX_WIDTH-(15*2), 13)
	editForm.SetTitleAlign(cview.AlignCenter)
	editForm.SetTitleColor(config.COLOR_TITLE_FG)
	editForm.SetTitle(" Edit Transaction ")
	editForm.SetWrapAround(true)
	editForm.SetLabelColor(config.COLOR_FORM_LABLE_FG)
	editForm.SetFieldBackgroundColor(config.COLOR_FORMFIELD_BG)
	editForm.SetFieldBackgroundColorFocused(config.COLOR_FORMFIELD_FOCUSED_BG)

	c := inputBindings.Configuration(HandleInputFormCustomBindings)
	c.SetKey(tcell.ModNone, tcell.KeyEnter, handleEditTransaction)

	editInputFieldDate.SetInputCapture(c.Capture)
	editInputFieldCategory.SetInputCapture(c.Capture)
	editInputFieldDescription.SetInputCapture(c.Capture)
	editInputFieldCents.SetInputCapture(c.Capture)
	editInputFieldNotes.SetInputCapture(c.Capture)

	return editForm
}
