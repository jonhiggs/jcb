package ui

import (
	"fmt"
	"jcb/config"
	"jcb/lib/transaction"
	inputBindings "jcb/ui/input-bindings"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var insertForm *cview.Form
var insertInputFieldDate *cview.InputField
var insertInputFieldDescription *cview.InputField
var insertInputFieldCents *cview.InputField
var insertInputFieldNotes *cview.InputField
var insertInputFieldCategory *cview.InputField

func handleOpenInsert(ev *tcell.EventKey) *tcell.EventKey {
	openInsert()
	return nil
}

func openInsert() {
	panels.ShowPanel("insert")
	panels.SendToFront("insert")

	curRow, _ := transactionsTable.GetSelection()
	curDate := new(transaction.Date)
	curDate.SetText(transactionsTable.GetCell(curRow, 1).GetText())

	insertInputFieldDate.SetText(curDate.GetText())
	insertInputFieldDescription.SetText("")
	insertInputFieldCents.SetText("")
	insertInputFieldNotes.SetText("")
	insertInputFieldCategory.SetText("")
}

func closeInsert() {
	panels.HidePanel("insert")
	insertForm.SetFocus(0)
}

func readInsertForm() (*transaction.Transaction, error) {
	var err error
	t := new(transaction.Transaction)

	err = t.SetText(
		[]string{
			insertInputFieldDate.GetText(),
			insertInputFieldCategory.GetText(),
			insertInputFieldDescription.GetText(),
			insertInputFieldCents.GetText(),
			insertInputFieldNotes.GetText(),
		})

	return t, err
}

func handleInsertTransaction(ev *tcell.EventKey) *tcell.EventKey {
	t, err := readInsertForm()

	if err != nil {
		printStatus(fmt.Sprint(err))
	} else {
		err := t.Save()
		if err == nil {
			updateTransactionsTable()
			selectTransaction(t.Id)
			closeInsert()
		}
	}

	return nil
}

func createInsertForm() *cview.Form {
	insertForm = cview.NewForm()
	insertForm.SetCancelFunc(closeInsert)

	lastCommitted, err := transaction.FindLastCommitted()
	if err != nil {
		panic("FIXME: handle when there are no committed transactions")
	}

	insertInputFieldDate = cview.NewInputField()
	insertInputFieldDate.SetLabel("Date:")
	insertInputFieldDate.SetText(lastCommitted.Date.GetText())
	insertInputFieldDate.SetFieldWidth(11)

	insertInputFieldCategory = cview.NewInputField()
	insertInputFieldCategory.SetLabel("Category:")
	insertInputFieldCategory.SetFieldWidth(0)

	insertInputFieldDescription = cview.NewInputField()
	insertInputFieldDescription.SetLabel("Description:")
	insertInputFieldDescription.SetFieldWidth(0)

	insertInputFieldCents = cview.NewInputField()
	insertInputFieldCents.SetLabel("Amount:")
	insertInputFieldCents.SetFieldWidth(10)

	insertInputFieldNotes = cview.NewInputField()
	insertInputFieldNotes.SetLabel("Notes:")
	insertInputFieldNotes.SetFieldWidth(0)

	insertForm.AddFormItem(insertInputFieldDate)
	insertForm.AddFormItem(insertInputFieldCategory)
	insertForm.AddFormItem(insertInputFieldDescription)
	insertForm.AddFormItem(insertInputFieldCents)
	insertForm.AddFormItem(insertInputFieldNotes)

	insertForm.SetBorder(true)
	insertForm.SetBorderAttributes(tcell.AttrBold)
	insertForm.SetRect(15, 4, config.MAX_WIDTH-(15*2), 13)
	insertForm.SetTitleAlign(cview.AlignCenter)
	insertForm.SetTitleColor(config.COLOR_TITLE_FG)
	insertForm.SetTitle(" Insert Transaction ")
	insertForm.SetWrapAround(true)
	insertForm.SetLabelColor(config.COLOR_FORM_LABLE_FG)
	insertForm.SetFieldBackgroundColor(config.COLOR_FORMFIELD_BG)
	insertForm.SetFieldBackgroundColorFocused(config.COLOR_FORMFIELD_FOCUSED_BG)

	c := inputBindings.Configuration(HandleInputFormCustomBindings)
	c.SetKey(0, tcell.KeyEnter, handleInsertTransaction)

	insertInputFieldDate.SetInputCapture(c.Capture)
	insertInputFieldDescription.SetInputCapture(c.Capture)
	insertInputFieldCents.SetInputCapture(c.Capture)
	insertInputFieldNotes.SetInputCapture(c.Capture)
	insertInputFieldCategory.SetInputCapture(c.Capture)

	return insertForm
}
