package ui

import gc "github.com/rthornton128/goncurses"

func initColorPairs() {
	gc.InitPair(0, gc.C_BLACK, gc.C_WHITE)
	gc.InitPair(1, gc.C_BLACK, gc.C_CYAN)
	gc.InitPair(1, gc.C_RED, gc.C_BLACK)
}
