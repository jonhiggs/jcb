package transaction

import (
	"fmt"
	"jcb/config"
	"jcb/lib/format"
	"jcb/lib/validate"
)

type Description struct {
	value string
}

// Get the string of Description
func (d *Description) GetText() string {
	return format.DescriptionString((*d).value)
}

// Set the text of Description and return ok, error.
func (d *Description) SetText(v string) error {
	err := validate.Description(v)
	if err != nil {
		return fmt.Errorf("setting text to %s: %w", v, err)
	}

	(*d).value = v
	return nil
}

// To support the Stringer interface
func (d *Description) String() string {
	return fmt.Sprintf("%-*s", config.DESCRIPTION_MAX_LENGTH, d.value)
}
