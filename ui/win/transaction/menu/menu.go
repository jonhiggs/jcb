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

	win, err = gc.NewWindow(y-4, x-1, 3, 1)
	menu.SubWindow(win)
	menu.Mark("")
	menu.Option(gc.O_SHOWDESC, false)
	menu.SetWindow(win)

	err = updateTransactions()
	if err != nil {
		statusWin.PrintError(err)
	}

	menu.UnPost()
	menu.Format(y-4, 1)
	menu.Post()

	err = scan()
	for err != nil {
		statusWin.PrintError(err)
		err = scan()
	}

	win.Refresh()

	return nil
}

func scan() error {
	win.Keypad(true)
	win.Refresh()

	for {
		ch := win.GetChar()
		switch ch {
		case 'x':
			id, _ := selectedTransaction()
			err := transaction.DeleteId(id)
			if err != nil {
				statusWin.PrintError(err)
			} else {
				menu.Driver(gc.DriverActions[gc.KEY_UP])
				updateTransactions()
			}
		//case 'U':
		//	updateTransactions()
		//	statusWin.PrintError(errors.New("updating"))
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
			menu.Driver(gc.DriverActions[gc.KEY_PAGEDOWN])
		case 'u':
			menu.Driver(gc.DriverActions[gc.KEY_PAGEUP])
		case 'j':
			menu.Driver(gc.DriverActions[gc.KEY_DOWN])
		case 'k':
			menu.Driver(gc.DriverActions[gc.KEY_UP])
		//case 'J':
		//	curItem := menu.Current(nil)
		//	curMonth, _ := strconv.ParseInt(curItem.Name()[5:7], 10, 64)
		//	for _, item := range menu.Items() {
		//		thisMonth, _ := strconv.ParseInt(item.Name()[5:7], 10, 64)
		//		if thisMonth > curMonth {
		//			menu.Current(item)
		//			ui.MainWindow.MovePrint(0, 0, item.Index())
		//			break
		//		}
		//	}
		//case 'K':
		//	//curDate := uiStringify.Date(menu.Current(nil).Name())
		//	//items := menu.Items()
		//	//ui.MainWindow.MovePrint(0, 0, thisMonth)
		//	//ui.MainWindow.MovePrint(1, 0, curMonth)
		//	//ui.MainWindow.MovePrint(2, 0, i)
		//	//for i := len(items) - 1; items[i].Index() != menu.Current(nil).Index(); i = i - 1 {
		//	//}
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
	headingWin.MovePrint(2, 0, "DATE")
	headingWin.MovePrint(2, 12, "DESCRIPTION")
	headingWin.MovePrint(2, 46, "AMOUNT")
	headingWin.MovePrint(2, 55, "BALANCE")
	headingWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	headingWin.Refresh()
}

func updateTransactions() error {
	transactions, err := transaction.All()
	if err != nil {
		return err
	}

	menuItems = make([]*gc.MenuItem, len(transactions)+1)
	var balance int64
	balance = 12345
	balanceStr, _ := stringf.Cents(balance)
	str := fmt.Sprintf("%s  %-30s  %8s  %8s", "2022-01-01", "Opening Balance", balanceStr, balanceStr)
	menuItems[0], _ = gc.NewItem(str, "opening_balance")
	for i, n := range transactions {
		ft, _ := stringf.Transaction(n)
		cents, _ := dataf.Cents(ft.Cents)
		balance = balance + cents
		balanceStr, _ = stringf.Cents(balance)
		str := fmt.Sprintf("%s  %-30s  %8s  %8s", ft.Date, ft.Description, ft.Cents, balanceStr)
		menuItems[i+1], _ = gc.NewItem(str, ft.Id)
	}

	if len(menuItems) == 0 {
		return errors.New("No data to show. Press ? for help.")
	}

	id, _ := selectedTransaction()
	menu.UnPost()
	err = menu.SetItems(menuItems)
	menu.Post()
	selectTransaction(id)
	if err != nil {
		return err
	}
	return nil
}

func selectedTransaction() (int64, error) {
	return dataf.Id(menu.Current(nil).Description())
}

func selectTransaction(id int64) {
	for _, item := range menu.Items() {
		desc, _ := strconv.ParseInt(item.Description(), 10, 64)
		if desc == id {
			menu.Current(item)
		}
	}
}
