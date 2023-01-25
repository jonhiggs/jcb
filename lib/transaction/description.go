package transaction

import (
	"errors"
	"fmt"
	"jcb/config"
	"strings"
)

type Description struct {
	value string
}

// Get the string of Description
func (d *Description) GetText() string {
	return (*d).value
}

// Get the string of Description
func (d *Description) GetValue() string {
	return (*d).GetText()
}

// Set the text of Description
func (d *Description) SetText(s string) error {
	s = strings.Trim(s, " ")
	if !validDescription(s) {
		return errors.New("setting description from string")
	}

	(*d).value = s
	return nil
}

// Set the string of Description
func (d *Description) SetValue(s string) error {
	return (*d).SetText(s)
}

// To support the Stringer interface
func (d *Description) String() string {
	return fmt.Sprintf("%-*s", config.DESCRIPTION_MAX_LENGTH, d.value)
}

// return ok if input is valid
func validDescription(string) bool {
	return true
}
