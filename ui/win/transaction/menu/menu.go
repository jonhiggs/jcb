package menuWin

import (
	"errors"
	"fmt"
	"jcb/lib/transaction"
	dataf "jcb/ui/formatter/data"
	stringf "jcb/ui/formatter/string"
	statusWin "jcb/ui/win/status"
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
	menu.Format(y-4, 1)
	menu.Mark("")
	menu.Option(gc.O_SHOWDESC, false)
	menu.SetWindow(win)

	err = updateTransactions()
	if err != nil {
		statusWin.PrintError(err)
	}

	menu.Post()
	win.Refresh()

	return nil
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
		balance = balance - cents
		balanceStr, _ = stringf.Cents(balance)
		str := fmt.Sprintf("%s  %-30s  %8s  %8s", ft.Date, ft.Description, ft.Cents, balanceStr)
		menuItems[i+1], _ = gc.NewItem(str, ft.Id)
	}

	if len(menuItems) == 0 {
		return errors.New("No data to show. Press ? for help.")
	}

	//st, _ := selectedTransaction()
	menu.UnPost()
	err = menu.SetItems(menuItems)
	menu.Post()
	//selectTransaction(st.Id)
	if err != nil {
		return err
	}
	return nil
}

func selectTransaction(id int64) {
	for _, item := range menu.Items() {
		desc, _ := strconv.ParseInt(item.Description(), 10, 64)
		if desc == id {
			menu.Current(item)
		}
	}
}
