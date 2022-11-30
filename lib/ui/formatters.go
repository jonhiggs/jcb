package ui

import (
	"jcb/domain"
	sform "jcb/lib/ui/formatter/string"
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

// convert a struct of data to a struct of strings
func formatTransaction(t domain.Transaction) FormattedTransaction {
	id := sform.Id(t.Id)
	date := sform.Date(t.Date)
	description := t.Description
	amount := sform.Cents(t.Cents)

	return FormattedTransaction{id, date, description, amount}
}

// convert a struct of strings to real data
func unformatTransaction(t FormattedTransaction) domain.Transaction {
	id, _ := strconv.ParseInt(t.Id, 10, 64)
	date, _ := time.Parse("2006-01-02", strings.Trim(t.Date, " "))
	description := strings.Trim(t.Description, " ")
	cents, _ := strconv.ParseInt(strings.Replace(strings.Trim(t.Amount, " "), ".", "", 1), 10, 64)

	return domain.Transaction{id, date, description, cents}
}
