package ui

import (
	"fmt"
	"jcb/domain"
	"strconv"
	"strings"
	"time"
)

type FormattedTransaction struct {
	Id          string
	Date        string
	Description string
	Amount      string
}

func formatTransaction(t domain.Transaction) FormattedTransaction {
	id := fmt.Sprintf("%s", t.Id)
	date := t.Date.Format("2006-01-02")
	description := t.Description
	amount := formatCents(t.Cents)

	return FormattedTransaction{id, date, description, amount}
}

func unformatTransaction(t FormattedTransaction) domain.Transaction {
	id, _ := strconv.ParseInt(t.Id, 10, 64)
	date, _ := time.Parse("2006-01-02", t.Date)
	description := t.Description
	cents, _ := strconv.ParseInt(strings.Replace(t.Amount, ".", "", 1), 10, 64)

	return domain.Transaction{id, date, description, cents}
}

func formatCents(i int64) string {
	s := fmt.Sprintf("%d", i)
	var d string
	var c string
	if len(s) == 2 {
		d = "0"
		c = s
	} else if len(s) == 1 {
		d = "0"
		c = fmt.Sprintf("0%s", s)
	} else {
		d = s[0 : len(s)-2]
		c = s[len(s)-2 : len(s)]
	}
	return fmt.Sprintf("%s.%s", d, c)
}
