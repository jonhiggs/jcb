package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

const VERSION = "0.2.0"

const MAX_WIDTH = 80
const INFO_WIDTH = 20

const DESCRIPTION_MAX_LENGTH = 32
const CATEGORY_MAX_LENGTH = 10
const NOTES_MAX_LENGTH = 200

const ATTR_COLUMN = 0
const DATE_COLUMN = 1
const CATEGORY_COLUMN = 2
const DESCRIPTION_COLUMN = 3
const AMOUNT_COLUMN = 4
const BALANCE_COLUMN = 5

const COLOR_TITLE_FG = tcell.ColorYellow
const COLOR_TITLE_BG = tcell.Color25
const COLOR_LIGHT_BG = tcell.Color232
const COLOR_COMMITTED_FG = tcell.Color81
const COLOR_COMMITTED_BG = tcell.ColorBlack
const COLOR_UNCOMMITTED_FG = tcell.ColorWhite
const COLOR_UNCOMMITTED_BG = tcell.ColorBlack
const COLOR_MODIFIED_FG = tcell.Color184
const COLOR_MODIFIED_BG = tcell.Color234
const COLOR_TAGGED_FG = tcell.ColorGreen
const COLOR_TAGGED_BG = tcell.Color234
const COLOR_NEGATIVE_FG = tcell.Color160
const COLOR_POSITIVE_FG = tcell.Color40
const COLOR_INFO_FG = tcell.Color238
const COLOR_FORMFIELD_BG = tcell.Color234
const COLOR_FORMFIELD_FG = tcell.Color250
const COLOR_FORMFIELD_FOCUSED_FG = tcell.ColorWhite
const COLOR_FORMFIELD_FOCUSED_BG = tcell.Color23
const COLOR_FORM_LABLE_FG = tcell.ColorWhite

func DefaultFile() string {
	data_dir, ok := os.LookupEnv("XDG_DATA_HOME")

	if !ok {
		home, ok := os.LookupEnv("HOME")
		if !ok {
			log.Fatal("Cannot determine home directory set the default database file.")
		}
		data_dir = fmt.Sprintf("%s/.local/share/jcb", home)
	}
	return fmt.Sprintf("%s/data.db", data_dir)
}
