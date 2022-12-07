package db

import (
	"database/sql"
	"jcb/domain"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init(file string, year int) error {
	var err error
	makeConfigDir(file)
	db, err = sql.Open("sqlite3", file)

	if err != nil {
		log.Fatal(err)
	}

	sts := `
	    CREATE TABLE IF NOT EXISTS transactions(
	        id INTEGER PRIMARY KEY AUTOINCREMENT,
	        date TEXT,
			description TEXT,
			cents INTEGER,
			balance INTEGER,
			committedAt TEXT,
			UNIQUE(id)
	    );
		CREATE TABLE IF NOT EXISTS balances(
			year INTEGER,
			closing INTEGER,
			UNIQUE(year)
		);
	`
	_, err = db.Exec(sts)

	statement, err := db.Prepare("INSERT OR IGNORE INTO transactions (id, date, description, cents) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	t := domain.Transaction{0, time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC), "Opening Balance", 0}
	_, err = statement.Exec(t.Id, t.Date, t.Description, t.Cents)
	if err != nil {
		return err
	}

	return err
}

func makeConfigDir(file string) {
	dir := filepath.Dir(file)
	os.MkdirAll(dir, 0700)
}
