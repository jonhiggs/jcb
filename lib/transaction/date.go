package transaction

import (
	"fmt"
	"jcb/db"
	"strings"
	"time"
)

type Date struct {
	value time.Time
	Saved bool
}

func NewDate() *Date {
	d := new(Date)
	d.Saved = true
	return d
}

func (d *Date) GetValue() time.Time { return (*d).value }

// Set the date.
func (d *Date) SetValue(t time.Time) error {
	(*d).value = t
	(*d).Saved = false
	return nil
}

// Get the string of Date
func (d *Date) GetText() string {
	return (*d).value.Format("2006-01-02")
}

// Set the date from a string.
func (d *Date) SetText(s string) error {
	var ts time.Time
	var err error

	if len(strings.Fields(s)) == 1 {
		if !ValidDBDateString(s) {
			fmt.Errorf("invalid date")
		}
		ts, err = time.Parse("2006-01-02", s)
	} else {
		if !ValidDateString(s) {
			fmt.Errorf("invalid date")
		}
		ts, err = time.Parse(db.TimeLayout, s)
	}

	if err != nil {
		return fmt.Errorf("setting date from string: %w", err)
	}

	(*d).value = ts
	(*d).Saved = false
	return nil
}

func (d *Date) Year() int {
	return int((*d).value.Year())
}

func (d *Date) Unix() int64 {
	return (*d).value.Unix()
}

// return ok if input is valid
func ValidDateString(string) bool {
	// TODO
	return true
}

func ValidDBDateString(string) bool {
	// TODO
	return true
}

func ValidDate(t time.Time) bool {
	lastCommitted, err := FindLastCommitted()
	if err != nil && t.Before(lastCommitted.Date.GetValue()) {
		return false
	}

	return true
}
