package balance

import (
	"errors"
	"jcb/db"
	"jcb/lib/transaction"
)

func Id(id int64) (int64, error) {
	switch transaction.Type(id) {
	case transaction.TYPE_OPENING:
		t, err := transaction.Find(id)
		return t.Cents, err

	case transaction.TYPE_COMMITTED:
		return db.TransactionBalance(id)

	case transaction.TYPE_UNCOMMITTED:
		b, _ := FinalCommitted()
		uc, _ := transaction.Uncommitted(-1)
		for _, t := range uc {
			b += t.Cents
		}

		return b, nil
	default:
		return 0, errors.New("Unknown transaction")
	}
}

func FinalCommitted() (int64, error) {
	committed, _ := transaction.Committed(-1)
	return db.TransactionBalance(committed[len(committed)-1].Id)
}

func Opening(year int) int64 {
	var b int64
	opening, _ := transaction.Find(0)
	committed, _ := transaction.Committed(year - 1)
	if len(committed) > 0 {
		b, _ = Id(committed[len(committed)-1].Id)
	} else if opening.Date.Year() == year {
		b = opening.Cents
	}

	return b
}
