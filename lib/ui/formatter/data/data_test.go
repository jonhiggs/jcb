package dataFormatter

import (
	"testing"
)

func TestCents(t *testing.T) {
	testA := Cents("0")

	if testA != 0 {
		t.Error("testA")
	}
}

func TestTime(t *testing.T) {
	testA := Date("2022-04-30")

	if testA.Format("2006-01-02") != "2022-04-30" {
		t.Error("testA")
	}
}

func TestDescription(t *testing.T) {
	testA := Description("   testing    ")

	if testA != "testing" {
		t.Error("testA")
	}
}

func TestId(t *testing.T) {
	testA := Id("042")

	if testA != 42 {
		t.Error("testA")
	}
}
