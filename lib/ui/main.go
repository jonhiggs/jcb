package ui

import (
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
	initFooter()

	var err error
	mainWin, err = gc.NewWindow(maxY-1, maxX-2, 0, 1)
	if err != nil {
		printError(err)
	}
	mainWin.Keypad(true)

	InitTransactions()

	scanMain()
}

func initColorPairs() {
	gc.InitPair(0, gc.C_BLACK, gc.C_WHITE)
	gc.InitPair(1, gc.C_BLACK, gc.C_CYAN)
	gc.InitPair(1, gc.C_RED, gc.C_BLACK)
}

func printError(e error) {
	if e != nil {
		footerWin.Clear()
		footerWin.MovePrint(0, 0, e)
		footerWin.Refresh()
	}
}

//func newWindow() {
//	Window, err := gc.NewWindow(Height-1, Width-2, 0, 1)
//	if err != nil {
//		PrintError(err)
//	}
//	mainWin.Keypad(true)
//}
//
