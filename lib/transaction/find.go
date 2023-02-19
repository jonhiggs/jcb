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
	var prevId = -1

	var records []*Transaction

	statement, err := db.Conn.Prepare(`
		SELECT
			id,
			date,
			description,
			cents,
			notes,
			category,
			IFNULL(committedAt, "2999-12-31") AS committedAt,
			IFNULL((SELECT tagged FROM cache WHERE id = 0),0) AS tagged
		FROM transactions
		WHERE id == 0
		UNION
		SELECT id, date, description, cents, notes, category, committedAt,tagged
		FROM (
			SELECT
				id,
				date,
				description,
				cents,
				notes,
				category,
				IFNULL(committedAt, "2999-12-31") AS committedAt,
				IFNULL((SELECT tagged FROM cache WHERE id = transactions.id),0) AS tagged
			FROM transactions
			WHERE id != 0
		)
		WHERE date >= ? AND date <= ?
		ORDER BY committedAt ASC, date ASC, cents DESC
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
		var tagged int

		err = rows.Scan(
			&id,
			&date,
			&description,
			&cents,
			&notes,
			&category,
			&committedAt,
			&tagged,
		)

		if err != nil {
			log.Fatal(fmt.Sprintf("All(): %s", err))
		}

		t := NewTransaction()
		t.Id = id
		if tagged == 0 {
			t.Tagged = false
		} else {
			t.Tagged = true
		}
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
		t.PrevId = prevId
		t.NextId = -1

		records = append(records, t)
		if len(records) > 1 {
			records[len(records)-2].NextId = t.Id
		}
		prevId = t.Id
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
	for _, t := range All(time.Unix(0, 0), time.Unix(32503554000, 0)) {
		if t.Id == id {
			return t, nil
		}
	}

	return nil, errors.New("transaction not found")
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
