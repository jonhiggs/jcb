package ui

import (
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
	initFooter()

	var err error
	mainWin, err = gc.NewWindow(maxY-1, maxX-2, 0, 1)
	if err != nil {
		printError(err)
	}

	initTransactions()

	scanMain()
}

func printError(e error) {
	if e != nil {
		footerWin.Clear()
		footerWin.MovePrint(0, 0, e)
		footerWin.Refresh()
	}
}
