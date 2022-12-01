package openingBalanceWin

import (
	openingBalance "jcb/lib/openingbalance"
	fieldReader "jcb/lib/ui/fieldreader"
	dataFormatter "jcb/lib/ui/formatter/data"

	gc "github.com/rthornton128/goncurses"
)

var win *gc.Window

var form gc.Form
var fields []*gc.Field

func Show() {
	win, _ = gc.NewWindow(9, 45, 6, 10)
	defer win.Delete()

	fields = make([]*gc.Field, 4)
	fields[0], _ = gc.NewField(1, 8, 6, 17, 0, 0)
	fields[0].SetBuffer("0.00")
	fields[0].SetOptionsOn(gc.FO_BLANK)
	defer fields[0].Free()

	form, _ = gc.NewForm(fields)
	defer form.UnPost()
	defer form.Free()
	form.SetSub(win)
	form.Post()

	win.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	win.MovePrint(1, 2, "Opening Balance 2022")
	win.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	win.MovePrint(3, 4, "Please enter the opening balance to")
	win.MovePrint(4, 4, "begin budgeting for the year of 2022.")

	win.MovePrint(6, 2, "Balance:")

	win.Box(0, 0)

	err := scan()
	for err != nil {
		win.MovePrint(0, 0, err)
		err = scan()
	}
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
			err := form.Driver(gc.REQ_VALIDATION)
			s, err := fieldReader.AsAmount(fields[0])
			if err != nil {
				return err
			}

			i, err := dataFormatter.Cents(s)
			if err != nil {
				return err
			}

			win.MovePrint(0, 0, s)

			return openingBalance.Save(i, 2022)
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
