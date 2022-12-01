package ui

import (
	openingBalance "jcb/lib/openingbalance"
	openingBalanceWin "jcb/lib/ui/win/openingbalance"
	statusWin "jcb/lib/ui/win/status"

	gc "github.com/rthornton128/goncurses"
)

var maxY int
var maxX int
var mainWin *gc.Window

const (
	CONTINUE = 0
	ABORT    = 1
	INSERT   = 2
)

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

	var err error
	mainWin, err = gc.NewWindow(maxY-1, maxX-2, 0, 1)
	if err != nil {
		statusWin.PrintError(err)
	}

	_, err = openingBalance.Find(2022)
	if err != nil {
		openingBalanceWin.Show()
	}

	initTransactions()
	scanMain()
}
