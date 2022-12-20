package ui

import (
	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var helpTextView *cview.TextView

func handleCloseHelp() {
	panels.HidePanel("help")
	handleCloseStatus()
}

func handleHelpScroll(ev *tcell.EventKey) *tcell.EventKey {
	_, _, _, h := helpTextView.GetInnerRect()
	offset, _ := helpTextView.GetScrollOffset()
	switch ev.Rune() {
	case ' ', 'd':
		pos := offset + (h / 2)
		helpTextView.ScrollTo(pos, 0)
	case 'u':
		pos := offset - (h / 2)
		helpTextView.ScrollTo(pos, 0)
	}

	return nil
}

func openHelp() {
	panels.ShowPanel("help")
	printStatus("do blah for blah")
}

func createHelp() *cview.TextView {
	helpTextView = cview.NewTextView()
	helpTextView.SetText(
		`
 Commands:

	h, help    Print this help
	q, quit    Quit
	w          Write changes to disk
	wq         Write changes and quit
	q!         Quit without writing
	version    Show version


 Key Bindings:

	i          Insert new transaction
	x          Delete selected transaction
	<enter>    Edit selected transaction
	j          Select next transaction
	k          Select previous transaction
	d          Scroll half a page down
	u          Scroll half a page up
	0          Select oldest uncommmitted transaction
	*          Select similar transaction
	}          Scroll to next month
	{          Scroll to previous month
	C          Commit all transactions to selection
	:          Enter command mode
	/          Enter find next query
	?          Enter find previous query
	n          Next matching transaction
	N          Previous matching transaction


 Input Field Key Bindings:

	<enter>    Submit form
	<esc>, C-c Cancel form
	C-a        Start of field
	C-e        End of field
	C-w        Delete word backwards
	C-u        Clear field
	M-f        Forward word
	M-b        Backwards word
`)
	helpTextView.SetDoneFunc(func(key tcell.Key) { handleCloseHelp() })

	c := cbind.NewConfiguration()
	c.Set(" ", handleHelpScroll)
	c.Set("u", handleHelpScroll)
	c.Set("d", handleHelpScroll)
	helpTextView.SetInputCapture(c.Capture)

	return helpTextView
}