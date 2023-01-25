package transaction

import "testing"

func TestDescriptionNew(t *testing.T) {
	d := new(Description)

	got := d.value
	expect := ""

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}
}

func TestDescriptionSetText(t *testing.T) {
	var got string
	var expect string

	d := new(Description)
	d.SetText("a new description")
	got = d.GetText()
	expect = "a new description"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}
}
