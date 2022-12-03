package transactionInsertWin

import (
	"fmt"
	"jcb/domain"
	"jcb/lib/transaction"
	dataf "jcb/ui/formatter/data"
	"jcb/ui/repeater"
	statusWin "jcb/ui/win/status"

	gc "github.com/rthornton128/goncurses"
)

var win *gc.Window
var form gc.Form
var fields []*gc.Field

func Show() int64 {
	gc.Cursor(1)
	defer gc.Cursor(0)
	win, _ = gc.NewWindow(10, 60, 8, 10)
	fields = make([]*gc.Field, 7)

	// month field
	fields[0], _ = gc.NewField(1, 2, 3, 22, 0, 0)
	fields[0].SetOptionsOn(gc.FO_BLANK)
	defer fields[0].Free()

	// day field
	fields[1], _ = gc.NewField(1, 2, 3, 25, 0, 0)
	fields[1].SetOptionsOn(gc.FO_BLANK)
	defer fields[1].Free()

	// description field
	fields[2], _ = gc.NewField(1, 30, 4, 17, 0, 0)
	//fields[1].SetBuffer(t.Description)
	defer fields[2].Free()

	// amount field
	fields[3], _ = gc.NewField(1, 8, 5, 17, 0, 0)
	defer fields[2].Free()

	// repeat pattern
	fields[4], _ = gc.NewField(1, 10, 6, 17, 0, 0)
	fields[4].SetBuffer("0d")
	defer fields[4].Free()

	// repeat until month
	fields[5], _ = gc.NewField(1, 2, 7, 22, 0, 0)
	fields[5].SetBuffer("12")
	fields[5].SetOptionsOn(gc.FO_BLANK)
	defer fields[5].Free()

	// repeat until day
	fields[6], _ = gc.NewField(1, 2, 7, 25, 0, 0)
	fields[6].SetBuffer("31")
	fields[6].SetOptionsOn(gc.FO_BLANK)
	defer fields[6].Free()

	form, _ = gc.NewForm(fields)
	defer form.UnPost()
	defer form.Free()
	form.SetSub(win)
	form.Post()

	win.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	win.MovePrint(1, 2, "New Transaction")
	win.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	win.MovePrint(3, 2, "Date:")
	win.MovePrint(3, 17, "2022-")
	win.MovePrint(3, 24, "-")
	win.MovePrint(4, 2, "Description:")
	win.MovePrint(5, 2, "Amount:")
	win.MovePrint(6, 2, "Repeat every:")
	win.MovePrint(7, 2, "Repeat until:")
	win.MovePrint(7, 17, "2022-")
	win.MovePrint(7, 24, "-")

	win.Box(0, 0)

	id, err := scan()
	for err != nil {
		statusWin.PrintError(err)
		id, err = scan()
	}

	statusWin.Clear()
	return id
}

func readForm() ([]domain.Transaction, error) {
	var trans []domain.Transaction
	err := form.Driver(gc.REQ_VALIDATION)
	if err != nil {
		return trans, err
	}

	// date
	dateMonth := fields[0].Buffer()
	dateDay := fields[1].Buffer()
	date, err := dataf.Date(fmt.Sprintf("2022-%s-%s", dateMonth, dateDay))
	if err != nil {
		return trans, err
	}

	description, err := dataf.Description(fields[2].Buffer())
	if err != nil {
		return trans, err
	}

	cents, err := dataf.Cents(fields[3].Buffer())
	if err != nil {
		return trans, err
	}

	rule, err := dataf.RepeatRule(fields[4].Buffer())
	if err != nil {
		return trans, err
	}

	repeatUntilMonth := fields[5].Buffer()
	repeatUntilDay := fields[6].Buffer()
	repeatUntil, err := dataf.Date(fmt.Sprintf("2022-%s-%s", repeatUntilMonth, repeatUntilDay))
	if err != nil {
		return trans, err
	}

	repeatFrom, err := transaction.CommittedUntil()

	timestamps, err := repeater.Expand(date, repeatFrom, repeatUntil, rule)
	if err != nil {
		return trans, err
	}

	for _, ts := range timestamps {
		trans = append(trans, domain.Transaction{-1, ts, description, cents})
	}

	return trans, err
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
			transactions, err := readForm()
			if err == nil {
				for _, t := range transactions {
					id, err = transaction.Insert(t)
					if err != nil {
						statusWin.PrintError(err)
						//} else {
						//	updateTransactions()
						//	selectTransaction(id)
					}
				}
			} else {
				return -1, err
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
		case 3:
			return -1, nil
		default:
			form.Driver(ch)
		}
	}
}
