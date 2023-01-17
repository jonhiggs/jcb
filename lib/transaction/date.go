package transaction

import (
	"fmt"
	"strings"
	"time"
)

type Date struct {
	value time.Time
}

func (d *Date) GetValue() time.Time { return (*d).value }

// Set the date.
func (d *Date) SetValue(t time.Time) error {
	(*d).value = t
	return nil
}

// Get the string of Date
func (d *Date) GetText() string { return (*d).value.Format("2006-01-02") }

// Set the date from a string.
func (d *Date) SetText(s string) error {
	var ts time.Time
	var err error

	if len(strings.Fields(s)) == 1 {
		ts, err = time.Parse("2006-01-02", s)
	} else {
		ts, err = time.Parse("2006-01-02 15:04:05.999999999-07:00", s)
	}

	if err != nil {
		return fmt.Errorf("setting date from string: %w", err)
	}

	(*d).value = ts
	return nil
}
