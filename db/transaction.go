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

const timeLayout = "2006-01-02 15:04:05.999999999-07:00"

func CommittedTransactions() []domain.Transaction {
	var rows *sql.Rows
	var err error

	var records []domain.Transaction

	rows, err = db.Query("SELECT id, date, description, cents, notes, category FROM transactions WHERE committedAt NOT NULL ORDER BY committedAt ASC", "")

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
		var category string

		err = rows.Scan(&id, &dateString, &description, &cents, &notes, &category)
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, domain.Transaction{id, parseDate(dateString), description, cents, notes, category})
	}

	return records
}

func UncommittedTransactions() []domain.Transaction {
	var rows *sql.Rows
	var err error

	var records []domain.Transaction

	rows, err = db.Query("SELECT id, date, description, cents, notes, category FROM transactions WHERE committedAt IS NULL ORDER BY date ASC, cents DESC", "")

	if err != nil {
		log.Fatal(fmt.Sprintf("UncommitTransaction(): %s", err))
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var dateString string
		var description string
		var cents int64
		var notes string
		var category string

		err = rows.Scan(&id, &dateString, &description, &cents, &notes, &category)
		if err != nil {
			log.Fatal(fmt.Sprintf("UncommitTransaction(): %s", err))
		}

		records = append(records, domain.Transaction{id, parseDate(dateString), description, cents, notes, category})
	}

	return records
}

func InsertTransaction(t domain.Transaction) (int64, error) {
	if t.Id != -1 {
		return -1, errors.New(fmt.Sprintf("Transaction %d already exists", t.Id))
	}

	statement, err := db.Prepare("INSERT INTO transactions (date, description, cents, notes, updatedAt, category) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(t.Date, t.Description, t.Cents, t.Notes, timeNow(), t.Category)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func EditTransaction(t domain.Transaction) error {
	statement, err := db.Prepare("UPDATE transactions SET date = ?, description = ?, cents = ?, notes = ?, category = ?, updatedAt = ? WHERE id = ? AND committedAt IS NULL")
	if err != nil {
		return err
	}

	_, err = statement.Exec(t.Date, t.Description, t.Cents, t.Notes, t.Category, timeNow(), t.Id)
	return err
}

func CommitTransaction(id int64, cents int64) {
	balance := CommittedBalance() + cents
	statement, _ := db.Prepare("UPDATE transactions SET balance = ?, committedAt = ?, updatedAt = ? WHERE id = ? AND committedAt IS NULL")

	_, err := statement.Exec(balance, timeNow(), timeNow(), id)
	if err != nil {
		log.Fatal(fmt.Sprintf("CommitTransaction(): %s", err))
	}
}

func CommittedBalance() int64 {
	if len(CommittedTransactions()) == 0 {
		t, _ := FindTransaction(0)
		return t.Cents
	} else {
		var balance int64
		statement, _ := db.Prepare("SELECT IFNULL(balance, (SELECT cents FROM transactions WHERE id=0)) FROM transactions WHERE committedAt IS NOT NULL ORDER BY committedAt LIMIT 1;")
		err := statement.QueryRow().Scan(&balance)
		if err != nil {
			log.Fatal(fmt.Sprintf("CommittedBalance(): %s", err))
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
	var category string

	statement, _ := db.Prepare("SELECT id, date, description, cents, notes, category FROM transactions WHERE id = ?")
	err := statement.QueryRow(id).Scan(&id, &date, &description, &cents, &notes, &category)
	if err != nil {
		log.Fatal(fmt.Sprintf("FindTransaction(): %s", err))
	}
	ts, _ := time.Parse(timeLayout, date)

	return domain.Transaction{id, ts, description, cents, notes, category}, err
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
	var cents int64

	statement, _ := db.Prepare("SELECT IFNULL(committedAt,''), updatedAt, notes, cents FROM transactions WHERE id = ?")
	err := statement.QueryRow(id).Scan(&committedAt, &updatedAt, &notes, &cents)
	if err != nil {
		log.Fatal(err)
	}

	return domain.Attributes{
		committedAt != "",
		notes != "",
		parseDate(updatedAt).UnixMicro() < SaveTime.UnixMicro(),
		cents < 0,
	}
}

func TransactionSum(startTime time.Time, endTime time.Time) int64 {
	var sum int64

	statement, _ := db.Prepare("SELECT COALESCE(SUM(cents),0) FROM transactions WHERE date >= ? AND date < ?")
	err := statement.QueryRow(startTime, endTime).Scan(&sum)

	if err != nil {
		log.Fatal(err)
	}

	return sum
}

func timeNow() string {
	return time.Now().Format(timeLayout)
}
