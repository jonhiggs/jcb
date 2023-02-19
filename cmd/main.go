package main

import (
	"fmt"
	"jcb/config"
	"jcb/db"
	"jcb/lib/tsv"
	"jcb/ui"
	"log"
	"os"
	"runtime/pprof"

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
	fmt.Println("  -i, --import-tsv <file>   Import transactions from TSV file.")
	fmt.Println("  -e, --export-tsv          Export transactions as TSV to stdout.")
}

func main() {
	options := []optparse.Option{
		{"file", 'f', optparse.KindRequired},
		{"import-tsv", 'i', optparse.KindRequired},
		{"export-tsv", 'e', optparse.KindNone},
		{"help", 'h', optparse.KindNone},
		{"version", 'v', optparse.KindNone},
		{"cpu-profile", 'p', optparse.KindRequired},
	}

	file := config.DefaultFile()
	tsvFile := ""
	exportTsv := false
	var cpuProfile string

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
			println("jcb v" + config.VERSION)
			return
		case "import-tsv":
			tsvFile = result.Optarg
		case "export-tsv":
			exportTsv = true
		case "cpu-profile":
			cpuProfile = result.Optarg
		}
	}

	if cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	err = db.Init(file)
	if err != nil {
		log.Fatal(err)
	}

	if exportTsv {
		tsv.Export()
		return
	}

	if tsvFile != "" {
		tsv.Import(tsvFile)
		db.Save()
		db.RemoveWorkingFile()
		return
	}

	ui.Start()

	db.RemoveWorkingFile()
}
