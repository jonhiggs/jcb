package main

import (
	"fmt"
	"jcb/config"
	"jcb/db"
	"jcb/lib/importer"
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
	fmt.Println("  -h, --help                This help.")
	fmt.Println("  -f, --file <file>         Set the savefile.")
	fmt.Println("  -v, --version             Show the version.")
	fmt.Println("  -i, --import-tsv <file>   Import transactions from TSV.")
}

func main() {
	options := []optparse.Option{
		{"file", 'f', optparse.KindRequired},
		{"import-tsv", 'i', optparse.KindRequired},
		{"help", 'h', optparse.KindNone},
		{"version", 'v', optparse.KindNone},
	}

	file := config.DefaultFile()
	tsvFile := ""
	fmt.Printf("Loading file %s\n", file)

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
		case "import-tsv":
			tsvFile = result.Optarg
		}
	}

	err = db.Init(file)
	if err != nil {
		log.Fatal(err)
	}

	if tsvFile != "" {
		importer.Tsv(tsvFile)
		db.Save()
		db.RemoveWorkingFile()
		return
	}

	ui.Start()

	db.RemoveWorkingFile()
}
