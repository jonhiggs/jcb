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

func Balance(id int64) (int64, error) {
	return db.TransactionBalance(id)
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
		errStr := fmt.Sprintln(bset)
		return bset, errors.New(errStr)
		return bset, nil
	} else {
		return []balance{}, errors.New(fmt.Sprintf("No uncommitted transaction with id %d was found", id))
	}
}

func balanceSet(tset []domain.Transaction, initialBalance int64) []balance {
	bset := make([]balance, len(tset))
	bset[len(tset)-1].Balance = initialBalance
	bset[len(tset)-1].Cents = tset[len(tset)-1].Cents
	bset[len(tset)-1].Id = tset[len(tset)-1].Id

	for i := len(bset) - 1; i > 0; i-- {
		bset[i-1].Id = tset[i-1].Id
		bset[i-1].Cents = tset[i-1].Cents
		bset[i-1].Balance = bset[i].Cents + tset[i].Cents
	}

	//startingBalance := finalBalance
	//for _, t := range tset {
	//	startingBalance += t.Cents
	//}
	//for i, t := range tset {
	//	if i == 0 {
	//		//bset[i] = startingBalance
	//		log.Printf("[%d]: %d", i, bset[i])
	//	} else if i == len(tset) {
	//		i[len(tset)] = finalBalance
	//	} else {
	//		bset[i-1] = bset[i-1] - t.Cents
	//		log.Printf("[%d]: %d - %d = %d", i-1, bset[i-1], t.Cents, bset[i])
	//	}
	//}

	//bset[len(tset)-1] = finalBalance
	////log.Printf("setting balance to %d", bset[len(tset)-1])
	//for i := len(tset) - 2; i >= 0; i-- {
	//	bset[i] = bset[i+1] + tset[i].Cents
	//	//log.Printf("[%d]: %d + %d = %d", i, bset[i+1], tset[i].Cents, bset[i])
	//}

	return bset
}
