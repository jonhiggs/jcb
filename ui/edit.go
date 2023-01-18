package ui

import (
	"fmt"
	"jcb/config"

	"jcb/lib/transaction"
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

	editForm.SetTitle(fmt.Sprintf(" Edit Transaction (%d) ", t.Id))
	editInputFieldDate.SetText(t.Date.GetText())
	editInputFieldDescription.SetText(t.Description.GetText())
	editInputFieldCents.SetText(t.Cents.GetText())
	editInputFieldNotes.SetText(t.Note.GetText())
	editInputFieldCategory.SetText(t.Category.GetText())
}

func closeEdit() {
	panels.HidePanel("edit")
	editForm.SetFocus(0)
}

func readEditForm() (*transaction.Transaction, error) {
	var err error
	t := new(transaction.Transaction)

	err = t.Date.SetText(editInputFieldDate.GetText())
	if err != nil {
		return t, fmt.Errorf("invalid date")
	}

	err = t.Description.SetText(editInputFieldDescription.GetText())
	if err != nil {
		return t, fmt.Errorf("invalid description")
	}

	err = t.Cents.SetText(editInputFieldCents.GetText())
	if err != nil {
		return t, fmt.Errorf("invalid cents")
	}

	err = t.Note.SetText(editInputFieldNotes.GetText())
	if err != nil {
		return t, fmt.Errorf("invalid note")
	}

	err = t.Category.SetText(editInputFieldCategory.GetText())
	if err != nil {
		return t, fmt.Errorf("invalid category")
	}

	return t, nil
}

func handleEditTransaction(ev *tcell.EventKey) *tcell.EventKey {
	t, err := readEditForm()
	if err != nil {
		printStatus(fmt.Sprint(err))
	} else {
		err := t.Save()
		if err == nil {
			updateTransactionsTable()
			selectTransaction(t.Id)
			closeEdit()
		}
	}

	return nil
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
