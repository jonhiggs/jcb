package db

import (
	"log"
	"time"
)

func TransactionSum(startTime time.Time, endTime time.Time) int64 {
	var sum int64

	statement, _ := db.Prepare("SELECT COALESCE(SUM(cents),0) FROM transactions WHERE date >= ? AND date < ?")
	err := statement.QueryRow(startTime, endTime).Scan(&sum)

	if err != nil {
		log.Fatal(err)
	}

	return sum
}
