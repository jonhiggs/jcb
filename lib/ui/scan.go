package ui

func scanMain() bool {
	mainWin.Refresh()
	ch := transactionWin.GetChar()
	switch ch {
	//case 'W':
	//	ui.Commit()
	//	ui.MainWindow.Touch()
	//	ui.MainWindow.Refresh()
	//case 'x':
	//	idStr := uiTransaction.TransactionMenu.Current(nil).Description()
	//	id, _ := strconv.ParseInt(idStr, 10, 64)

	//	db.DeleteTransaction(id)
	//	err := uiTransaction.DeleteTransactionMenuItem(id)
	//	if err != nil {
	//		ui.PrintError(err)
	//	}
	//case 'i':
	//	err := ui.NewTransaction()
	//	if err != nil {
	//		ui.PrintError(err)
	//	}
	//	uiTransaction.UpdateTransactions()
	//	uiTransaction.TransactionWindow.Touch()
	//	uiTransaction.TransactionWindow.Refresh()
	//	ui.MainWindow.Touch()
	//	ui.MainWindow.Refresh()
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
	//case 'g':
	//	uiTransaction.TransactionMenu.Driver(gc.DriverActions[gc.KEY_HOME])
	//case 'G':
	//	uiTransaction.TransactionMenu.Driver(gc.DriverActions[gc.KEY_END])
	//case 'd':
	//	uiTransaction.TransactionMenu.Driver(gc.DriverActions[gc.KEY_PAGEDOWN])
	//case 'u':
	//	uiTransaction.TransactionMenu.Driver(gc.DriverActions[gc.KEY_PAGEUP])
	//case 'j':
	//	uiTransaction.TransactionMenu.Driver(gc.DriverActions[gc.KEY_DOWN])
	//case 'J':
	//	curItem := uiTransaction.TransactionMenu.Current(nil)
	//	curMonth, _ := strconv.ParseInt(curItem.Name()[5:7], 10, 64)
	//	for _, item := range uiTransaction.TransactionMenu.Items() {
	//		thisMonth, _ := strconv.ParseInt(item.Name()[5:7], 10, 64)
	//		if thisMonth > curMonth {
	//			uiTransaction.TransactionMenu.Current(item)
	//			ui.MainWindow.MovePrint(0, 0, item.Index())
	//			break
	//		}
	//	}
	//case 'k':
	//	uiTransaction.TransactionMenu.Driver(gc.DriverActions[gc.KEY_UP])
	//case 'K':
	//	//curDate := uiStringify.Date(uiTransaction.TransactionMenu.Current(nil).Name())
	//	//items := uiTransaction.TransactionMenu.Items()
	//	//ui.MainWindow.MovePrint(0, 0, thisMonth)
	//	//ui.MainWindow.MovePrint(1, 0, curMonth)
	//	//ui.MainWindow.MovePrint(2, 0, i)
	//	//for i := len(items) - 1; items[i].Index() != uiTransaction.TransactionMenu.Current(nil).Index(); i = i - 1 {
	//	//}
	case '?':
		ui.Help()
		ui.MainWindow.Touch()
		ui.MainWindow.Refresh()
	case 'q':
		return false
	case 3:
		return false
	default:
		return true //transactionMenu.Driver(gc.DriverActions[ch])
	}
	return true
}
