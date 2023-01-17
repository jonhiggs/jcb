package transaction

import (
	"fmt"
	"jcb/config"
	"jcb/lib/validate"
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
	err := validate.Description(s)
	if err != nil {
		return fmt.Errorf("setting description from string: %w", err)
	}

	(*d).value = s
	return nil
}

// To support the Stringer interface
func (d *Description) String() string {
	return fmt.Sprintf("%-*s", config.DESCRIPTION_MAX_LENGTH, d.value)
}
