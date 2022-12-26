package config

import (
	"fmt"
	"log"
	"os"
)

const VERSION = "0.0.0"

const MAX_WIDTH = 80

const DESC_MAX_LENGTH = 32
const NOTES_MAX_LENGTH = 200

const ATTR_COLUMN = 0
const DATE_COLUMN = 1
const CATEGORY_COLUMN = 2
const DESCRIPTION_COLUMN = 3
const AMOUNT_COLUMN = 4
const BALANCE_COLUMN = 5

func DefaultFile() string {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		log.Fatal("Cannot determine home directory set the default database file.")
	}
	return fmt.Sprintf("%s/.config/jcb/data.db", home)
}
