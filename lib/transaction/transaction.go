// Package transactions2 is a refactor of transactions which will soon be
// discarded. It's purpose is for fetching and manipulating transactions which
// are stored in the transactionbase.
package transaction

import (
	"fmt"
	"jcb/db"
	"log"
	"time"
)

// A transaction is an event that either has happened or you predict will
// happen.
type Transaction struct {
	Id          int64       `default:"-1"`
	Date        Date        ``
	Description Description `default:""`
	Cents       Cents       `default:"0"`
	Note        Note        `default:""`
	Category    Category    `default:""`
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
func (t *Transaction) IsCommittable() bool {
	lc, _ := FindLastCommitted()
	for _, tt := range All(lc.Date.GetValue(), t.Date.GetValue()) {
		if tt.Date.GetValue().Unix() > t.Date.GetValue().Unix() {
			return false
		}
	}

	return true
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
