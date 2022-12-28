package ui

import (
	"fmt"
	"jcb/config"
	"jcb/db"
	"jcb/domain"
	"jcb/lib/repeater"
	"jcb/lib/validator"
	"log"

	dataf "jcb/lib/formatter/data"
	"jcb/lib/transaction"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var repeatForm *cview.Form
var repeatInputFieldRule *cview.InputField
var repeatInputFieldUntil *cview.InputField

func handleOpenRepeat(ev *tcell.EventKey) *tcell.EventKey {
	openRepeat()
	return nil
}

func openRepeat() {
	panels.ShowPanel("repeat")

	repeatInputFieldRule.SetText("1m")
	repeatInputFieldUntil.SetText(fmt.Sprintf("%d-12-31", db.DateLastUncommitted().Year()))
}

func closeRepeat() {
	panels.HidePanel("repeat")
	repeatForm.SetFocus(0)
}

func checkRepeatForm() bool {
	var err error

	err = validator.RepeatRule(repeatInputFieldRule.GetText())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return false
	}

	err = validator.Date(repeatInputFieldUntil.GetText())
	if err != nil {
		printStatus(fmt.Sprint(err))
		return false
	}

	return true
}

func readRepeatForm() []domain.Transaction {
	var transactions []domain.Transaction

	curRow, _ := transactionsTable.GetSelection()
	date := dataf.Date(transactionsTable.GetCell(curRow, config.DATE_COLUMN).GetText())
	description := dataf.Description(transactionsTable.GetCell(curRow, config.DESCRIPTION_COLUMN).GetText())
	cents := dataf.Cents(transactionsTable.GetCell(curRow, config.AMOUNT_COLUMN).GetText())
	repeatRule := dataf.RepeatRule(repeatInputFieldRule.GetText())
	repeatUntil := dataf.Date(repeatInputFieldUntil.GetText())
	notes := transaction.Notes(transactionIds[curRow])
	category := dataf.Category(transactionsTable.GetCell(curRow, config.CATEGORY_COLUMN).GetText())

	timestamps, err := repeater.Expand(date, repeatUntil, repeatRule)
	if err != nil {
		log.Fatal(err)
	}

	for _, ts := range timestamps {
		transactions = append(transactions, domain.Transaction{
			-1,
			ts,
			description,
			cents,
			notes,
			category,
		})
	}

	return transactions
}

func handleRepeatTransaction(ev *tcell.EventKey) *tcell.EventKey {
	if !checkRepeatForm() {
		return nil
	}

	var id int64
	var err error
	for _, t := range readRepeatForm() {
		id, err = transaction.Insert(t)
		if err != nil {
			printStatus(fmt.Sprint(err))
			return nil
		}
	}

	updateTransactionsTable()
	selectTransaction(id)

	closeRepeat()
	return nil
}

func createRepeatForm() *cview.Form {
	repeatForm = cview.NewForm()
	repeatForm.SetCancelFunc(closeRepeat)

	repeatInputFieldRule = cview.NewInputField()
	repeatInputFieldRule.SetLabel("Frequency:")
	repeatInputFieldRule.SetFieldWidth(14)
	repeatInputFieldRule.SetText("1m")
	repeatInputFieldRule.SetFieldNote(`<number>(dwm)`)

	repeatInputFieldUntil = cview.NewInputField()
	repeatInputFieldUntil.SetLabel("End:")
	repeatInputFieldUntil.SetFieldWidth(11)
	repeatInputFieldUntil.SetText(fmt.Sprintf("%d-12-31"))
	repeatInputFieldUntil.SetAcceptanceFunc(dateInputFieldAcceptance)

	repeatForm.AddFormItem(repeatInputFieldRule)
	repeatForm.AddFormItem(repeatInputFieldUntil)

	repeatForm.SetBorder(true)
	repeatForm.SetBorderAttributes(tcell.AttrBold)
	repeatForm.SetRect(6, 4, 45, 8)
	repeatForm.SetTitleAlign(cview.AlignCenter)
	repeatForm.SetTitle(" Repeat Transaction ")
	repeatForm.SetWrapAround(true)
	repeatForm.SetFieldBackgroundColor(tcell.Color242)
	repeatForm.SetFieldBackgroundColorFocused(tcell.ColorRed)

	c := cbind.NewConfiguration()
	c.SetKey(0, tcell.KeyEnter, handleRepeatTransaction)
	repeatInputFieldRule.SetInputCapture(c.Capture)
	repeatInputFieldUntil.SetInputCapture(c.Capture)

	return repeatForm
}
