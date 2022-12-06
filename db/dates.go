package db

import (
	"log"
	"time"
)

//func DateFirstCommitted() time.Time {
//	var dateString string
//
//	statement, _ := db.Prepare("SELECT date FROM transactions WHERE committedAt NOT NULL ORDER BY date, id LIMIT 1")
//	err := statement.QueryRow(date).Scan(&dateString)
//	if err != nil {
//		return time.Unix(0, 0)
//	}
//
//	return ParseDate(dateString)
//}

func DateLastCommitted(year int) time.Time {
	if year == -1 {
		var dateString string
		statement, err := db.Prepare("SELECT date FROM transactions WHERE committedAt NOT NULL ORDER BY date DESC, id DESC LIMIT 1")
		if err != nil {
			log.Fatal(err)
		}
		statement.QueryRow().Scan(&dateString)
		return ParseDate(dateString)
	}

	return time.Unix(0, 0)
}

func DateLastUncommitted(year int) time.Time {
	if year == -1 {
		var dateString string
		statement, err := db.Prepare("SELECT date FROM transactions WHERE committedAt IS NULL ORDER BY date DESC, id DESC LIMIT 1")
		if err != nil {
			log.Fatal(err)
		}
		statement.QueryRow().Scan(&dateString)
		return ParseDate(dateString)
	}

	return time.Unix(0, 0)
}

//func DateFirstUncommitted() time.Time {
//	var dateString string
//
//	statement, _ := db.Prepare("SELECT date FROM transactions WHERE committedAt IS NULL ORDER BY date, id LIMIT 1")
//	err := statement.QueryRow(id).Scan(&dateString)
//	if err != nil {
//		return time.Unix(0, 0)
//	}
//
//	return ParseDate(dateString)
//}
//
func ParseDate(s string) time.Time {
	d, _ := time.Parse("2006-01-02 15:04:05.999999999-07:00", s)
	return d
}
