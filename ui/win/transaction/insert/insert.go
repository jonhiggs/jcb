package transactionInsertWin

import (
	"jcb/domain"
	"jcb/lib/transaction"
	dataf "jcb/ui/formatter/data"
	statusWin "jcb/ui/win/status"
	"time"

	gc "github.com/rthornton128/goncurses"
)

var win *gc.Window
var form gc.Form
var fields []*gc.Field

func Show() int64 {
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

	id, err := scan()
	for err != nil {
		statusWin.PrintError(err)
		id, err = scan()
	}

	statusWin.Clear()
	return id
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

func scan() (int64, error) {
	win.Keypad(true)
	win.Refresh()

	form.Driver(gc.REQ_FIRST_FIELD)
	form.Driver(gc.REQ_END_LINE)

	for {
		ch := win.GetChar()
		switch ch {
		case gc.KEY_RETURN:
			var id int64
			var err error
			t, err := readForm()
			if err == nil {
				id, err = transaction.Save(t)
				if err != nil {
					statusWin.PrintError(err)
					//} else {
					//	updateTransactions()
					//	selectTransaction(id)
				}
			}
			return id, err
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
			return -1, nil
		default:
			form.Driver(ch)
		}
	}
}
