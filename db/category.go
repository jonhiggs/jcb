package db

import (
	"database/sql"
	"log"
	"time"
)

func Categories() []string {
	var rows *sql.Rows
	var err error
	var categories []string

	rows, err = db.Query("SELECT DISTINCT category FROM transactions ORDER BY category", "")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category string

		err = rows.Scan(&category)
		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func CategorySum(category string, startTime time.Time, endTime time.Time) int64 {
	var sum int64

	statement, _ := db.Prepare("SELECT COALESCE(SUM(cents),0) FROM transactions WHERE category = ? AND date >= ? AND date < ?")
	err := statement.QueryRow(category, startTime, endTime).Scan(&sum)

	if err != nil {
		log.Fatal(err)
	}

	return sum
}
