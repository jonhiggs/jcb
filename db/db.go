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
var SaveTime time.Time
var dirty bool

func Init(file string) error {
	var err error
	makeConfigDir(file)
	saveFile = file
	SaveTime = time.Now()
	workingFile = makeWorkingFile()

	db, err = sql.Open("sqlite3", workingFile)

	if err != nil {
		log.Fatal(err)
	}

	sts := `
	    CREATE TABLE IF NOT EXISTS transactions(
	        id INTEGER PRIMARY KEY AUTOINCREMENT,
	        date TEXT,
			category TEXT default '',
			description TEXT DEFAULT '',
			notes TEXT DEFAULT '',
			cents INTEGER,
			balance INTEGER,
			committedAt TEXT,
			updatedAt TEXT,
			UNIQUE(id)
	    );
	`
	_, err = db.Exec(sts)

	statement, err := db.Prepare("INSERT OR IGNORE INTO transactions (id, date, description, cents, committedAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	year := time.Now().Year()

	t := domain.Transaction{0, time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC), "Opening Balance", 0, "", ""}
	_, err = statement.Exec(t.Id, t.Date, t.Description, t.Cents, time.Now(), time.Now())
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
	_, err := os.Stat(saveFile)
	if err != nil {
		os.Create(saveFile)
	}

	src, err := os.Open(saveFile)
	check(err)
	defer src.Close()

	dstFile := fmt.Sprintf("%s/.%s.tmp", filepath.Dir(saveFile), filepath.Base(saveFile))
	_, err = os.Stat(dstFile)
	if err == nil {
		fmt.Println("An unsaved file has been found. Would you like to restore it? [y|n]")
		var choice rune
		fmt.Scanf("%c", &choice)

		switch choice {
		case 'y':
			return dstFile
		case 'n':
			os.Remove(dstFile)
		default:
			fmt.Println("Invalid choice")
			os.Exit(1)
		}
	}

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
	db.Close()

	err := os.Rename(workingFile, saveFile)
	check(err)
	makeWorkingFile()
	dirty = false
	SaveTime = time.Now()

	db, err = sql.Open("sqlite3", workingFile)
}

func RemoveWorkingFile() {
	os.Remove(workingFile)
}

func Dirty() bool {
	var count int

	// this is to handle deletes
	if dirty {
		return true
	}

	statement, err := db.Prepare("SELECT COUNT(*) FROM transactions WHERE updatedAt > ?")
	if err != nil {
		log.Fatal(err)
	}

	err = statement.QueryRow(SaveTime).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}
