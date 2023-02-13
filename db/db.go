package db

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const TimeLayout = "2006-01-02 15:04:05.999999999-07:00"

var Conn *sql.DB
var workingFile string
var saveFile string
var SaveTime time.Time
var Dirty bool

func Init(file string) error {
	var err error
	makeConfigDir(file)
	saveFile = file
	SaveTime = time.Now()
	workingFile = makeWorkingFile()

	fmt.Fprintf(os.Stderr, "Loading file %s\n", file)

	Conn, err = sql.Open("sqlite3", workingFile)

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
	_, err = Conn.Exec(sts)

	statement, err := Conn.Prepare("INSERT OR IGNORE INTO transactions (id, date, description, cents, balance, committedAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		0,
		time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC),
		"Opening Balance",
		0,
		0,
		time.Now(),
		time.Now(),
	)
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
		fmt.Fprint(os.Stderr, "An unsaved file has been found. Would you like to restore it? [y|n] ")
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
	Conn.Close()

	err := os.Rename(workingFile, saveFile)
	check(err)
	makeWorkingFile()
	Dirty = false
	SaveTime = time.Now()

	Conn, err = sql.Open("sqlite3", workingFile)
}

func RemoveWorkingFile() {
	os.Remove(workingFile)
}

func IsDirty() bool {
	var count int

	// this is to handle deletes
	if Dirty {
		return true
	}

	statement, err := Conn.Prepare("SELECT COUNT(*) FROM transactions WHERE updatedAt > ?")
	if err != nil {
		log.Fatal(err)
	}

	err = statement.QueryRow(SaveTime).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

func TimeNow() string {
	return time.Now().Format(TimeLayout)
}
