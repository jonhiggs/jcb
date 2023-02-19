package transaction

import "testing"

func TestNoteNew(t *testing.T) {
	n := NewNote()

	got := n.value
	expect := ""

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	if !n.Saved {
		t.Errorf("got %v, expected %v", n.Saved, true)
	}
}

func TestNoteSetText(t *testing.T) {
	var got string
	var expect string

	n := NewNote()
	n.SetText("a new note")
	got = n.GetText()
	expect = "a new note"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	if n.Saved {
		t.Errorf("got %v, expected %v", n.Saved, false)
	}
}
