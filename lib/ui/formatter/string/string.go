// convert data between data and form strings

package stringFormatter

import (
	"fmt"
	"time"
)

func Cents(i int64) string {
	s := fmt.Sprintf("%d", i)
	var d string
	var c string
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
	return fmt.Sprintf("%s.%s", d, c)
}

func Date(d time.Time) string {
	return d.Format("2006-01-02")
}

func Id(d int64) string {
	return fmt.Sprintf("%d", d)
}
