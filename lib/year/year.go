package year

import (
	"jcb/db"
	"jcb/lib/balance"
)

func IsCommitted(year int) bool {
	t, _ := db.CommittedTransactions(year)
	return len(t) == 0
}

func Opening(year int) int64 {
	opening, _ := db.Find(0)
	if opening.Date.Year() == year {
		return opening.Cents
	} else {
		t, _ := transactions.Committed(year - 1)
		return balance.Id(t.Id)
	}
}
