package ui

import (
	"errors"
	"jcb/lib/transaction"

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
				printError(err)
			}
			err = transaction.Delete(t)
			if err != nil {
				printError(err)
			}
			transactionMenu.Driver(gc.DriverActions[gc.KEY_UP])
			updateTransactions()
		case 'U':
			updateTransactions()
			printError(errors.New("updating"))
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
			renderTransactionInsert()
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

func scanTransactionInsert() int {
	transactionInsertWin.Keypad(true)
	transactionInsertWin.Refresh()

	transactionInsertForm.Driver(gc.REQ_FIRST_FIELD)
	transactionInsertForm.Driver(gc.REQ_END_LINE)

	for {
		ch := transactionInsertWin.GetChar()
		switch ch {
		case gc.KEY_RETURN:
			return INSERT
		case 1: // ctrl-a
			transactionInsertForm.Driver(gc.REQ_BEG_FIELD)
		case 5: // ctrl-e
			transactionInsertForm.Driver(gc.REQ_END_FIELD)
		case 11: // ctrl-k
			transactionInsertForm.Driver(gc.REQ_DEL_LINE)
		case 4, 33: // ctrl-d, delete
			transactionInsertForm.Driver(gc.REQ_DEL_CHAR)
		case 23, 27: // ctrl-w, esc/alt-backspace
			transactionInsertForm.Driver(gc.REQ_DEL_WORD)
		case gc.KEY_BACKSPACE:
			transactionInsertForm.Driver(gc.REQ_DEL_PREV)
		case gc.KEY_DOWN, gc.KEY_TAB:
			transactionInsertForm.Driver(gc.REQ_NEXT_FIELD)
			transactionInsertForm.Driver(gc.REQ_END_LINE)
		case 2, gc.KEY_LEFT:
			transactionInsertForm.Driver(gc.REQ_LEFT_CHAR)
		case 6, gc.KEY_RIGHT:
			transactionInsertForm.Driver(gc.REQ_RIGHT_CHAR)
		case 'q', 3:
			return ABORT
		default:
			transactionInsertForm.Driver(ch)
		}
	}
}
