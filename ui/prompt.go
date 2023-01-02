package ui

import (
	"jcb/config"
	inputBindings "jcb/ui/input-bindings"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var promptForm *cview.Form
var promptInputField *cview.InputField

func createPromptForm() *cview.Form {
	promptForm = cview.NewForm()
	promptForm.SetBorder(false)
	promptForm.SetCancelFunc(closePrompt)
	promptForm.SetItemPadding(0)
	promptForm.SetPadding(0, 0, 0, 0)
	promptForm.SetLabelColor(tcell.ColorWhite)
	promptForm.SetFieldBackgroundColor(tcell.ColorBlack)
	promptForm.SetFieldBackgroundColorFocused(tcell.ColorBlack)

	promptInputField = cview.NewInputField()

	promptForm.AddFormItem(promptInputField)

	return promptForm
}

func closePrompt() {
	panels.HidePanel("prompt")
}

func openPrompt(
	label string,
	text string,
	enterFunc func(ev *tcell.EventKey) *tcell.EventKey,
	acceptanceFunc func(textToCheck string, lastChar rune) bool,
) {
	panels.ShowPanel("prompt")
	panels.SendToFront("prompt")

	promptInputField.SetLabel(label)
	promptInputField.SetText(text)

	c := inputBindings.Configuration(HandleInputFormCustomBindings)
	c.SetKey(0, tcell.KeyEnter, enterFunc)

	promptInputField.SetInputCapture(c.Capture)
	promptInputField.SetFieldWidth(config.MAX_WIDTH - len(label))
	promptInputField.SetAcceptanceFunc(acceptanceFunc)
	promptForm.SetFocus(0)
	return
}
