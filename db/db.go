package db

import (
	"database/sql"
	"fmt"
	"io"
	"jcb/domain"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var workingFile string
var saveFile string
var Dirty bool

func Init(file string) error {
	var err error
	Dirty = false
	makeConfigDir(file)
	saveFile = file
	workingFile = makeWorkingFile()

	db, err = sql.Open("sqlite3", workingFile)

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

	year := time.Now().Year()

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

func makeWorkingFile() string {
	src, err := os.Open(saveFile)
	check(err)
	defer src.Close()

	dstFile := fmt.Sprintf("%s/.%s.tmp", filepath.Dir(saveFile), filepath.Base(saveFile))
	dest, err := os.Create(dstFile)
	check(err)
	defer dest.Close()

	_, err = io.Copy(dest, src) // check first var for number of bytes copied
	check(err)

	err = dest.Sync()
	check(err)

	return dstFile
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Save() {
	err := os.Remove(saveFile)
	check(err)
	err = os.Rename(workingFile, saveFile)
	check(err)
	makeWorkingFile()
	Dirty = false
}
