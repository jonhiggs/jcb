package main

import (
	"fmt"
	"jcb/config"
	"jcb/db"
	"jcb/ui"
	"log"
	"os"

	"nullprogram.com/x/optparse"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  jcb [OPTIONS]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -h, --help       This help.")
	fmt.Println("  -f, --file       Set the savefile.")
	fmt.Println("  -v, --version    Show the version.")
}

func main() {
	options := []optparse.Option{
		{"file", 'f', optparse.KindRequired},
		{"help", 'h', optparse.KindNone},
		{"version", 'v', optparse.KindNone},
	}

	file := config.DefaultFile()

	results, _, err := optparse.Parse(options, os.Args)
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		switch result.Long {
		case "help":
			usage()
			return
		case "file":
			file = result.Optarg
		case "version":
			println(config.VERSION)
			return
		}
	}

	err = db.Init(file)
	if err != nil {
		log.Fatal(err)
	}

	ui.Start()

}
