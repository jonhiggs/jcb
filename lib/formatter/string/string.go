// convert data to strings

package stringFormatter

import (
	"fmt"
	"jcb/domain"
	"log"
	"strings"
)

func Category(d string) string {
	return strings.Trim(d, " ")
}

func Id(d int64) string {
	s := fmt.Sprintf("%d", d)

	if d < 0 {
		s = "0"
		log.Fatal("Id cannot be less than 0")
	}
	return s
}

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
