package transaction

import (
	"testing"
	"time"
)

func TestRepeatDaily(t *testing.T) {
	d := new(Transaction)
	d.Date.SetText("2022-01-01")
	d.Description.SetText("the description")
	d.Cents.SetValue(1200)
	d.Category.SetText("test")

	var expect string

	got, err := d.Repeat("1d", 1, time.Date(2022, 12, 31, 23, 59, 59, 59, time.UTC))

	if len(got) != 364 {
		t.Errorf("expected 364 but got %d", len(got))
	}

	expect = "2022-01-02"
	if got[0].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[0].Date.GetText())
	}

	expect = "2022-01-04"
	if got[2].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[2].Date.GetText())
	}

	if err != nil {
		t.Errorf("expected no errors but got %s", err)
	}
}

func TestRepeatFourDaily(t *testing.T) {
	d := new(Transaction)
	d.Date.SetText("2022-01-01")
	d.Description.SetText("the description")
	d.Cents.SetValue(1200)
	d.Category.SetText("test")

	var expect string

	got, err := d.Repeat("4d", 1, time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC))

	if len(got) != 7 {
		t.Errorf("expected 7 but got %d", len(got))
	}

	expect = "2022-01-05"
	if got[0].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[0].Date.GetText())
	}

	if err != nil {
		t.Errorf("expected no errors but got %s", err)
	}
}

func TestRepeatZeroDaily(t *testing.T) {
	d := new(Transaction)
	d.Date.SetText("2022-01-01")
	d.Description.SetText("the description")
	d.Cents.SetValue(1200)
	d.Category.SetText("test")

	got, err := d.Repeat("0d", 1, time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC))

	if len(got) != 0 {
		t.Errorf("expected 0 but got %d", len(got))
	}

	if err != nil {
		t.Errorf("expected no errors but got %s", err)
	}

}

func TestRepeatWeekly(t *testing.T) {
	d := new(Transaction)
	d.Date.SetText("2022-01-01")
	d.Description.SetText("the description")
	d.Cents.SetValue(1200)
	d.Category.SetText("test")

	var expect string

	got, err := d.Repeat("1w", 1, time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC))

	if len(got) != 8 {
		t.Errorf("expected 8 but got %d", len(got))
	}

	expect = "2022-01-08"
	if got[0].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[0].Date.GetText())
	}

	if err != nil {
		t.Errorf("expected no errors but got %s", err)
	}
}

func TestRepeatMonthly(t *testing.T) {
	d := new(Transaction)
	d.Date.SetText("2022-01-31")
	d.Description.SetText("the description")
	d.Cents.SetValue(1200)
	d.Category.SetText("test")

	var expect string

	got, err := d.Repeat("1m", 1, time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC))

	if len(got) != 12 {
		t.Errorf("expected 12 but got %d", len(got))
	}

	expect = "2022-03-03"
	if got[0].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[0].Date.GetText())
	}

	expect = "2022-03-31"
	if got[1].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[1].Date.GetText())
	}

	expect = "2023-01-31"
	if got[11].Date.GetText() != expect {
		t.Errorf("expected %s but got %s", expect, got[11].Date.GetText())
	}

	if err != nil {
		t.Errorf("expected no errors but got %s", err)
	}
}
