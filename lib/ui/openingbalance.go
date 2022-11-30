package ui

import (
	gc "github.com/rthornton128/goncurses"
)

var openingBalanceWin *gc.Window

var openingBalanceForm gc.Form
var openingBalanceFormFields []*gc.Field

func renderOpeningBalance() {
	openingBalanceWin, _ = gc.NewWindow(9, 45, 6, 10)
	defer openingBalanceWin.Delete()

	openingBalanceFormFields = make([]*gc.Field, 4)
	openingBalanceFormFields[0], _ = gc.NewField(1, 8, 6, 17, 0, 0)
	openingBalanceFormFields[0].SetBuffer("hi")
	defer openingBalanceFormFields[0].Free()

	openingBalanceForm, _ = gc.NewForm(openingBalanceFormFields)
	defer openingBalanceForm.UnPost()
	defer openingBalanceForm.Free()
	openingBalanceForm.SetSub(openingBalanceWin)
	openingBalanceForm.Post()

	openingBalanceWin.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	openingBalanceWin.MovePrint(1, 2, "Opening Balance 2022")
	openingBalanceWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	openingBalanceWin.MovePrint(3, 4, "Please enter the opening balance to")
	openingBalanceWin.MovePrint(4, 4, "begin budgeting for the year of 2022.")

	openingBalanceWin.MovePrint(6, 2, "Balance:")

	openingBalanceWin.Box(0, 0)

	scanOpeningBalance()

	mainWin.Touch()
	mainWin.Refresh()
	footerWin.Touch()
	footerWin.Refresh()
}

func initOpeningBalance() error {

	renderOpeningBalance()

	return nil
}

//func renderOpeningBalanceSet() {
//	gc.Cursor(1)
//	openingBalanceSetWin, _ = gc.NewWindow(9, 60, 8, 10)
//
//	defer gc.Cursor(0)
//	defer openingBalanceSetWin.Delete()
//
//	openingBalanceSetFormFields = make([]*gc.Field, 4)
//	openingBalanceSetFormFields[0], _ = gc.NewField(1, 8, 5, 17, 0, 0)
//	defer openingBalanceSetFormFields[0].Free()
//
//	openingBalanceSetForm, _ = gc.NewForm(openingBalanceSetFormFields)
//	defer openingBalanceSetForm.UnPost()
//	defer openingBalanceSetForm.Free()
//	openingBalanceSetForm.SetSub(openingBalanceSetWin)
//	openingBalanceSetForm.Post()
//
//	openingBalanceSetWin.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
//	openingBalanceSetWin.MovePrint(1, 2, "Opening Balance")
//	openingBalanceSetWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
//
//	openingBalanceSetWin.MovePrint(3, 2, "Amount:")
//
//	openingBalanceSetWin.Box(0, 0)
//
//	action := CONTINUE
//	for action == CONTINUE {
//		action = scanTransactionInsert()
//		switch action {
//		case ABORT:
//			printError(errors.New("aborting"))
//			break
//		case INSERT:
//			//t, err := openingBalanceSetRead()
//			//if err == nil {
//			//	id, err := transaction.Save(t)
//			//	if err != nil {
//			//		printError(err)
//			//	} else {
//			//		updateTransactions()
//			//		selectTransaction(id)
//			//	}
//			//	break
//			//} else {
//			//	printError(err)
//			//	action = CONTINUE
//			//}
//		}
//	}
//
//	clearError()
//
//	mainWin.Touch()
//	mainWin.Refresh()
//	footerWin.Touch()
//	footerWin.Refresh()
//}
//
//// construct a domain.Transaction from the data in the form
//func openingBalanceSetRead() (domain.Transaction, error) {
//	err := openingBalanceSetForm.Driver(gc.REQ_VALIDATION)
//	if err != nil {
//		return domain.Transaction{}, err
//	}
//
//	idStr := "0"
//	dateStr := openingBalanceSetFormFields[0].Buffer()
//	descriptionStr := openingBalanceSetFormFields[1].Buffer()
//	amountStr := strings.Trim(openingBalanceSetFormFields[2].Buffer(), " ")
//
//	amountSplit := strings.Split(amountStr, ".")
//	if len(amountSplit) > 2 {
//		return domain.Transaction{}, errors.New("Amount has too many dots")
//	}
//	if len(amountSplit) == 2 && len(amountSplit[1]) > 2 {
//		return domain.Transaction{}, errors.New(fmt.Sprintf("Amount has to many decimal places [%d]", len(amountSplit[1])))
//	}
//
//	t := unformatTransaction(FormattedTransaction{idStr, dateStr, descriptionStr, amountStr})
//
//	err = transaction.Validate(t)
//	return t, err
//}
