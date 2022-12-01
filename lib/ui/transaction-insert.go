package ui

import (
	"errors"
	"fmt"
	"jcb/domain"
	"jcb/lib/transaction"
	dformat "jcb/lib/ui/formatter/data"
	sformat "jcb/lib/ui/formatter/string"
	statusWin "jcb/lib/ui/win/status"
	"strings"
	"time"

	gc "github.com/rthornton128/goncurses"
)

var transactionInsertWin *gc.Window
var transactionInsertForm gc.Form
var transactionInsertFormFields []*gc.Field

func renderTransactionInsert() {
	t, _ := sformat.Transaction(domain.Transaction{0, time.Now(), "", 0})

	gc.Cursor(1)
	transactionInsertWin, _ = gc.NewWindow(9, 60, 8, 10)

	defer gc.Cursor(0)
	defer transactionInsertWin.Delete()

	// date field
	transactionInsertFormFields = make([]*gc.Field, 4)
	transactionInsertFormFields[0], _ = gc.NewField(1, 10, 3, 17, 0, 0)
	transactionInsertFormFields[0].SetOptionsOn(gc.FO_BLANK)
	transactionInsertFormFields[0].SetBuffer(t.Date[0:8])
	defer transactionInsertFormFields[0].Free()

	// description field
	transactionInsertFormFields[1], _ = gc.NewField(1, 30, 4, 17, 0, 0)
	transactionInsertFormFields[1].SetBuffer(t.Description)
	defer transactionInsertFormFields[1].Free()

	// amount field
	transactionInsertFormFields[2], _ = gc.NewField(1, 8, 5, 17, 0, 0)
	defer transactionInsertFormFields[1].Free()

	// repetition field
	transactionInsertFormFields[3], _ = gc.NewField(1, 10, 6, 17, 0, 0)
	transactionInsertFormFields[3].SetBuffer("0d")
	defer transactionInsertFormFields[3].Free()

	transactionInsertForm, _ = gc.NewForm(transactionInsertFormFields)
	defer transactionInsertForm.UnPost()
	defer transactionInsertForm.Free()
	transactionInsertForm.SetSub(transactionInsertWin)
	transactionInsertForm.Post()

	transactionInsertWin.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	transactionInsertWin.MovePrint(1, 2, "New Transaction")
	transactionInsertWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	transactionInsertWin.MovePrint(3, 2, "Date:")
	transactionInsertWin.MovePrint(4, 2, "Description:")
	transactionInsertWin.MovePrint(5, 2, "Amount:")
	transactionInsertWin.MovePrint(6, 2, "Repeat every")

	transactionInsertWin.Box(0, 0)

	action := CONTINUE
	for action == CONTINUE {
		action = scanTransactionInsert()
		switch action {
		case ABORT:
			statusWin.PrintError(errors.New("aborting"))
			break
		case INSERT:
			t, err := transactionInsertRead()
			if err == nil {
				id, err := transaction.Save(t)
				if err != nil {
					statusWin.PrintError(err)
				} else {
					updateTransactions()
					selectTransaction(id)
				}
				break
			} else {
				statusWin.PrintError(err)
				action = CONTINUE
			}
		}
	}

	statusWin.Clear()

	mainWin.Touch()
	mainWin.Refresh()
	statusWin.Refresh()
}

// construct a domain.Transaction from the data in the form
func transactionInsertRead() (domain.Transaction, error) {
	err := transactionInsertForm.Driver(gc.REQ_VALIDATION)
	if err != nil {
		return domain.Transaction{}, err
	}

	idStr := "0"
	dateStr := transactionInsertFormFields[0].Buffer()
	descriptionStr := transactionInsertFormFields[1].Buffer()
	amountStr := strings.Trim(transactionInsertFormFields[2].Buffer(), " ")

	amountSplit := strings.Split(amountStr, ".")
	if len(amountSplit) > 2 {
		return domain.Transaction{}, errors.New("Amount has too many dots")
	}
	if len(amountSplit) == 2 && len(amountSplit[1]) > 2 {
		return domain.Transaction{}, errors.New(fmt.Sprintf("Amount has to many decimal places [%d]", len(amountSplit[1])))
	}

	t, _ := dformat.Transaction(domain.StringTransaction{idStr, dateStr, descriptionStr, amountStr})

	err = transaction.Validate(t)
	return t, err
}
