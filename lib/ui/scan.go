package ui

import gc "github.com/rthornton128/goncurses"

func scanMain() {
	mainWin.Keypad(true)
	mainWin.Refresh()
	for {
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
		case 'i':
			renderTransactionAdd()
		case '?':
			renderHelp()
		case 3, 'q':
			return
		default:
			continue //transactionMenu.Driver(gc.DriverActions[ch])
		}
	}
}

func scanHelp() bool {
	helpWin.GetChar()
	return false
}

func scanTransactionAdd() {
	transactionAddWin.Keypad(true)
	transactionAddWin.Refresh()

	transactionAddForm.Driver(gc.REQ_FIRST_FIELD)
	transactionAddForm.Driver(gc.REQ_END_LINE)

	for {
		ch := transactionAddWin.GetChar()
		switch ch {
		case 1: // ctrl-a
			transactionAddForm.Driver(gc.REQ_BEG_FIELD)
		case 5: // ctrl-e
			transactionAddForm.Driver(gc.REQ_END_FIELD)
		case 11: // ctrl-k
			transactionAddForm.Driver(gc.REQ_DEL_LINE)
		case 4, 33: // ctrl-d, delete
			transactionAddForm.Driver(gc.REQ_DEL_CHAR)
		case 23, 27: // ctrl-w, esc/alt-backspace
			transactionAddForm.Driver(gc.REQ_DEL_WORD)
		case gc.KEY_BACKSPACE:
			transactionAddForm.Driver(gc.REQ_DEL_PREV)
		case gc.KEY_DOWN, gc.KEY_TAB:
			transactionAddForm.Driver(gc.REQ_NEXT_FIELD)
			transactionAddForm.Driver(gc.REQ_END_LINE)
		case 2, gc.KEY_LEFT:
			transactionAddForm.Driver(gc.REQ_LEFT_CHAR)
		case 6, gc.KEY_RIGHT:
			transactionAddForm.Driver(gc.REQ_RIGHT_CHAR)
		case 'q', 3:
			return
		default:
			transactionAddForm.Driver(ch)
		}
	}
}
