package stringFormatter

import (
	"fmt"
	"jcb/domain"
	"testing"
	"time"
)

func TestCents(t *testing.T) {
	var got string
	var expect string

	got, _ = Cents(0)
	expect = "0.00"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got, _ = Cents(10)
	expect = "0.10"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got, _ = Cents(100)
	expect = "1.00"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got, _ = Cents(-1)
	expect = "-0.01"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

func TestDate(t *testing.T) {
	var ts time.Time
	var got string
	var expect string
	var err error

	expect = "2022-11-03"
	ts, _ = time.Parse("2006-01-02", expect)
	got, _ = Date(ts)
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	expect = "2022-13-03"
	ts, _ = time.Parse("2006-01-02", expect)
	got, err = Date(ts)
	if got != "" {
		t.Error(fmt.Sprintf("got %s, expected %s", got, ""))
	}
	if err == nil {
		t.Error(fmt.Sprintf("expected error for %s", expect))
	}
}

func TestDescription(t *testing.T) {
	var got string
	var expect string
	var err error

	got, err = Description("   testing    ")
	expect = "testing"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %s", expect))
	}
}

func TestId(t *testing.T) {
	var got string
	var expect string
	var err error

	got, err = Id(42)
	expect = "42"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %s", expect))
	}

	got, err = Id(-42)
	expect = "0"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
	if err == nil {
		t.Error(fmt.Sprintf("error expected for %d", -42))
	}
}

func TestTransaction(t *testing.T) {
	_, err := Transaction(domain.Transaction{1, time.Now(), "testing", 1200})
	if err != nil {
		t.Error("no error expected")
	}
}
