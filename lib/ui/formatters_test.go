package ui

import (
	"fmt"
	"jcb/domain"
	"testing"
	"time"
)

func TestFormatTransactionA(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2002-01-02")
	trans := domain.Transaction{0, date, "test new", 100}
	ftrans := formatTransaction(trans)

	if ftrans.Id != "0" {
		t.Error("incorrect id")
	}

	if ftrans.Date != "2002-01-02" {
		t.Error("incorrect date")
	}

	if ftrans.Description != "test new" {
		t.Error("incorrect description")
	}

	if ftrans.Amount != "1.00" {
		t.Error("incorrect amount")
	}
}

func TestFormatTransactionB(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2002-01-02")
	trans := domain.Transaction{0, date, "test new", -100}
	ftrans := formatTransaction(trans)

	if ftrans.Id != "0" {
		t.Error("incorrect id")
	}

	if ftrans.Date != "2002-01-02" {
		t.Error("incorrect date")
	}

	if ftrans.Description != "test new" {
		t.Error("incorrect description")
	}

	if ftrans.Amount != "-1.00" {
		t.Error("incorrect amount")
	}
}

func TestUnformatTransactionA(t *testing.T) {
	ftrans := FormattedTransaction{"0", "2002-03-04", "test new", "1.00"}
	trans := unformatTransaction(ftrans)

	if trans.Id != 0 {
		t.Error("incorrect id")
	}

	if trans.Date.Format("2006-01-02") != "2002-03-04" {
		t.Error("incorrect date")
	}

	if trans.Description != "test new" {
		t.Error("incorrect description")
	}

	if trans.Cents != 100 {
		t.Error("incorrect cents")
	}
}

func TestUnformatTransactionB(t *testing.T) {
	ftrans := FormattedTransaction{"  0  ", "  2002-03-04  ", "  test new  ", "  0.2   "}
	trans := unformatTransaction(ftrans)

	if trans.Id != 0 {
		t.Error("incorrect id")
	}

	if trans.Date.Format("2006-01-02") != "2002-03-04" {
		t.Error("incorrect date")
	}

	if trans.Description != "test new" {
		t.Error("incorrect description")
	}

	if trans.Cents != 2 {
		t.Error(fmt.Sprintf("incorrect cents [%d]", trans.Cents))
	}
}
