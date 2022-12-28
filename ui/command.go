package ui

import (
	"fmt"
	"jcb/config"
	"jcb/db"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func handleCommand(ev *tcell.EventKey) *tcell.EventKey {
	openPrompt(":", "", func(ev *tcell.EventKey) *tcell.EventKey {
		panels.HidePanel("prompt")
		runCommand(promptInputField.GetText())
		return nil
	})

	return nil
}

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
		printStatus("File saved")
	case "refresh":
		updateTransactionsTable()
		printStatus("Refreshed the transactions")
	case "quit", "q":
		if db.Dirty() {
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
