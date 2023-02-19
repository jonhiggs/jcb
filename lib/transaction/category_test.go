package transaction

import "testing"

func TestCategoryNew(t *testing.T) {
	c := NewCategory()

	got := c.value
	expect := ""

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	if !c.Saved {
		t.Errorf("got %v, expected %v", c.Saved, true)
	}
}

func TestCategorySetText(t *testing.T) {
	var got string
	var expect string
	var err error

	{
		c := NewCategory()
		_ = c.SetText("category")
		got = c.GetText()
		expect = "category"

		if got != expect {
			t.Errorf("got %s, expected %s", got, expect)
		}
	}

	{
		c := NewCategory()
		err = c.SetText("one two")
		got = c.GetText()
		expect = ""

		if got != expect {
			t.Errorf("got %s, expected %s", got, expect)
		}

		if err == nil {
			t.Errorf("didn't receive expected error: %s", err)
		}
	}

	{
		c := NewCategory()
		err = c.SetText("a-really-long-long-category")
		got = c.GetText()
		expect = ""

		if got != expect {
			t.Errorf("got %s, expected %s", got, expect)
		}

		if err == nil {
			t.Errorf("didn't receive expected error: %s", err)
		}
	}

	{
		c := NewCategory()
		err = c.SetText("&&***")
		got = c.GetText()
		expect = ""

		if got != expect {
			t.Errorf("got %s, expected %s", got, expect)
		}

		if err == nil {
			t.Errorf("didn't receive expected error: %s", err)
		}
	}
}
