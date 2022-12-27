package ui

import "code.rocketnine.space/tslocum/cview"

var status *cview.TextView

func handleCloseStatus() {
	panels.HidePanel("status")
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
