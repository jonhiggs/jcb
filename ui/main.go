package ui

import (
	openingBalance "jcb/lib/openingbalance"
	openingBalanceWin "jcb/ui/win/openingbalance"
	statusWin "jcb/ui/win/status"
	transactionMenuWin "jcb/ui/win/transaction/menu"

	gc "github.com/rthornton128/goncurses"
)

var maxY int
var maxX int
var mainWin *gc.Window

func Start() {
	stdscr, _ := gc.Init()
	defer gc.End()
	stdscr.Refresh()
	maxY, maxX = stdscr.MaxYX()

	if gc.HasColors() {
		gc.StartColor()
	}

	gc.Echo(false)
	gc.Raw(true)
	gc.Cursor(0)

	stdscr.Keypad(true)

	initColorPairs()
	statusWin.Show(maxY, maxX)

	_, err := openingBalance.Find(2022)
	if err != nil {
		openingBalanceWin.Show()
	}

	transactionMenuWin.Show(maxY, maxX)

	stdscr.GetChar()
}
