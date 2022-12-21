package transaction

import (
	"errors"
	"fmt"
	"jcb/db"
	"jcb/domain"
	"jcb/lib/dates"
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
	if t.Date.Unix() < dates.LastCommitted().Unix() {
		return -1, errors.New("Date is too early")
	}
	return db.InsertTransaction(t)
}

func Edit(t domain.Transaction) error {
	if t.Id == 0 && t.Date.Unix() > dates.FirstUncommitted().Unix() {
		return errors.New("Date of opening balance must be before the first transaction")
	}
	if t.Date.Unix() < dates.LastCommitted().Unix() {
		return errors.New("Date must be after the latest committed transaction")
	}
	return db.EditTransaction(t)
}

func Delete(id int64) error {
	if id == 0 {
		return errors.New("You cannot delete the opening balance")
	}
	return db.DeleteTransaction(id)
}

func DeleteId(id int64) error {
	if id == 0 {
		return errors.New("You cannot delete the opening balance")
	}
	return db.DeleteTransaction(id)
}

func Uncommitted() ([]domain.Transaction, error) {
	return db.UncommittedTransactions()
}

func Committed() ([]domain.Transaction, error) {
	return db.CommittedTransactions()
}

func Commit(id int64, initialBalance int64) error {
	balance, err := commitSet(id, initialBalance)
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

func IsCommitted(id int64) bool {
	return db.TransactionIsCommitted(id)
}

// set of transactions that need to be committed before committing provided id
func commitSet(id int64, initialBalance int64) ([]balance, error) {
	var found bool

	uncommitted, err := db.UncommittedTransactions()
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
