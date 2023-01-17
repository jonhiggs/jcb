// convert a string from an input box and format into data
package format

import (
	"fmt"
)

// return cents as a string
func CentsString(i int) string {
	var d string
	var c string

	negative := ""
	if i < 0 {
		negative = "-"
		i = i * -1
	}

	s := fmt.Sprintf("%d", i)

	if len(s) == 2 {
		d = "0"
		c = s
	} else if len(s) == 1 {
		d = "0"
		c = fmt.Sprintf("0%s", s)
	} else {
		d = s[0 : len(s)-2]
		c = s[len(s)-2 : len(s)]
	}

	return fmt.Sprintf("%s%s.%s", negative, d, c)
}
