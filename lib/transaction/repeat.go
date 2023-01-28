package transaction

import (
	"strconv"
	"time"
)

// take a transaction and a some repeat parameters, then return an expanded
// list of repeated transactions.
func (t *Transaction) Expand(rule string, multiplier int, end time.Time) ([]*Transaction, error) {
	var transactions []*Transaction

	// TODO: validate the rule

	f, _ := strconv.Atoi(rule[0 : len(rule)-1])
	u := rule[len(rule)-1 : len(rule)]

	i := 0
	for f > 0 {
		var d time.Time
		i++
		newT := *t

		switch u {
		case "d":
			d = t.Date.GetValue().AddDate(0, 0, (multiplier * i * f))
		case "w":
			d = t.Date.GetValue().AddDate(0, 0, (multiplier * i * (f * 7)))
		case "m":
			d = t.Date.GetValue().AddDate(0, (multiplier * i * f), 0)
		case "y":
			d = t.Date.GetValue().AddDate((multiplier * i * f), 0, 0)
		}

		if d.After(end) {
			break
		}

		newT.Date.SetValue(d)
		transactions = append(transactions, &newT)

	}

	return transactions, nil
}
