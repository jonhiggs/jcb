package ui

import (
	"fmt"
	"jcb/config"
	"jcb/db"
	"strings"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var commandForm *cview.Form
var commandInputField *cview.InputField

func handleCloseCommand() {
	panels.HidePanel("command")
}

func handleOpenCommand(ev *tcell.EventKey) *tcell.EventKey {
	panels.ShowPanel("command")
	panels.SendToFront("command")

	commandInputField.SetText("")
	commandForm.SetFocus(0)
	return nil
}

func handleCommand(ev *tcell.EventKey) *tcell.EventKey {
	runCommand(commandInputField.GetText())

	commandInputField.SetText("")
	panels.HidePanel("command")

	return nil
}

func runCommand(command string) {
	cmd := strings.Fields(command)
	switch cmd[0] {
	case "version":
		printStatus(config.VERSION)
	case "savetime":
		printStatus(fmt.Sprintf("%s", db.SaveTime))
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
		printStatus(fmt.Sprintf("Unknown command '%s'", commandInputField.GetText()))
	}
}

func createCommandForm() *cview.Form {
	commandForm = cview.NewForm()
	commandForm.SetBorder(false)
	commandForm.SetCancelFunc(handleCloseCommand)
	commandForm.SetItemPadding(0)
	commandForm.SetPadding(0, 0, 0, 0)
	commandForm.SetLabelColor(tcell.ColorWhite)
	commandForm.SetFieldBackgroundColor(tcell.ColorBlack)
	commandForm.SetFieldBackgroundColorFocused(tcell.ColorBlack)

	commandInputField = cview.NewInputField()
	commandInputField.SetFieldWidth(24)
	commandInputField.SetLabel(":")

	commandForm.AddFormItem(commandInputField)

	c := cbind.NewConfiguration()
	c.SetKey(0, tcell.KeyEnter, handleCommand)
	commandInputField.SetInputCapture(c.Capture)

	return commandForm
}
