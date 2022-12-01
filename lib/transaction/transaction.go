package transaction

import (
	"jcb/db"
	"jcb/domain"
)

func Find(id int64) (domain.Transaction, error) {
	return db.FindTransaction(id)
}

func Save(t domain.Transaction) (int64, error) {
	return db.SaveTransaction(t)
}

func Delete(t domain.Transaction) error {
	return db.DeleteTransaction(t.Id)
}

func All() ([]domain.Transaction, error) {
	return db.AllTransactions()
}
