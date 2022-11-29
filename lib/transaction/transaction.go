package transaction

import (
	"errors"
	"jcb/db"
	"jcb/domain"
)

func Save(t domain.Transaction) error {
	return db.SaveTransaction(t)
}

func Delete(t domain.Transaction) error {
	return db.DeleteTransaction(t.Id)
}

func All() ([]domain.Transaction, error) {
	return db.AllTransactions()
}

func Validate(t domain.Transaction) (domain.Transaction, error) {
	return t, errors.New("It always fails")
	//return t, nil
}
