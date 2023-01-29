package transaction

import (
	"testing"
)

func TestUpdateCategorySingle(t *testing.T) {
	var got string
	var expect string

	d := new(Transaction)
	d.Category.SetText("a")

	got = d.Category.GetText()
	expect = "a"
	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	resp := UpdateCategory("b", []*Transaction{d})
	got = resp[0].Category.GetText()
	expect = "b"
	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	if len(resp) != 1 {
		t.Errorf("got %d, expected %d", len(resp), 1)
	}
}

func TestUpdateCategoryNoop(t *testing.T) {
	d := new(Transaction)
	d.Category.SetText("a")

	resp := UpdateCategory("a", []*Transaction{d})

	if len(resp) != 0 {
		t.Errorf("got %d, expected %d", len(resp), 0)
	}
}

func TestUpdateCategoryMultiple(t *testing.T) {
	var got string
	var expect string

	d := new(Transaction)
	d.Category.SetText("a")
	e := new(Transaction)
	e.Category.SetText("b")

	resp := UpdateCategory("c", []*Transaction{d, e})

	got = resp[0].Category.GetText()
	expect = "c"
	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	got = resp[1].Category.GetText()
	expect = "c"
	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}
}
