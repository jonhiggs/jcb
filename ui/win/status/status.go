package statusWin

import (
	gc "github.com/rthornton128/goncurses"
)

var win *gc.Window

func Show(y int) {
	win, _ = gc.NewWindow(1, 50, y-1, 2)
	win.ColorOn(1)
	Clear()
}

func Clear() {
	win.Clear()
}

func PrintError(e error) {
	if e != nil {
		win.Clear()
		win.AttrOn(gc.ColorPair(1) | gc.A_BOLD)
		defer win.AttrOff(gc.ColorPair(1) | gc.A_BOLD)
		win.MovePrint(0, 0, e)
		win.Refresh()
		win.Clear()
	}
}

func Refresh() {
	win.Touch()
	win.Refresh()
}
