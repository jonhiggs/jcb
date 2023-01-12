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
