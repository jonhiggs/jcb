// convert data to strings

package stringFormatter

import (
	"jcb/domain"
)

func Attributes(a domain.Attributes) string {
	s := ""
	if a.Committed {
		s += "C"
	} else {
		s += " "
	}

	if a.Notes {
		s += "n"
	} else {
		s += " "
	}

	if a.Saved {
		s += " "
	} else {
		s += "+"
	}

	return s
}
