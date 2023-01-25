package transaction

import "testing"

func TestNoteNew(t *testing.T) {
	n := new(Note)

	got := n.value
	expect := ""

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}
}

func TestNoteSetText(t *testing.T) {
	var got string
	var expect string

	n := new(Note)
	n.SetText("a new note")
	got = n.GetText()
	expect = "a new note"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}
}
