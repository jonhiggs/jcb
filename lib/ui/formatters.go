package ui

import (
	"jcb/domain"
	dform "jcb/lib/ui/formatter/data"
	sform "jcb/lib/ui/formatter/string"
	"strconv"
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
	date := dform.Date(t.Date)
	description := dform.Description(t.Description)
	cents := dform.Cents(t.Amount)

	return domain.Transaction{id, date, description, cents}
}
