package ui

import (
	"errors"
	"fmt"
	"jcb/lib/transaction"
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
	mainWin.MovePrint(2, 49, "AMOUNT")
	mainWin.MovePrint(2, 57, "BALANCE")
	mainWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	updateTransactions()

	transactionMenu.Post()

	return nil
}

func selectedTransactionId() int64 {
	idStr := transactionMenu.Current(nil).Description()
	i, _ := strconv.ParseInt(idStr, 10, 64)
	return i
}

//func deletetransactionMenuItem(id int64) error {
//	var newtransactionMenuItems []*gc.MenuItem
//	found := false
//	var pos *gc.MenuItem
//
//	for _, item := range transactionMenuItems {
//		i, _ := strconv.ParseInt(item.Description(), 10, 64)
//		if i != id {
//			newtransactionMenuItems = append(newtransactionMenuItems, item)
//		} else {
//			found = true
//			pos = transactionMenuItems[item.Index()-1]
//			item.Free()
//		}
//	}
//
//	if !found {
//		return errors.New("Failed to delete any items")
//	}
//
//	transactionMenu.UnPost()
//	transactionMenuItems = newtransactionMenuItems
//	err := transactionMenu.SetItems(transactionMenuItems)
//	if err != nil {
//		return err
//	}
//	transactionMenu.Current(pos)
//
//	transactionMenu.Post()
//	return err
//}

func updateTransactions() error {
	transactions, err := transaction.All()
	if err != nil {
		return err
	}

	transactionMenuItems = make([]*gc.MenuItem, len(transactions))
	for i, n := range transactions {
		ft := formatTransaction(n)
		str := fmt.Sprintf("%s  %s  %s", ft.Date, ft.Description, ft.Amount)
		transactionMenuItems[i], _ = gc.NewItem(str, ft.Id)
	}

	if len(transactionMenuItems) == 0 {
		return errors.New("No data to show. Press ? for help.")
	}

	transactionMenu.UnPost()
	err = transactionMenu.SetItems(transactionMenuItems)
	transactionMenu.Post()
	if err != nil {
		return err
	}
	return nil
}
