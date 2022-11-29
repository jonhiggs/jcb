package ui

import (
	"jcb/domain"
	"jcb/lib/transaction"
	"time"

	gc "github.com/rthornton128/goncurses"
)

var transactionAddWin *gc.Window
var transactionAddForm gc.Form

func renderTransactionAdd() {
	t := formatTransaction(domain.Transaction{0, time.Now(), "", 0})

	gc.Cursor(1)
	transactionAddWin, _ = gc.NewWindow(9, 60, 8, 10)

	defer gc.Cursor(0)
	defer transactionAddWin.Delete()

	// date field
	fields := make([]*gc.Field, 4)
	fields[0], _ = gc.NewField(1, 10, 3, 17, 0, 0)
	fields[0].SetBuffer(t.Date)
	defer fields[0].Free()

	// description field
	fields[1], _ = gc.NewField(1, 30, 4, 17, 0, 0)
	fields[1].SetBuffer(t.Description)
	defer fields[1].Free()

	// amount field
	fields[2], _ = gc.NewField(1, 8, 5, 17, 0, 0)
	fields[2].SetBuffer(t.Amount)
	defer fields[1].Free()

	// repetition field
	fields[3], _ = gc.NewField(1, 10, 6, 17, 0, 0)
	fields[3].SetBuffer("0d")
	defer fields[3].Free()

	transactionAddForm, _ = gc.NewForm(fields)
	defer transactionAddForm.UnPost()
	defer transactionAddForm.Free()
	transactionAddForm.SetSub(transactionAddWin)
	transactionAddForm.Post()

	transactionAddWin.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	transactionAddWin.MovePrint(1, 2, "New Transaction")
	transactionAddWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	transactionAddWin.MovePrint(3, 2, "Date:")
	transactionAddWin.MovePrint(4, 2, "Description:")
	transactionAddWin.MovePrint(5, 2, "Amount:")
	transactionAddWin.MovePrint(6, 2, "Repeat every")

	transactionAddWin.Box(0, 0)

	var err error
	switch scanTransactionAdd() {
	case INSERT:
		err = transactionInsert(fields)
	}
	if err != nil {
		printError(err)
	}

	mainWin.Touch()
	mainWin.Refresh()
	footerWin.Touch()
	footerWin.Refresh()
}

func transactionInsert(fields []*gc.Field) error {
	err := transactionAddForm.Driver(gc.REQ_VALIDATION)
	if err != nil {
		return err
	}

	t := unformatTransaction(FormattedTransaction{
		"0",
		fields[0].Buffer(),
		fields[1].Buffer(),
		fields[2].Buffer(),
	})

	return transaction.Save(t)
}
