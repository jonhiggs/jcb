package ui

import (
	"fmt"
	"jcb/domain"
	"regexp"
	"strings"

	dataf "jcb/lib/formatter/data"
	"jcb/lib/transaction"
	"jcb/lib/validator"

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

func handleSaveInsert() {
	return
}

func validateInsertForm(s string) {
	err := validator.Date(s)
	if err != nil {
		status.SetText(fmt.Sprint(err))
		insertInputFieldDate.SetLabelColor(tcell.ColorRed)
	} else {
		insertInputFieldDate.SetLabelColor(tcell.ColorGreen)
		status.SetText("")
	}
}

func dateInputFieldAcceptance(s string, c rune) bool {
	valid_char, _ := regexp.MatchString(`[\d\-]*`, string(c))
	if !valid_char {
		return false
	}

	switch len(s) {
	case 1:
		return c == '2'
	case 2:
		return c == '0'
	case 3:
		return c == '2'
	case 4:
		return c == '2'
	case 5, 8:
		return c == '-'
	case 6:
		return c == '0' || c == '1'
	case 7:
		var v bool
		if strings.HasPrefix(s, "0") {
			v, _ = regexp.MatchString(`[1-9]`, string(c))
		} else {
			v, _ = regexp.MatchString(`[0-2]`, string(c))
		}
		return v
	case 9:
		v, _ := regexp.MatchString(`[0123]`, string(c))
		return v
	case 10:
		v, _ := regexp.MatchString(`[0-9]`, string(c))
		return v
	}

	return false
}

func dateInputFieldChanged(s string) {
	return
}

func readInsertForm() domain.Transaction {
	return domain.Transaction{
		-1,
		dataf.Date(insertInputFieldDate.GetText()),
		dataf.Description(insertInputFieldDescription.GetText()),
		dataf.Cents(insertInputFieldCents.GetText()),
	}
}

func handleSaveTransaction() {
	t := readInsertForm()
	id, err := transaction.Insert(t)
	if err == nil {
		status.SetText(fmt.Sprintf("saving transaction %d", id))
	} else {
		status.SetText(fmt.Sprint(err))
	}
}

func createInsertForm() *cview.Form {
	insertForm = cview.NewForm()
	insertForm.SetCancelFunc(handleCloseInsert)

	insertInputFieldDate = cview.NewInputField()
	insertInputFieldDate.SetLabel("Date")
	insertInputFieldDate.SetText("2022-")
	insertInputFieldDate.SetFieldWidth(11)
	insertInputFieldDate.SetAcceptanceFunc(dateInputFieldAcceptance)
	insertInputFieldDate.SetChangedFunc(dateInputFieldChanged)

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

	insertForm.AddButton("Save", handleSaveTransaction)
	insertForm.AddButton("Quit", handleCloseInsert)
	insertForm.SetBorder(true)
	insertForm.SetBorderAttributes(tcell.AttrBold)
	insertForm.SetRect(4, 4, 45, 18)
	insertForm.SetTitleAlign(cview.AlignCenter)
	insertForm.SetTitle(" New Transaction ")
	insertForm.SetWrapAround(true)
	insertForm.SetFieldBackgroundColor(tcell.Color242)
	insertForm.SetFieldBackgroundColorFocused(tcell.ColorRed)

	return insertForm
}
