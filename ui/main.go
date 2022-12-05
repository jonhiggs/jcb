package ui

import (
	"fmt"
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
	if maxX < 72 {
		gc.End()
		fmt.Println("Your terminal must be at least 72 chars wide.")
		return
	}

	if gc.HasColors() {
		gc.StartColor()
	}

	gc.Echo(false)
	gc.Raw(true)
	gc.Cursor(0)

	stdscr.Keypad(true)

	initColorPairs()
	statusWin.Show(maxY, maxX)

	transactionMenuWin.Show(maxY, maxX)
}
