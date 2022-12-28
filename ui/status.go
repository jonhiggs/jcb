package ui

import "code.rocketnine.space/tslocum/cview"

var status *cview.TextView

func closeStatus() {
	panels.HidePanel("status")
}

func clearStatus() {
	status.SetText("")
}

func printStatus(message string) {
	status.SetText(message)
	panels.ShowPanel("status")
	panels.HidePanel("command")
	panels.HidePanel("prompt")
	panels.SendToBack("status")
}

func createStatusTextView() *cview.TextView {
	status = cview.NewTextView()
	return status
}
