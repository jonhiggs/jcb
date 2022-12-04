package db

import (
	"errors"
	"fmt"
	"jcb/domain"
	"time"
)

const timeLayout = "2006-01-02 15:04:05.999999999-07:00"

func CommittedTransactions() ([]domain.Transaction, error) {
	rows, err := db.Query("SELECT id, date, description, cents FROM transactions WHERE committedAt NOT NULL ORDER BY committedAt ASC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var records []domain.Transaction
	for rows.Next() {
		t := domain.Transaction{}
		var date string

		err = rows.Scan(&t.Id, &date, &t.Description, &t.Cents)

		ts, _ := time.Parse(timeLayout, date)
		t.Date = ts

		if err != nil {
			return nil, err
		}

		records = append(records, t)
	}

	return records, nil
}

func UncommittedTransactions() ([]domain.Transaction, error) {
	rows, err := db.Query("SELECT id, date, description, cents FROM transactions WHERE committedAt IS NULL ORDER BY date, id ASC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var records []domain.Transaction
	for rows.Next() {
		t := domain.Transaction{}
		var date string

		err = rows.Scan(&t.Id, &date, &t.Description, &t.Cents)

		ts, err := time.Parse(timeLayout, date)
		if err != nil {
			return nil, err
		}

		t.Date = ts

		if err != nil {
			return nil, err
		}

		records = append(records, t)
	}

	return records, nil
}

func InsertTransaction(t domain.Transaction) (int64, error) {
	if t.Id == -1 {
		statement, err := db.Prepare("INSERT INTO transactions (date, description, cents) VALUES (?, ?, ?)")
		if err != nil {
			return -1, err
		}

		res, err := statement.Exec(t.Date, t.Description, t.Cents)
		if err != nil {
			return -1, err
		}
		return res.LastInsertId()
	}
	return -1, errors.New(fmt.Sprintf("Transaction %d already exists", t.Id))
}

func EditTransaction(t domain.Transaction) error {
	statement, err := db.Prepare("UPDATE transactions SET date = ?, description = ?, cents = ? WHERE id = ? AND committedAt IS NULL")
	if err != nil {
		return err
	}

	_, err = statement.Exec(t.Date, t.Description, t.Cents, t.Id)
	return err
}

func CommitTransaction(id int64, balance int64) error {
	statement, _ := db.Prepare("UPDATE transactions SET balance = ?, committedAt = ? WHERE id = ? AND committedAt IS NULL")
	_, err := statement.Exec(balance, time.Now().Format(timeLayout), id)
	return err
}

func UncommitTransaction(id int64) error {
	var committedAt string
	statement, _ := db.Prepare("SELECT committedAt FROM transactions WHERE id = ?")
	err := statement.QueryRow(id).Scan(&committedAt)
	ts, _ := time.Parse(timeLayout, committedAt)

	statement, _ = db.Prepare("UPDATE transactions SET committedAt = NULL, balance = NULL WHERE committedAt >= ?")
	_, err = statement.Exec(ts)
	return err
}

func FindTransaction(id int64) (domain.Transaction, error) {
	var date string
	var description string
	var cents int64

	statement, _ := db.Prepare("SELECT id, date, description, cents FROM transactions WHERE id = ?")
	err := statement.QueryRow(id).Scan(&id, &date, &description, &cents)
	ts, _ := time.Parse(timeLayout, date)

	return domain.Transaction{id, ts, description, cents}, err
}

func DeleteTransaction(id int64) error {
	statement, err := db.Prepare("DELETE FROM transactions WHERE id = ? AND committedAt IS NULL")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	return err
}

func TransactionBalance(id int64) (int64, error) {
	var balance int64

	statement, _ := db.Prepare("SELECT balance FROM transactions WHERE id = ? AND committedAt NOT NULL")
	err := statement.QueryRow(id).Scan(&balance)
	if err != nil {
		return -1, err
	}

	return balance, err
}

func TransactionCommittedUntil() (time.Time, error) {
	var date string
	statement, _ := db.Prepare("SELECT date FROM transactions WHERE committedAt NOT NULL ORDER BY date DESC LIMIT 1")
	err := statement.QueryRow().Scan(&date)
	if err != nil {
		return time.Unix(0, 0), err
	}

	ts, _ := time.Parse(timeLayout, date)

	return ts, err
}
