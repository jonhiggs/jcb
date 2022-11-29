package main

import (
	"jcb/db"
	"jcb/lib/ui"
)

func main() {
	db.Init()
	ui.Start()
}
