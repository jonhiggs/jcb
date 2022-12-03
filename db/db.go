package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init(file string) error {
	var err error
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
			committedAt TEXT
	    );
	`
	_, err = db.Exec(sts)
	return err
}
