package db

import (
	"errors"
	"fmt"
	"jcb/domain"
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
