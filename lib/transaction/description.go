package transaction

import (
	"fmt"
	"jcb/config"
	"jcb/lib/validate"
)

type Description struct {
	Text string
}

// Get the string of Description
func (d *Description) GetText() string { return (*d).Text }

// Set the text of Description and return ok, error.
func (d *Description) SetText(v string) (bool, error) {
	_, err := validate.Description(v)
	if err != nil {
		return false, fmt.Errorf("setting text to %s: %w")
	}

	if (*d).Text == v {
		return false, nil
	} else {
		(*d).Text = v
		return true, nil
	}
}

// To support the Stringer interface
func (d *Description) String() string {
	return fmt.Sprintf("%-*s", config.DESCRIPTION_MAX_LENGTH, d.Text)
}
