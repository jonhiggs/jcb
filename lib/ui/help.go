package ui

import gc "github.com/rthornton128/goncurses"

var helpWin *gc.Window

func renderHelp() {
	helpWin, _ = gc.NewWindow(23, 45, 6, 10)
	defer helpWin.Delete()
	helpWin.Box(0, 0)

	helpWin.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	helpWin.MovePrint(1, 3, "General")
	helpWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	helpWin.MovePrint(3, 3, "^C   Close")
	helpWin.MovePrint(4, 3, "j    Down")
	helpWin.MovePrint(5, 3, "k    Up")
	helpWin.MovePrint(6, 3, "j    Down Month")
	helpWin.MovePrint(7, 3, "k    Up Month")
	helpWin.MovePrint(8, 3, "Tab  Next Input")
	helpWin.MovePrint(9, 3, "[    Prior Year")
	helpWin.MovePrint(10, 3, "]    Next Year")
	helpWin.MovePrint(11, 3, "u    Page Up")
	helpWin.MovePrint(12, 3, "d    Page Down")
	helpWin.MovePrint(13, 3, "g    Start of Year")
	helpWin.MovePrint(14, 3, "G    End of Year")

	helpWin.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	helpWin.MovePrint(10, 25, "Transaction")
	helpWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	helpWin.MovePrint(12, 25, "i    Create")
	helpWin.MovePrint(13, 25, "l    Lock")
	helpWin.MovePrint(14, 25, "L    Unlock")
	helpWin.MovePrint(15, 25, "x    Delete")
	helpWin.MovePrint(16, 25, "e    Edit")
	helpWin.MovePrint(17, 25, "I    Import File")
	helpWin.MovePrint(18, 25, "W    Write")

	helpWin.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	helpWin.MovePrint(01, 25, "Input")
	helpWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	helpWin.MovePrint(3, 25, "^A   Begining")
	helpWin.MovePrint(4, 25, "^E   End")
	helpWin.MovePrint(5, 25, "^K   Clear")
	helpWin.MovePrint(6, 25, "^D   Abort")
	helpWin.MovePrint(7, 25, "^B   Backward")
	helpWin.MovePrint(8, 25, "^F   Forward")

	helpWin.AttrOn(gc.ColorPair(0) | gc.A_BOLD)
	helpWin.MovePrint(21, 33, "jcb v0.0.0")
	helpWin.AttrOff(gc.ColorPair(0) | gc.A_BOLD)

	scanHelp()

	mainWin.Touch()
	mainWin.Refresh()
	footerWin.Touch()
	footerWin.Refresh()
}
