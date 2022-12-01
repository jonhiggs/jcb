package repeater

import (
	"fmt"
	"testing"
	"time"
)

func TestFrequency(t *testing.T) {
	var got int
	var expect int
	var err error

	got, err = frequency("1d")
	expect = 1
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %d", expect))
	}

	got, err = frequency("zzd")
	expect = 0
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
	if err == nil {
		t.Error(fmt.Sprintf("error expected for %d", expect))
	}

	got, err = frequency("1w")
	expect = 1
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %d", expect))
	}
}

func TestUnit(t *testing.T) {
	var got string
	var expect string
	var err error

	got, err = unit("1d")
	expect = "d"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %s", expect))
	}

	got, err = unit("1w")
	expect = "w"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %s", expect))
	}

	got, err = unit("1m")
	expect = "m"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
	if err != nil {
		t.Error(fmt.Sprintf("no error expected for %s", expect))
	}

	got, err = unit("1z")
	expect = "x"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
	if err == nil {
		t.Error(fmt.Sprintf("no error expected for %s", expect))
	}
}

func TestExpandOneDay(t *testing.T) {
	var res []time.Time
	var err error
	var got string

	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 12, 31, 23, 59, 59, 59, time.UTC)

	res, err = Expand(startDate, endDate, "1d")
	got = res[0].Format("2006-01-02")
	if got != "2022-01-01" {
		t.Error(fmt.Sprintf("first element expected 2022-01-01 but was %s", got))
	}

	got = res[len(res)-1].Format("2006-01-02")
	if got != "2022-12-31" {
		t.Error(fmt.Sprintf("last element expected 2022-12-31 but was %s", got))
	}

	if len(res) != 365 {
		t.Error(fmt.Sprintf("expected 365 elements, got %d.", len(res)))
	}

	if err != nil {
		t.Error("error was not expected")
	}
}

func TestExpandOneWeek(t *testing.T) {
	var res []time.Time
	var err error
	var got string

	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 12, 31, 23, 59, 59, 59, time.UTC)

	res, err = Expand(startDate, endDate, "1w")
	got = res[0].Format("2006-01-01")
	if got != "2022-01-01" {
		t.Error(fmt.Sprintf("first element expected 2022-01-01 but was %s", got))
	}

	got = res[1].Format("2006-01-02")
	if got != "2022-01-08" {
		t.Error(fmt.Sprintf("second element expected 2022-01-08 but was %s", got))
	}

	got = res[len(res)-1].Format("2006-01-02")
	if got != "2022-12-31" {
		t.Error(fmt.Sprintf("last element expected 2022-12-31 but was %s", got))
	}

	if len(res) != 53 {
		t.Error(fmt.Sprintf("expected 53 elements, got %d.", len(res)))
	}

	if err != nil {
		t.Error("error was not expected")
	}
}

func TestExpandOneMonth(t *testing.T) {
	var res []time.Time
	var err error
	var got string

	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 12, 31, 23, 59, 59, 59, time.UTC)

	res, err = Expand(startDate, endDate, "1m")
	got = res[0].Format("2006-01-02")
	if got != "2022-01-01" {
		t.Error(fmt.Sprintf("first element expected 2022-01-01 but was %s", got))
	}

	got = res[1].Format("2006-01-02")
	if got != "2022-02-01" {
		t.Error(fmt.Sprintf("second element expected 2022-02-01 but was %s", got))
	}

	got = res[11].Format("2006-01-02")
	if got != "2022-12-01" {
		t.Error(fmt.Sprintf("last element expected 2022-12-01 but was %s", got))
	}

	if len(res) != 12 {
		t.Error(fmt.Sprintf("expected 12 elements, got %d.", len(res)))
	}

	if err != nil {
		t.Error("error was not expected")
	}
}

func TestExpandOneMonth31(t *testing.T) {
	var res []time.Time
	var err error
	var got string

	startDate := time.Date(2022, 1, 31, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 12, 31, 23, 59, 59, 59, time.UTC)

	res, err = Expand(startDate, endDate, "1m")
	got = res[0].Format("2006-01-02")
	if got != "2022-01-31" {
		t.Error(fmt.Sprintf("first element expected 2022-01-31 but was %s", got))
	}

	got = res[1].Format("2006-01-02")
	if got != "2022-03-03" {
		t.Error(fmt.Sprintf("second element expected 2022-03-03 but was %s", got))
	}

	got = res[2].Format("2006-01-02")
	if got != "2022-03-31" {
		t.Error(fmt.Sprintf("second element expected 2022-03-31 but was %s", got))
	}

	if len(res) != 12 {
		t.Error(fmt.Sprintf("expected 12 elements, got %d.", len(res)))
	}

	if err != nil {
		t.Error("error was not expected")
	}
}

func TestExpandShortEnd(t *testing.T) {
	var res []time.Time
	var err error
	var got string

	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 1, 10, 23, 59, 59, 59, time.UTC)

	res, err = Expand(startDate, endDate, "2d")
	got = res[0].Format("2006-01-02")
	if got != "2022-01-01" {
		t.Error(fmt.Sprintf("first element expected 2022-01-01 but was %s", got))
	}

	got = res[1].Format("2006-01-02")
	if got != "2022-01-03" {
		t.Error(fmt.Sprintf("second element expected 2022-01-03 but was %s", got))
	}

	if len(res) != 5 {
		t.Error(fmt.Sprintf("expected 5 elements, got %d.", len(res)))
	}

	if err != nil {
		t.Error("error was not expected")
	}
}
