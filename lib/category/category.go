package category

import (
	"jcb/db"
	"log"
	"time"
)

type Category struct {
	Name string `default:uncategorised`
}

// Returns all the categories
func All() []*Category {
	var categories []*Category

	rows, err := db.Conn.Query(`
		SELECT DISTINCT IFNULL(category, "uncategorised")
		FROM transactions
		ORDER BY category
		`, "",
	)
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

		categories = append(categories, &Category{category})
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
