package category

import (
	"fmt"
	"jcb/db"
	"log"
	"time"
)

type Category struct {
	Name      string `default:uncategorised`
	StartTime time.Time
	EndTime   time.Time
}

// Returns all the categories
func All(startTime time.Time, endTime time.Time) []*Category {
	var categories []*Category

	statement, err := db.Conn.Prepare(`
		SELECT DISTINCT IFNULL(category, "uncategorised")
		FROM transactions
		WHERE date >= ? AND date <= ?
		ORDER BY category
	`)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := statement.Query(startTime, endTime)
	if err != nil {
		log.Fatal(fmt.Sprintf("All(): %s", err))
	}

	defer rows.Close()

	for rows.Next() {
		var category string

		err = rows.Scan(&category)
		if err != nil {
			log.Fatal(err)
		}

		categories = append(
			categories,
			&Category{
				Name:      category,
				StartTime: startTime,
				EndTime:   endTime,
			})
	}

	return categories
}

// Returns a categories the total cents of transactions within a time range.
func (cat *Category) SumCents(startTime time.Time, endTime time.Time) int64 {
	var sum int64

	statement, _ := db.Conn.Prepare(`
		SELECT COALESCE(SUM(cents),0)
		FROM transactions
		WHERE category = ? AND date >= ? AND date < ?
	`)

	err := statement.QueryRow(cat.Name, startTime, endTime).Scan(&sum)
	if err != nil {
		log.Fatal(err)
	}

	return sum
}
