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

func Committed() ([]domain.Transaction, error) {
	return db.CommittedTransactions()
}

func Commit(id int64, balance int64) error {
	set, err := commitSet(id)

	for i := len(set) - 1; i != 0; i-- {
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

	//return set, errors.New(fmt.Sprintf("%d", len(set)))

	return set, nil
}

func Balance(id int64) (int64, error) {
	return db.TransactionBalance(id)
}
