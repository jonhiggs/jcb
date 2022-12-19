package ui

import "code.rocketnine.space/tslocum/cview"

var status *cview.TextView

func handleCloseStatus() {
	panels.HidePanel("status")
}

func createStatusTextView() *cview.TextView {
	status = cview.NewTextView()
	return status
}
