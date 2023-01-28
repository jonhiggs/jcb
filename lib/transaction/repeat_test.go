package transaction

import (
	"testing"
	"time"
)

func TestExpandDaily(t *testing.T) {
	d := new(Transaction)
	d.Date.SetText("2022-01-01")
	d.Description.SetText("the description")
	d.Cents.SetValue(1200)
	d.Category.SetText("test")

	var expect string

	got, err := Expand(d, "1d", 1, time.Date(2022, 12, 31, 23, 59, 59, 59, time.UTC))

	if len(got) != 365 {
		t.Errorf("expected 365 but got %d", len(got))
	}

	expect = "2022-01-01"
	if got[0].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[0].Date.GetText())
	}

	expect = "2022-01-02"
	if got[1].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[1].Date.GetText())
	}

	expect = "2022-01-04"
	if got[3].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[3].Date.GetText())
	}

	if err != nil {
		t.Errorf("expected no errors but got %s", err)
	}
}

func TestExpandFourDaily(t *testing.T) {
	d := new(Transaction)
	d.Date.SetText("2022-01-01")
	d.Description.SetText("the description")
	d.Cents.SetValue(1200)
	d.Category.SetText("test")

	var expect string

	got, err := Expand(d, "4d", 1, time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC))

	if len(got) != 8 {
		t.Errorf("expected 4 but got %d", len(got))
	}

	expect = "2022-01-05"
	if got[1].Date.GetText() != expect {
		t.Errorf("expected 2022-01-05 but got %s", got[1].Date.GetText())
	}

	if err != nil {
		t.Errorf("expected no errors but got %s", err)
	}
}

func TestExpandZeroDaily(t *testing.T) {
	d := new(Transaction)
	d.Date.SetText("2022-01-01")
	d.Description.SetText("the description")
	d.Cents.SetValue(1200)
	d.Category.SetText("test")

	var expect string

	got, err := Expand(d, "0d", 1, time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC))

	if len(got) != 1 {
		t.Errorf("expected 1 but got %d", len(got))
	}

	expect = "2022-01-01"
	if got[0].Date.GetText() != expect {
		t.Errorf("expected 2022-01-01 but got %s", got[0].Date.GetText())
	}

	if err != nil {
		t.Errorf("expected no errors but got %s", err)
	}

}

func TestExpandWeekly(t *testing.T) {
	d := new(Transaction)
	d.Date.SetText("2022-01-01")
	d.Description.SetText("the description")
	d.Cents.SetValue(1200)
	d.Category.SetText("test")

	var expect string

	got, err := Expand(d, "1w", 1, time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC))

	if len(got) != 9 {
		t.Errorf("expected 4 but got %d", len(got))
	}

	expect = "2022-01-01"
	if got[0].Date.GetText() != expect {
		t.Errorf("expected 2022-01-01 but got %s", got[0].Date.GetText())
	}

	expect = "2022-01-08"
	if got[1].Date.GetText() != expect {
		t.Errorf("expected 2022-01-08 but got %s", got[1].Date.GetText())
	}

	if err != nil {
		t.Errorf("expected no errors but got %s", err)
	}
}

func TestExpandMonthly(t *testing.T) {
	d := new(Transaction)
	d.Date.SetText("2022-01-31")
	d.Description.SetText("the description")
	d.Cents.SetValue(1200)
	d.Category.SetText("test")

	var expect string

	got, err := Expand(d, "1m", 1, time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC))

	if len(got) != 13 {
		t.Errorf("expected 13 but got %d", len(got))
	}

	expect = "2022-01-31"
	if got[0].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[0].Date.GetText())
	}

	expect = "2022-03-03"
	if got[1].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[1].Date.GetText())
	}

	expect = "2022-03-31"
	if got[2].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[2].Date.GetText())
	}

	expect = "2023-01-31"
	if got[12].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[12].Date.GetText())
	}

	if err != nil {
		t.Errorf("expected no errors but got %s", err)
	}
}
