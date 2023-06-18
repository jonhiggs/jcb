package db

import (
	"database/sql"
	"errors"
	"fmt"
	"jcb/domain"
	"log"
)

func InsertBudget(b domain.Budget) (int64, error) {
	if b.Id != -1 {
		return -1, errors.New(fmt.Sprintf("Budget %d already exists", b.Id))
	}

	statement, err := db.Prepare("INSERT INTO budgets (date, category, cumulative, notes, cents, updatedAt) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(b.Date, b.Category, b.Cumulative, b.Notes, b.Cents, timeNow())
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func Budgets() []domain.Budget {
	var rows *sql.Rows
	var err error

	var records []domain.Budget

	rows, err = db.Query("SELECT id, date, category, cents, notes, cumulative FROM budgets ORDER BY date ASC", "")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var dateString string
		var category string
		var cumulative int
		var cents int64
		var notes string

		err = rows.Scan(&id, &dateString, &category, &cents, &notes, &cumulative)
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, domain.Budget{id, parseDate(dateString), category, cents, notes, cumulative == 1})
	}

	return records
}
