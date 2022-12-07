package balance

import (
	"errors"
	"jcb/db"
	"jcb/lib/transaction"
)

func Id(id int64) (int64, error) {
	switch transaction.Type() {
	case transaction.TYPE_OPENING:
		return transaction.Find(id).Cents

	case transaction.TYPE_COMMITTED:
		return db.TransactionBalance(id)

	case transaction.TYPE_UNCOMMITTED:
		balance, _ := FinalCommitted()
		uncommitted, _ := transaction.Uncommitted()
		for _, t := range uncommitted {
			balance += t.Amount()
		}

		return balance, nil
	default:
		return 0, errors.New("Unknown transaction")
	}
}

func FinalCommitted() (int64, error) {
	return db.TransactionBalance(committed[len(committed)-1].Id)
}
