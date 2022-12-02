package transaction

import (
	"errors"
	"fmt"
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

func Committed() ([]domain.Transaction, error) {
	return db.CommittedTransactions()
}

func Commit(id int64, balance int64) error {
	set, err := commitSet(id)
	if err != nil {
		return err
	}

	for i := len(set) - 1; i >= 0; i-- {
		err = db.CommitTransaction(set[i].Id, balance)
		if err != nil {
			return err
		}

		balance -= set[i].Cents
	}
	return nil
}

// set of transactions that need to be committed before committing provided id
func commitSet(id int64) ([]domain.Transaction, error) {
	var set []domain.Transaction
	var found bool
	uncommitted, err := db.UncommittedTransactions()
	if err != nil {
		return set, err
	}
	for _, t := range uncommitted {
		set = append(set, t)

		if t.Id == id {
			found = true
			break
		}
	}

	if found {
		return set, nil
	} else {
		return []domain.Transaction{}, errors.New(fmt.Sprintf("No uncommitted transaction with id %d was found", id))
	}
}

func Balance(id int64) (int64, error) {
	return db.TransactionBalance(id)
}
