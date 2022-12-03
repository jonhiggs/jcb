package repeater

import (
	"fmt"
	"testing"
	"time"
)

func TestNextDateDay(t *testing.T) {
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	res, _ := nextDate(date, "1d")
	got := res.Format("2006-01-02")
	if got != "2022-01-02" {
		t.Error(fmt.Sprintf("expected 2022-01-02 but got %s", got))
	}
}

func TestNextDateWeek(t *testing.T) {
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	res, _ := nextDate(date, "1w")
	got := res.Format("2006-01-02")
	if got != "2022-01-08" {
		t.Error(fmt.Sprintf("expected 2022-01-08 but got %s", got))
	}
}

func TestNextDateMonth(t *testing.T) {
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	res, _ := nextDate(date, "1m")
	got := res.Format("2006-01-02")
	if got != "2022-02-01" {
		t.Error(fmt.Sprintf("expected 2022-02-01 but got %s", got))
	}
}

func TestExpandSingle(t *testing.T) {
	var res []time.Time
	var err error
	var got string

	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	startDate := date
	endDate := time.Date(2022, 12, 31, 23, 59, 59, 59, time.UTC)

	res, err = Expand(date, startDate, endDate, "0d")
	got = res[0].Format("2006-01-02")
	if got != "2022-01-01" {
		t.Error(fmt.Sprintf("first element expected 2022-01-01 but was %s", got))
	}

	if len(res) != 1 {
		t.Error(fmt.Sprintf("expected 1 element, got %d.", len(res)))
	}

	if err != nil {
		t.Error("error was not expected")
	}
}

//func TestExpandOneDay(t *testing.T) {
//	var res []time.Time
//	var err error
//	var got string
//
//	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
//	startDate := date
//	endDate := time.Date(2022, 12, 31, 23, 59, 59, 59, time.UTC)
//
//	res, err = Expand(date, startDate, endDate, "1d")
//	got = res[0].Format("2006-01-02")
//	if got != "2022-01-01" {
//		t.Error(fmt.Sprintf("first element expected 2022-01-01 but was %s", got))
//	}
//
//	got = res[len(res)-1].Format("2006-01-02")
//	if got != "2022-12-31" {
//		t.Error(fmt.Sprintf("last element expected 2022-12-31 but was %s", got))
//	}
//
//	if len(res) != 365 {
//		t.Error(fmt.Sprintf("expected 365 elements, got %d.", len(res)))
//	}
//
//	if err != nil {
//		t.Error("error was not expected")
//	}
//}

