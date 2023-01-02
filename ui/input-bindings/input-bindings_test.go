package inputBindings

import (
	"fmt"
	"jcb/ui"
	inputBindings "jcb/ui/input-bindings"
	"testing"

	"code.rocketnine.space/tslocum/cview"
)

var field = cview.NewInputField()
var c = inputBindings.Configuration(ui.HandleInputFormCustomBindings)

func TestDeleteChar(t *testing.T) {
	var got string
	var expect string
	field.SetInputCapture(c.Capture)

	field.SetText("abcd")
	field.SetCursorPosition(2)
	char := inputBindings.DeleteChar(field)

	got = char
	expect = "c"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = field.GetText()
	expect = "abd"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

func TestBackwardChar(t *testing.T) {
	var got int
	var expect int
	field.SetInputCapture(c.Capture)

	field.SetText("abcd")
	field.SetCursorPosition(2)
	inputBindings.BackwardChar(field)
	got = field.GetCursorPosition()
	expect = 1
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}

	field.SetText("abcd")
	field.SetCursorPosition(0)
	inputBindings.BackwardChar(field)
	got = field.GetCursorPosition()
	expect = 0
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
}

func TestBackwardWord(t *testing.T) {
	var got int
	var expect int
	field.SetInputCapture(c.Capture)

	field.SetText("one two")
	field.SetCursorPosition(6)
	inputBindings.BackwardWord(field)
	got = field.GetCursorPosition()
	expect = 4
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}

	field.SetText("one-two")
	field.SetCursorPosition(6)
	inputBindings.BackwardWord(field)
	got = field.GetCursorPosition()
	expect = 4
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}

	field.SetText("one_two")
	field.SetCursorPosition(6)
	inputBindings.BackwardWord(field)
	got = field.GetCursorPosition()
	expect = 0
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}

	field.SetText("one    two")
	field.SetCursorPosition(7)
	inputBindings.BackwardWord(field)
	got = field.GetCursorPosition()
	expect = 0
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
}

func TestForwardChar(t *testing.T) {
	var got int
	var expect int
	field.SetInputCapture(c.Capture)

	field.SetText("abcd")
	field.SetCursorPosition(2)
	inputBindings.ForwardChar(field)
	got = field.GetCursorPosition()
	expect = 3
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}

	field.SetText("abcd")
	field.SetCursorPosition(4)
	inputBindings.ForwardChar(field)
	got = field.GetCursorPosition()
	expect = 4
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
}

func TestForwardWord(t *testing.T) {
	var got int
	var expect int
	field.SetInputCapture(c.Capture)

	field.SetText("one two")
	field.SetCursorPosition(0)
	inputBindings.ForwardWord(field)
	got = field.GetCursorPosition()
	expect = 3
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}

	field.SetText("one-two")
	field.SetCursorPosition(0)
	inputBindings.ForwardWord(field)
	got = field.GetCursorPosition()
	expect = 3
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}

	field.SetText("one_two")
	field.SetCursorPosition(0)
	inputBindings.ForwardWord(field)
	got = field.GetCursorPosition()
	expect = 7
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}

	field.SetText("one    two")
	field.SetCursorPosition(3)
	inputBindings.ForwardWord(field)
	got = field.GetCursorPosition()
	expect = 10
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
}

func TestUnixWordRubout(t *testing.T) {
	var got string
	var expect string
	field.SetInputCapture(c.Capture)

	field.SetText("one two")
	field.SetCursorPosition(7)
	inputBindings.UnixWordRubout(field)
	got = field.GetText()
	expect = "one "
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	field.SetText("one    ")
	field.SetCursorPosition(7)
	inputBindings.UnixWordRubout(field)
	got = field.GetText()
	expect = ""
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	field.SetText("one-two")
	field.SetCursorPosition(7)
	inputBindings.UnixWordRubout(field)
	got = field.GetText()
	expect = ""
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

func TestOtherUnixWordRubout(t *testing.T) {
	var got string
	var expect string
	field.SetInputCapture(c.Capture)

	field.SetText("one two")
	field.SetCursorPosition(7)
	inputBindings.OtherUnixWordRubout(field)
	got = field.GetText()
	expect = "one "
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	field.SetText("one    ")
	field.SetCursorPosition(7)
	inputBindings.OtherUnixWordRubout(field)
	got = field.GetText()
	expect = ""
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	field.SetText("one-two")
	field.SetCursorPosition(7)
	inputBindings.OtherUnixWordRubout(field)
	got = field.GetText()
	expect = "one-"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

func TestDeleteWord(t *testing.T) {
	var got string
	var expect string
	field.SetInputCapture(c.Capture)

	field.SetText("one two")
	field.SetCursorPosition(0)
	inputBindings.DeleteWord(field)
	got = field.GetText()
	expect = " two"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	field.SetText("one    two")
	field.SetCursorPosition(0)
	inputBindings.DeleteWord(field)
	got = field.GetText()
	expect = "    two"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	field.SetText("    two")
	field.SetCursorPosition(0)
	inputBindings.DeleteWord(field)
	got = field.GetText()
	expect = ""
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	field.SetText("one-two")
	field.SetCursorPosition(0)
	inputBindings.DeleteWord(field)
	got = field.GetText()
	expect = "-two"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

func TestKillLine(t *testing.T) {
	var got string
	var expect string
	field.SetInputCapture(c.Capture)

	field.SetText("one two")
	field.SetCursorPosition(0)
	inputBindings.KillLine(field)
	got = field.GetText()
	expect = ""
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	field.SetText("one two")
	field.SetCursorPosition(4)
	inputBindings.KillLine(field)
	got = field.GetText()
	expect = "one "
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

func TestYank(t *testing.T) {
	var got string
	var expect string
	field.SetInputCapture(c.Capture)

	field.SetText("one two")
	field.SetCursorPosition(7)
	inputBindings.UnixWordRubout(field)
	field.SetCursorPosition(0)
	inputBindings.Yank(field)
	got = field.GetText()
	expect = "twoone "
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}
