package ui

import (
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

	transactionWin = mainWin.Derived(0, maxY-3, 3, 1)
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
//
//func updateTransactions() error {
//	var menuNames []string
//	var menuDescriptions []string
//
//	for _, r := range db.AllTransactions() {
//		date := uiStringify.FormatDate(r.Date)
//		description := uiStringify.FormatDescription(r.Description)
//		cents := uiStringify.FormatCents(r.Cents)
//		itemStr := fmt.Sprintf("%s  %-30s  %s", date, description, cents)
//		menuNames = append(menuNames, itemStr)
//		menuDescriptions = append(menuDescriptions, strconv.FormatInt(r.Id, 10))
//	}
//
//	transactionMenuItems = make([]*gc.MenuItem, len(menuNames))
//	for i, n := range menuNames {
//		transactionMenuItems[i], _ = gc.NewItem(n, menuDescriptions[i])
//	}
//
//	if len(transactionMenuItems) == 0 {
//		return errors.New("No data to show. Press ? for help.")
//	}
//
//	transactionMenu.UnPost()
//	err := transactionMenu.SetItems(transactionMenuItems)
//	transactionMenu.Post()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//