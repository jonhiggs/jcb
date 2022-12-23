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

	rows, err = db.Query("SELECT id, date, description, cents, notes FROM transactions WHERE committedAt NOT NULL ORDER BY committedAt ASC", "")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var dateString string
		var description string
		var cents int64
		var notes string

		err = rows.Scan(&id, &dateString, &description, &cents, &notes)
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, domain.Transaction{id, parseDate(dateString), description, cents, notes})
	}

	return records, nil
}

func UncommittedTransactions() ([]domain.Transaction, error) {
	var rows *sql.Rows
	var err error

	var records []domain.Transaction

	rows, err = db.Query("SELECT id, date, description, cents, notes FROM transactions WHERE committedAt IS NULL ORDER BY date, description ASC", "")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var dateString string
		var description string
		var cents int64
		var notes string

		err = rows.Scan(&id, &dateString, &description, &cents, &notes)
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, domain.Transaction{id, parseDate(dateString), description, cents, notes})
	}

	return records, nil
}

func InsertTransaction(t domain.Transaction) (int64, error) {
	if t.Id != -1 {
		return -1, errors.New(fmt.Sprintf("Transaction %d already exists", t.Id))
	}

	statement, err := db.Prepare("INSERT INTO transactions (date, description, cents, notes, updatedAt) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(t.Date, t.Description, t.Cents, t.Notes, timeNow())
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func EditTransaction(t domain.Transaction) error {
	statement, err := db.Prepare("UPDATE transactions SET date = ?, description = ?, cents = ?, notes = ?, updatedAt = ? WHERE id = ? AND committedAt IS NULL")
	if err != nil {
		return err
	}

	_, err = statement.Exec(t.Date, t.Description, t.Cents, t.Notes, timeNow(), t.Id)
	return err
}

func CommitTransaction(id int64, cents int64) error {
	balance := CommittedBalance() + cents
	statement, _ := db.Prepare("UPDATE transactions SET balance = ?, committedAt = ?, updatedAt = ? WHERE id = ? AND committedAt IS NULL")

	_, err := statement.Exec(balance, timeNow(), timeNow(), id)
	return err
}

func CommittedBalance() int64 {
	ct, _ := CommittedTransactions()
	if len(ct) == 0 {
		t, _ := FindTransaction(0)
		return t.Cents
	} else {
		var balance int64
		statement, _ := db.Prepare("SELECT balance FROM transactions WHERE committedAt IS NOT NULL ORDER BY committedAt LIMIT 1")
		err := statement.QueryRow().Scan(&balance)
		if err != nil {
			log.Fatal(err)
		}
		return balance
	}
}

func UncommitTransaction(id int64) error {
	var committedAt string
	statement, _ := db.Prepare("SELECT committedAt FROM transactions WHERE id = ?")
	err := statement.QueryRow(id).Scan(&committedAt)
	ts, _ := time.Parse(timeLayout, committedAt)

	statement, _ = db.Prepare("UPDATE transactions SET committedAt = NULL, updatedAt = ?, balance = NULL WHERE committedAt >= ?")
	_, err = statement.Exec(timeNow(), ts)
	return err
}

func FindTransaction(id int64) (domain.Transaction, error) {
	var date string
	var description string
	var cents int64
	var notes string

	statement, _ := db.Prepare("SELECT id, date, description, cents, notes FROM transactions WHERE id = ?")
	err := statement.QueryRow(id).Scan(&id, &date, &description, &cents, &notes)
	ts, _ := time.Parse(timeLayout, date)

	return domain.Transaction{id, ts, description, cents, notes}, err
}

func DeleteTransaction(id int64) error {
	statement, err := db.Prepare("DELETE FROM transactions WHERE id = ? AND committedAt IS NULL")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)

	dirty = true

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

func TransactionNotes(id int64) string {
	var notes string
	statement, _ := db.Prepare("SELECT notes FROM transactions WHERE id == ?")
	err := statement.QueryRow(id).Scan(&notes)
	if err != nil {
		log.Fatal(err)
	}

	return notes
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

func TransactionAttributes(id int64) domain.Attributes {
	var committedAt string
	var updatedAt string
	var notes string

	statement, _ := db.Prepare("SELECT IFNULL(committedAt,''), updatedAt, notes FROM transactions WHERE id = ?")
	err := statement.QueryRow(id).Scan(&committedAt, &updatedAt, &notes)
	if err != nil {
		log.Fatal(err)
	}

	return domain.Attributes{
		committedAt != "",
		notes != "",
		parseDate(updatedAt).UnixMicro() < SaveTime.UnixMicro(),
	}
}

func timeNow() string {
	return time.Now().Format(timeLayout)
}
