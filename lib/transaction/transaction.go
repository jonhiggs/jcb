package transaction

import (
	"jcb/db"
	"jcb/domain"
)

func Save(t domain.Transaction) error {
	return db.SaveTransaction(t)
}

func Delete(t domain.Transaction) error {
	return db.DeleteTransaction(t.Id)
}
