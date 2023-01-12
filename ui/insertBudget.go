package ui

import (
	"fmt"
	"jcb/config"
	"jcb/domain"
	"jcb/lib/budget"
	"jcb/lib/dates"
	"jcb/lib/validator"
	inputBindings "jcb/ui/input-bindings"

	dataf "jcb/lib/formatter/data"
	stringf "jcb/lib/formatter/string"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var insertBudgetForm *cview.Form
var insertBudgetInputFieldDate *cview.InputField
var insertBudgetInputFieldCategory *cview.InputField
var insertBudgetCheckBoxCumulative *cview.CheckBox
var insertBudgetInputFieldCents *cview.InputField
var insertBudgetInputFieldNotes *cview.InputField

func handleOpenInsertBudget(ev *tcell.EventKey) *tcell.EventKey {
	openInsertBudget()
	return nil
}

func openInsertBudget() {
	panels.ShowPanel("insertBudget")
	panels.SendToFront("insertBudget")

	curRow, _ := transactionsTable.GetSelection()
	curDate := dataf.Date(transactionsTable.GetCell(curRow, 1).GetText())

	insertBudgetInputFieldDate.SetText(stringf.Date(curDate))
	insertBudgetInputFieldCents.SetText("")
	insertBudgetInputFieldNotes.SetText("")
	insertBudgetInputFieldCategory.SetText("")
	insertBudgetCheckBoxCumulative.SetChecked(false)
}

func closeInsertBudget() {
	panels.HidePanel("insertBudget")
	insertBudgetForm.SetFocus(0)
}

func checkInsertBudgetForm() bool {
	var err error

	err = validator.Date(insertBudgetInputFieldDate.GetText())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return false
	}

	err = validator.Cents(insertBudgetInputFieldCents.GetText())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return false
	}

	return true
}

func readInsertBudgetForm() domain.Budget {
	date := dataf.Date(insertBudgetInputFieldDate.GetText())
	cents := dataf.Cents(insertBudgetInputFieldCents.GetText())
	notes := dataf.Notes(insertBudgetInputFieldNotes.GetText())
	category := dataf.Category(insertBudgetInputFieldCategory.GetText())
	cumulative := false

	return domain.Budget{-1, date, category, cents, notes, cumulative}
}

func handleInsertBudgetTransaction(ev *tcell.EventKey) *tcell.EventKey {
	if !checkInsertBudgetForm() {
		return nil
	}

	var id int64
	var err error

	id, err = budget.Insert(readInsertBudgetForm())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return nil
	}

	updateTransactionsTable()
	selectTransaction(id)

	closeInsertBudget()
	return nil
}

func createInsertBudgetForm() *cview.Form {
	insertBudgetForm = cview.NewForm()
	insertBudgetForm.SetCancelFunc(closeInsertBudget)

	insertBudgetInputFieldDate = cview.NewInputField()
	insertBudgetInputFieldDate.SetLabel("Date:")
	insertBudgetInputFieldDate.SetText(dates.LastCommitted().Format("2006-01-02"))
	insertBudgetInputFieldDate.SetFieldWidth(11)

	insertBudgetInputFieldCategory = cview.NewInputField()
	insertBudgetInputFieldCategory.SetLabel("Category:")
	insertBudgetInputFieldCategory.SetFieldWidth(0)

	insertBudgetInputFieldCents = cview.NewInputField()
	insertBudgetInputFieldCents.SetLabel("Amount:")
	insertBudgetInputFieldCents.SetFieldWidth(10)

	insertBudgetCheckBoxCumulative = cview.NewCheckBox()
	insertBudgetCheckBoxCumulative.SetLabel("Cumulative:")

	insertBudgetInputFieldNotes = cview.NewInputField()
	insertBudgetInputFieldNotes.SetLabel("Notes:")
	insertBudgetInputFieldNotes.SetFieldWidth(0)

	insertBudgetForm.AddFormItem(insertBudgetInputFieldDate)
	insertBudgetForm.AddFormItem(insertBudgetInputFieldCategory)
	insertBudgetForm.AddFormItem(insertBudgetInputFieldCents)
	insertBudgetForm.AddFormItem(insertBudgetCheckBoxCumulative)
	insertBudgetForm.AddFormItem(insertBudgetInputFieldNotes)

	insertBudgetForm.SetBorder(true)
	insertBudgetForm.SetBorderAttributes(tcell.AttrBold)
	insertBudgetForm.SetRect(15, 4, config.MAX_WIDTH-(15*2), 13)
	insertBudgetForm.SetTitleAlign(cview.AlignCenter)
	insertBudgetForm.SetTitleColor(config.COLOR_TITLE_FG)
	insertBudgetForm.SetTitle(" Insert Budget ")
	insertBudgetForm.SetWrapAround(true)
	insertBudgetForm.SetLabelColor(config.COLOR_FORM_LABLE_FG)
	insertBudgetForm.SetFieldBackgroundColor(config.COLOR_FORMFIELD_BG)
	insertBudgetForm.SetFieldBackgroundColorFocused(config.COLOR_FORMFIELD_FOCUSED_BG)

	c := inputBindings.Configuration(HandleInputFormCustomBindings)
	c.SetKey(0, tcell.KeyEnter, handleInsertBudgetTransaction)

	insertBudgetInputFieldDate.SetInputCapture(c.Capture)
	insertBudgetInputFieldCents.SetInputCapture(c.Capture)
	insertBudgetInputFieldNotes.SetInputCapture(c.Capture)
	insertBudgetInputFieldCategory.SetInputCapture(c.Capture)

	return insertBudgetForm
}
