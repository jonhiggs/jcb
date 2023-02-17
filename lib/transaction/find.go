// Functions to collect slices of transactions.
package transaction

import (
	"database/sql"
	"errors"
	"fmt"
	"jcb/db"
	"log"
	"time"
)

var ErrNoResult = errors.New("No results")

// Returns every transaction within time range.
func All(startTime time.Time, endTime time.Time) []*Transaction {
	var rows *sql.Rows
	var err error

	var records []*Transaction

	statement, err := db.Conn.Prepare(`
		SELECT id, date, description, cents, notes, category, IFNULL(committedAt, "2999-12-31") AS committedAt
		FROM transactions
		WHERE id == 0
		UNION
		SELECT id, date, description, cents, notes, category, committedAt
		FROM (
			SELECT id, date, description, cents, notes, category, IFNULL(committedAt, "2999-12-31") AS committedAt
			FROM transactions
			WHERE id != 0
			ORDER BY committedAt ASC, date ASC, cents DESC
		)
		WHERE date >= ? AND date <= ?
	`)
	if err != nil {
		log.Fatal(fmt.Sprintf("All(): %s", err))
	}

	rows, err = statement.Query(startTime, endTime)
	if err != nil {
		log.Fatal(fmt.Sprintf("All(): %s", err))
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var date string
		var description string
		var cents int
		var notes string
		var category string
		var committedAt string

		err = rows.Scan(
			&id,
			&date,
			&description,
			&cents,
			&notes,
			&category,
			&committedAt,
		)

		if err != nil {
			log.Fatal(fmt.Sprintf("All(): %s", err))
		}

		t := NewTransaction()
		t.Id = id
		t.Date.SetText(date)
		t.Description.SetValue(description)
		t.Cents.SetValue(cents)
		t.Note.SetValue(notes)
		t.Category.SetValue(category)

		t.Committed = (committedAt != "2999-12-31")

		t.Date.Saved = true
		t.Description.Saved = true
		t.Cents.Saved = true
		t.Note.Saved = true
		t.Category.Saved = true

		records = append(records, t)
	}

	return records
}

// Returns transactions within a category
func FindByCategory(category string, start time.Time, end time.Time) []*Transaction {
	var records []*Transaction

	for _, t := range All(start, end) {
		if t.Category.GetText() == category {
			records = append(records, t)
		}
	}

	return records
}

// Return a transaction from a transaction ID.
func Find(id int) (*Transaction, error) {
	var date string
	var description string
	var cents int
	var notes string
	var category string
	var committedAt string

	statement, _ := db.Conn.Prepare(`
		SELECT id, date, description, cents, notes, category, IFNULL(committedAt, "2999-12-31") AS committedAt
		FROM transactions WHERE id = ?
	`)

	err := statement.QueryRow(id).Scan(&id, &date, &description, &cents, &notes, &category, &committedAt)
	if err != nil {
		log.Fatal(fmt.Sprintf("Find(): %s", err))
	}

	t := new(Transaction)
	t.Id = id
	t.Date.SetText(date)
	t.Description.SetValue(description)
	t.Cents.SetValue(cents)
	t.Note.SetValue(notes)
	t.Category.SetValue(category)

	t.Committed = (committedAt != "2999-12-31")

	t.Date.Saved = true
	t.Description.Saved = true
	t.Cents.Saved = true
	t.Note.Saved = true
	t.Category.Saved = true

	return t, nil
}

func FindLastCommitted() (*Transaction, error) {
	all := All(time.Unix(0, 0), time.Unix(32503554000, 0))

	for i := len(all) - 1; i >= 0; i-- {
		if !all[i].Committed {
			continue
		}

		return all[i], nil
	}

	return nil, errors.New("no committed transactions were found")
}

func FindLastUncommitted() (*Transaction, error) {
	all := All(time.Unix(0, 0), time.Unix(32503554000, 0))
	last := all[len(all)-1]

	if last.Committed {
		return nil, errors.New("no uncommitted transactions were found")
	}

	return last, nil
}
