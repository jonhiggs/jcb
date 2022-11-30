package transaction

import (
	"errors"
	"jcb/db"
	"jcb/domain"
	"strconv"
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

func Validate(t domain.Transaction) error {
	var err error
	err = validateDate(t)
	if err != nil {
		return err
	}

	err = validateDescription(t)
	if err != nil {
		return err
	}

	err = validateAmount(t)
	if err != nil {
		return err
	}

	return nil
}

func validateDate(t domain.Transaction) error {
	year, err := strconv.ParseInt(t.Date.Format("2006"), 10, 64)
	if err != nil {
		return errors.New("Unable to parse the year")
	}
	if year < 2000 {
		return errors.New("Invalid Year")
	}
	return nil
}

func validateDescription(t domain.Transaction) error {
	if len(t.Description) <= 0 {
		return errors.New("Description cannot be empty")
	}
	return nil
}

func validateAmount(t domain.Transaction) error {
	if t.Cents == 0 {
		return errors.New("Amount cannot be zero")
	}
	return nil
}
