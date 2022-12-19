package ui

import (
	"fmt"
	"jcb/domain"
	"jcb/lib/repeater"
	"log"
	"regexp"
	"strings"

	dataf "jcb/lib/formatter/data"
	"jcb/lib/transaction"

	"code.rocketnine.space/tslocum/cbind"
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

	insertInputFieldDate.SetText(fmt.Sprintf("%d-", year))
	insertInputFieldDescription.SetText("")
	insertInputFieldCents.SetText("")
	insertInputFieldRepeatRule.SetText("0d")
	insertInputFieldRepeatUntil.SetText(fmt.Sprintf("%d-12-31", year))

	return nil
}

func handleCloseInsert() {
	panels.HidePanel("insert")
	insertForm.SetFocus(0)
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
		status.SetText(s[len(s)-2:])
		if strings.HasPrefix(s[len(s)-2:], "0") {
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

func readInsertForm() []domain.Transaction {
	var transactions []domain.Transaction
	repeatFrom, err := transaction.CommittedUntil()

	date := dataf.Date(insertInputFieldDate.GetText())
	description := dataf.Description(insertInputFieldDescription.GetText())
	cents := dataf.Cents(insertInputFieldCents.GetText())
	repeatRule := dataf.RepeatRule(insertInputFieldRepeatRule.GetText())
	repeatUntil := dataf.Date(insertInputFieldRepeatUntil.GetText())
	repeatFrom, err = transaction.CommittedUntil()

	timestamps, err := repeater.Expand(date, repeatFrom, repeatUntil, repeatRule)
	if err != nil {
		log.Fatal(err)
	}

	for _, ts := range timestamps {
		transactions = append(transactions, domain.Transaction{
			-1,
			ts,
			description,
			cents,
		})
	}

	return transactions
}

func handleSaveTransaction() {
	for _, t := range readInsertForm() {
		id, err := transaction.Insert(t)
		if err == nil {
			updateTransactionsTable()
			selectTransaction(id)
		} else {
			log.Fatal(err)
		}
	}

	handleCloseInsert()
}

func createInsertForm() *cview.Form {
	insertForm = cview.NewForm()
	insertForm.SetCancelFunc(handleCloseInsert)

	insertInputFieldDate = cview.NewInputField()
	insertInputFieldDate.SetLabel("Date:")
	insertInputFieldDate.SetText(fmt.Sprintf("%d-", year))
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
	insertInputFieldRepeatUntil.SetText(fmt.Sprintf("%d-12-31"))
	insertInputFieldRepeatUntil.SetAcceptanceFunc(dateInputFieldAcceptance)

	insertForm.AddFormItem(insertInputFieldDate)
	insertForm.AddFormItem(insertInputFieldDescription)
	insertForm.AddFormItem(insertInputFieldCents)
	insertForm.AddFormItem(insertInputFieldRepeatRule)
	insertForm.AddFormItem(insertInputFieldRepeatUntil)

	insertForm.SetBorder(true)
	insertForm.SetBorderAttributes(tcell.AttrBold)
	insertForm.SetRect(6, 4, 45, 14)
	insertForm.SetTitleAlign(cview.AlignCenter)
	insertForm.SetTitle(" Insert Transaction ")
	insertForm.SetWrapAround(true)
	insertForm.SetFieldBackgroundColor(tcell.Color242)
	insertForm.SetFieldBackgroundColorFocused(tcell.ColorRed)

	c := cbind.NewConfiguration()
	c.SetKey(0, tcell.KeyEnter, handleEditTransaction)
	insertInputFieldDate.SetInputCapture(c.Capture)
	insertInputFieldDescription.SetInputCapture(c.Capture)
	insertInputFieldCents.SetInputCapture(c.Capture)
	insertInputFieldRepeatRule.SetInputCapture(c.Capture)
	insertInputFieldRepeatUntil.SetInputCapture(c.Capture)

	return insertForm
}
