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

	got = Cents(0)
	expect = "0.00"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = Cents(10)
	expect = "0.10"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = Cents(100)
	expect = "1.00"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = Cents(-1)
	expect = "-0.01"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = Cents(-1)
	expect = "-0.01"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

func TestDate(t *testing.T) {
	var ts time.Time
	var got string
	var expect string

	expect = "2022-11-03"
	ts, _ = time.Parse("2006-01-02", expect)
	got = Date(ts)
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

func TestDescription(t *testing.T) {
	var got string
	var expect string
	var err error

	got = Description("   testing    ")
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

	got = Id(42)
	expect = "42"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

func TestTransaction(t *testing.T) {
	var got domain.StringTransaction
	var expect string

	got = Transaction(domain.Transaction{1, time.Now(), "testing", 1200, ""})
	expect = "testing"
	if got.Description != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got.Description, expect))
	}

	expect = "12.00"
	if got.Cents != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got.Cents, expect))
	}

}
