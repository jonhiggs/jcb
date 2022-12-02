package db

import (
	"database/sql"
	"jcb/domain"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() error {
	var err error
	db, err = sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	sts := `
	    CREATE TABLE IF NOT EXISTS transactions(
	        id INTEGER PRIMARY KEY AUTOINCREMENT,
	        date TEXT, 
			description TEXT, 
			cents INTEGER,
	        UNIQUE(date, description, cents)
	    );
		CREATE TABLE IF NOT EXISTS opening_balance(
			year INTEGER,
			cents INTEGER,
			UNIQUE(year)
		);
		CREATE TABLE IF NOT EXISTS locks(
			id INTEGER,
			UNIQUE(id)
		);
		CREATE TABLE IF NOT EXISTS locks(
			id INTEGER
		);
	`
	_, err = db.Exec(sts)
	return err
}

func FindTransaction(id int64) (domain.Transaction, error) {
	var date string
	var description string
	var cents int64

	statement, _ := db.Prepare("SELECT id, date, description, cents FROM transactions WHERE id = ?")
	err := statement.QueryRow(id).Scan(&id, &date, &description, &cents)

	return domain.Transaction{id, parseDate(date), description, cents}, err
}

func SaveTransaction(t domain.Transaction) (int64, error) {
	if t.Id == 0 {
		statement, err := db.Prepare("INSERT OR IGNORE INTO transactions (date, description, cents) VALUES (?, ?, ?)")
		if err != nil {
			return -1, err
		}
		res, err := statement.Exec(t.Date, t.Description, t.Cents)
		if err != nil {
			return -1, err
		}
		return res.LastInsertId()
	} else {
		statement, err := db.Prepare("UPDATE transactions SET date = ?, description = ?, cents = ? WHERE id = ?")
		if err != nil {
			return -1, err
		}
		_, err = statement.Exec(t.Date, t.Description, t.Cents, t.Id)
		return t.Id, err
	}
}

func AllTransactions() ([]domain.Transaction, error) {
	rows, err := db.Query("SELECT id, date, description, cents FROM transactions ORDER BY date ASC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	const timeLayout = "2006-01-02 03:04:05-07:00"

	var records []domain.Transaction
	for rows.Next() {
		var id int64
		var date string
		var desc string
		var cents int64

		err = rows.Scan(&id, &date, &desc, &cents)
		if err != nil {
			return nil, err
		}

		dateTime, _ := time.Parse(timeLayout, date)
		records = append(records, domain.Transaction{id, dateTime, desc, cents})
	}

	return records, nil
}

func DeleteTransaction(id int64) error {
	statement, err := db.Prepare("DELETE FROM transactions WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	return err
}

func parseDate(timeStr string) time.Time {
	const timeLayout = "2006-01-02 03:04:05-07:00"
	dateTime, _ := time.Parse(timeLayout, timeStr)
	return dateTime
}

func FindOpeningBalance(year int64) (int64, error) {
	var cents int64

	statement, _ := db.Prepare("SELECT cents FROM opening_balance WHERE year = ?")
	err := statement.QueryRow(year).Scan(&cents)

	return cents, err
}

func SaveOpeningBalance(cents int64, year int64) error {
	statement, err := db.Prepare("INSERT INTO opening_balance (year, cents) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(year, cents)
	return err
}

func LockCreateId(id int64) error {
	statement, err := db.Prepare("INSERT OR IGNORE INTO locks (id) VALUES (?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return LockClearId(id + 1)
}

func LockClearId(id int64) error {
	statement, err := db.Prepare("DELETE FROM locks WHERE id >= ?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	return err
}
