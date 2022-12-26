package config

import (
	"fmt"
	"log"
	"os"
)

const VERSION = "0.0.0"
const DESC_MAX_LENGTH = 26
const NOTES_MAX_LENGTH = 200

func DefaultFile() string {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		log.Fatal("Cannot determine home directory set the default database file.")
	}
	return fmt.Sprintf("%s/.config/jcb/data.db", home)
}
