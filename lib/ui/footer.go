package ui

import gc "github.com/rthornton128/goncurses"

var footerWin *gc.Window

func initFooter() {
	footerWin, _ = gc.NewWindow(1, maxX-2, maxY-1, 2)
	footerWin.ColorOn(1)
	footerWin.AttrOn(gc.ColorPair(1))
	footerWin.MovePrint(0, 0, "[Min Balance: 2022-09-24 $203.33]")
	footerWin.Refresh()
	footerWin.AttrOff(gc.ColorPair(1))
}
