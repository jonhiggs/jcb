package transaction

import "testing"

func TestCentsGetText(t *testing.T) {
	c := NewCents()

	if !c.Saved {
		t.Errorf("got %v, expected %v", c.Saved, true)
	}

	c.SetValue(123)

	got := c.GetText()
	expect := "1.23"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	if c.Saved {
		t.Errorf("got %v, expected %v", c.Saved, false)
	}

}

func TestCentsSetText(t *testing.T) {
	c := NewCents()
	var got string
	var expect string

	_ = c.SetText("1.23")
	got = c.GetText()
	expect = "1.23"

	if c.Saved {
		t.Errorf("got %v, expected %v", c.Saved, false)
	}

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

	c := NewCents()

	c.SetValue(1000)
	got = c.IsCredit()
	expect = true

	if got != expect {
		t.Errorf("got %v, expected %v", got, expect)
	}

	c.SetValue(-1000)
	got = c.IsCredit()
	expect = false

	if got != expect {
		t.Errorf("got %v, expected %v", got, expect)
	}
}

func TestCentsIsDebit(t *testing.T) {
	var got bool
	var expect bool

	c := NewCents()

	c.SetValue(1000)
	got = c.IsDebit()
	expect = false

	if got != expect {
		t.Errorf("got %v, expected %v", got, expect)
	}

	c.SetValue(-1000)
	got = c.IsDebit()
	expect = true

	if got != expect {
		t.Errorf("got %v, expected %v", got, expect)
	}
}

func TestCentsAdd(t *testing.T) {
	var got int
	var expect int

	c := NewCents()

	got = c.Add(20)
	expect = 20

	if got != expect {
		t.Errorf("got %v, expected %v", got, expect)
	}

	got = c.Add(20)
	expect = 40

	if got != expect {
		t.Errorf("got %v, expected %v", got, expect)
	}

	if c.Saved {
		t.Errorf("got %v, expected %v", c.Saved, false)
	}
}
