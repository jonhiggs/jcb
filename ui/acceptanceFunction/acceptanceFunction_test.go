package acceptanceFunction

import (
	"jcb/ui/acceptanceFunction"
	"testing"

	"code.rocketnine.space/tslocum/cview"
)

var field = cview.NewInputField()

func TestDescription(t *testing.T) {
	var got bool
	var expect bool

	field.SetText("this is a description")
	got = acceptanceFunction.Description(field)
	expect = true
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("01234567890123456789012345678901")
	got = acceptanceFunction.Description(field)
	expect = true
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("012345678901234567890123456789012")
	got = acceptanceFunction.Date(field)
	expect = false
	if got != expect {
		t.Error("didn't get expected result")
	}
}

func TestNotes(t *testing.T) {
	var got bool
	var expect bool

	field.SetText("this is a description")
	got = acceptanceFunction.Notes(field)
	expect = true
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("01234567890123456789012345678901012345678901234567890123456789010123456789012345678901234567890101234567890123456789012345678901012345678901234567890123456789010123456789012345678901234567890101234567")
	got = acceptanceFunction.Notes(field)
	expect = true
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("012345678901234567890123456789010123456789012345678901234567890101234567890123456789012345678901012345678901234567890123456789010123456789012345678901234567890101234567890123456789012345678901012345678")
	got = acceptanceFunction.Notes(field)
	expect = false
	if got != expect {
		t.Error("didn't get expected result")
	}
}

func TestCategory(t *testing.T) {
	var got bool
	var expect bool

	field.SetText("category")
	got = acceptanceFunction.Category(field)
	expect = true
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("")
	got = acceptanceFunction.Category(field)
	expect = true
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("no spaces")
	got = acceptanceFunction.Category(field)
	expect = false
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("0123456789")
	got = acceptanceFunction.Category(field)
	expect = true
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("01234567890")
	got = acceptanceFunction.Category(field)
	expect = false
	if got != expect {
		t.Error("didn't get expected result")
	}
}

func TestDate(t *testing.T) {
	var got bool
	var expect bool

	field.SetText("2022-01-01")
	got = acceptanceFunction.Date(field)
	expect = true
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("2022-1-1")
	got = acceptanceFunction.Date(field)
	expect = true
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("abcd")
	got = acceptanceFunction.Date(field)
	expect = false
	if got != expect {
		t.Error("didn't get expected result")
	}
}

func TestCents(t *testing.T) {
	var got bool
	var expect bool

	field.SetText("12.34")
	got = acceptanceFunction.Cents(field)
	expect = true
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("12.1234")
	got = acceptanceFunction.Date(field)
	expect = false
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText(".12")
	got = acceptanceFunction.Date(field)
	expect = false
	if got != expect {
		t.Error("didn't get expected result")
	}

	field.SetText("12")
	got = acceptanceFunction.Date(field)
	expect = true
	if got != expect {
		t.Error("didn't get expected result")
	}
}
