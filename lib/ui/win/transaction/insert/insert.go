package transactionInsertWin

import (
	"jcb/domain"
	"jcb/lib/transaction"
	dataf "jcb/lib/ui/formatter/data"
	statusWin "jcb/lib/ui/win/status"
	"time"

	gc "github.com/rthornton128/goncurses"
)

var win *gc.Window
var form gc.Form
var fields []*gc.Field

func Show() {
	gc.Cursor(1)
	defer gc.Cursor(0)
	win, _ = gc.NewWindow(9, 60, 8, 10)
	fields = make([]*gc.Field, 4)

	// date field
	fields[0], _ = gc.NewField(1, 10, 3, 17, 0, 0)
	fields[0].SetOptionsOn(gc.FO_BLANK)
	fields[0].SetBuffer(time.Now().Format("2006-01-"))
	defer fields[0].Free()

	// description field
	fields[1], _ = gc.NewField(1, 30, 4, 17, 0, 0)
	//fields[1].SetBuffer(t.Description)
	defer fields[1].Free()

	// amount field
	fields[2], _ = gc.NewField(1, 8, 5, 17, 0, 0)
	defer fields[1].Free()

	// repetition field
	fields[3], _ = gc.NewField(1, 10, 6, 17, 0, 0)
	fields[3].SetBuffer("0d")
	defer fields[3].Free()

	form, _ = gc.NewForm(fields)
	defer form.UnPost()
	defer form.Free()
	form.SetSub(win)
	form.Post()

	win.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	win.MovePrint(1, 2, "New Transaction")
	win.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	win.MovePrint(3, 2, "Date:")
	win.MovePrint(4, 2, "Description:")
	win.MovePrint(5, 2, "Amount:")
	win.MovePrint(6, 2, "Repeat every")

	win.Box(0, 0)

	err := scan()
	for err != nil {
		statusWin.PrintError(err)
		err = scan()
	}

	statusWin.Clear()
}

func readForm() (domain.Transaction, error) {
	err := form.Driver(gc.REQ_VALIDATION)
	if err != nil {
		return domain.Transaction{}, err
	}

	date, err := dataf.Date(fields[0].Buffer())
	if err != nil {
		return domain.Transaction{}, err
	}

	description, err := dataf.Description(fields[1].Buffer())
	if err != nil {
		return domain.Transaction{}, err
	}

	cents, err := dataf.Cents(fields[2].Buffer())
	if err != nil {
		return domain.Transaction{}, err
	}

	return domain.Transaction{0, date, description, cents}, err
}

func scan() error {
	win.Keypad(true)
	win.Refresh()

	form.Driver(gc.REQ_FIRST_FIELD)
	form.Driver(gc.REQ_END_LINE)

	for {
		ch := win.GetChar()
		switch ch {
		case gc.KEY_RETURN:
			t, err := readForm()
			if err == nil {
				_, err := transaction.Save(t)
				if err != nil {
					statusWin.PrintError(err)
					//} else {
					//	updateTransactions()
					//	selectTransaction(id)
				}
			}
			return err
		case 1: // ctrl-a
			form.Driver(gc.REQ_BEG_FIELD)
		case 5: // ctrl-e
			form.Driver(gc.REQ_END_FIELD)
		case 11: // ctrl-k
			form.Driver(gc.REQ_DEL_LINE)
		case 4, 33: // ctrl-d, delete
			form.Driver(gc.REQ_DEL_CHAR)
		case 23, 27: // ctrl-w, esc/alt-backspace
			form.Driver(gc.REQ_DEL_WORD)
		case gc.KEY_BACKSPACE:
			form.Driver(gc.REQ_DEL_PREV)
		case gc.KEY_DOWN, gc.KEY_TAB:
			form.Driver(gc.REQ_NEXT_FIELD)
			form.Driver(gc.REQ_END_LINE)
		case 2, gc.KEY_LEFT:
			form.Driver(gc.REQ_LEFT_CHAR)
		case 6, gc.KEY_RIGHT:
			form.Driver(gc.REQ_RIGHT_CHAR)
		case 'q', 3:
			return nil
		default:
			form.Driver(ch)
		}
	}
}

