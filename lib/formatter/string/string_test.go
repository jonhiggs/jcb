package stringFormatter

import (
	"fmt"
	"jcb/domain"
	"testing"
)

//func TestCents(t *testing.T) {
//	var got string
//	var expect string
//
//	got = Cents(0)
//	expect = "0.00"
//	if got != expect {
//		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
//	}
//
//	got = Cents(10)
//	expect = "0.10"
//	if got != expect {
//		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
//	}
//
//	got = Cents(100)
//	expect = "1.00"
//	if got != expect {
//		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
//	}
//
//	got = Cents(-1)
//	expect = "-0.01"
//	if got != expect {
//		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
//	}
//
//	got = Cents(-1)
//	expect = "-0.01"
//	if got != expect {
//		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
//	}
//}

//func TestDate(t *testing.T) {
//	var ts time.Time
//	var got string
//	var expect string
//
//	expect = "2022-11-03"
//	ts, _ = time.Parse("2006-01-02", expect)
//	got = Date(ts)
//	if got != expect {
//		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
//	}
//}

//func TestDescription(t *testing.T) {
//	var got string
//	var expect string
//	var err error
//
//	got = Description("   testing    ")
//	expect = "testing"
//	if got != expect {
//		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
//	}
//	if err != nil {
//		t.Error(fmt.Sprintf("no error expected for %s", expect))
//	}
//}

func TestNotes(t *testing.T) {
	var got string
	var expect string
	var err error

	got = Category("   notes    ")
	expect = "notes"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %s", expect))
	}
}

func TestCategory(t *testing.T) {
	var got string
	var expect string
	var err error

	got = Category("   category    ")
	expect = "category"
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

func TestAttributes(t *testing.T) {
	var got string
	var expect string

	got = Attributes(domain.Attributes{true, true, true, true})
	expect = "Cn "
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = Attributes(domain.Attributes{true, true, false, true})
	expect = "Cn+"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = Attributes(domain.Attributes{false, false, true, true})
	expect = "   "
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

//func TestTransaction(t *testing.T) {
//	var got domain.StringTransaction
//	var expect string
//
//	got = Transaction(domain.Transaction{1, time.Now(), "testing", 1200, "notes", "category"})
//	expect = "testing"
//	if got.Description != expect {
//		t.Error(fmt.Sprintf("got %s, expected %s", got.Description, expect))
//	}
//
//	expect = "12.00"
//	if got.Cents != expect {
//		t.Error(fmt.Sprintf("got %s, expected %s", got.Cents, expect))
//	}
//
//	expect = "notes"
//	if got.Notes != expect {
//		t.Error(fmt.Sprintf("got %s, expected %s", got.Notes, expect))
//	}
//
//	expect = "category"
//	if got.Category != expect {
//		t.Error(fmt.Sprintf("got %s, expected %s", got.Category, expect))
//	}
//}
