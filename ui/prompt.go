package ui

import (
	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var promptForm *cview.Form
var promptInputField *cview.InputField

func createPromptForm() *cview.Form {
	promptForm = cview.NewForm()
	promptForm.SetBorder(false)
	promptForm.SetCancelFunc(handleClosePrompt)
	promptForm.SetItemPadding(0)
	promptForm.SetPadding(0, 0, 0, 0)
	promptForm.SetLabelColor(tcell.ColorWhite)
	promptForm.SetFieldBackgroundColor(tcell.ColorBlack)
	promptForm.SetFieldBackgroundColorFocused(tcell.ColorBlack)

	promptInputField = cview.NewInputField()
	promptInputField.SetFieldWidth(24)

	promptForm.AddFormItem(promptInputField)

	return promptForm
}

func handleClosePrompt() {
	panels.HidePanel("prompt")
}

func openPrompt(label string, text string, enterFunc func(ev *tcell.EventKey) *tcell.EventKey) {
	panels.ShowPanel("prompt")
	panels.SendToFront("prompt")

	promptInputField.SetLabel(label)
	promptInputField.SetText(text)

	c := cbind.NewConfiguration()
	c.SetKey(0, tcell.KeyEnter, enterFunc)
	promptInputField.SetInputCapture(c.Capture)
	promptForm.SetFocus(0)
	return
}
