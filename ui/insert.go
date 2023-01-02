package ui

import (
	"fmt"
	"jcb/config"
	"jcb/domain"
	"jcb/lib/dates"
	"jcb/lib/validator"
	acceptanceFunction "jcb/ui/acceptance-functions"
	inputBindings "jcb/ui/input-bindings"

	dataf "jcb/lib/formatter/data"
	stringf "jcb/lib/formatter/string"
	"jcb/lib/transaction"

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
	curDate := dataf.Date(transactionsTable.GetCell(curRow, 1).GetText())

	insertInputFieldDate.SetText(stringf.Date(curDate))
	insertInputFieldDescription.SetText("")
	insertInputFieldCents.SetText("")
	insertInputFieldNotes.SetText("")
	insertInputFieldCategory.SetText("")
}

func closeInsert() {
	panels.HidePanel("insert")
	insertForm.SetFocus(0)
}

func checkInsertForm() bool {
	var err error

	err = validator.Date(insertInputFieldDate.GetText())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return false
	}

	err = validator.Description(insertInputFieldDescription.GetText())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return false
	}

	err = validator.Cents(insertInputFieldCents.GetText())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return false
	}

	return true
}

func dateInputFieldChanged(s string) {
	return
}

func readInsertForm() domain.Transaction {
	date := dataf.Date(insertInputFieldDate.GetText())
	description := dataf.Description(insertInputFieldDescription.GetText())
	cents := dataf.Cents(insertInputFieldCents.GetText())
	notes := dataf.Notes(insertInputFieldNotes.GetText())
	category := dataf.Category(insertInputFieldCategory.GetText())

	return domain.Transaction{-1, date, description, cents, notes, category}
}

func handleInsertTransaction(ev *tcell.EventKey) *tcell.EventKey {
	if !checkInsertForm() {
		return nil
	}

	var id int64
	var err error

	id, err = transaction.Insert(readInsertForm())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return nil
	}

	updateTransactionsTable()
	selectTransaction(id)

	closeInsert()
	return nil
}

func createInsertForm() *cview.Form {
	insertForm = cview.NewForm()
	insertForm.SetCancelFunc(closeInsert)

	insertInputFieldDate = cview.NewInputField()
	insertInputFieldDate.SetLabel("Date:")
	insertInputFieldDate.SetText(dates.LastCommitted().Format("2006-01-02"))
	insertInputFieldDate.SetFieldWidth(11)
	insertInputFieldDate.SetAcceptanceFunc(acceptanceFunction.Date)
	insertInputFieldDate.SetChangedFunc(dateInputFieldChanged)

	insertInputFieldCategory = cview.NewInputField()
	insertInputFieldCategory.SetLabel("Category:")
	insertInputFieldCategory.SetFieldWidth(0)
	insertInputFieldCategory.SetAcceptanceFunc(acceptanceFunction.Category)

	insertInputFieldDescription = cview.NewInputField()
	insertInputFieldDescription.SetLabel("Description:")
	insertInputFieldDescription.SetFieldWidth(0)
	insertInputFieldDescription.SetAcceptanceFunc(acceptanceFunction.Description)

	insertInputFieldCents = cview.NewInputField()
	insertInputFieldCents.SetLabel("Amount:")
	insertInputFieldCents.SetFieldWidth(10)
	insertInputFieldCents.SetAcceptanceFunc(acceptanceFunction.Cents)

	insertInputFieldNotes = cview.NewInputField()
	insertInputFieldNotes.SetLabel("Notes:")
	insertInputFieldNotes.SetFieldWidth(0)
	insertInputFieldNotes.SetAcceptanceFunc(acceptanceFunction.Notes)

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
