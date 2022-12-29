package ui

import (
	"jcb/config"
	promptBindings "jcb/ui/prompt-bindings"
	"time"

	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var app *cview.Application
var panels *cview.Panels
var lowestBalance int64
var lowestBalanceDate time.Time
var findQuery string

func Start() {
	app = cview.NewApplication()

	box0 := cview.NewTextView()
	box0.SetDynamicColors(true)
	box0.SetRegions(true)
	box0.SetWordWrap(true)
	box0.SetText("")
	box0.SetTextAlign(cview.AlignCenter)

	balance := cview.NewTextView()
	balance.SetDynamicColors(true)
	balance.SetRegions(true)
	balance.SetWordWrap(true)
	balance.SetText("balance")
	balance.SetTextAlign(cview.AlignRight)

	panels = cview.NewPanels()
	panels.AddPanel("transactions", createTransactionsTable(), false, true)
	panels.AddPanel("report", createReportTable(), false, false)
	panels.AddPanel("insert", createInsertForm(), false, false)
	panels.AddPanel("edit", createEditForm(), false, false)
	panels.AddPanel("prompt", createPromptForm(), false, false)
	panels.AddPanel("status", createStatusTextView(), false, false)
	panels.AddPanel("info", createInfoTextView(), false, false)
	panels.AddPanel("help", createHelp(), false, false)

	panels.ShowPanel("info")
	panels.SendToBack("info")
	updateInfo()

	c := cbind.NewConfiguration()
	handleExit := func(ev *tcell.EventKey) *tcell.EventKey {
		pn, _ := panels.GetFrontPanel()
		if pn == "transactions" {
			printStatus("To quit, use the command ':q'.")
		} else {
			closeInsert()
			closeEdit()
			closePrompt()
			closeHelp()
		}
		return nil
	}

	c.SetRune(tcell.ModCtrl, 'c', handleExit)

	app.SetInputCapture(c.Capture)

	app.SetAfterResizeFunc(func(w int, h int) {
		transactionsTable.SetRect(0, 0, config.MAX_WIDTH, h-1)
		status.SetRect(0, h-1, config.MAX_WIDTH-config.INFO_WIDTH, h)
		info.SetRect(config.MAX_WIDTH-config.INFO_WIDTH, h-1, config.INFO_WIDTH, h)
		helpTextView.SetRect(0, 0, config.MAX_WIDTH, h-1)
		reportTable.SetRect(0, 0, w, h-1)
		promptForm.SetRect(0, h-1, config.MAX_WIDTH, h)
	})

	app.SetRoot(panels, true)
	app.Run()
}

func handleInputFormCustomBindings(ev *tcell.EventKey) *tcell.EventKey {
	pn, _ := panels.GetFrontPanel()
	var field *cview.InputField
	switch pn {
	case "edit":
		fieldId, _ := editForm.GetFocusedItemIndex()
		field = editForm.GetFormItem(fieldId).(*cview.InputField)
	case "insert":
		fieldId, _ := insertForm.GetFocusedItemIndex()
		field = insertForm.GetFormItem(fieldId).(*cview.InputField)
	case "prompt":
		fieldId, _ := promptForm.GetFocusedItemIndex()
		field = promptForm.GetFormItem(fieldId).(*cview.InputField)
	}

	switch ev.Key() {
	case tcell.KeyCtrlD:
		promptBindings.DeleteChar(field)
	case tcell.KeyCtrlF:
		promptBindings.ForwardChar(field)
	case tcell.KeyCtrlB:
		promptBindings.BackwardChar(field)
	case tcell.KeyCtrlW:
		promptBindings.UnixWordRubout(field)
	}
	return nil
}
