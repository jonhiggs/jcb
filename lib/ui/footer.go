package ui

import gc "github.com/rthornton128/goncurses"

var footerWin *gc.Window

func initFooter() {
	footerWin, _ = gc.NewWindow(1, maxX-2, maxY-1, 2)
	footerWin.ColorOn(1)
	clearError()
}
