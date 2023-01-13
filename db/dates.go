package db

import (
	"log"
	"time"
)

const timeLayout = "2006-01-02 15:04:05.999999999-07:00"

func DateLastCommitted() time.Time {
	var dateString string
	statement, err := Conn.Prepare("SELECT date FROM transactions WHERE committedAt NOT NULL ORDER BY date DESC, id DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	statement.QueryRow().Scan(&dateString)
	return parseDate(dateString)
}

func DateLastUncommitted() time.Time {
	var dateString string
	statement, err := Conn.Prepare("SELECT date FROM transactions WHERE committedAt IS NULL ORDER BY date DESC, id DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	statement.QueryRow().Scan(&dateString)
	return parseDate(dateString)
}

func DateFirstUncommitted() time.Time {
	var dateString string
	statement, err := Conn.Prepare("SELECT date FROM transactions WHERE id != 0 AND committedAt IS NULL ORDER BY date ASC, id ASC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	statement.QueryRow().Scan(&dateString)
	return parseDate(dateString)
}

func parseDate(s string) time.Time {
	d, _ := time.Parse(timeLayout, s)
	return d
}

func ParseDate(s string) time.Time {
	d, _ := time.Parse(timeLayout, s)
	return d
}

func TimeNow() string {
	return time.Now().Format(timeLayout)
}
