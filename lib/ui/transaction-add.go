package ui

import (
	model "jcb/lib/model"
	"time"

	gc "github.com/rthornton128/goncurses"
)

var transactionAddWin *gc.Window
var transactionAddForm gc.Form

func renderTransactionAdd() {
	t := formatTransaction(model.Transaction{0, time.Now(), "", 0})

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

	scanTransactionAdd()

	mainWin.Touch()
	mainWin.Refresh()
	footerWin.Touch()
	footerWin.Refresh()

	//for {
	//	ch := win.GetChar()
	//	win.MovePrint(0, 0, ch)
	//	switch ch {
	//	case gc.KEY_RETURN:
	//		errStack := make([]error, 2)
	//		hasErr := false

	//		errStack[0] = form.Driver(gc.REQ_VALIDATION)
	//		errStack[1] = uiValidator.Date(fields[0].Buffer())

	//		for _, e := range errStack {
	//			if e != nil {
	//				PrintError(e)
	//				hasErr = true
	//			}
	//		}

	//		if !hasErr {
	//			_, err := db.SaveTransaction(db.Transaction{
	//				0,
	//				uiStringify.Date(fields[0].Buffer()),
	//				fields[1].Buffer(),
	//				uiStringify.Cents(fields[2].Buffer()),
	//			})

	//			return err
	//		}
	//	case 1: // ctrl-a
	//		form.Driver(gc.REQ_BEG_FIELD)
	//	case 5: // ctrl-e
	//		form.Driver(gc.REQ_END_FIELD)
	//	case 11: // ctrl-k
	//		form.Driver(gc.REQ_DEL_LINE)
	//	case 3: // ctrl-c
	//		return nil
	//	case 4, 33: // ctrl-d, delete
	//		form.Driver(gc.REQ_DEL_CHAR)
	//	case 23, 27: // ctrl-w, esc/alt-backspace
	//		form.Driver(gc.REQ_DEL_WORD)
	//	case gc.KEY_BACKSPACE:
	//		form.Driver(gc.REQ_DEL_PREV)
	//	case gc.KEY_DOWN, gc.KEY_TAB:
	//		form.Driver(gc.REQ_NEXT_FIELD)
	//		form.Driver(gc.REQ_END_LINE)
	//	case 2, gc.KEY_LEFT:
	//		form.Driver(gc.REQ_LEFT_CHAR)
	//	case 6, gc.KEY_RIGHT:
	//		form.Driver(gc.REQ_RIGHT_CHAR)
	//	default:
	//		form.Driver(ch)
	//	}
	//}
}
