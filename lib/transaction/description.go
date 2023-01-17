package transaction

import (
	"errors"
	"fmt"
	"jcb/config"
	"jcb/lib/validator"
)

type Description struct {
	Text string
}

func (d *Description) GetText() string { return (*d).Text }

func (d *Description) SetText(v string) (bool, error) {
	err := validator.Description(v)
	if err != nil {
		return false, errors.New("Invalid description")
	}

	if d.String() == v {
		return false, nil
	}

	(*d).Text = v
	return true, nil
}

func (d *Description) String() string {
	return fmt.Sprintf("%-*s", config.DESCRIPTION_MAX_LENGTH, d.Text)
}
