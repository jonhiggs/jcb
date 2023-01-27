// Package transactions2 is a refactor of transactions which will soon be
// discarded. It's purpose is for fetching and manipulating transactions which
// are stored in the transactionbase.
package transaction

import (
	"errors"
	"fmt"
	"jcb/db"
	"log"
	"time"
)

// A transaction is an event that either has happened or you predict will
// happen.
type Transaction struct {
	Id          int64 `default:"-1"`
	Date        Date
	Description Description
	Cents       Cents
	Note        Note
	Category    Category
}

type TextSetter interface {
	SetText(s string) error
}

// Set fields from text, returns error
func (t *Transaction) SetText(data []string) error {
	dataFields := []TextSetter{
		&t.Date,
		&t.Category,
		&t.Description,
		&t.Cents,
		&t.Note,
	}

	for colN, f := range dataFields {
		err := f.SetText(data[colN])
		if err != nil {
			return err
		}
	}

	return nil
}

// Returns true if the transaction has been committed. A committed transaction
// is one that has been reconciled against the bank statement.
func (t *Transaction) IsCommitted() bool {
	var field string
	if t.Id == -1 {
		return false
	}
	statement, _ := db.Conn.Prepare(`
		SELECT IFNULL(committedAt,"") FROM transactions
		WHERE id = ?
	`)
	err := statement.QueryRow(t.Id).Scan(&field)
	if err != nil {
		log.Fatal(fmt.Sprintf("IsCommitted(): %s", err))
	}

	return field != ""
}

// Returns true if the transaction has been saved to the permanent transactionbase. A
// working transactionbase is used on startup which is only flushed to the save file
// with the `:w` command.
func (t *Transaction) IsSaved() bool {
	var field string
	statement, _ := db.Conn.Prepare("SELECT updatedAt FROM transactions WHERE id = ?")
	err := statement.QueryRow(t.Id).Scan(&field)
	if err != nil {
		log.Fatal(fmt.Sprintf("IsSaved(): %s", err))
	}

	saveTime, _ := time.Parse(db.TimeLayout, field)
	return saveTime.UnixMicro() < db.SaveTime.UnixMicro()
}

// Returns the attributes string
func (t *Transaction) GetAttributeString() string {
	s := ""
	if t.IsCommitted() {
		s += "C"
	} else {
		s += " "
	}

	if t.Note.Exists() {
		s += "n"
	} else {
		s += " "
	}

	if t.IsSaved() {
		s += " "
	} else {
		s += "+"
	}

	return s
}

// Returns true if transaction is immediately ready to be committed.
func (t *Transaction) IsCommittable() error {
	if t.IsCommitted() {
		return errors.New("Transaction is already committed")
	}

	lastCommitted, _ := FindLastCommitted()
	startTime := lastCommitted.Date.GetValue()
	endTime := time.Date(startTime.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, tt := range All(startTime, endTime) {
		if tt.IsCommitted() {
			continue
		}

		if tt.Id == t.Id {
			break
		}

		// return false if there are any transactions before 't'.
		if t.Date.GetValue().After(tt.Date.GetValue()) {
			return errors.New("you must first commit all older transactions")
		}
	}

	return nil
}

// Return false if a similar transaction already exists.
func (t *Transaction) IsUniq() bool {
	var count int
	statement, err := db.Conn.Prepare(`
		SELECT COUNT(*)
		FROM transactions
		WHERE substr(date, 0,11) == ?
		  AND description == ?
		  AND cents == ?
  	`)
	if err != nil {
		log.Fatal(fmt.Sprintf("IsUniq(): %s", err))
	}

	err = statement.QueryRow(t.Date.GetValue(), t.Description.GetText(), t.Cents.GetValue()).Scan(&count)
	if err != nil {
		log.Fatal(fmt.Sprintf("IsUniq(): %s", err))
	}

	return count == 0
}

func DateRange() (time.Time, time.Time) {
	first, _ := FindFirst()
	last, _ := FindLast()
	return first.Date.GetValue(), last.Date.GetValue()
}

func SumCents(ts []*Transaction) int {
	var sum int
	for _, t := range ts {
		sum += t.Cents.GetValue()
	}
	return sum
}

func (t *Transaction) Balance() *Cents {
	b := new(Cents)

	// the opening balance
	if t.Id == 0 {
		return &t.Cents
	}

	if t.IsCommitted() {
		var balance int

		statement, _ := db.Conn.Prepare(`
			SELECT balance 
			FROM transactions WHERE id = ?
		`)

		err := statement.QueryRow(t.Id).Scan(&balance)
		if err != nil {
			panic(err)
		}

		b.SetValue(balance)
		return b
	}

	lastCommitted, err := FindLastCommitted()
	if err != nil {
		b.SetValue(0)
	} else {
		var balance int

		statement, _ := db.Conn.Prepare(`
			SELECT balance 
			FROM transactions WHERE id = ?
		`)

		err := statement.QueryRow(lastCommitted.Id).Scan(&balance)
		if err != nil {
			panic(err)
		}

		b.SetValue(balance)
	}

	startTime := lastCommitted.Date.GetValue()
	endTime := time.Date(startTime.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, t := range All(startTime, endTime) {
		if t.IsCommitted() {
			continue
		}
		b.Add(t.Cents.GetValue())
	}

	return b
}
