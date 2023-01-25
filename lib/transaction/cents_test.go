package transaction

import "testing"

func TestCentsNew(t *testing.T) {
	c := new(Cents)

	got := c.value
	expect := 0

	if got != expect {
		t.Errorf("got %d, expected %d", got, expect)
	}
}

func TestCentsGetText(t *testing.T) {
	c := Cents{123}

	got := c.GetText()
	expect := "1.23"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

}

func TestCentsSetText(t *testing.T) {
	c := new(Cents)
	var got string
	var expect string

	_ = c.SetText("1.23")
	got = c.GetText()
	expect = "1.23"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	_ = c.SetText("1.234")
	got = c.GetText()
	expect = "1.23"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	_ = c.SetText("01.23")
	got = c.GetText()
	expect = "1.23"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	_ = c.SetText("1")
	got = c.GetText()
	expect = "1.00"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	_ = c.SetText("-01")
	got = c.GetText()
	expect = "-1.00"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}
}

func TestCentsIsCredit(t *testing.T) {
	var got bool
	var expect bool

	got = (&Cents{1000}).IsCredit()
	expect = true

	if got != expect {
		t.Errorf("got %v, expected %v", got, expect)
	}

	got = (&Cents{-1000}).IsCredit()
	expect = false

	if got != expect {
		t.Errorf("got %v, expected %v", got, expect)
	}
}

func TestCentsIsDebit(t *testing.T) {
	var got bool
	var expect bool

	got = (&Cents{1000}).IsDebit()
	expect = false

	if got != expect {
		t.Errorf("got %v, expected %v", got, expect)
	}

	got = (&Cents{-1000}).IsDebit()
	expect = true

	if got != expect {
		t.Errorf("got %v, expected %v", got, expect)
	}
}

func TestCentsAdd(t *testing.T) {
	var got int
	var expect int

	c := Cents{1000}
	got = c.Add(20)
	expect = 1020

	if got != expect {
		t.Errorf("got %v, expected %v", got, expect)
	}
}
