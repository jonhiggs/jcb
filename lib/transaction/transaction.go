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
	if t.Id == 0 {
		tOriginal, _ := Find(0)
		u, _ := Uncommitted(tOriginal.Date.Year())
		// using u[1] because u[0] is the opening balance.
		if len(u) > 1 {
			if t.Date.Unix() > u[1].Date.Unix() {
				return errors.New(fmt.Sprintf("Opening balance must be before %s", u[1].Date.Format("2006-01-02")))
			}
		}
	}
	return db.EditTransaction(t)
}

func Delete(t domain.Transaction) error {
	if t.Id == 0 {
		return errors.New("You cannot delete the opening balance")
	}
	return db.DeleteTransaction(t.Id)
}

func DeleteId(id int64) error {
	if id == 0 {
		return errors.New("You cannot delete the opening balance")
	}
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

func ClosingBalance(year int) (int64, error) {
	ey, err := EarliestYear()
	if err != nil {
		return -1, err
	}

	if year < ey {
		return -1, errors.New(fmt.Sprintf("Cannot determine closing balance of %d", year))
	}

	committed, err := Committed(year)
	if err != nil {
		return -1, err
	}

	uncommitted, err := Uncommitted(year)
	if err != nil {
		return -1, err
	}

	initialBalance := committed[len(committed)-1].Id
	bset := balanceSet(uncommitted, initialBalance)
	return bset[len(bset)-1].Balance, nil
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

func EarliestYear() (int, error) {
	return db.EarliestYear()
}

func LatestYear() (int, error) {
	return db.LatestYear()
}
