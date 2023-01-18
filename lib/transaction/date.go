package transaction

import (
	"errors"
	"fmt"
	"jcb/db"
	"strings"
	"time"
)

type Date struct {
	value time.Time
}

func (d *Date) GetValue() time.Time { return (*d).value }

// Set the date.
func (d *Date) SetValue(t time.Time) error {
	if !ValidDate(t) {
		return errors.New("invalid date")
	}

	(*d).value = t
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

	return (*d).SetValue(ts)
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

func ValidDate(time.Time) bool {
	// TODO
	return true
}
