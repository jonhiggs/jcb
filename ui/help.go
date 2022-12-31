package ui

import (
	"code.rocketnine.space/tslocum/cbind"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

var helpTextView *cview.TextView

func closeHelp() {
	panels.HidePanel("help")
	closeStatus()
}

func openHelp() {
	panels.ShowPanel("help")
	panels.SendToFront("help")
	panels.HidePanel("report")
	printStatus("Press <space> to scroll further")
}

func createHelp() *cview.TextView {
	helpTextView = cview.NewTextView()
	helpTextView.SetScrollBarVisibility(cview.ScrollBarNever)
	helpTextView.SetText(
		`
 Commands:

	h, help      Print this help
	q!           Quit without writing
	q, quit      Quit
	refresh      Refresh transactions
	version      Show version
	w            Write changes to disk
	wq           Write changes and quit


 Key Bindings:

	i            Insert new transaction
	x            Delete selected transaction
	<Enter>      Edit selected transaction
	j            Select next transaction
	k            Select previous transaction
	^d           Scroll half a page down
	t            Toggle tag for selected transaction
	T            Tag matching transactions
	<C-t>        Untag matching transactions
	<C-u>        Scroll half a page up
	0            Select oldest uncommmitted transaction
	*            Select similar transaction
	}            Scroll to next month
	{            Scroll to previous month
	]            Scroll next modified transaction
	[            Scroll previous modified transaction
	>            Scroll to next year
	<            Scroll to previous year
	C            Commit all prior transactions
	c            Commit single transaction
	r            Repeat transaction
	D            Edit category
	d            Edit description
	=            Edit amount
	@            Edit date
	:            Enter command mode
	/            Enter find next query
	?            Enter find previous query
	n            Next matching transaction
	N            Previous matching transaction
	<F1>         Show the help page
	<F2>         Show the transactions page
	<F3>         Show the report page
	;x           Delete tagged transactions
	;t, ;<C-t>   Untag tagged transactions
	;D           Edit category of tagged transactions
	;d           Edit description of tagged transactions
	;=           Edit amount of tagged transactions
	;@           Edit date of tagged transactions


 Input Field Key Bindings:

	<Enter>      Submit form
	<Esc>, <C-c> Cancel form
	C-a          Start of field
	C-e          End of field
	C-w          Delete word backwards
	C-u          Clear field
	M-f          Forward word
	M-b          Backwards word
`)
	helpTextView.SetDoneFunc(func(key tcell.Key) { closeHelp() })

	c := cbind.NewConfiguration()
	c.Set(" ", handleHelpScroll)
	c.Set("u", handleHelpScroll)
	c.Set("d", handleHelpScroll)
	c.Set("F1", handleOpenHelp)
	c.Set("F2", handleOpenTransactions)
	c.Set("F3", handleOpenReport)
	c.Set("q", func(ev *tcell.EventKey) *tcell.EventKey { closeHelp(); return nil })
	helpTextView.SetInputCapture(c.Capture)

	return helpTextView
}
