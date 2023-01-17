// convert a string from an input box and format into data
package format

import (
	"fmt"
	"jcb/lib/validate"
	"strings"
)

func Description(s string) (string, error) {
	s = strings.Trim(s, " ")

	_, err := validate.Description(s)
	if err != nil {
		return s, fmt.Errorf("formatting description: %w", err)
	}

	return s, nil
}
