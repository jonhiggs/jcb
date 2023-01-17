package dataFormatter

import (
	"fmt"
	"testing"
)

//func TestCents(t *testing.T) {
//	var got int64
//	var expect int64
//
//	got = Cents("0")
//	expect = 0
//	if got != expect {
//		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
//	}
//
//	got = Cents("30")
//	expect = 3000
//	if got != expect {
//		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
//	}
//
//	got = Cents("-10")
//	expect = -1000
//	if got != expect {
//		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
//	}
//
//	got = Cents("30.40")
//	expect = 3040
//	if got != expect {
//		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
//	}
//}

//func TestDate(t *testing.T) {
//	testA := Date("2022-04-30")
//
//	if testA.Format("2006-01-02") != "2022-04-30" {
//		t.Error("testA")
//	}
//}

//func TestDescription(t *testing.T) {
//	testA := Description("   testing    ")
//
//	if testA != "testing" {
//		t.Error("testA")
//	}
//}

func TestCategory(t *testing.T) {
	testA := Category("   testing    ")

	if testA != "testing" {
		t.Error("testA")
	}
}

//func TestNotes(t *testing.T) {
//	testA := Notes("   testing    ")
//
//	if testA != "testing" {
//		t.Error("testA")
//	}
//}

func TestId(t *testing.T) {
	var got int64
	var expect int64

	got = Id("042")
	expect = 42
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
}

func TestRepeatRule(t *testing.T) {
	var got string
	var expect string

	got = RepeatRule("0d")
	expect = "0d"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = RepeatRule("32w")
	expect = "32w"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = RepeatRule("2m")
	expect = "2m"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = RepeatRule("1y")
	expect = "1y"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

func TestRepeatRuleUnit(t *testing.T) {
	var got string
	var expect string

	got = RepeatRuleUnit("0d")
	expect = "d"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = RepeatRuleUnit("7w")
	expect = "w"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = RepeatRuleUnit("3m")
	expect = "m"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}

	got = RepeatRuleUnit("3y")
	expect = "y"
	if got != expect {
		t.Error(fmt.Sprintf("got %s, expected %s", got, expect))
	}
}

func TestRepeatRuleFrequency(t *testing.T) {
	var got int
	var expect int

	got = RepeatRuleFrequency("0d")
	expect = 0
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}

	got = RepeatRuleFrequency("7w")
	expect = 7
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}

	got = RepeatRuleFrequency("3m")
	expect = 3
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}

	got = RepeatRuleFrequency("1y")
	expect = 1
	if got != expect {
		t.Error(fmt.Sprintf("got %d, expected %d", got, expect))
	}
}