//func TestExpandOneWeek(t *testing.T) {
//	var res []time.Time
//	var err error
//	var got string
//
//	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
//	startDate := date
//	endDate := time.Date(2022, 12, 31, 23, 59, 59, 59, time.UTC)
//
//	res, err = Expand(date, startDate, endDate, "1w")
//	got = res[0].Format("2006-01-01")
//	if got != "2022-01-01" {
//		t.Error(fmt.Sprintf("first element expected 2022-01-01 but was %s", got))
//	}
//
//	got = res[1].Format("2006-01-02")
//	if got != "2022-01-08" {
//		t.Error(fmt.Sprintf("second element expected 2022-01-08 but was %s", got))
//	}
//
//	got = res[len(res)-1].Format("2006-01-02")
//	if got != "2022-12-31" {
//		t.Error(fmt.Sprintf("last element expected 2022-12-31 but was %s", got))
//	}
//
//	if len(res) != 53 {
//		t.Error(fmt.Sprintf("expected 53 elements, got %d.", len(res)))
//	}
//
//	if err != nil {
//		t.Error("error was not expected")
//	}
//}
//
//func TestExpandOneMonth(t *testing.T) {
//	var res []time.Time
//	var err error
//	var got string
//
//	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
//	startDate := date
//	endDate := time.Date(2022, 12, 31, 23, 59, 59, 59, time.UTC)
//
//	res, err = Expand(date, startDate, endDate, "1m")
//	got = res[0].Format("2006-01-02")
//	if got != "2022-01-01" {
//		t.Error(fmt.Sprintf("first element expected 2022-01-01 but was %s", got))
//	}
//
//	got = res[1].Format("2006-01-02")
//	if got != "2022-02-01" {
//		t.Error(fmt.Sprintf("second element expected 2022-02-01 but was %s", got))
//	}
//
//	got = res[11].Format("2006-01-02")
//	if got != "2022-12-01" {
//		t.Error(fmt.Sprintf("last element expected 2022-12-01 but was %s", got))
//	}
//
//	if len(res) != 12 {
//		t.Error(fmt.Sprintf("expected 12 elements, got %d.", len(res)))
//	}
//
//	if err != nil {
//		t.Error("error was not expected")
//	}
//}
//
//func TestExpandOneMonth31(t *testing.T) {
//	var res []time.Time
//	var err error
//	var got string
//
//	date := time.Date(2022, 1, 31, 0, 0, 0, 0, time.UTC)
//	startDate := date
//	endDate := time.Date(2022, 12, 31, 23, 59, 59, 59, time.UTC)
//
//	res, err = Expand(date, startDate, endDate, "1m")
//	got = res[0].Format("2006-01-02")
//	if got != "2022-01-31" {
//		t.Error(fmt.Sprintf("first element expected 2022-01-31 but was %s", got))
//	}
//
//	got = res[1].Format("2006-01-02")
//	if got != "2022-03-03" {
//		t.Error(fmt.Sprintf("second element expected 2022-03-03 but was %s", got))
//	}
//
//	got = res[2].Format("2006-01-02")
//	if got != "2022-03-31" {
//		t.Error(fmt.Sprintf("second element expected 2022-03-31 but was %s", got))
//	}
//
//	if len(res) != 12 {
//		t.Error(fmt.Sprintf("expected 12 elements, got %d.", len(res)))
//	}
//
//	if err != nil {
//		t.Error("error was not expected")
//	}
//}
//
//func TestExpandShortEnd(t *testing.T) {
//	var res []time.Time
//	var err error
//	var got string
//
//	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
//	startDate := date
//	endDate := time.Date(2022, 1, 10, 23, 59, 59, 59, time.UTC)
//
//	res, err = Expand(date, startDate, endDate, "2d")
//	got = res[0].Format("2006-01-02")
//	if got != "2022-01-01" {
//		t.Error(fmt.Sprintf("first element expected 2022-01-01 but was %s", got))
//	}
//
//	got = res[1].Format("2006-01-02")
//	if got != "2022-01-03" {
//		t.Error(fmt.Sprintf("second element expected 2022-01-03 but was %s", got))
//	}
//
//	if len(res) != 5 {
//		t.Error(fmt.Sprintf("expected 5 elements, got %d.", len(res)))
//	}
//
//	if err != nil {
//		t.Error("error was not expected")
//	}
//}
//
//func TestExpandNarrowDay(t *testing.T) {
//	var res []time.Time
//	var err error
//	var got string
//
//	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
//	startDate := time.Date(2022, 1, 5, 0, 0, 0, 0, time.UTC)
//	endDate := time.Date(2022, 1, 10, 23, 59, 59, 59, time.UTC)
//
//	res, err = Expand(date, startDate, endDate, "1d")
//	got = res[0].Format("2006-01-02")
//	if got != "2022-01-05" {
//		t.Error(fmt.Sprintf("first element expected 2022-01-05 but was %s", got))
//	}
//
//	if len(res) != 6 {
//		t.Error(fmt.Sprintf("expected 5 elements, got %d.", len(res)))
//	}
//
//	if err != nil {
//		t.Error("error was not expected")
//	}
//}
//
//func TestExpandNarrowWeek(t *testing.T) {
//	var res []time.Time
//	var err error
//	var got string
//
//	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
//	startDate := time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)
//	endDate := time.Date(2022, 4, 30, 23, 59, 59, 59, time.UTC)
//
//	res, err = Expand(date, startDate, endDate, "1w")
//	got = res[0].Format("2006-01-02")
//	if got != "2022-04-02" {
//		t.Error(fmt.Sprintf("first element expected 2022-04-02 but was %s", got))
//	}
//
//	got = res[len(res)-1].Format("2006-01-02")
//	if got != "2022-04-30" {
//		t.Error(fmt.Sprintf("last element expected 2022-04-30 but was %s", got))
//	}
//
//	if len(res) != 5 {
//		t.Error(fmt.Sprintf("expected 6 elements, got %d.", len(res)))
//	}
//
//	if err != nil {
//		t.Error("error was not expected")
//	}
//}
//
//func TestExpandNarrowMonth(t *testing.T) {
//	var res []time.Time
//	var err error
//	var got string
//
//	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
//	startDate := time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)
//	endDate := time.Date(2022, 9, 30, 23, 59, 59, 59, time.UTC)
//
//	res, err = Expand(date, startDate, endDate, "1m")
//	got = res[0].Format("2006-01-02")
//	if got != "2022-04-01" {
//		t.Error(fmt.Sprintf("first element expected 2022-04-01 but was %s", got))
//	}
//
//	got = res[len(res)-1].Format("2006-01-02")
//	if got != "2022-09-01" {
//		t.Error(fmt.Sprintf("last element expected 2022-09-01 but was %s", got))
//	}
//
//	if len(res) != 6 {
//		t.Error(fmt.Sprintf("expected 6 elements, got %d.", len(res)))
//	}
//
//	if err != nil {
//		t.Error("error was not expected")
//	}
//}

//func TestExpandOutOfRange(t *testing.T) {
//	var res []time.Time
//	var err error
//
//	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
//	startDate := time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)
//	endDate := time.Date(2022, 9, 30, 23, 59, 59, 59, time.UTC)
//
//	res, err = Expand(date, startDate, endDate, "0d")
//	if len(res) != 0 {
//		t.Error(fmt.Sprintf("expected 0 elements, got %d.", len(res)))
//	}
//
//	if err != nil {
//		t.Error("error was not expected")
//	}
//}
