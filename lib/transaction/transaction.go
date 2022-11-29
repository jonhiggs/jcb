package transaction

import (
	"errors"
	"jcb/db"
	"jcb/domain"
	"strconv"
)

func Save(t domain.Transaction) error {
	return db.SaveTransaction(t)
}

func Delete(id int64) error {
	return db.DeleteTransaction(id)
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

func validateAmount(t domain.Transaction) error {
	if t.Cents == 0 {
		return errors.New("Amount cannot be zero")
	}
	return nil
}
