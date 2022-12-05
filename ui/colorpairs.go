package ui

import gc "github.com/rthornton128/goncurses"

func initColorPairs() {
	gc.InitPair(1, gc.C_RED, gc.C_BLACK)    //error
	gc.InitPair(2, gc.C_YELLOW, gc.C_BLACK) //titles
	gc.InitPair(3, gc.C_RED, gc.C_BLACK)    //negative
	gc.InitPair(4, gc.C_BLUE, gc.C_BLACK)   //positive

}
