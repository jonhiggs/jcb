package dataFormatter

import (
	"fmt"
	"jcb/domain"
	"testing"
)

func TestCents(t *testing.T) {
	var got int64
	var expect int64
	var err error

	got, err = Cents("0")
	expect = 0
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %d", expect))
	}

	got, err = Cents("30")
	expect = 3000
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %d", expect))
	}

	got, err = Cents("30.40")
	expect = 3040
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %d", expect))
	}

	got, err = Cents("0.2345")
	expect = 0
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
	if err == nil {
		t.Error(fmt.Sprintf("error expected for %d", expect))
	}

	got, err = Cents("0.23.45")
	expect = 0
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
	if err == nil {
		t.Error(fmt.Sprintf("error expected for %d", expect))
	}
}

func TestDate(t *testing.T) {
	testA, _ := Date("2022-04-30")

	if testA.Format("2006-01-02") != "2022-04-30" {
		t.Error("testA")
	}
}

func TestDescription(t *testing.T) {
	testA, _ := Description("   testing    ")

	if testA != "testing" {
		t.Error("testA")
	}
}

func TestId(t *testing.T) {
	var got int64
	var expect int64
	var err error

	got, err = Id("042")
	expect = 42
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %d", expect))
	}

	got, err = Id("-42")
	expect = 0
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
	if err == nil {
		t.Error(fmt.Sprintf("error expected for %d", -42))
	}
}

func TestTransaction(t *testing.T) {
	_, err := Transaction(domain.StringTransaction{"1", "2022-03-22", "testing", "12.00"})
	if err != nil {
		t.Error("no error expected")
	}
}
