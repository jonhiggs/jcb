package transaction

import (
	"errors"
	"fmt"
	"jcb/config"
	"jcb/lib/validator"
)

type Description string

func (d *Description) Get() string { return string(d) }

func (d *Description) Set(v string) (bool, error) {
	err := validator.Description(v)
	if err != nil {
		return false, errors.New("Invalid description")
	}

	if d.String() == v {
		return false, nil
	}

	d = Description(v)
	return true, nil
}

func (d *Description) String() string { return string(d) }

func (d *Description) PaddedString() string {
	return fmt.Sprintf("%-*s", config.DESCRIPTION_MAX_LENGTH, d.Get())
}
