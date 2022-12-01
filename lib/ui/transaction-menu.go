package ui

import (
	"errors"
	"fmt"
	"jcb/domain"
	"jcb/lib/transaction"
	sformat "jcb/lib/ui/formatter/string"
	"strconv"

	gc "github.com/rthornton128/goncurses"
)

var transactionWin *gc.Window
var transactionMenuItems []*gc.MenuItem
var transactionMenu *gc.Menu

func initTransactions() error {
	var err error
	transactionMenu, err = gc.NewMenu(make([]*gc.MenuItem, 0))
	if err != nil {
		printError(err)
	}

	transactionWin = mainWin.Derived(0, 0, 3, 1)
	transactionMenu.SubWindow(transactionWin)
	transactionMenu.Format(maxY-4, 1)
	transactionMenu.Mark("")
	transactionMenu.Option(gc.O_SHOWDESC, false)
	transactionMenu.SetWindow(mainWin)

	mainWin.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	mainWin.MovePrint(2, 1, "DATE")
	mainWin.MovePrint(2, 13, "DESCRIPTION")
	mainWin.MovePrint(2, 47, "AMOUNT")
	mainWin.MovePrint(2, 56, "BALANCE")
	mainWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	err = updateTransactions()
	if err != nil {
		printError(err)
	}

	transactionMenu.Post()

	return nil
}

func selectedTransaction() (domain.Transaction, error) {
	idStr := transactionMenu.Current(nil).Description()
	i, _ := strconv.ParseInt(idStr, 10, 64)
	return transaction.Find(i)
}

func updateTransactions() error {
	transactions, err := transaction.All()
	if err != nil {
		return err
	}

	transactionMenuItems = make([]*gc.MenuItem, len(transactions))
	for i, n := range transactions {
		ft, _ := sformat.Transaction(n)
		str := fmt.Sprintf("%s  %-30s  %8s", ft.Date, ft.Description, ft.Cents)
		transactionMenuItems[i], _ = gc.NewItem(str, ft.Id)
	}

	if len(transactionMenuItems) == 0 {
		return errors.New("No data to show. Press ? for help.")
	}

	st, _ := selectedTransaction()
	transactionMenu.UnPost()
	err = transactionMenu.SetItems(transactionMenuItems)
	transactionMenu.Post()
	selectTransaction(st.Id)
	if err != nil {
		return err
	}
	return nil
}

func selectTransaction(id int64) {
	for _, item := range transactionMenu.Items() {
		desc, _ := strconv.ParseInt(item.Description(), 10, 64)
		if desc == id {
			transactionMenu.Current(item)
		}
	}
}
