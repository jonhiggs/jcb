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
	        date TEXT, description TEXT, cents INTEGER,
	        UNIQUE(date, description, cents)
	    );
	    CREATE TABLE IF NOT EXISTS tags(
	        id INTEGER PRIMARY KEY AUTOINCREMENT,
	        name TEXT,
	        UNIQUE(name)
	    );
	    CREATE TABLE IF NOT EXISTS tag_associations(
	        transaction_id INTEGER,
	        tag_id INTEGER,
	        UNIQUE(transaction_id, tag_id)
	    );
	`
	_, err = db.Exec(sts)
	return err
}

func FindTransaction(id int64) (domain.Transaction, error) {
	var date time.Time
	var description string
	var cents int64

	statement, _ := db.Prepare("SELECT id, date, description, cents FROM transactions WHERE id = ?")
	err := statement.QueryRow(id).Scan(&id, &date, &description, &cents)
	return domain.Transaction{id, date, description, cents}, err
}

func SaveTransaction(t domain.Transaction) error {
	if t.Id == 0 {
		statement, err := db.Prepare("INSERT OR IGNORE INTO transactions (date, description, cents) VALUES (?, ?, ?)")
		if err != nil {
			return err
		}
		_, err = statement.Exec(t.Date, t.Description, t.Cents)
		return err
	} else {
		statement, err := db.Prepare("UPDATE transactions SET date = ?, description = ?, cents = ? WHERE id = ?")
		if err != nil {
			return err
		}
		_, err = statement.Exec(t.Date, t.Description, t.Cents, t.Id)
		return err
	}
}

func AllTransactions() ([]domain.Transaction, error) {
	rows, err := db.Query("SELECT * FROM transactions")
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
