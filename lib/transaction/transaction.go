package transaction

import (
	"jcb/db"
	"jcb/domain"
)

func Find(id int64) (domain.Transaction, error) {
	return db.FindTransaction(id)
}

func Insert(t domain.Transaction) (int64, error) {
	return db.InsertTransaction(t)
}

func Delete(t domain.Transaction) error {
	return db.DeleteTransaction(t.Id)
}

func DeleteId(id int64) error {
	return db.DeleteTransaction(id)
}

func Uncommitted() ([]domain.Transaction, error) {
	return db.UncommittedTransactions()
}

// set of transactions that need to be committed before committing provided id
func CommitSet(id int64, balance int64) ([]domain.Transaction, error) {
	var set []domain.Transaction
	uncommitted, err := db.UncommittedTransactions()
	if err != nil {
		return set, err
	}
	for _, t := range uncommitted {
		set = append(set, t)

		if t.Id == id {
			break
		}

	}

	return set, nil
}
