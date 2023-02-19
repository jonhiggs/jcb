package ui

import (
	"fmt"
	"jcb/config"
	"jcb/db"
	"strings"
)

func runCommand(command string) {
	if command == "" {
		return
	}
	cmd := strings.Fields(command)
	switch cmd[0] {
	case "version":
		printStatus(config.VERSION)
	case "w":
		db.Save()
		for _, t := range transactions {
			t.Date.Saved = true
			t.Description.Saved = true
			t.Cents.Saved = true
			t.Note.Saved = true
			t.Category.Saved = true
		}
		updateTransactionsTable()
		updateInfo()
		printStatus("File saved")
	case "refresh":
		updateTransactionsTable()
		printStatus("Refreshed the transactions")
	case "quit", "q":
		if db.IsDirty() {
			printStatus("You have unsaved changes. Use ':q!' to quit without saving.")
		} else {
			app.Stop()
		}
	case "help":
		openHelp()
	case "wq":
		db.Save()
		app.Stop()
	case "q!":
		app.Stop()
	default:
		printStatus(fmt.Sprintf("Unknown command '%s'", promptInputField.GetText()))
	}
}
