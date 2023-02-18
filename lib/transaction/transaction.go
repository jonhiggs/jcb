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
	Id          int
	Date        Date
	Description Description
	Cents       Cents
	Note        Note
	Category    Category
	Tagged      bool
	FindMatch   bool
	Committed   bool
}

type TextSetter interface {
	SetText(s string) error
}

func NewTransaction() *Transaction {
	t := new(Transaction)
	t.Id = -1
	t.Tagged = false
	// TODO: set find match
	return t
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

// Returns true if the transaction has been saved to the permanent transactionbase. A
// working transactionbase is used on startup which is only flushed to the save file
// with the `:w` command.
func (t *Transaction) IsSaved() bool {
	if t.Id == -1 {
		return false
	}

	if t.Description.Saved && t.Cents.Saved && t.Category.Saved && t.Date.Saved && t.Note.Saved {
		return true
	}

	return false
}

// Returns the attributes string
func (t *Transaction) GetAttributeString() string {
	s := ""
	if t.Committed {
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
	if t.Committed {
		return errors.New("transaction is already committed")
	}

	lastCommitted, err := FindLastCommitted()
	if err != nil {
		if t.Id != 0 {
			return errors.New("commit the opening balance transaction first")
		}
		return nil
	}

	startTime := lastCommitted.Date.GetValue()
	endTime := time.Date(startTime.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, tt := range All(startTime, endTime) {
		if tt.Committed {
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
	all := All(time.Unix(0, 0), time.Unix(32503554000, 0))
	return all[0].Date.GetValue(), all[len(all)-1].Date.GetValue()
}

func SumCents(ts []*Transaction) int {
	var sum int
	for _, t := range ts {
		sum += t.Cents.GetValue()
	}
	return sum
}

func (t *Transaction) Balance() *Cents {
	b := NewCents()

	// the opening balance
	if t.Id == 0 {
		return &t.Cents
	}

	if t.Committed {
		var balance int

		statement, _ := db.Conn.Prepare(`
			SELECT balance
			FROM transactions WHERE id = ?
		`)

		err := statement.QueryRow(t.Id).Scan(&balance)
		if err != nil {
			panic(fmt.Sprintf("%s for %d", err, t.Id))
		}

		b.SetValue(balance)
		return b
	}

	// set balance to that of the last committed transaction
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
			panic(fmt.Sprintf("%s for %d", err, t.Id))
		}

		b.SetValue(balance)
	}

	// Add the rolling tally of the uncommitted transactions to the balance.
	for _, x := range All(time.Unix(0, 0), t.Date.GetValue()) {
		if x.Committed {
			b = x.Balance()
			continue
		}

		b.Add(x.Cents.GetValue())

		if x.Id == t.Id {
			break
		}
	}

	return b
}
