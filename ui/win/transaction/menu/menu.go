package menuWin

import (
	"errors"
	"fmt"
	"jcb/lib/transaction"
	dataf "jcb/ui/formatter/data"
	stringf "jcb/ui/formatter/string"
	helpWin "jcb/ui/win/help"
	statusWin "jcb/ui/win/status"
	transactionInsertWin "jcb/ui/win/transaction/insert"
	"strconv"

	gc "github.com/rthornton128/goncurses"
)

var headingWin *gc.Window
var win *gc.Window
var menuItems []*gc.MenuItem
var menu *gc.Menu

func Show(y int, x int) error {
	heading(y, x)

	var err error
	menu, err = gc.NewMenu(make([]*gc.MenuItem, 0))
	if err != nil {
		statusWin.PrintError(err)
	}

	win, err = gc.NewWindow(y-2, x-1, 1, 1)
	menu.SubWindow(win)
	menu.Mark("")
	menu.Option(gc.O_SHOWDESC, false)
	menu.SetWindow(win)

	err = updateTransactions()
	if err != nil {
		statusWin.PrintError(err)
	}

	menu.UnPost()
	menu.Format(y-2, 1)
	menu.Post()

	err = scan(y, x)
	for err != nil {
		statusWin.PrintError(err)
		err = scan(y, x)
	}

	win.Refresh()

	return nil
}

func scan(y int, x int) error {
	win.Keypad(true)
	win.Refresh()

	for {
		ch := win.GetChar()
		switch ch {
		case 'x':
			err := transaction.DeleteId(selectedTransaction())
			if err != nil {
				statusWin.PrintError(err)
			} else {
				if menu.Current(nil).Index() == menu.Count()-1 {
					menu.Driver(gc.DriverActions[gc.KEY_UP])
				} else {
					menu.Driver(gc.DriverActions[gc.KEY_DOWN])
				}
				updateTransactions()
			}
		//case 'e':
		//	err := ui.EditTransaction(uiTransaction.SelectedTransactionId())
		//	if err != nil {
		//		ui.PrintError(err)
		//	}
		//	uiTransaction.UpdateTransactions()
		//	uiTransaction.TransactionWindow.Touch()
		//	uiTransaction.TransactionWindow.Refresh()
		//	ui.MainWindow.Touch()
		//	ui.MainWindow.Refresh()
		case 'g':
			menu.Driver(gc.DriverActions[gc.KEY_HOME])
		case 'G':
			menu.Driver(gc.DriverActions[gc.KEY_END])
		case 'd':
			for i := 0; i < (y / 2); i++ {
				menu.Driver(gc.DriverActions[gc.KEY_DOWN])
			}
		case 'u':
			for i := 0; i < (y / 2); i++ {
				menu.Driver(gc.DriverActions[gc.KEY_UP])
			}
		case 'j':
			menu.Driver(gc.DriverActions[gc.KEY_DOWN])
		case 'k':
			menu.Driver(gc.DriverActions[gc.KEY_UP])
		case 'J':
			curMonth := menu.Current(nil).Name()[0:7]
			for curMonth == menu.Current(nil).Name()[0:7] {
				menu.Driver(gc.DriverActions[gc.KEY_DOWN])
				if menu.Current(nil).Index() == menu.Count()-1 {
					break
				}
			}
		case 'K':
			curMonth := menu.Current(nil).Name()[0:7]
			for curMonth == menu.Current(nil).Name()[0:7] {
				menu.Driver(gc.DriverActions[gc.KEY_UP])
				if menu.Current(nil).Index() == menu.Count()-1 {
					break
				}
			}
		case 'i':
			id := transactionInsertWin.Show()
			updateTransactions()
			selectTransaction(id)
			win.Touch()
			win.Refresh()
			statusWin.Refresh()
		case '?':
			helpWin.Show()
			win.Touch()
			win.Refresh()
			statusWin.Refresh()
		case 3, 'q':
			return nil
		default:
			continue //menu.Driver(gc.DriverActions[ch])
		}
	}
}

func heading(y int, x int) {
	var err error
	headingWin, err = gc.NewWindow(y-1, x-2, 0, 1)
	if err != nil {
		statusWin.PrintError(err)
	}

	headingWin.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	headingWin.MovePrint(0, 0, "DATE")
	headingWin.MovePrint(0, 12, "DESCRIPTION")
	headingWin.MovePrint(0, 46, "AMOUNT")
	headingWin.MovePrint(0, 55, "BALANCE")
	headingWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	headingWin.Refresh()
}

func updateTransactions() error {
	uncommitted, err := transaction.Uncommitted()
	if err != nil {
		return err
	}

	menuItems = make([]*gc.MenuItem, len(uncommitted))
	for i, n := range uncommitted {
		ft, err := stringf.Transaction(n)
		if err != nil {
			return err
		}
		str := fmt.Sprintf("%s  %-30s  %8s  %8s", ft.Date, ft.Description, ft.Cents)
		menuItems[i], _ = gc.NewItem(str, ft.Id)
	}

	if len(menuItems) == 0 {
		return errors.New("No data to show. Press ? for help.")
	}

	id := selectedTransaction()
	menu.UnPost()
	err = menu.SetItems(menuItems)
	menu.Post()
	selectTransaction(id)
	if err != nil {
		return err
	}
	return nil
}

func selectedTransaction() int64 {
	id, _ := dataf.Id(menu.Current(nil).Description())
	return id
}

func selectTransaction(id int64) {
	for _, item := range menu.Items() {
		desc, _ := strconv.ParseInt(item.Description(), 10, 64)
		if desc == id {
			menu.Current(item)
		}
	}
}
