package transaction

import (
	"errors"
	"jcb/db"
	"time"
)

// Update or Create a transaction in the Database.
func (t *Transaction) Save() error {
	if t.IsCommitted() {
		return errors.New("Cannot modify committed transactions")
	}

	if t.id == -1 {
		statement, err := db.Conn.Prepare(`
			INSERT INTO transactions (
				date,
				description,
				cents,
				notes,
				updatedAt,
				category
			) VALUES (?, ?, ?, ?, ?, ?)
		`)
		if err != nil {
			return err
		}

		res, err := statement.Exec(t.date, t.Description, t.cents, t.notes, db.TimeNow(), t.category)
		if err != nil {
			return err
		}

		t.id, _ = res.LastInsertId()
	} else {
		statement, err := db.Conn.Prepare(`
			UPDATE transactions
			SET
				date = ?,
				description = ?,
				cents = ?,
				notes = ?,
				category = ?,
				updatedAt = ?
			WHERE id = ?
				AND committedAt IS NULL
		`)
		if err != nil {
			return err
		}
		_, err = statement.Exec(t.date, t.Description, t.cents, t.notes, t.category, db.TimeNow(), t.id)

		if err != nil {
			return err
		}
	}
	return nil
}

// Delete a transaction
func (t *Transaction) Delete() error {
	if t.IsCommitted() {
		return errors.New("Cannot delete committed transactions")
	}
	statement, err := db.Conn.Prepare(`
		DELETE FROM transactions WHERE id = ? AND committedAt IS NULL
	`)
	if err != nil {
		return err
	}
	_, err = statement.Exec(t.GetID())

	db.Dirty = true

	return err
}

// Commit a transaction
func (t *Transaction) Commit() error {
	if !t.IsCommittable() {
		return errors.New("Commit older transactions first")
	}

	return errors.New("Not implemented")

	//balance := CommittedBalance() + cents
	//statement, _ := db.Prepare("UPDATE transactions SET balance = ?, committedAt = ?, updatedAt = ? WHERE id = ? AND committedAt IS NULL")

	//_, err := statement.Exec(balance, timeNow(), timeNow(), id)
	//if err != nil {
	//	log.Fatal(fmt.Sprintf("Committransaction(): %s", err))
	//}

	return nil
}

// Uncommit a transaction
func (t *Transaction) Uncommit() error {
	return errors.New("Not implemented")
	//balance := CommittedBalance() + cents
	//statement, _ := db.Prepare("UPDATE transactions SET balance = ?, committedAt = ?, updatedAt = ? WHERE id = ? AND committedAt IS NULL")

	//_, err := statement.Exec(balance, timeNow(), timeNow(), id)
	//if err != nil {
	//	log.Fatal(fmt.Sprintf("Committransaction(): %s", err))
	//}

	return nil
}

func DateRange() (time.Time, time.Time) {
	first, _ := FindFirst()
	last, _ := FindLast()
	return first.GetDate(), last.GetDate()
}

func SumCents(ts []*Transaction) int64 {
	var sum int64
	for _, t := range ts {
		sum += t.GetCents()
	}
	return sum
}
