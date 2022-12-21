package db

import (
	"database/sql"
	"errors"
	"fmt"
	"jcb/domain"
	stringf "jcb/lib/formatter/string"
	"log"
	"time"
)

const (
	TYPE_OPENING     = 0
	TYPE_COMMITTED   = 1
	TYPE_UNCOMMITTED = 2
)

const timeLayout = "2006-01-02 15:04:05.999999999-07:00"

func CommittedTransactions() ([]domain.Transaction, error) {
	var rows *sql.Rows
	var err error

	var records []domain.Transaction

	rows, err = db.Query("SELECT id, date, description, cents FROM transactions WHERE committedAt NOT NULL ORDER BY committedAt ASC", "")

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

func UncommittedTransactions() ([]domain.Transaction, error) {
	var rows *sql.Rows
	var err error

	var records []domain.Transaction

	rows, err = db.Query("SELECT id, date, description, cents FROM transactions WHERE committedAt IS NULL ORDER BY date, description ASC", "")

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
	if t.Id != -1 {
		return -1, errors.New(fmt.Sprintf("Transaction %d already exists", t.Id))
	}

	statement, err := db.Prepare("INSERT INTO transactions (date, description, cents) VALUES (?, ?, ?)")
	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(t.Date, t.Description, t.Cents)
	if err != nil {
		return -1, err
	}
	Dirty = true
	return res.LastInsertId()
}

func EditTransaction(t domain.Transaction) error {
	statement, err := db.Prepare("UPDATE transactions SET date = ?, description = ?, cents = ? WHERE id = ? AND committedAt IS NULL")
	if err != nil {
		return err
	}

	_, err = statement.Exec(t.Date, t.Description, t.Cents, t.Id)
	Dirty = true
	return err
}

func CommitTransaction(id int64, balance int64) error {
	statement, _ := db.Prepare("UPDATE transactions SET balance = ?, committedAt = ? WHERE id = ? AND committedAt IS NULL")
	_, err := statement.Exec(balance, time.Now().Format(timeLayout), id)
	Dirty = true
	return err
}

func UncommitTransaction(id int64) error {
	var committedAt string
	statement, _ := db.Prepare("SELECT committedAt FROM transactions WHERE id = ?")
	err := statement.QueryRow(id).Scan(&committedAt)
	ts, _ := time.Parse(timeLayout, committedAt)

	statement, _ = db.Prepare("UPDATE transactions SET committedAt = NULL, balance = NULL WHERE committedAt >= ?")
	_, err = statement.Exec(ts)
	Dirty = true
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
	Dirty = true
	return err
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

func TransactionUniq(t domain.Transaction) bool {
	var count int
	statement, err := db.Prepare("SELECT COUNT(*) FROM transactions WHERE substr(date, 0,11) == ? AND description == ? AND cents == ?")
	if err != nil {
		log.Fatal(err)
	}
	err = statement.QueryRow(stringf.Date(t.Date), t.Description, t.Cents).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		return true
	} else {
		return false
	}
}
