package ui

import (
	"errors"
	"jcb/lib/transaction"
	helpWin "jcb/lib/ui/win/help"
	statusWin "jcb/lib/ui/win/status"
	transactionInsertWin "jcb/lib/ui/win/transaction/insert"

	gc "github.com/rthornton128/goncurses"
)

func scanMain() {
	mainWin.Keypad(true)
	mainWin.Refresh()
	for {
		ch := transactionWin.GetChar()
		switch ch {
		case 'x':
			t, err := selectedTransaction()
			if err != nil {
				statusWin.PrintError(err)
			}
			err = transaction.Delete(t)
			if err != nil {
				statusWin.PrintError(err)
			}
			transactionMenu.Driver(gc.DriverActions[gc.KEY_UP])
			updateTransactions()
		case 'U':
			updateTransactions()
			statusWin.PrintError(errors.New("updating"))
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
			transactionMenu.Driver(gc.DriverActions[gc.KEY_HOME])
		case 'G':
			transactionMenu.Driver(gc.DriverActions[gc.KEY_END])
		case 'd':
			transactionMenu.Driver(gc.DriverActions[gc.KEY_PAGEDOWN])
		case 'u':
			transactionMenu.Driver(gc.DriverActions[gc.KEY_PAGEUP])
		case 'j':
			transactionMenu.Driver(gc.DriverActions[gc.KEY_DOWN])
		case 'k':
			transactionMenu.Driver(gc.DriverActions[gc.KEY_UP])
		//case 'J':
		//	curItem := transactionMenu.Current(nil)
		//	curMonth, _ := strconv.ParseInt(curItem.Name()[5:7], 10, 64)
		//	for _, item := range transactionMenu.Items() {
		//		thisMonth, _ := strconv.ParseInt(item.Name()[5:7], 10, 64)
		//		if thisMonth > curMonth {
		//			transactionMenu.Current(item)
		//			ui.MainWindow.MovePrint(0, 0, item.Index())
		//			break
		//		}
		//	}
		//case 'K':
		//	//curDate := uiStringify.Date(transactionMenu.Current(nil).Name())
		//	//items := transactionMenu.Items()
		//	//ui.MainWindow.MovePrint(0, 0, thisMonth)
		//	//ui.MainWindow.MovePrint(1, 0, curMonth)
		//	//ui.MainWindow.MovePrint(2, 0, i)
		//	//for i := len(items) - 1; items[i].Index() != transactionMenu.Current(nil).Index(); i = i - 1 {
		//	//}
		case 'i':
			transactionInsertWin.Show()
			mainWin.Touch()
			mainWin.Refresh()
			statusWin.Refresh()
		case '?':
			helpWin.Show()
			mainWin.Touch()
			mainWin.Refresh()
			statusWin.Refresh()
		case 3, 'q':
			return
		default:
			continue //transactionMenu.Driver(gc.DriverActions[ch])
		}
	}
}
