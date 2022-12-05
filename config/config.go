package config

import (
	"fmt"
	"log"
	"os"
)

const VERSION = "0.0.0"
const WIDTH = 72

func DefaultFile() string {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		log.Fatal("Cannot determine home directory set the default database file.")
	}
	return fmt.Sprintf("%s/.config/jcb/data.db", home)
}
