package transaction

import (
	"errors"
	"jcb/db"
	"jcb/domain"
	"jcb/lib/dates"
	stringf "jcb/lib/formatter/string"
)

func Find(id int64) (domain.Transaction, error) {
	return db.FindTransaction(id)
}

func Insert(t domain.Transaction) (int64, error) {
	if t.Date.Unix() < dates.LastCommitted().Unix() {
		return -1, errors.New("Cannot insert before a committed transation")
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

	if Attributes(id).Committed {
		return errors.New("Cannot delete committed transactions")
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
	ids := commitSet(id)

	for _, i := range ids {
		t, err := Find(i)
		if err != nil {
			return err
		}

		err = db.CommitTransaction(i, t.Cents)
		if err != nil {
			return err
		}
	}
	return nil
}

func CommitSingle(id int64) error {
	t, err := Find(id)
	if err != nil {
		return err
	}

	ut, err := Uncommitted()
	if err != nil {
		return err
	}

	found := false
	for i := len(ut) - 1; i >= 0; i-- {
		if ut[i].Id == id {
			found = true
		}

		if !found {
			continue
		}

		if stringf.Date(ut[i].Date) != stringf.Date(t.Date) {
			return errors.New("Commit older transactions first")
		}
	}

	return db.CommitTransaction(id, t.Cents)
}

func UncommitSingle(id int64) error {
	ct, _ := Committed()
	if ct[len(ct)-1].Id != id {
		return errors.New("Only the final transaction can be uncommitted")
	}

	return db.UncommitTransaction(id)
}

func Uncommit(id int64) error {
	return db.UncommitTransaction(id)
}

func Notes(id int64) string {
	return db.TransactionNotes(id)
}

func Uniq(t domain.Transaction) bool {
	return db.TransactionUniq(t)
}

func Attributes(id int64) domain.Attributes {
	return db.TransactionAttributes(id)
}

// set of transactions that need to be committed before committing provided id
func commitSet(id int64) []int64 {
	uncommitted, err := db.UncommittedTransactions()
	if err != nil {
		return []int64{}
	}

	var ids []int64
	for _, t := range uncommitted {
		ids = append(ids, t.Id)

		if t.Id == id {
			break
		}
	}

	return ids
}
