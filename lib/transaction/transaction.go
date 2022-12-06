package transaction

import (
	"errors"
	"fmt"
	"jcb/db"
	"jcb/domain"
	"time"
)

type balance struct {
	Id      int64
	Cents   int64
	Balance int64
}

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

func Uncommitted(year int) ([]domain.Transaction, error) {
	return db.UncommittedTransactions(year)
}

func Committed(year int) ([]domain.Transaction, error) {
	return db.CommittedTransactions(year)
}

func Commit(id int64, initialBalance int64, year int) error {
	balance, err := commitSet(id, initialBalance, year)
	if err != nil {
		return err
	}

	for _, b := range balance {
		err = db.CommitTransaction(b.Id, b.Balance)
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
func commitSet(id int64, initialBalance int64, year int) ([]balance, error) {
	var found bool

	uncommitted, err := db.UncommittedTransactions(year)
	if err != nil {
		return []balance{}, err
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
		bset := balanceSet(tset, initialBalance)
		return bset, nil
	} else {
		return []balance{}, errors.New(fmt.Sprintf("No uncommitted transaction with id %d was found", id))
	}
}

func balanceSet(tset []domain.Transaction, initialBalance int64) []balance {
	bset := make([]balance, len(tset))
	bset[len(tset)-1].Balance = initialBalance

	for i := len(bset) - 1; i > 0; i-- {
		bset[i].Id = tset[i].Id
		bset[i].Cents = tset[i].Cents
		bset[i-1].Balance = bset[i].Balance - bset[i].Cents
	}
	bset[0].Id = tset[0].Id
	bset[0].Cents = tset[0].Cents

	return bset
}
