package transaction

import (
	"testing"
)

func TestDescriptionNew(t *testing.T) {
	d := NewDescription()

	got := d.value
	expect := ""

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	if !d.Saved {
		t.Errorf("got %v, expected %v", d.Saved, true)
	}
}

func TestDescriptionSetText(t *testing.T) {
	var got string
	var expect string

	d := NewDescription()
	d.SetText("a new description")
	got = d.GetText()
	expect = "a new description"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	if d.Saved {
		t.Errorf("got %v, expected %v", d.Saved, false)
	}
}
