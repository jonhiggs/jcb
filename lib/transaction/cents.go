package transaction

import (
	"fmt"
	"jcb/lib/format"
	"jcb/lib/validate"
	"strconv"
	"strings"
)

type Cents struct {
	value int
}

// Get the cents.
func (c *Cents) GetValue() int { return (*c).value }

// Set the cents.
func (c *Cents) SetValue(i int) error {
	(*c).value = i
	return nil
}

// Get cents as a string
func (c *Cents) GetText() string {
	return format.CentsString((*c).value)
}

// Set the cents from a string.
func (c *Cents) SetText(s string) error {
	err := validate.Cents(s)
	if err != nil {
		return fmt.Errorf("setting cents from string: %w", err)
	}

	s = strings.Trim(s, " ")

	if len(strings.Split(s, ".")) == 1 {
		s = fmt.Sprintf("%s.00", s)
	}

	s = strings.Replace(s, ".", "", 1)
	i, _ := strconv.Atoi(s)

	(*c).value = i

	return nil
}

func (c *Cents) String() string {
	return fmt.Sprintf("%10s", c.GetText())
}
