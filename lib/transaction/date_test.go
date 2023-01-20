package transaction

import (
	"fmt"
	"testing"
	"time"
)

func TestDateNew(t *testing.T) {
	d := new(Date)

	got := d.value.Format("2006-01-02")
	expect := "0001-01-01"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}
}

func TestDateGetText(t *testing.T) {
	d := new(Date)
	d.value = time.Date(2020, 2, 3, 0, 0, 0, 0, time.UTC)

	got := d.GetText()
	expect := "2020-02-03"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}
}

func TestDateSetText(t *testing.T) {
	d := new(Date)
	var err error
	var got string
	var expect string

	// a form date
	err = d.SetText("2020-02-03")
	got = d.GetText()
	expect = "2020-02-03"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	if err != nil {
		t.Errorf("received unexpected error: %s", err)
	}

	// a database date
	err = d.SetText("2023-11-20 00:00:00+00:00")
	got = d.GetText()
	expect = "2023-11-20"

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}

	if err != nil {
		t.Errorf("received unexpected error: %s", err)
	}

	// an invalid date
	got = fmt.Sprint(d.SetText("abcd"))
	expect = `setting date from string: parsing time "abcd" as "2006-01-02": cannot parse "abcd" as "2006"`

	if got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}
}

func TestDateYear(t *testing.T) {
	d := new(Date)
	d.value = time.Date(2020, 2, 3, 0, 0, 0, 0, time.UTC)

	got := d.Year()
	expect := 2020

	if got != expect {
		t.Errorf("got %d, expected %d", got, expect)
	}
}

func TestDateUnix(t *testing.T) {
	d := new(Date)
	d.value = time.Date(2020, 2, 3, 0, 0, 0, 0, time.UTC)

	got := d.Unix()
	expect := int64(1580688000)

	if got != expect {
		t.Errorf("got %d, expected %d", got, expect)
	}
}