//var transactionInsertWin *gc.Window
//var transactionInsertForm gc.Form
//var fields []*gc.Field
//
//func renderTransactionInsert() {
//	t, _ := sformat.Transaction(domain.Transaction{0, time.Now(), "", 0})
//
//	gc.Cursor(1)
//	transactionInsertWin, _ = gc.NewWindow(9, 60, 8, 10)
//
//	defer gc.Cursor(0)
//	defer transactionInsertWin.Delete()
//
//	// date field
//	fields = make([]*gc.Field, 4)
//	fields[0], _ = gc.NewField(1, 10, 3, 17, 0, 0)
//	fields[0].SetOptionsOn(gc.FO_BLANK)
//	fields[0].SetBuffer(t.Date[0:8])
//	defer fields[0].Free()
//
//	// description field
//	fields[1], _ = gc.NewField(1, 30, 4, 17, 0, 0)
//	fields[1].SetBuffer(t.Description)
//	defer fields[1].Free()
//
//	// amount field
//	fields[2], _ = gc.NewField(1, 8, 5, 17, 0, 0)
//	defer fields[1].Free()
//
//	// repetition field
//	fields[3], _ = gc.NewField(1, 10, 6, 17, 0, 0)
//	fields[3].SetBuffer("0d")
//	defer fields[3].Free()
//
//	transactionInsertForm, _ = gc.NewForm(fields)
//	defer transactionInsertForm.UnPost()
//	defer transactionInsertForm.Free()
//	transactionInsertForm.SetSub(transactionInsertWin)
//	transactionInsertForm.Post()
//
//	transactionInsertWin.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
//	transactionInsertWin.MovePrint(1, 2, "New Transaction")
//	transactionInsertWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
//
//	transactionInsertWin.MovePrint(3, 2, "Date:")
//	transactionInsertWin.MovePrint(4, 2, "Description:")
//	transactionInsertWin.MovePrint(5, 2, "Amount:")
//	transactionInsertWin.MovePrint(6, 2, "Repeat every")
//
//	transactionInsertWin.Box(0, 0)
//
//	action := CONTINUE
//	for action == CONTINUE {
//		action = scanTransactionInsert()
//		switch action {
//		case ABORT:
//			statusWin.PrintError(errors.New("aborting"))
//			break
//		case INSERT:
//			t, err := transactionInsertRead()
//			if err == nil {
//				id, err := transaction.Save(t)
//				if err != nil {
//					statusWin.PrintError(err)
//				} else {
//					updateTransactions()
//					selectTransaction(id)
//				}
//				break
//			} else {
//				statusWin.PrintError(err)
//				action = CONTINUE
//			}
//		}
//	}
//
//	statusWin.Clear()
//
//	mainWin.Touch()
//	mainWin.Refresh()
//	statusWin.Refresh()
//}
//
//// construct a domain.Transaction from the data in the form
//func transactionInsertRead() (domain.Transaction, error) {
//	err := transactionInsertForm.Driver(gc.REQ_VALIDATION)
//	if err != nil {
//		return domain.Transaction{}, err
//	}
//
//	idStr := "0"
//	dateStr := fields[0].Buffer()
//	descriptionStr := fields[1].Buffer()
//	amountStr := strings.Trim(fields[2].Buffer(), " ")
//
//	amountSplit := strings.Split(amountStr, ".")
//	if len(amountSplit) > 2 {
//		return domain.Transaction{}, errors.New("Amount has too many dots")
//	}
//	if len(amountSplit) == 2 && len(amountSplit[1]) > 2 {
//		return domain.Transaction{}, errors.New(fmt.Sprintf("Amount has to many decimal places [%d]", len(amountSplit[1])))
//	}
//
//	t, _ := dformat.Transaction(domain.StringTransaction{idStr, dateStr, descriptionStr, amountStr})
//
//	err = transaction.Validate(t)
//	return t, err
//}
