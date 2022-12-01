package main

import (
	"jcb/db"
	"jcb/ui"
	"log"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal(err)
	}
	ui.Start()
}
