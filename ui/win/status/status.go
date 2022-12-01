package statusWin

import gc "github.com/rthornton128/goncurses"

var win *gc.Window

func Show(y int, x int) {
	win, _ = gc.NewWindow(1, x-2, y-1, 2)
	win.ColorOn(1)
	Clear()
}

func Clear() {
	win.ColorOn(1)
	win.AttrOn(gc.ColorPair(1))
	win.MovePrint(0, 0, "[Min Balance: 2022-09-24 $203.33]")
	win.AttrOff(gc.ColorPair(1))
	win.Refresh()
}

func PrintError(e error) {
	if e != nil {
		win.Clear()
		win.MovePrint(0, 0, e)
		win.Refresh()
		win.GetChar()
		win.Clear()
	}
}

func Refresh() {
	win.Touch()
	win.Refresh()
}
