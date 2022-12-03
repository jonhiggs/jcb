package win

import (
	"fmt"
	"jcb/config"

	gc "github.com/rthornton128/goncurses"
)

var win *gc.Window

func Show() {
	win, _ = gc.NewWindow(23, 45, 6, 10)
	defer win.Delete()
	win.Box(0, 0)

	win.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	win.MovePrint(1, 3, "General")
	win.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	win.MovePrint(3, 3, "^C   Close")
	win.MovePrint(4, 3, "j    Down")
	win.MovePrint(5, 3, "k    Up")
	win.MovePrint(6, 3, "J    Down Month")
	win.MovePrint(7, 3, "K    Up Month")
	win.MovePrint(8, 3, "Tab  Next Input")
	win.MovePrint(9, 3, "[    Prior Year")
	win.MovePrint(10, 3, "]    Next Year")
	win.MovePrint(11, 3, "u    Page Up")
	win.MovePrint(12, 3, "d    Page Down")
	win.MovePrint(13, 3, "g    Start of Year")
	win.MovePrint(14, 3, "G    End of Year")

	win.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	win.MovePrint(10, 25, "Transaction")
	win.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	win.MovePrint(12, 25, "i    Create")
	win.MovePrint(13, 25, "C    Commit")
	win.MovePrint(14, 25, "L    Unlock")
	win.MovePrint(15, 25, "x    Delete")
	win.MovePrint(16, 25, "e    Edit")
	win.MovePrint(17, 25, "I    Import File")
	win.MovePrint(18, 25, "W    Write")

	win.AttrOn(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)
	win.MovePrint(01, 25, "Input")
	win.AttrOff(gc.ColorPair(0) | gc.A_BOLD | gc.A_UNDERLINE)

	win.MovePrint(3, 25, "^A   Begining")
	win.MovePrint(4, 25, "^E   End")
	win.MovePrint(5, 25, "^K   Clear")
	win.MovePrint(6, 25, "^D   Abort")
	win.MovePrint(7, 25, "^B   Backward")
	win.MovePrint(8, 25, "^F   Forward")

	win.AttrOn(gc.ColorPair(0) | gc.A_BOLD)
	win.MovePrint(21, 38-len(config.VERSION), fmt.Sprintf("jcb v%s", config.VERSION))
	win.AttrOff(gc.ColorPair(0) | gc.A_BOLD)

	scan()
}

func scan() {
	win.GetChar()
}
