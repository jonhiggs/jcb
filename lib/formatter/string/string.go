// convert data to strings

package stringFormatter

import (
	"fmt"
	"jcb/domain"
	"log"
	"strings"
	"time"
)

func Cents(i int64) string {
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
	s = fmt.Sprintf("%s%s.%s", negative, d, c)
	return s
}

func Date(d time.Time) string {
	return d.Format("2006-01-02")
}

func Description(d string) string {
	return strings.Trim(d, " ")
}

func Notes(d string) string {
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

func Transaction(d domain.Transaction) domain.StringTransaction {
	id := Id(d.Id)
	date := Date(d.Date)
	description := Description(d.Description)
	cents := Cents(d.Cents)
	notes := Notes(d.Notes)

	return domain.StringTransaction{id, date, description, cents, notes}
}
