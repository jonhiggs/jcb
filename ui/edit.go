package ui

import (
	"fmt"
	"jcb/config"
	"jcb/domain"

	dataf "jcb/lib/formatter/data"
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"
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

func closeEdit() {
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
		selectTransaction(t.Id)
		closeEdit()
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
