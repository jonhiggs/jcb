package transaction

import (
	"errors"
	"fmt"
	"jcb/db"
)

// Update or Create a transaction in the Database.
func (t *Transaction) Save() error {
	if t.Committed {
		return errors.New("Cannot modify committed transactions")
	}

	if t.Id == -1 {
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

		res, err := statement.Exec(
			t.Date.GetValue(),
			t.Description.GetValue(),
			t.Cents.GetValue(),
			t.Note.GetValue(),
			db.TimeNow(),
			t.Category.GetValue(),
		)
		if err != nil {
			return err
		}

		id, _ := res.LastInsertId()
		t.Id = int(id)
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
		_, err = statement.Exec(
			t.Date.GetValue(),
			t.Description.GetValue(),
			t.Cents.GetValue(),
			t.Note.GetValue(),
			t.Category.GetValue(),
			db.TimeNow(),
			t.Id,
		)

		if err != nil {
			return err
		}
	}
	return nil
}

// Delete a transaction
func (t *Transaction) Delete() error {
	if t.Committed {
		return errors.New("Cannot delete committed transactions")
	}
	statement, err := db.Conn.Prepare(`
		DELETE FROM transactions WHERE id = ? AND committedAt IS NULL
	`)
	if err != nil {
		return err
	}
	_, err = statement.Exec(t.Id)

	db.Dirty = true

	return err
}

// Commit a transaction
func (t *Transaction) Commit() error {
	err := t.IsCommittable()
	if err != nil {
		return fmt.Errorf("commit failed: %w", err)
	}

	statement, _ := db.Conn.Prepare("UPDATE transactions SET balance = ?, committedAt = ?, updatedAt = ? WHERE id = ? AND committedAt IS NULL")

	_, err = statement.Exec(
		t.Balance().GetValue(),
		db.TimeNow(),
		db.TimeNow(),
		t.Id,
	)
	if err != nil {
		panic(err)
	}

	return nil
}

// Uncommit a transaction
func (t *Transaction) Uncommit() error {
	statement, _ := db.Conn.Prepare("UPDATE transactions SET balance = NULL, committedAt = NULL, updatedAt = ? WHERE id = ? AND committedAt IS NOT NULL")

	_, err := statement.Exec(db.TimeNow(), t.Id)
	if err != nil {
		panic(err)
	}

	return nil
}

// Toggle whether transaction is committed
func (t *Transaction) ToggleCommit() error {
	if t.Committed {
		return t.Uncommit()
	} else {
		return t.Commit()
	}
}
