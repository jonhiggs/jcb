package db

import (
	"database/sql"
	"errors"
	"fmt"
	"jcb/domain"
	"log"
	"time"
)

const (
	TYPE_OPENING     = 0
	TYPE_COMMITTED   = 1
	TYPE_UNCOMMITTED = 2
)

const timeLayout = "2006-01-02 15:04:05.999999999-07:00"

func CommittedTransactions(year int) ([]domain.Transaction, error) {
	var rows *sql.Rows
	var err error

	var records []domain.Transaction

	if year == -1 {
		rows, err = db.Query("SELECT id, date, description, cents FROM transactions WHERE committedAt NOT NULL ORDER BY committedAt ASC", fmt.Sprintf("%d%%", year))
	} else {
		rows, err = db.Query("SELECT id, date, description, cents FROM transactions WHERE date LIKE ? AND committedAt NOT NULL ORDER BY committedAt ASC", fmt.Sprintf("%d%%", year))
	}

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var dateString string
		var description string
		var cents int64

		err = rows.Scan(&id, &dateString, &description, &cents)
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, domain.Transaction{id, parseDate(dateString), description, cents})
	}

	return records, nil
}

func UncommittedTransactions(year int) ([]domain.Transaction, error) {
	var rows *sql.Rows
	var err error

	var records []domain.Transaction

	if year == -1 {
		rows, err = db.Query("SELECT id, date, description, cents FROM transactions WHERE committedAt IS NULL ORDER BY date, description ASC", fmt.Sprintf("%d%%", year))
	} else {
		rows, err = db.Query("SELECT id, date, description, cents FROM transactions WHERE date LIKE ? AND committedAt IS NULL ORDER BY date, description ASC", fmt.Sprintf("%d%%", year))
	}

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var dateString string
		var description string
		var cents int64

		err = rows.Scan(&id, &dateString, &description, &cents)
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, domain.Transaction{id, parseDate(dateString), description, cents})
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

func TransactionCommittedUntil() (time.Time, error) {
	rows, err := db.Query("SELECT date FROM transactions WHERE committedAt NOT NULL ORDER BY date DESC LIMIT 1")
	if err != nil {
		return time.Unix(0, 0), err
	}

	defer rows.Close()
	for rows.Next() {
		var date string
		err = rows.Scan(&date)
		ts, _ := time.Parse(timeLayout, date)
		return ts, nil
	}

	return time.Unix(0, 0), nil
}

func TransactionIsCommitted(id int64) bool {
	var count int
	statement, _ := db.Prepare("SELECT COUNT(*) FROM transactions WHERE id == ? AND committedAt NOT NULL")
	statement.QueryRow(id).Scan(&count)

	switch count {
	case 0:
		return false
	case 1:
		return true
	}

	log.Fatal("There should never be more than two items sharing an id")
	return false
}
