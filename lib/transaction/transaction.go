// Package transactions2 is a refactor of transactions which will soon be
// discarded. It's purpose is for fetching and manipulating transactions which
// are stored in the transactionbase.
package transaction

import (
	"fmt"
	"jcb/config"
	"jcb/db"
	"log"
	"strings"
)

// A transaction is an event that either has happened or you predict will
// happen.
type Transaction struct {
	id          int64 `default:-1`
	Date        Date
	Description Description
	Cents       Cents
	notes       string
	category    string
}

// Set the cents. Returns true if value was changed.
//func (t *Transaction) SetCents(i int64) bool {
//	if t.cents == i {
//		return false
//	}
//
//	t.cents = i
//	return true
//}

// Set the amount. Returns true if value was changed.
//func (t *Transaction) SetAmount(s string) bool {
//	if validator.Cents(s) != nil {
//		log.Fatal(fmt.Sprintf("cannot convert invalid cents '%s' to data", s))
//	}
//
//	s = strings.Trim(s, " ")
//
//	if len(strings.Split(s, ".")) == 1 {
//		s = fmt.Sprintf("%s.00", s)
//	}
//
//	i, _ := strconv.ParseInt(strings.Replace(s, ".", "", 1), 10, 64)
//
//	return t.Cents.SetValue(i)
//}

func (t *Transaction) SetNotes(s string) {
	t.notes = s
}

// Set the category. Returns true if value was changed.
func (t *Transaction) SetCategory(s string) bool {
	if t.category == s {
		return false
	}

	t.category = s
	return true
}

func (t *Transaction) GetID() int64 {
	return t.id
}

// Returns the transaction description. Expects a bool argument that when true
// will pad the string for presentation in a table.
func (t *Transaction) GetCategory(pad bool) string {
	s := strings.Trim(t.category, " ")

	if len(s) > config.CATEGORY_MAX_LENGTH {
		s = s[0:config.CATEGORY_MAX_LENGTH]
	}

	if pad {
		return fmt.Sprintf("%-*s", config.CATEGORY_MAX_LENGTH, s)
	} else {
		return s
	}
}

func (t *Transaction) Balance() int64 {
	return 0
}

// Returns true if the transaction has been committed. A committed transaction
// is one that has been reconciled against the bank statement.
func (t *Transaction) IsCommitted() bool {
	var field string
	if t.id == -1 {
		return false
	}
	statement, _ := db.Conn.Prepare(`
		SELECT IFNULL(committedAt,"") FROM transactions
		WHERE id = ?
	`)
	err := statement.QueryRow(t.id).Scan(&field)
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
	err := statement.QueryRow(t.id).Scan(&field)
	if err != nil {
		log.Fatal(fmt.Sprintf("IsSaved(): %s", err))
	}

	return db.ParseDate(field).UnixMicro() < db.SaveTime.UnixMicro()
}

// Returns true if the transaction has notes associated with it.
func (t *Transaction) HasNotes() bool {
	var field string
	statement, _ := db.Conn.Prepare("SELECT notes FROM transactions WHERE id = ?")
	err := statement.QueryRow(t.id).Scan(&field)
	if err != nil {
		log.Fatal(fmt.Sprintf("HasNotes(): %s", err))
	}

	return field != ""
}

// Returns the notes as a well-formed string.
func (t *Transaction) GetNotes() string {
	return t.notes
}

// Returns the attributes string
func (t *Transaction) GetAttributeString() string {
	s := ""
	if t.IsCommitted() {
		s += "C"
	} else {
		s += " "
	}

	if t.HasNotes() {
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

// Returns true if the transaction is a negative amount.
func (t *Transaction) IsDebit() bool {
	return t.Cents.GetValue() < 0
}

// Returns true if the transaction is a positive amount.
func (t *Transaction) IsCredit() bool {
	return t.Cents.GetValue() >= 0
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
