package transaction

import (
	"errors"
	"fmt"
	"jcb/db"
	"jcb/domain"
	"time"
)

func Find(id int64) (domain.Transaction, error) {
	return db.FindTransaction(id)
}

func Insert(t domain.Transaction) (int64, error) {
	return db.InsertTransaction(t)
}

func Edit(t domain.Transaction) error {
	return db.EditTransaction(t)
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
	tset, bset, err := commitSet(id, balance)
	if err != nil {
		return err
	}

	for i, t := range tset {
		err = db.CommitTransaction(t.Id, bset[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func Uncommit(id int64) error {
	return db.UncommitTransaction(id)
}

func CommittedUntil() (time.Time, error) {
	return db.TransactionCommittedUntil()
}

func Balance(id int64) (int64, error) {
	return db.TransactionBalance(id)
}

// set of transactions that need to be committed before committing provided id
func commitSet(id int64, balance int64) ([]domain.Transaction, []int64, error) {
	var found bool

	uncommitted, err := db.UncommittedTransactions()
	if err != nil {
		return []domain.Transaction{}, []int64{}, err
	}

	var tset []domain.Transaction
	for _, t := range uncommitted {
		tset = append(tset, t)

		if t.Id == id {
			found = true
			break
		}
	}

	if found {
		return tset, balanceSet(tset, balance), nil
	} else {
		return []domain.Transaction{}, []int64{}, errors.New(fmt.Sprintf("No uncommitted transaction with id %d was found", id))
	}
}

func balanceSet(tset []domain.Transaction, finalBalance int64) []int64 {
	bset := make([]int64, len(tset))
	for i := len(tset) - 1; i >= 0; i-- {
		if i == len(tset)-1 {
			bset[i] = finalBalance
		} else {
			bset[i] = bset[i+1] - tset[i].Cents
		}
	}

	return bset
}
