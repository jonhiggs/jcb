package transactionEditWin

import (
	"fmt"
	"jcb/domain"
	"jcb/lib/transaction"
	dataf "jcb/ui/formatter/data"
	stringf "jcb/ui/formatter/string"
	statusWin "jcb/ui/win/status"

	gc "github.com/rthornton128/goncurses"
)

var win *gc.Window
var form gc.Form
var fields []*gc.Field

func Show(id int64) int64 {
	t, _ := transaction.Find(id)
	tstr, _ := stringf.Transaction(t)
	gc.Cursor(1)
	defer gc.Cursor(0)
	win, _ = gc.NewWindow(8, 49, 4, 10)
	fields = make([]*gc.Field, 4)

	// month field
	fields[0], _ = gc.NewField(1, 2, 3, 22, 0, 0)
	//fields[0].SetOptionsOn(gc.FO_BLANK)
	fields[0].SetBuffer(t.Date.Format("01"))
	defer fields[0].Free()

	// day field
	fields[1], _ = gc.NewField(1, 2, 3, 25, 0, 0)
	//fields[1].SetOptionsOn(gc.FO_BLANK)
	fields[1].SetBuffer(t.Date.Format("02"))
	defer fields[1].Free()

	// description field
	fields[2], _ = gc.NewField(1, 30, 4, 17, 0, 0)
	fields[2].SetBuffer(tstr.Description)
	defer fields[2].Free()

	// amount field
	fields[3], _ = gc.NewField(1, 8, 5, 17, 0, 0)
	fields[3].SetBuffer(tstr.Cents)
	defer fields[3].Free()

	form, _ = gc.NewForm(fields)
	defer form.UnPost()
	defer form.Free()
	form.SetSub(win)
	form.Post()

	win.AttrOn(gc.ColorPair(2) | gc.A_BOLD | gc.A_UNDERLINE)
	win.MovePrint(1, 2, "Edit Transaction")
	win.AttrOff(gc.ColorPair(2) | gc.A_BOLD | gc.A_UNDERLINE)

	win.MovePrint(3, 2, "Date:")
	win.MovePrint(3, 17, "2022-")
	win.MovePrint(3, 24, "-")
	win.MovePrint(4, 2, "Description:")
	win.MovePrint(5, 2, "Amount:")

	win.Box(0, 0)

	err := scan(id)
	for err != nil {
		statusWin.PrintError(err)
		err = scan(id)
	}

	return id
}

func readForm(id int64) (domain.Transaction, error) {
	err := form.Driver(gc.REQ_VALIDATION)
	if err != nil {
		return domain.Transaction{}, err
	}

	// date
	dateMonth := fields[0].Buffer()
	dateDay := fields[1].Buffer()
	date, err := dataf.Date(fmt.Sprintf("2022-%s-%s", dateMonth, dateDay))
	if err != nil {
		return domain.Transaction{}, err
	}

	description, err := dataf.Description(fields[2].Buffer())
	if err != nil {
		return domain.Transaction{}, err
	}

	cents, err := dataf.Cents(fields[3].Buffer())
	if err != nil {
		return domain.Transaction{}, err
	}

	return domain.Transaction{id, date, description, cents}, nil
}

func scan(id int64) error {
	win.Keypad(true)
	win.Refresh()

	form.Driver(gc.REQ_FIRST_FIELD)
	form.Driver(gc.REQ_END_LINE)

	for {
		ch := win.GetChar()
		switch ch {
		case gc.KEY_RETURN:
			t, err := readForm(id)
			if err != nil {
				return err
			}
			return transaction.Edit(t)
		case 1: // ctrl-a
			form.Driver(gc.REQ_BEG_FIELD)
		case 5: // ctrl-e
			form.Driver(gc.REQ_END_FIELD)
		case 11: // ctrl-k
			form.Driver(gc.REQ_DEL_LINE)
		case 4, 33: // ctrl-d, delete
			form.Driver(gc.REQ_DEL_CHAR)
		case 23, 27: // ctrl-w, esc/alt-backspace
			form.Driver(gc.REQ_DEL_WORD) // FIXME
		case gc.KEY_BACKSPACE:
			form.Driver(gc.REQ_DEL_PREV)
		case gc.KEY_DOWN, gc.KEY_TAB:
			form.Driver(gc.REQ_NEXT_FIELD)
			form.Driver(gc.REQ_END_LINE)
		case 2, gc.KEY_LEFT:
			form.Driver(gc.REQ_LEFT_CHAR)
		case 6, gc.KEY_RIGHT:
			form.Driver(gc.REQ_RIGHT_CHAR)
		case 3:
			return nil
		default:
			form.Driver(ch)
		}
	}
}
