package db

import (
	"database/sql"
	"errors"
	"fmt"
	"jcb/domain"
	"log"
	"strconv"
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
	rows, err := db.Query("SELECT id, date, description, cents FROM transactions WHERE date LIKE ? AND committedAt IS NULL ORDER BY date, id ASC", fmt.Sprintf("%d%%", year))
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

func LatestYear() (int, error) {
	var year string
	statement, _ := db.Prepare("SELECT DISTINCT substring(date,1,4) FROM transactions WHERE committedAt NOT NULL ORDER BY date DESC LIMIT 1")
	err := statement.QueryRow().Scan(&year)
	if err != nil {
		return -1, err
	}
	y, err := strconv.Atoi(year)
	return y + 1, err
}

func DateSpan() (time.Time, time.Time, time.Time, error) {
	var oS string
	var fS string
	var lS string

	var o time.Time
	var f time.Time
	var l time.Time

	statement, _ := db.Prepare("SELECT date FROM transactions WHERE id = 0")
	err := statement.QueryRow().Scan(&oS)
	if err != nil {
		return time.Unix(0, 0), time.Unix(0, 0), time.Unix(0, 0), err
	}
	o, _ = time.Parse(timeLayout, oS)

	statement, _ = db.Prepare("SELECT date FROM transactions WHERE id != 0 ORDER BY date ASC LIMIT 1")
	err = statement.QueryRow().Scan(&lS)
	if err != nil {
		l = time.Date(o.Year()+1, 12, 31, 23, 59, 59, 59, time.UTC)
	} else {
		l, _ = time.Parse(timeLayout, lS)
	}

	statement, _ = db.Prepare("SELECT date FROM transactions WHERE id != 0 ORDER BY date DESC LIMIT 1")
	err = statement.QueryRow().Scan(&fS)
	if err != nil {
		f = l
	} else {
		f, _ = time.Parse(timeLayout, fS)
	}

	return o, f, l, nil
}
