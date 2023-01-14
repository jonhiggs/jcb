// Functions to collect slices of transactions.
package transaction

import (
	"database/sql"
	"fmt"
	"jcb/db"
	"log"
	"time"
)

// Returns every transaction within time range.
func All(startTime time.Time, endTime time.Time) []*Transaction {
	var rows *sql.Rows
	var err error

	var records []*Transaction

	statement, err := db.Conn.Prepare(`
		SELECT id, date, description, cents, notes, category
		FROM (
			SELECT id, date, description, cents, notes, category, IFNULL(committedAt, "2999") AS committedAt
			FROM transactions
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
		var id int64
		var dateString string
		var description string
		var cents int64
		var notes string
		var category string

		err = rows.Scan(
			&id,
			&dateString,
			&description,
			&cents,
			&notes,
			&category,
		)

		if err != nil {
			log.Fatal(fmt.Sprintf("All(): %s", err))
		}

		t := new(Transaction)
		t.id = id
		t.SetDate(db.ParseDate(dateString))
		t.SetDescription(Description(description))
		t.SetCents(cents)
		t.SetNotes(notes)
		t.SetCategory(category)

		records = append(records, t)
	}

	return records
}

// Returns transactions within a category
func FindByCategory(category string, start time.Time, end time.Time) []*Transaction {
	var records []*Transaction

	for _, t := range All(start, end) {
		if t.GetCategory(false) == category {
			records = append(records, t)
		}
	}

	return records
}

// Return a transaction from a transaction ID.
func Find(id int64) (*Transaction, error) {
	var dateString string
	var description string
	var cents int64
	var notes string
	var category string

	statement, _ := db.Conn.Prepare(`
		SELECT id, date, description, cents, notes, category 
		FROM transactions WHERE id = ?
	`)

	err := statement.QueryRow(id).Scan(&id, &dateString, &description, &cents, &notes, &category)
	if err != nil {
		log.Fatal(fmt.Sprintf("Find(): %s", err))
	}

	t := new(Transaction)
	t.id = id
	t.SetDate(db.ParseDate(dateString))
	t.SetDescription(Description(description))
	t.SetCents(cents)
	t.SetNotes(notes)
	t.SetCategory(category)

	return t, nil
}

func FindFirst() (*Transaction, error) {
	var id int64
	var dateString string
	var description string
	var cents int64
	var notes string
	var category string

	statement, _ := db.Conn.Prepare(`
		SELECT id, date, description, cents, notes, category 
		FROM transactions ORDER BY date LIMIT 1
	`)

	err := statement.QueryRow().Scan(&id, &dateString, &description, &cents, &notes, &category)
	if err != nil {
		log.Fatal(fmt.Sprintf("FindFirst(): %s", err))
	}

	t := new(Transaction)
	t.id = id
	t.SetDate(db.ParseDate(dateString))
	t.SetDescription(Description(description))
	t.SetCents(cents)
	t.SetNotes(notes)
	t.SetCategory(category)

	return t, nil
}

func FindLast() (*Transaction, error) {
	var id int64
	var dateString string
	var description string
	var cents int64
	var notes string
	var category string

	statement, _ := db.Conn.Prepare(`
		SELECT id, date, description, cents, notes, category 
		FROM transactions
		ORDER BY date
		DESC LIMIT 1
	`)

	err := statement.QueryRow().Scan(&id, &dateString, &description, &cents, &notes, &category)
	if err != nil {
		log.Fatal(fmt.Sprintf("FindLast(): %s", err))
	}

	t := new(Transaction)
	t.id = id
	t.SetDate(db.ParseDate(dateString))
	t.SetDescription(Description(description))
	t.SetCents(cents)
	t.SetNotes(notes)
	t.SetCategory(category)

	return t, nil
}

func FindLastCommitted() (*Transaction, error) {
	var id int64
	var dateString string
	var description string
	var cents int64
	var notes string
	var category string

	statement, _ := db.Conn.Prepare(`
		SELECT id, date, description, cents, notes, category 
		FROM transactions
		WHERE committedAt IS NOT NULL
		ORDER BY committedAt
		DESC LIMIT 1
	`)

	err := statement.QueryRow().Scan(&id, &dateString, &description, &cents, &notes, &category)
	if err != nil {
		log.Fatal(fmt.Sprintf("FindLastCommitted(): %s", err))
	}

	t := new(Transaction)
	t.id = id
	t.SetDate(db.ParseDate(dateString))
	t.SetDescription(Description(description))
	t.SetCents(cents)
	t.SetNotes(notes)
	t.SetCategory(category)

	return t, nil
}
