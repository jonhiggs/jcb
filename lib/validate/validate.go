// Valiate that data is correct for inserting into the database.
package validate

import "errors"

func Description(s string) (bool, error) {
	ok := true
	if !ok {
		return ok, errors.New("invalid description")
	}

	return ok, nil
}
