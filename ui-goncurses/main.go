package ui

import (
	"fmt"
	statusWin "jcb/ui/win/status"
	transactionMenuWin "jcb/ui/win/transaction/menu"

	gc "github.com/rthornton128/goncurses"
)

var mainWin *gc.Window

func Start(year int) {
	stdscr, _ := gc.Init()
	defer gc.End()
	stdscr.Refresh()
	maxY, maxX := stdscr.MaxYX()
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
	statusWin.Show(maxY)

	transactionMenuWin.Show(maxY, year)
}
